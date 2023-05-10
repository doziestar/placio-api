import { v4 as uuidv4 } from 'uuid';
import mongoose, { Schema } from 'mongoose';
import { ILogin } from '@interfaces/login.interface';

const LoginSchema = new Schema<ILogin>({
  id: { type: String, required: true, unique: true },
  user_id: { type: String, required: true },
  ip: { type: String, required: true },
  time: { type: Date, required: true },
  browser: { type: String },
  device: { type: String },
});

const Login = mongoose.model<ILogin>('Login', LoginSchema, 'login');
export const schema = Login;

export async function create(user: string, req: any): Promise<ILogin> {
  const ip =
    (req.headers['x-forwarded-for'] || '').split(',').pop().trim() ||
    req.connection?.remoteAddress ||
    req.socket?.remoteAddress ||
    req.connection?.socket?.remoteAddress;

  const ua = req.headers['user-agent'];
  const device = ua.substring(ua.indexOf('(') + 1, ua.indexOf(')')).replace(/_/g, '.');
  const uarr = ua.split(' ');
  const browser = uarr[uarr.length - 1];

  const newLogin = new Login({
    id: uuidv4(),
    user_id: user,
    ip: ip,
    time: new Date(),
    browser: browser,
    device: device,
  });

  return await newLogin.save();
}

export async function verify(user: string, current: ILogin) {
  let riskLevel = 0;

  const flag = {
    ip: current.ip,
    device: current.device,
    browser: current.browser.split('/')[0],
  };

  const history = await Login.find({ user_id: user, id: { $ne: current.id } }).select({});

  if (history.length) {
    if (history.findIndex(x => x.ip === current.ip) < 0) riskLevel++;
    if (history.findIndex(x => x.browser === current.browser) < 0) riskLevel++;
    const devices = history.filter(x => x.device !== current.device)?.length;
    if (devices > 1) riskLevel++;
  }

  let time: string[] = new Date(current.time).toISOString().split('T');
  time = [`${time[0]} ${time[1].split('.')[0]}`];

  return {
    flag: flag,
    level: riskLevel,
    time: time[0],
  };
}
