import { Request, Response, NextFunction } from 'express';

const authService = new AuthService();

export class AuthController {
  async signin(req: Request, res: Response, next: NextFunction) {
    try {
      const result = await authService.signin(req, res);
      res.status(result.status).send(result.data);
    } catch (error) {
      next(error);
    }
  }

  async signinOtp(req: Request, res: Response, next: NextFunction) {
    try {
      const result = await authService.signinOtp(req);
      res.status(result.status).send(result.data);
    } catch (error) {
      next(error);
    }
  }

  async signup(req: Request, res: Response, next: NextFunction) {
    try {
      const result = await authService.signup(req);
      res.status(result.status).send(result.data);
    } catch (error) {
      next(error);
    }
  }

  async social(req: Request, res: Response, next: NextFunction) {
    try {
      const result = await authService.social(req, res);
      res.status(result.status).send(result.data);
    } catch (error) {
      next(error);
    }
  }

  async magic(req: Request, res: Response, next: NextFunction) {
    try {
      const result = await authService.magic(req);
      res.status(result.status).send(result.data);
    } catch (error) {
      next(error);
    }
  }

  async magicVerify(req: Request, res: Response, next: NextFunction) {
    try {
      const result = await authService.magicVerify(req);
      res.status(result.status).send(result.data);
    } catch (error) {
      next(error);
    }
  }

  async get(req: Request, res: Response, next: NextFunction) {
    try {
      const result = await authService.get(req);
      res.status(result.status).send(result.data);
    } catch (error) {
      next(error);
    }
  }

  async impersonate(req: Request, res: Response, next: NextFunction) {
    try {
      const result = await authService.impersonate(req);
      res.status(result.status).send(result.data);
    } catch (error) {
      next(error);
    }
  }

  async switch(req: Request, res: Response, next: NextFunction) {
    try {
      const result = await authService.switch(req);
      res.status(result.status).send(result.data);
    } catch (error) {
      next(error);
    }
  }

  async signout(req: Request, res: Response, next: NextFunction) {
    try {
      const result = await authService.signout(req);
      res.status(result.status).send(result.data);
    } catch (error) {
      next(error);
    }
  }
}



export async function signin(req: Request, res: Response, next: NextFunction) {
  const data = req.body;
  let userData: User;
  let useEmail = false; // determine if flow is email or social
  console.log('data', data);
  if (data.email) {
    useEmail = true;
    data.provider = 'app';
    utility.validate(data, ['email', 'password']);
  } else {
    // using social, extra fields from jwt
    utility.validate(data, ['token']);
    const decode = auth.token.verify(data.token);
    data.provider = decode.provider;
    data.provider_id = decode.provider_id;
    data.email = decode.email;
  }

  // check user exists
  userData = useEmail
    ? await user.get(null, data.email)
    : await user.get(null, null, null, {
        provider: data.provider,
        id: data.provider_id,
      });

  utility.assert(userData, 'Please enter the correct login details', 'email');

  // verify password
  if (useEmail) {
    const verified = await user.password.verify(
      userData.id,
      userData.account_id,
      data.password
    );
    utility.assert(
      verified,
      'Please enter the correct login details',
      'password'
    );
  }

  // get the account
  const accountData = await account.get(userData.account_id);
  utility.assert(
    accountData?.active,
    'Your account has been deactivated. Please contact support.'
  );

  // log the sign in and check if it's suspicious
  const log = await login.create(userData.id, req);
  const risk: Risk = await login.verify(userData.id, log);

  // block the signin & send a magic link if risk level is 3 or user account is disabled
  if (useEmail) {
    if (risk.level === 3 || userData.disabled) {
      await user.update(userData.id, userData.account_id, { disabled: true });
      const token = auth.token({ id: userData.id }, null, 300);

      await mail.send({
        to: userData.email,
        template: 'blocked_signin',
        content: {
          token: token,
          domain:
            utility.validateNativeURL(data.magic_view_url) || `${domain}/magic`,
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

    // notify the user of suspicious log
