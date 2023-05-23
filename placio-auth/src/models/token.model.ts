import mongoose, { Schema, Document } from 'mongoose';
import Cryptr from 'cryptr';
import { CRYPTO_SECRET } from '@config';
import { v4 as uuidv4 } from 'uuid';
import { IToken } from '@interfaces/token.interface';
import { TokenResponsePayload } from '@/interfaces/auth.interface';
import exp from 'constants';

const crypto = new Cryptr(CRYPTO_SECRET);

const TokenSchema = new Schema<IToken>({
  id: { type: String, required: true, unique: true },
  provider: { type: String, required: true },
  jwt: { type: String },
  access: { type: String },
  refresh: { type: String },
  user_id: { type: String, required: true },
});

export const Token = mongoose.model<IToken>('Token', TokenSchema, 'token');

export async function saveToken(provider: string, data: TokenResponsePayload, user: string) {
  const tokenData = await Token.findOne({ provider: provider, user_id: user });

  if (tokenData) {
    await Token.findOneAndUpdate({ id: tokenData.id, user_id: user }, data);
  } else {
    const newToken = new Token({
      id: uuidv4(),
      provider: provider,
      access: data.accessToken.token,
      refresh: data.refreshToken.token,
      user_id: user,
    });

    await newToken.save();
  }

  return data;
}

export async function getToken(id?: string, provider?: string, user?: string, skipDecryption?: boolean) {
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

export async function verifyToken(provider: string, user: string) {
  console.log('verifyToken', provider, user);
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
