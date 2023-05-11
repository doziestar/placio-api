import { hash, compare } from 'bcrypt';
import { sign } from 'jsonwebtoken';
import { Service } from 'typedi';
import { SECRET_KEY } from '@config';
import { HttpException } from '@exceptions/httpException';
import { DataStoredInToken, TokenData } from '@interfaces/auth.interface';
import { User } from '@interfaces/users.interface';
import { UserModel } from '@models/users.model';
import { v4 as uuidv4 } from 'uuid';
import _ from 'lodash';
import passport from '@services/social.service';
import { sendEmailVerification, sendPasswordReset, sendWelcomeEmail } from '@services/email.service';
import { verifyPhoneNumber } from './sms.service';

const createToken = (user: User): TokenData => {
  const dataStoredInToken: DataStoredInToken = { _id: user._id };
  const expiresIn: number = 60 * 60;

  return { expiresIn, token: sign(dataStoredInToken, SECRET_KEY, { expiresIn }) };
};

@Service()
export class AuthService {
  public async signup(userData: User, provider?: string): Promise<{ tokenData: TokenData; user: User }> {
    const findUser: User = await UserModel.findOne({ email: userData.email });
    if (findUser) throw new HttpException(409, `This email ${userData.email} already exists`);

    const hashedPassword = provider ? undefined : await hash(userData.password, 10);
    userData.id = uuidv4();
    userData.has_password = provider ? false : true;
    const createUserData: User = await UserModel.create({ ...userData, password: hashedPassword });

    const tokenData = createToken(createUserData);

    return {
      tokenData,
      user: _.pick(createUserData, ['_id', 'name', 'email']),
    };
  }

  public async login(userData: User, provider?: string): Promise<{ tokenData: TokenData; user: User }> {
    const findUser: User = await UserModel.findOne({ email: userData.email });
    if (!findUser) throw new HttpException(409, `This email ${userData.email} was not found`);

    if (!provider) {
      const isPasswordMatching: boolean = await compare(userData.password, findUser.password);
      if (!isPasswordMatching) throw new HttpException(409, 'Password is not matching');
    }

    const tokenData = createToken(findUser);

    return {
      tokenData,
      user: _.pick(findUser, ['_id', 'name', 'email']),
    };
  }

  public async oauthLogin(provider: string, accessToken: string, refreshToken: string, profile: any): Promise<{ tokenData: TokenData; user: User }> {
    const email = profile.emails[0].value;
    const name = profile.displayName || `${profile.name.givenName} ${profile.name.familyName}`;
    const id = profile.id;

    const user = await passport.upsertUser(provider, id, email, name, accessToken, refreshToken);
    const tokenData = createToken(user);

    return {
      tokenData,
      user: _.pick(user, ['_id', 'name', 'email']),
    };
  }

  public async updateAccount(userId: string, updates: Partial<User>): Promise<User> {
    const updatedUser = await UserModel.findOneAndUpdate({ _id: userId }, updates, { new: true });
    if (!updatedUser) throw new HttpException(404, 'User not found');

    return _.pick(updatedUser, ['_id', 'name', 'email']);
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
