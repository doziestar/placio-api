import { Document } from 'mongoose';

interface IAccount extends Document {
  id: string;
  plan?: string;
  name?: string;
  active: boolean;
  owner_email?: string;
  owner_name?: string;
  stripe_subscription_id?: string;
  stripe_customer_id?: string;
  date_created: Date;
}

export { IAccount };
