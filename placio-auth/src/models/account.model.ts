import { IAccount } from '@/interfaces/account.interface';
import { UserModel } from '@/models/users.model';
import mongoose, { Schema } from 'mongoose';
import { v4 as uuidv4 } from 'uuid';
import { subscription } from '@/models/subscription.model';

const AccountSchema: Schema = new Schema<IAccount>({
  id: { type: String, required: true, unique: true },
  plan: { type: String },
  name: { type: String },
  owner_email: { type: String },
  owner_name: { type: String },
  active: { type: Boolean, required: true },
  stripe_subscription_id: { type: String },
  stripe_customer_id: { type: String },
  date_created: { type: Date, required: true },
});

const Account = mongoose.model<IAccount>('Account', AccountSchema, 'account');

export async function createAccount(plan: string): Promise<IAccount> {
  const data = new Account({
    id: uuidv4(),
    name: 'My Account',
    active: true,
    date_created: new Date(),
    plan,
  });

  await data.save();
  return data;
}

export async function getAccount(id: string): Promise<Partial<IAccount> | null> {
  const accountData = await Account.findOne({ id }).lean();

  if (accountData) {
    const userData = await UserModel.findOne({
      'account.id': id,
      $or: [{ 'account.permission': 'owner' }, { 'account.permission': 'master' }],
    }).select({ name: 1, email: 1 });

    if (userData) {
      accountData.owner_email = userData.email;
      accountData.owner_name = userData.name;
    }
  }

  return accountData;
}

export async function sub(id: string): Promise<{ status?: string; data?: any }> {
  let subs, status;

  const accountData = await Account.findOne({ id });
  if (!accountData) throw { message: `Account doesn't exist` };

  if (accountData.plan !== 'free' && accountData.stripe_subscription_id) {
    subs = await subscription.retrieve(accountData.stripe_subscription_id);

    status = subs?.status !== 'active' ? subs?.latest_invoice?.payment_intent?.status : subs.status;

    if (status !== 'active' && status !== 'trialing') await Account.findOneAndUpdate({ id }, { active: false });
  } else if (accountData.plan === 'free') {
    status = 'active';
  }

  return {
    status,
    data: subscription,
  };
}

export async function update(id: string, data: Partial<IAccount>): Promise<IAccount | null> {
  return await Account.findOneAndUpdate({ id }, data);
}

export async function deleteAccount(id: string): Promise<IAccount | null> {
  return await Account.findOneAndRemove({ id });
}
