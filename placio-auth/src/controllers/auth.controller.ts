import { verify } from 'jsonwebtoken';
import { NextFunction, Request, Response } from 'express';
import { AuthService } from '@services/auth.service';
import { User } from '@interfaces/users.interface';
import { RequestWithUser } from '@interfaces/auth.interface';
import Container from 'typedi';
import { TokenType } from '@/interfaces/token.interface';
import { getUser, getUserAccount, updateUser, verifyPassword } from '@/models/users.model';
import { assert, validate, validateNativeURL } from '@/models/utility.model';
import { createLogin, verifyLogin } from '@/models/login.model';
import passport from 'passport';
import { sendMail } from '@/models/mail.model';
import { HttpException } from '@/exceptions/httpException';
import { saveToken, verifyToken } from '@/models/token.model';
import Auth from '@/models/auth.model';
import { getAccount, getSubscription } from '@/models/account.model';

interface Data {
  email?: string;
  password?: string;
  token?: string;
  provider?: string;
  provider_id?: string;
  magic_view_url?: string;
}

class AuthController {
  public authService = Container.get(AuthService);

  public signUp = async (req: Request, res: Response) => {
    try {
      // const userData: User = plainToClass(CreateUserDto, req.body);
      // const { tokenData, user } = await this.authService.signup(req.body, null, req);

      // res.status(201).json({ tokenData, user });

      const data = req.body;
      // utility.validate(data, ['email']);

      // check user exists
      const userData = await getUser(null, data.email, data.account_id);
      assert(userData, `You're not registered`, 401);

      // log the sign in and check if it's suspicious
      await createLogin(userData.id, req);
      return this.authenticate(req, res, userData, data);
    } catch (error) {
      throw new HttpException(error.statusCode || 500, error.message);
    }
  };

  public signin = async (req: Request, res: Response, next: NextFunction) => {
    const data: Data = req.body;
    let useEmail = false;

    if (data.email) {
      useEmail = true;
      data.provider = 'app';
      validate(data, ['email', 'password']);
    } else {
      // validate(data, ['token']);
      // const decode = Auth.verifyToken(data.token);
      // data.provider = decode.provider;
      // data.provider_id = decode.provider_id;
      // data.email = decode.email;
    }

    const userData = useEmail ? await getUser(null, data.email) : await getUser(null, null, null, { provider: data.provider, id: data.provider_id });
    console.log('userData', userData);

    assert(userData, 'Please enter the correct login details', 'email');

    if (useEmail) {
      const verified = await verifyPassword(userData.id, userData.default_account, data.password);
      assert(verified, 'Please enter the correct login details', 'password');
    }

    const accountData = await getAccount(userData.default_account);
    assert(accountData?.active, 'Your account has been deactivated. Please contact support.', 403);

    const log = await createLogin(userData.id, req);
    const risk = await verifyLogin(userData.id, log);

    if (useEmail) {
      if (risk.level === 3 || userData.disabled) {
        await updateUser(userData.id, userData.default_account, { disabled: true });
        const token = Auth.token({ id: userData.id }, null, 300);

        await sendMail({
          to: userData.email,
          template: 'blocked_signin',
          content: {
            token: token.accessToken.token,
            // domain: validateNativeURL(data.magic_view_url) || `${domain}/magic`,
            domain: 'http://localhost:3000/magic',
          },
        });

        const msg =
          risk.level === 3
            ? 'Your sign in attempt has been blocked due to suspicious activity. '
            : 'Your account has been disabled due to suspicious activity. ';

        return res.status(403).send({
          message: msg + 'Please check your email for further instructions.',
        });
      }

      if (risk.level > 0) {
        await sendMail({
          to: userData.email,
          template: 'new_signin',
          content: {
            ip: risk.flag.ip,
            time: risk.time,
            device: risk.flag.device,
            browser: risk.flag.browser,
          },
        });
      }
    }

    if (userData['2fa_enabled']) {
      const jwt = Auth.token({ email: userData.email, provider: data.provider }, null, 300);
      return res.status(200).send({ '2fa_required': true, token: jwt });
    }

    // done
    return this.authenticate(req, res, userData, data);
  };

  async social(req, res, next) {
    const signinURL = req.session.deep_signin_url || `${process.env.CLIENT_URL}/signin`;
    const socialURL = req.session.deep_social_url || `${process.env.CLIENT_URL}/signin/social`;

    passport.authenticate(req.params.provider, { failureRedirect: signinURL }, async (err, profile) => {
      if (err || !profile.id) {
        console.log(err);
        return res.redirect(`${signinURL}?error=${encodeURIComponent(err?.message || 'Unauthorized')}`);
      }

      // authenticate the user
      const provider = req.params.provider;
      const email = profile.emails[0]?.value;
      const userData = await getUser(null, email, null, {
        provider: provider,
        id: profile.id,
      });

      if (userData) {
        const jwt = Auth.token({ provider: provider, provider_id: profile.id, email: email }, null, 300);
        res.redirect(`${socialURL}?provider=${provider}&token=${jwt}`);
      } else {
        res.redirect(`${signinURL}?error=${encodeURIComponent(`You're not registered`)}`);
      }
    })(req, res, next);
  }

  async magic(req, res) {
    const userData = await getUser(null, req.body.email);

    if (userData) {
      // generate a token that expires in 5 mins
      const token = await Auth.token({ id: userData.id }, null, 300);

      // send welcome email
      await sendMail({
        to: userData.email,
        template: `magic_signin`,
        content: {
          token: token.accessToken.token,
          // domain: validateNativeURL(req.body.magic_view_url) || `${d}/magic`,
          domain: `${process.env.CLIENT_URL}/magic`,
        },
      });
    }

    // always return a positive response to avoid hinting if user exists
    return res.status(200).send();
  }

  // async verifyMagicLink(req, res) {
  //   const data = req.body;
  //   validate(data, ['token']);
  //   const magicToken = Auth.verifyToken(data.token);

  //   // check user exists
  //   const userData = await getUser(magicToken.id);

  //   // authenticated
  //   if (userData) {
  //     // log the sign in and check if it's suspicious
  //     const log = await createLogin(userData.id, req);
  //     const risk = await verifyLogin(userData.id, log);

  //     // notify the user of suspicious logins
  //     if (risk.level > 0) {
  //       await sendMail({
  //         to: userData.email,
  //         template: 'new_signin',
  //         content: {
  //           ip: risk.flag.ip,
  //           time: risk.time,
  //           device: risk.flag.device,
  //           browser: risk.flag.browser,
  //         },
  //       });
  //     }

  //     // 2fa is required
  //     if (userData['2fa_enabled']) {
  //       // notify the client and use email to identify the user when sending otp
  //       // send a token so the otp password screen can't be accessed directly without a password
  //       const jwt = Auth.token({ email: userData.email, provider: 'app' }, null, 300);
  //       return res.status(200).send({ '2fa_required': true, token: jwt });
  //     }

  //     return this.authenticate(req, res, userData, data);
  //   }

  //   // error
  //   return res.status(401).send();
  // }

  async getAuthStatus(req, res) {
    // is there an account/user?
    let hasJWT = false,
      hasSocialToken = false,
      usingSocialSignin;

    // does the user have an active jwt?
    if (req.provider === 'app') {
      usingSocialSignin = false;
      hasJWT = await verifyToken(req.provider, req.user);
    }

    // is there an active access_token if the user is
    // signed in via social network or was their account de-authed
    if (req.provider !== 'app') {
      usingSocialSignin = true;
      hasSocialToken = await verifyToken(req.provider, req.user);
    }

    // does this user have an active subscription?
    const subscription = await getSubscription(req.account);
    const userAccounts = await getUserAccount(req.user);
    updateUser(req.user, req.account, { last_active: new Date() });

    return res.status(200).send({
      data: {
        jwt_token: hasJWT,
        social_token: hasSocialToken,
        subscription: subscription.status,
        accounts: userAccounts,
        account_id: req.account,
        authenticated: usingSocialSignin ? hasSocialToken : hasJWT,
      },
    });
  }

  // async impersonate(req, res, next) {
  //   assert(req.body.token, 'Authorization token required');
  //   const data = auth.token.verify(req.body.token);

  //   assert(data.user_id && data.permission === 'master', 'Invalid token');

  //   // check user exists
  //   const userData = await getUser(data.user_id);
  //   assert(userData, 'User does not exist');

  //   return this.authenticate(req, res, userData);
  // }

  /*
   * auth.switch()
   * let the user switch account
   */

  // async switch(req, res) {
  //   const data = req.body;
  //   validate(data, ['account']);

  //   // check user belongs to this account
  //   const userData = await getUser(req.user, null, data.account);
  //   assert(userData, `You don't belong to this account.`);

  //   return authenticate(req, res, userData, data);
  // }

  /*
   * auth.signout()
   * sign the user out
   * destroy any tokens
   */
  // async signout(req, res) {
  //   // destroy social tokens
  //   await token.delete(null, req.provider, req.user);
  //   return res.status(200).send();
  // }

  /*
   * authenticate()
   * call this function to finalise the auth process
   */

  async authenticate(req, res, userData, data) {
    console.log('authenticate', userData);
    const accountData = await getAccount(userData.account_id);
    console.log('accountData', accountData);
    const subscription = await getSubscription(userData.account_id);
    console.log('subscription', subscription);
    const userAccounts = await getUserAccount(userData.id);
    console.log('userAccounts', userAccounts);

    // create & store the token
    const jwt = Auth.token({
      accountId: userData.account_id,
      userId: userData.id,
      permission: userData.permission,
      provider: data?.provider || 'app',
    });

    await saveToken(data?.provider || 'app', jwt, userData.id);
    await updateUser(userData.id, userData.account_id, {
      last_active: new Date(),
      disabled: false,
    });

    // return user to server
    return res.status(200).send({
      token: jwt,
      subscription: subscription.status,
      plan: accountData.plan,
      permission: userData.permission,
      name: userData.name,
      accounts: userAccounts,
      account_id: userData.account_id,
      has_password: userData.has_password,
      onboarded: userData.onboarded,
    });
  }

  public logIn = async (req: Request, res: Response, next: NextFunction) => {
    try {
      const { tokenData, user } = await this.authService.login(req.body, null, req);

      res.status(200).json({ tokenData, user });
    } catch (error) {
      next(error);
    }
  };

  public logOut = async (req: RequestWithUser, res: Response, next: NextFunction) => {
    try {
      const userId = req.user.id;
      await this.authService.logout(userId);

      res.status(200).json({ message: 'logout' });
    } catch (error) {
      next(error);
    }
  };

  // public refreshTokens = async (req: Request, res: Response, next: NextFunction) => {
  //   try {
  //     const { tokenData, user } = await this.authService.refreshTokens(req.body.refreshToken, req);

  //     res.status(200).json({ tokenData, user });
  //   } catch (error) {
  //     next(error);
  //   }
  // };

  // public forgotPassword = async (req: Request, res: Response, next: NextFunction) => {
  //   try {
  //     const email = req.body.email;
  //     await this.authService.forgotPassword(email);
  //     res.status(204).send();
  //   } catch (error) {
  //     next(error);
  //   }
  // };

  // public verifyResetPassword = async (req: Request, res: Response, next: NextFunction) => {
  //   try {
  //     const token = req.body.token;
  //     const { tokenData, user } = await this.authService.verifyResetPassword(token);
  //     res.status(200).json({ tokenData, user });
  //   } catch (error) {
  //     next(error);
  //   }
  // };

  // public sendOtp = async (req: Request, res: Response, next: NextFunction) => {
  //   try {
  //     const phone = req.body.phone;
  //     await this.authService.sendOtp(phone);
  //     res.status(204).send();
  //   } catch (error) {
  //     next(error);
  //   }
  // };

  public authorizeUser = async (req: Request, res: Response, next: NextFunction) => {
    try {
      const token = req.query.token;
      const type = req.query.type;
      const user = await this.authService.authorizeUser(token as string, type as TokenType);
      res.status(200).json({ user });
    } catch (error) {
      next(error);
    }
  };

  public oauthLogin = async (provider: string) => {
    return async (req: Request, res: Response, next: NextFunction) => {
      try {
        const { accessToken, refreshToken, profile } = req.body;
        const { tokenData, user } = await this.authService.oauthLogin(provider, accessToken, refreshToken, profile);

        res.status(200).json({ tokenData, user });
      } catch (error) {
        next(error);
      }
    };
  };

  public updateAccount = async (req: RequestWithUser, res: Response, next: NextFunction) => {
    try {
      const userId = req.user._id;
      const updates: Partial<User> = req.body;
      const updatedUser = await this.authService.updateAccount(userId, updates);

      res.status(200).json({ user: updatedUser });
    } catch (error) {
      next(error);
    }
  };

  public resetPassword = async (req: RequestWithUser, res: Response, next: NextFunction) => {
    try {
      const userId = req.user._id;
      const newPassword = req.body.newPassword;
      await this.authService.resetPassword(userId, newPassword);

      res.status(204).send();
    } catch (error) {
      next(error);
    }
  };

  public verifyEmail = async (req: RequestWithUser, res: Response, next: NextFunction) => {
    try {
      const userId = req.user._id;
      const email = req.body.email;
      await this.authService.verifyEmail(userId, email);

      res.status(204).send();
    } catch (error) {
      next(error);
    }
  };

  public verifyPhone = async (req: RequestWithUser, res: Response, next: NextFunction) => {
    try {
      const userId = req.user._id;
      const phoneNumber = req.body.phoneNumber;
      await this.authService.verifyPhone(userId, phoneNumber);

      res.status(204).send();
    } catch (error) {
      next(error);
    }
  };

  public sendWelcome = async (req: RequestWithUser, res: Response, next: NextFunction) => {
    try {
      const userId = req.user._id;
      await this.authService.sendWelcome(userId);

      res.status(204).send();
    } catch (error) {
      next(error);
    }
  };
}

export { AuthController };
