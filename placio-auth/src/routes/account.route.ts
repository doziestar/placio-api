import { Router } from 'express';
import { AuthController } from '@controllers/auth.controller';
import { CreateUserDto, LoginUserDto } from '@dtos/users.dto';
import { Routes } from '@interfaces/routes.interface';
import { AuthMiddleware } from '@middlewares/auth.middleware';
import { ValidationMiddleware } from '@middlewares/validation.middleware';
import accountController from '@/controllers/account.controller';

class AccountRoute implements Routes {
  public path = '/account';
  public router = Router();
  public accountController = accountController;

  constructor() {
    this.initializeRoutes();
  }

  private initializeRoutes() {
    this.router.post(`${this.path}`, ValidationMiddleware(CreateUserDto, 'body' as any, true), this.accountController.create);
  }
}

export { AccountRoute };
