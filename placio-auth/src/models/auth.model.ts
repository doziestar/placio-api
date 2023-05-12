import jwt from 'jsonwebtoken';
import httpStatus from 'http-status';
import { assert, decode } from './utility.model';
import { LogController } from './log.model';
import { IKey, KeyController } from './key.model';
import config from './config.model';
import { TokenData, TokenResponsePayload } from '@/interfaces/auth.interface';
import { CRYPTO_SECRET, TOKEN_SECRET } from '@config';

const log = new LogController();
const key = new KeyController();

const permissions = config.permissions;

const settings = config.token;

interface VerifyPayload {
  accountId?: string;
  userId?: string;
  permission?: string;
  provider?: string;
}

class Auth {
  token(data: object, secret: string = CRYPTO_SECRET || 'iamasecrte', duration = settings.duration): TokenResponsePayload {
    const accessToken = jwt.sign(data, secret, { expiresIn: duration });
    const refreshToken = jwt.sign(data, secret, { expiresIn: duration * 2 });

    return {
      accessToken: {
        token: accessToken,
        expiresIn: duration,
      },
      refreshToken: {
        token: refreshToken,
        expiresIn: duration * 2,
      },
    };
  }

  verifyToken(token: string, secret: string = process.env.TOKEN_SECRET || ''): TokenData {
    const data = jwt.verify(token, secret) as TokenData;
    return data;
  }

  async verifyMiddleware(permission: string, scope: string) {
    return async function (req: any, res: any, next: Function) {
      try {
        const header = req.headers['authorization'];

        if (!header) {
          if (permission === 'public') {
            return next();
          } else {
            throw { message: 'No authorization header provided' };
          }
        }

        const type = header.split(' ')[0];
        const token = header.split(' ')[1];

        if (type === 'Basic') {
          const apikey = token.includes('key-') ? token : decode(token).replace(':', '');

          const verified: IKey = await key.verify(apikey);
          assert(verified, 'Invalid API key', 401);
          assert(verified.scope.includes(scope), `You don't have permission to use this scope`, 401);

          if (process.env.ENABLE_API_LOGS === 'true') log.create(null, req.body, req, req.user, req.user);

          req.account = verified.account_id;
          next();
        } else if (type === 'Bearer') {
          const decode = jwt.verify(token, process.env.TOKEN_SECRET || '') as VerifyPayload;

          if (decode.accountId && decode.userId && decode.permission && decode.provider) {
            if (permission === 'public' || permissions[decode.permission][permission]) {
              req.account = decode.accountId;
              req.user = decode.userId;
              req.permission = decode.permission;
              req.provider = decode.provider;
              next();
            } else throw new Error();
          } else throw { message: 'Invalid token' };
        } else throw { message: 'Unrecognised header type' };
      } catch (err) {
        res.status(httpStatus.UNAUTHORIZED).send({
          message: err.message || 'You do not have permission to perform this action.',
        });
      }
    };
  }
}

export default new Auth();
