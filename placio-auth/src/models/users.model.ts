import { model, Schema, Document } from 'mongoose';
import { SocialProvider, UpdateUserData, User } from '@interfaces/users.interface';
import { v4 as uuidv4 } from 'uuid';
import bcrypt from 'bcrypt';
// import Cryptr from 'cryptr';

// const crypto = new Cryptr(process.env.CRYPTO_SECRET);

const UserSchema: Schema = new Schema<User>({
  id: { type: String, required: true, unique: true },
  fingerprint: { type: String },
  name: { type: String, required: false },
  username: { type: String, required: false },
  email: { type: String, required: false },
  password: { type: String },
  date_created: Date,
  last_active: Date,
  disabled: { type: Boolean },
  email_verified: { type: Boolean },
  phone_verified: { type: Boolean },
  support_enabled: { type: Boolean, required: false },
  '2fa_enabled': { type: Boolean, required: false },
  '2fa_secret': { type: String, required: false },
  '2fa_backup_code': { type: String, required: false },
  default_account: { type: String, required: false },
  facebook_id: { type: String },
  twitter_id: { type: String },
  apple_id: { type: String },
  ip: { type: String },
  user_agent: { type: String },
  twitter: {
    accessToken: { type: String },
    refreshToken: { type: String },
    userId: { type: String },
    userName: { type: String },
    codeVerifier: { type: String },
    state: { type: String },
    name: { type: String },
    dateCreated: { type: Date },
    expiresIn: { type: Date },
  },
  google: {
    accessToken: { type: String },
    refreshToken: { type: String },
    userId: { type: String },
    email: { type: String },
    dateCreated: { type: Date },
  },
  facebook: {
    accessToken: { type: String },
    refreshToken: { type: String },
    userId: { type: String },
    email: { type: String },
  },
  apple: {
    accessToken: { type: String },
    refreshToken: { type: String },
    userId: { type: String },
    email: { type: String },
  },
  has_password: { type: Boolean },
  onboarded: { type: Boolean },
  permission: { type: String },
});

export const UserModel = model<User & Document>('User', UserSchema);

// export async function decryptFingerprint(fingerprint: string): Promise<string> {
//   return crypto.decrypt(fingerprint);
// }

export async function create(user: Partial<User>): Promise<Partial<User>> {
  const data: Partial<User> = {
    id: uuidv4(),
    name: user.name,
    email: user.email,
    date_created: new Date(),
    last_active: new Date(),
    support_enabled: false,
    '2fa_enabled': false,
    facebook_id: user.facebook_id,
    twitter_id: user.twitter_id,
  };

  if (user.password) {
    const salt = await bcrypt.genSalt(10);
    data.password = await bcrypt.hash(user.password, salt);
  }

  const newUser = new UserModel(data);
  await newUser.save();

  if (data.password) {
    delete data.password;
    data.has_password = true;
  }

  return data;
}

export async function get(id?: string, email?: string, social?: SocialProvider, permission?: string): Promise<User | null> {
  let data;
  const cond = {
    ...(permission && { 'account.permission': permission }),
  };

  if (social) {
    cond[`${social.provider}_id`] = social.id;
    data = await UserModel.find({ $or: [{ email: email }, cond] }).lean();
  } else {
    data = await UserModel.find({
      ...cond,
      ...(id && { id: id }),
      ...(email && { email: email }),
    }).lean();
  }

  if (data?.length) {
    data.forEach((u: User) => {
      u.has_password = u.password ? true : false;
      delete u.password;
    });
  }

  return id || email || social ? (data[0] as User) : null;
}

export async function addAccount(id: string, permission: string): Promise<User | null> {
  const user = await UserModel.findOne({ id: id });

  if (user) {
    // user.account.push({
    //   permission: permission,
    //   onboarded: false,
    // });
    user.markModified('account');
    return await user.save();
  }

  throw { message: 'No user with that ID' };
}

export async function deleteAccount(id: string): Promise<User | null> {
  const user = await UserModel.findOne({ id: id });

  if (user) {
    user.markModified('account');
    return await user.save();
  }

  throw { message: 'No user with that ID' };
}

export async function update(id: string, data: UpdateUserData): Promise<User | null> {
  return await UserModel.findOneAndUpdate({ id: id }, data);
}

// export async function deleteById(id: string): Promise<User | null> {
//   return await UserModel.deleteOne({ id: id });
// }

export async function updateTwitter(id: string, data: any): Promise<User | null> {
  return await UserModel.findOneAndUpdate({ id: id }, { twitter: data });
}

export async function password(id: string): Promise<{ password: string } | null | User> {
  return await UserModel.findOne({ id: id }).select({ password: 1 });
}

export async function verifyPassword(id: string, password: string): Promise<boolean> {
  const data = await UserModel.findOne({ id: id }).select({ name: 1, email: 1, password: 1 });

  const verified = data?.password ? await bcrypt.compare(password, data.password) : false;

  return verified;
}

export async function savePassword(id: string, newPassword: string, reset: boolean): Promise<User | null> {
  const salt = await bcrypt.genSalt(10);
  const hash = await bcrypt.hash(newPassword, salt);
  return await UserModel.findOneAndUpdate({ id: id }, { password: hash });
}

// 2FA related functions

export async function secret2FA(id?: string, email?: string): Promise<string | null> {
  const data = await UserModel.findOne({
    ...(id && { id: id }),
    ...(email && { email: email }),
  }).select({ '2fa_secret': 1 });

  // Replace this with the actual decryption logic
  const decrypted = data ? '' : null;

  return decrypted;
}

export async function saveBackup2FA(id: string, code: string): Promise<User | null> {
  const salt = await bcrypt.genSalt(10);
  const hash = await bcrypt.hash(code, salt);
  return await UserModel.findOneAndUpdate({ id: id }, { '2fa_backup_code': hash });
}

export async function verifyBackup2FA(id: string, email: string, code: string): Promise<boolean> {
  const data = await UserModel.findOne({
    ...(id && { id: id }),
    ...(email && { email: email }),
  }).select({ '2fa_backup_code': 1 });

  return data?.['2fa_backup_code'] ? await bcrypt.compare(code, data['2fa_backup_code']) : false;
}
