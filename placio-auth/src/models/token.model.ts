import mongoose, { Schema, Document } from 'mongoose';
import Cryptr from 'cryptr';
import { v4 as uuidv4 } from 'uuid';
import { IToken } from '@interfaces/token.interface';

const crypto = new Cryptr(process.env.CRYPTO_SECRET);

const TokenSchema = new Schema<IToken>({
  id: { type: String, required: true, unique: true },
  provider: { type: String, required: true },
  jwt: { type: String },
  access: { type: String },
  refresh: { type: String },
  user_id: { type: String, required: true },
});

const Token = mongoose.model<IToken>('Token', TokenSchema, 'token');

export async function save(provider: string, data: any, user: string) {
  if (data.access) data.access = crypto.encrypt(data.access);
  if (data.refresh) data.refresh = crypto.encrypt(data.refresh);

  const tokenData = await Token.findOne({ provider: provider, user_id: user });

  if (tokenData) {
    await Token.findOneAndUpdate({ id: tokenData.id, user_id: user }, data);
  } else {
    const newToken = new Token({
      id: uuidv4(),
      provider: provider,
      jwt: data.jwt,
      access: data.access,
      refresh: data.refresh,
      user_id: user,
    });

    await newToken.save();
  }

  return data;
}

export async function get(id?: string, provider?: string, user?: string, skipDecryption?: boolean) {
  const data = await Token.find({
    user_id: user,
    ...(id && { id: id }),
    ...(provider && { provider: provider }),
  });

  if (data.length && !skipDecryption) {
    data.forEach(token => {
      if (token.access) token.access = crypto.decrypt(token.access);
      if (token.refresh) token.refresh = crypto.decrypt(token.refresh);
    });
  }

  return data;
}

export async function verify(provider: string, user: string) {
  const data = await Token.find({ user_id: user, provider: provider });
  return data.length ? true : false;
}

export async function deleteToken(id?: string, provider?: string, user?: string) {
  return await Token.deleteOne({
    user_id: user,
    ...(provider && { provider: provider }),
    ...(id && { id: id }),
  });
}

export async function generateTemporaryToken() {
  const token = crypto.encrypt(uuidv4());
  const newToken = new Token({
    id: uuidv4(),
    provider: 'temporary',
    jwt: token,
    access: token,
    refresh: token,
    user_id: {
      ip: '',
      user_agent: '',
    },
  });

  await newToken.save();

  return token;
}

export async function verifyTemporaryToken(token: string) {
  const data = await Token.find({ provider: 'temporary', jwt: token });
  return data.length ? true : false;
}
