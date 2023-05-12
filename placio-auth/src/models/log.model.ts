import { v4 as uuidv4 } from 'uuid';
import mongoose, { Document, Schema } from 'mongoose';

interface ILog extends Document {
  id: string;
  time: Date;
  message?: string;
  body?: string;
  method?: string;
  endpoint?: string;
  account_id?: string;
  user_id?: string;
}

const LogSchema = new Schema<ILog>({
  id: { type: String, required: true, unique: true },
  time: { type: Date, required: true },
  message: { type: String },
  body: { type: String },
  method: { type: String },
  endpoint: { type: String },
  account_id: { type: String },
  user_id: { type: String },
});

const Log = mongoose.model<ILog>('Log', LogSchema, 'log');

interface IRequest {
  user?: string;
  account?: string;
  route?: {
    path?: string;
    methods?: Record<string, boolean>;
  };
}

class LogController {
  async create(message: string, body: string, req?: IRequest, user?: string, account?: string) {
    const newLog = new Log({
      id: uuidv4(),
      message: message,
      time: new Date(),
      user_id: req?.user || user,
      account_id: req?.account || account,
      endpoint: req?.route?.path,
      body: body && (typeof body === 'object' ? JSON.stringify(body, Object.getOwnPropertyNames(body)) : body),
      //   method: req?.route ? Object.keys(req.route.methods).reduce(key => req.route.methods[key] as string) : null,
    });

    return await newLog.save();
  }
}

export { Log, LogController };
