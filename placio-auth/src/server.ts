import { App } from '@/app';
import { AuthRoute } from '@routes/auth.route';
import { UserRoute } from '@routes/users.route';
import { ValidateEnv } from '@utils/validateEnv';
import { AccountRoute } from './routes/account.route';

ValidateEnv();

const app = new App([new UserRoute(), new AuthRoute(), new AccountRoute()]);

app.listen();
