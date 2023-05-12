import { Document } from 'mongoose';

interface ILogin extends Document {
  id: string;
  user_id: string;
  ip: string;
  time: Date;
  browser: string;
  device: string;
}

interface Risk {
  level: number;
  flag: {
    ip: string;
    device: string;
    browser: string;
  };
  time: string;
}

export { ILogin, Risk };
