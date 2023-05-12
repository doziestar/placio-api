import { hash, compare } from 'bcrypt';
import { sign, verify } from 'jsonwebtoken';
import { Service } from 'typedi';
import { SECRET_KEY } from '@config';
import { HttpException } from '@exceptions/httpException';
import { DataStoredInToken, TokenResponsePayload } from '@interfaces/auth.interface';
import { User } from '@interfaces/users.interface';
import { UserModel } from '@models/users.model';
import { v4 as uuidv4 } from 'uuid';
import _ from 'lodash';
import passport from '@services/social.service';
import { sendEmailVerification, sendPasswordReset, sendWelcomeEmail } from '@services/email.service';
import { verifyPhoneNumber } from './sms.service';
import { deleteToken, saveToken, verifyToken } from '@/models/token.model';
import { createLogin } from '@/models/login.model';
import { Request } from 'express';
import { TokenType } from '@/interfaces/token.interface';

const createToken = async (user: User): Promise<TokenResponsePayload> => {
  const dataStoredInToken: DataStoredInToken = { id: user.id };
  const expiresIn: number = 60 * 60 * 24 * 7;

  const accessToken = sign(dataStoredInToken, SECRET_KEY, { expiresIn });
  const refreshToken = sign(dataStoredInToken, SECRET_KEY, { expiresIn: 60 * 60 * 24 * 30 });

  const tokenData: TokenResponsePayload = {
    accessToken: {
      token: accessToken,
      expiresIn: expiresIn,
    },
    refreshToken: {
      token: refreshToken,
      expiresIn: expiresIn * 24 * 30,
    },
  };

  await saveToken('app', tokenData, user.id);

  return tokenData;
};

@Service()
export class AuthService {
  public async authenticate(
    userData: User,
    provider?: string,
    req?: Request,
    isSignUp = false,
  ): Promise<{ tokenData: TokenResponsePayload; user: User }> {
    let findUser: User;
    if (isSignUp) {
      // Check if user already exists with username or email
      findUser = await UserModel.findOne({
        $or: [{ username: userData.username }, { email: userData.email }],
      });
      if (findUser) throw new HttpException(409, `This email ${userData.email} already exists`);
    } else {
      findUser = await UserModel.findOne({
        $or: [{ username: userData.username }, { email: userData.email }],
      });
      if (!findUser) throw new HttpException(409, `This email ${userData.email} was not found`);
    }

    if (isSignUp) {
      const hashedPassword = provider ? undefined : await hash(userData.password, 10);
      userData.id = uuidv4();
      userData.has_password = provider ? false : true;
      findUser = await UserModel.create({ ...userData, password: hashedPassword });
    } else if (!provider) {
      const isPasswordMatching: boolean = await compare(userData.password, findUser.password);
      if (!isPasswordMatching) throw new HttpException(409, 'Password is not matching');
    }

    const tokenData = await createToken(findUser);

    const login = await createLogin(findUser.id, req);
    console.log(login);

    return {
      tokenData,
      user: _.pick(findUser, ['_id', 'name', 'username', 'email']),
    };
  }

  public async signup(userData: User, provider?: string, req?: Request): Promise<{ tokenData: TokenResponsePayload; user: User }> {
    return await this.authenticate(userData, provider, req, true);
  }

  public async login(userData: User, provider?: string, req?: Request): Promise<{ tokenData: TokenResponsePayload; user: User }> {
    return await this.authenticate(userData, provider, req, false);
  }

  public async oauthLogin(
    provider: string,
    accessToken: string,
    refreshToken: string,
    profile: any,
  ): Promise<{ tokenData: TokenResponsePayload; user: User }> {
    const email = profile.emails[0].value;
    const name = profile.displayName || `${profile.name.givenName} ${profile.name.familyName}`;
    const id = profile.id;

    const user = await passport.upsertUser(provider, id, email, name, accessToken, refreshToken);
    const tokenData = await createToken(user);

    return {
      tokenData,
      user: _.pick(user, ['_id', 'name', 'email']),
    };
  }

  public async logout(userId: string): Promise<void> {
    await deleteToken(null, 'app', userId);
  }

  public async logoutAll(userId: string): Promise<void> {
    await UserModel.updateOne({ _id: userId }, { $set: { token: null } });
  }

  public async updateAccount(userId: string, updates: Partial<User>): Promise<User> {
    const updatedUser = await UserModel.findOneAndUpdate({ _id: userId }, updates, { new: true });
    if (!updatedUser) throw new HttpException(404, 'User not found');

    return _.pick(updatedUser, ['_id', 'name', 'email']);
  }

  public async authorizeUser(token: string, type: TokenType): Promise<Partial<User>> {
    console.log(token);
    // decode token
    const decodedToken = await verify(token, SECRET_KEY);

    const exist = await verifyToken('app', decodedToken['id'] as string);
    if (!exist) throw new HttpException(401, 'Invalid token');

    const userdata = await UserModel.findOne({ id: decodedToken['id'] as string });
    if (!userdata) throw new HttpException(401, 'Invalid token');

    // if (type === 'email') {
    //   if (user.email_verified) throw new HttpException(401, 'Email already verified');
    //   await this.verifyEmail(user._id, user.email);
    // } else if (type === 'phone') {
    //   if (user.phone_verified) throw new HttpException(401, 'Phone already verified');
    //   await this.verifyPhone(user._id, user.phone_number);
    // }

    return _.pick(userdata, ['id', 'name', 'email', 'username', 'email_verified', 'phone_verified']);
  }

  public async resetPassword(userId: string, newPassword: string): Promise<void> {
    const hashedPassword = await hash(newPassword, 10);
    const updatedUser = await UserModel.findOneAndUpdate({ _id: userId }, { password: hashedPassword });

    if (!updatedUser) throw new HttpException(404, 'User not found');

    // Send password reset confirmation email
    sendPasswordReset(updatedUser.email, updatedUser.name);
  }

  public async verifyEmail(userId: string, email: string): Promise<void> {
    const user = await UserModel.findOne({ _id: userId });

    if (!user) throw new HttpException(404, 'User not found');
    if (user.email !== email) throw new HttpException(409, 'Email does not match user record');

    await UserModel.findOneAndUpdate({ _id: userId }, { email_verified: true });
    sendEmailVerification(email);
  }

  public async verifyPhone(userId: string, phoneNumber: string): Promise<void> {
    const user = await UserModel.findOne({ _id: userId });

    if (!user) throw new HttpException(404, 'User not found');

    // Use Twilio to verify phone number
    verifyPhoneNumber(phoneNumber);

    await UserModel.findOneAndUpdate({ _id: userId }, { phone_verified: true });
  }

  public async sendWelcome(userId: string): Promise<void> {
    const user = await UserModel.findOne({ _id: userId });

    if (!user) throw new HttpException(404, 'User not found');

    sendWelcomeEmail(user.email);
  }
}
