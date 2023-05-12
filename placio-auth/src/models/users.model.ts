import { model, Schema, Document } from 'mongoose';
import { SocialProvider, UpdateUserData, User } from '@interfaces/users.interface';
import { v4 as uuidv4 } from 'uuid';
import bcrypt from 'bcrypt';
import Cryptr from 'cryptr';

const crypto = new Cryptr(process.env.CRYPTO_SECRET);

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
  push_token: [
    {
      type: String,
    },
  ],
  account: [
    {
      type: Schema.Types.ObjectId,
      ref: 'Account',
    },
  ],
  ip: { type: String },
  user_agent: { type: String },
  twitter: {
    accessToken: { type: String },
    refreshToken: { type: String },
    UserModelId: { type: String },
    UserModelName: { type: String },
    codeVerifier: { type: String },
    state: { type: String },
    name: { type: String },
    dateCreated: { type: Date },
    expiresIn: { type: Date },
  },
  google: {
    accessToken: { type: String },
    refreshToken: { type: String },
    UserModelId: { type: String },
    email: { type: String },
    dateCreated: { type: Date },
  },
  facebook: {
    accessToken: { type: String },
    refreshToken: { type: String },
    UserModelId: { type: String },
    email: { type: String },
  },
  apple: {
    accessToken: { type: String },
    refreshToken: { type: String },
    UserModelId: { type: String },
    email: { type: String },
  },
  has_password: { type: Boolean },
  onboarded: { type: Boolean },
  permission: { type: String },
});

export const UserModel = model<User & Document>('UserModel', UserSchema);

// export async function decryptFingerprint(fingerprint: string): Promise<string> {
//   return crypto.decrypt(fingerprint);
// }

export async function decryptFingerprint(fingerprint: string) {
  return crypto.decrypt(fingerprint);
}

export async function createUser(user: Partial<User>, account: string) {
  const data = {
    id: uuidv4(),
    name: user.name,
    email: user.email,
    date_created: new Date(),
    last_active: new Date(),
    support_enabled: false,
    '2fa_enabled': false,
    facebook_id: user.facebook_id,
    twitter_id: user.twitter_id,
    default_account: account,
    password: user.password,
    has_password: false,
    account_id: account,
  };

  if (user.password) {
    const salt = await bcrypt.genSalt(10);
    data.password = await bcrypt.hash(user.password, salt);
  }

  const newUserModel = new UserModel(data);
  await newUserModel.save();

  if (data.password) {
    delete data.password;
    data.has_password = true;
  }

  data.account_id = account;
  return data;
}

export async function getUser(id?: string, email?: string, account?: string, social?: SocialProvider, permission?: string) {
  let data;
  const cond = {
    ...(account && { 'account.id': account }),
    ...(permission && { 'account.permission': permission }),
  };

  if (social) {
    cond[`${social.provider}_id`] = social.id;
    data = await UserModel.find({ $or: [{ email: email }, cond] }).lean();
  } else {
    console.log('cond', cond);
    data = await UserModel.find({
      ...cond,
      ...{
        ...(id && { id: id }),
        ...(email && { email: email }),
      },
    }).lean();
    console.log('data,', data);
  }

  if (data?.length) {
    data.forEach(u => {
      u.account_id = account || u.default_account;
      const a = u.account.find(x => x.id === u.account_id);
      u.permission = a.permission;
      u.onboarded = a.onboarded;

      u.has_password = u.password ? true : false;
      delete u.password;
      delete u.account;
    });
  }

  return id || email || social ? data[0] : data;
}

export async function account(id: string) {
  const data = await UserModel.aggregate([
    { $match: { id: id } },
    { $project: { id: 1, account: 1, email: 1 } },
    {
      $lookup: {
        from: 'account',
        localField: 'account.id',
        foreignField: 'id',
        as: 'account_data',
      },
    },
  ]);

  return data[0]?.account.map(a => {
    return {
      id: a.id,
      UserModel_id: data[0].id,
      permission: a.permission,
      name: data[0].account_data.find(x => x.id === a.id)?.name,
    };
  });
}

export async function addAccount(id: string, account: string, permission: string) {
  const data = await UserModel.findOne({ id: id });

  if (data) {
    data.account.push({
      id: account,
      permission: permission,
      onboarded: false,
    });
    data.markModified('account');
    return await data.save();
  }

  throw { message: `No UserModel with that ID` };
}

// export async function deleteAccount(id: string, account: string) {
//   const data = await UserModel.findOne({ id: id });

//   if (data) {
//     data.account.splice(
//       data.account.findIndex(x => x.id === account),
//       1,
//     );
//     data.markModified('account');
//     return await data.save();
//   }

//   throw { message: `No UserModel with that ID` };
// }

export async function password(id: string, account: string) {
  return await UserModel.findOne({ id: id, 'account.id': account }).select({
    password: 1,
  });
}

export async function verifyPassword(id: string, account: string, password: string) {
  const data = await UserModel.findOne({ id: id, 'account.id': account }).select({
    name: 1,
    email: 1,
    password: 1,
  });

  const verified = data?.password ? await bcrypt.compare(password, data.password) : false;

  delete data.password;
  return verified ? data : false;
}

export async function savePassword(id: string, password: string, reset: boolean) {
  const salt = await bcrypt.genSalt(10);
  const hash = await bcrypt.hash(password, salt);
  return await UserModel.findOneAndUpdate({ id: id }, { password: hash });
}

export async function get2faSecret(id?: string, email?: string) {
  const data = await UserModel.findOne({
    ...(id && { id: id }),
    ...(email && { email: email }),
  }).select({ '2fa_secret': 1 });

  return data ? crypto.decrypt(data['2fa_secret']) : null;
}

export async function save2faBackupCode(id: string, code: string) {
  const salt = await bcrypt.genSalt(10);
  const hash = await bcrypt.hash(code, salt);
  return await UserModel.findOneAndUpdate({ id: id }, { '2fa_backup_code': hash });
}

export async function verify2faBackupCode(id: string, email: string, account: string, code: string) {
  const data = await UserModel.findOne({
    ...(id && { id: id, 'account.id': account }),
    ...(email && { email: email }),
  }).select({ '2fa_backup_code': 1 });

  return data?.['2fa_backup_code'] ? await bcrypt.compare(code, data['2fa_backup_code']) : false;
}

export async function update(id: string, account: string, data: Partial<User>) {
  if (data.onboarded || data.permission) {
    UserModel.findOne({ id: id, 'account.id': account }, (err, doc) => {
      if (err) throw err;
      if (!doc) throw { message: `No UserModel with that ID` };

      const index = doc.account.findIndex(x => x.id === account);

      if (data.onboarded) doc.account[index].onboarded = data.onboarded;

      if (data.permission) doc.account[index].permission = data.permission;

      doc.markModified('account');
      doc.save();
    });
  } else {
    await UserModel.findOneAndUpdate({ id: id, 'account.id': account }, data);
  }

  return data;
}

export async function deleteUserModel(id: string, account: string) {
  return await UserModel.deleteMany({
    ...(id && { id: id }),
    'account.id': account,
  });
}
