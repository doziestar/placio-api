import { verify } from 'jsonwebtoken';
import { NextFunction, Request, Response } from 'express';
import { AuthService } from '@services/auth.service';
import { User } from '@interfaces/users.interface';
import { RequestWithUser } from '@interfaces/auth.interface';
import Container from 'typedi';
import { TokenType } from '@/interfaces/token.interface';
import { getUser } from '@/models/users.model';
import { assert, validateNativeURL } from '@/models/utility.model';
import { createLogin, verifyLogin } from '@/models/login.model';
import passport from 'passport';
import { sendMail } from '@/models/mail.model';

class AuthController {
  public authService = Container.get(AuthService);

  public signUp = async (req: Request, res: Response, next: NextFunction) => {
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
      next(error);
    }
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
        const jwt = auth.token({ provider: provider, provider_id: profile.id, email: email }, null, 300);
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
      const token = await auth.token({ id: userData.id }, null, 300);

      // send welcome email
      await sendMail({
        to: userData.email,
        template: `magic_signin`,
        content: {
          token: token,
          domain: validateNativeURL(req.body.magic_view_url) || `${domain}/magic`,
        },
      });
    }

    // always return a positive response to avoid hinting if user exists
    return res.status(200).send();
  }

  async verifyMagicLink(req, res) {
    const data = req.body;
    validate(data, ['token']);
    const magicToken = auth.token.verify(data.token);

    // check user exists
    const userData = await getUser(magicToken.id);

    // authenticated
    if (userData) {
      // log the sign in and check if it's suspicious
      const log = await createLogin(userData.id, req);
      const risk = await verifyLogin(userData.id, log);

      // notify the user of suspicious logins
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

      // 2fa is required
      if (userData['2fa_enabled']) {
        // notify the client and use email to identify the user when sending otp
        // send a token so the otp password screen can't be accessed directly without a password
        const jwt = auth.token({ email: userData.email, provider: 'app' }, null, 300);
        return res.status(200).send({ '2fa_required': true, token: jwt });
      }

      return this.authenticate(req, res, userData, data);
    }

    // error
    return res.status(401).send();
  }

  async getAuthStatus(req, res) {
    // is there an account/user?
    let hasJWT = false,
      hasSocialToken = false,
      usingSocialSignin;

    // does the user have an active jwt?
    if (req.provider === 'app') {
      usingSocialSignin = false;
      hasJWT = await token.verify(req.provider, req.user);
    }

    // is there an active access_token if the user is
    // signed in via social network or was their account de-authed
    if (req.provider !== 'app') {
      usingSocialSignin = true;
      hasSocialToken = await token.verify(req.provider, req.user);
    }

    // does this user have an active subscription?
    const subscription = await account.subscription(req.account);
    const userAccounts = await user.account(req.user);
    user.update(req.user, req.account, { last_active: new Date() });

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

  async impersonate(req, res, next) {
    assert(req.body.token, 'Authorization token required');
    const data = auth.token.verify(req.body.token);

    assert(data.user_id && data.permission === 'master', 'Invalid token');

    // check user exists
    const userData = await getUser(data.user_id);
    assert(userData, 'User does not exist');

    return authenticate(req, res, userData);
  }

  /*
   * auth.switch()
   * let the user switch account
   */

  async switch(req, res) {
    const data = req.body;
    validate(data, ['account']);

    // check user belongs to this account
    const userData = await getUser(req.user, null, data.account);
    assert(userData, `You don't belong to this account.`);

    return authenticate(req, res, userData, data);
  }

  /*
   * auth.signout()
   * sign the user out
   * destroy any tokens
   */
  async signout(req, res) {
    // destroy social tokens
    await token.delete(null, req.provider, req.user);
    return res.status(200).send();
  }

  /*
   * authenticate()
   * call this function to finalise the auth process
   */

  async authenticate(req, res, userData, data) {
    const accountData = await account.get(userData.account_id);
    const subscription = await account.subscription(userData.account_id);
    const userAccounts = await user.account(userData.id);

    // create & store the token
    const jwt = auth.token({
      accountId: userData.account_id,
      userId: userData.id,
      permission: userData.permission,
      provider: data?.provider || 'app',
    });

    await token.save(data?.provider || 'app', { access: jwt }, userData.id);
    user.update(userData.id, userData.account_id, {
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
