import { Router } from 'express';
import { AuthController } from '@controllers/auth.controller';
import { CreateUserDto } from '@dtos/users.dto';
import { Routes } from '@interfaces/routes.interface';
import { AuthMiddleware } from '@middlewares/auth.middleware';
import { ValidationMiddleware } from '@middlewares/validation.middleware';

export class AuthRoute implements Routes {
  public path = '/auth/';
  public router = Router();
  public auth = new AuthController();

  constructor() {
    this.initializeRoutes();
  }

  private initializeRoutes() {
    this.router.post(`${this.path}signup`, ValidationMiddleware(CreateUserDto, 'body' as any), this.auth.signUp);
    this.router.post(`${this.path}login`, ValidationMiddleware(CreateUserDto, 'body' as any), this.auth.logIn);
    this.router.post(`${this.path}logout`, AuthMiddleware, this.auth.logOut);
    this.router.post(`${this.path}`, limiter(throttle.signin), use(authController.signin));

    this.router.post(`${this.path}otp`, limiter(throttle.signin), use(authController.signin.otp));

    this.router.get(`${this.path}`, AuthMiddleware, use(authController.get));

    this.router.post(`${this.path}magic`, limiter(throttle.signin), use(authController.magic));

    this.router.post(`${this.path}magic/verify`, limiter(throttle.signin), use(authController.magic.verify));

    this.router.post(`${this.path}password/reset/request`, limiter(throttle.password_reset), use(userController.password.reset.request));

    this.router.post(`${this.path}password/reset`, limiter(throttle.password_reset), use(userController.password.reset));

    this.router.post(`${this.path}switch`, AuthMiddleware, use(authController.switch));

    this.router.post(`${this.path}impersonate`, use(authController.impersonate));

    this.router.delete(`${this.path}`, AuthMiddleware, use(authController.signout));
  }
}
