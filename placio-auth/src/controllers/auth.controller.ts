import { NextFunction, Request, Response } from 'express';
import { AuthService } from '@services/auth.service';
import { User } from '@interfaces/users.interface';
import { RequestWithUser } from '@interfaces/auth.interface';
import Container from 'typedi';

class AuthController {
  public authService = Container.get(AuthService);

  public signUp = async (req: Request, res: Response, next: NextFunction) => {
    try {
      // const userData: User = plainToClass(CreateUserDto, req.body);
      const { tokenData, user } = await this.authService.signup(req.body);

      res.status(201).json({ tokenData, user });
    } catch (error) {
      next(error);
    }
  };

  public logIn = async (req: Request, res: Response, next: NextFunction) => {
    try {
      // const userData: User = plainToClass(CreateUserDto, req.body);
      const { tokenData, user } = await this.authService.login(req.body);

      res.status(200).json({ tokenData, user });
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
