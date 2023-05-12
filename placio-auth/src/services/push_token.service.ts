import { UserModel } from '@/models/users.model';

export async function createPushToken(user: string, token: string) {
  return await UserModel.findOneAndUpdate({ id: user }, { $push: { push_token: token } });
}

export async function getPushToken(user: string, token: string) {
  const data = await UserModel.findOne({ id: user }).select({ push_token: 1 });
  return data.push_token?.length ? data.push_token : null;
}

export async function deletePushToken(token: string, user: string) {
  return await UserModel.findOneAndRemove({ user: user }, { $pull: { push_token: token } });
}
