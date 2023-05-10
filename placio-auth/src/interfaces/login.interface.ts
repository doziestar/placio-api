import { Document } from 'mongoose';

interface ILogin extends Document {
  id: string;
  user_id: string;
  ip: string;
  time: Date;
  browser: string;
  device: string;
}

export { ILogin };
