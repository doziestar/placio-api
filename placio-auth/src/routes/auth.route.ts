import { Router } from 'express';
import { AuthController } from '@controllers/auth.controller';
import { CreateUserDto, LoginUserDto } from '@dtos/users.dto';
import { Routes } from '@interfaces/routes.interface';
import { AuthMiddleware } from '@middlewares/auth.middleware';
import { ValidationMiddleware } from '@middlewares/validation.middleware';

class AuthRoute implements Routes {
  public path = '/auth';
  public router = Router();
  public authController = new AuthController();

  constructor() {
    this.initializeRoutes();
  }

  private initializeRoutes() {
    // this.router.post(`${this.path}/signup`, ValidationMiddleware(CreateUserDto, 'body' as any, true), this.authController.signUp);
    this.router.post(`${this.path}`, ValidationMiddleware(LoginUserDto, 'body' as any, true), this.authController.signin);
    this.router.post(`${this.path}/logout`, AuthMiddleware, this.authController.logOut);

    // this.router.post(`${this.path}/refresh`, this.authController.refreshTokens);
    // this.router.post(`${this.path}/forgot-password`, this.authController.forgotPassword);
    this.router.get(`${this.path}/authorize`, this.authController.authorizeUser);

    this.router.put(`${this.path}/update`, AuthMiddleware, this.authController.updateAccount);
    this.router.put(`${this.path}/reset-password`, AuthMiddleware, this.authController.resetPassword);
    this.router.put(`${this.path}/verify-email`, AuthMiddleware, this.authController.verifyEmail);
    this.router.put(`${this.path}/verify-phone`, AuthMiddleware, this.authController.verifyPhone);
    this.router.post(`${this.path}/send-welcome`, AuthMiddleware, this.authController.sendWelcome);
  }
}

export { AuthRoute };
