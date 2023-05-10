import { Document, Schema } from 'mongoose';

interface User extends Document {
  id: string;
  fingerprint?: string;
  name?: string;
  email?: string;
  password?: string;
  date_created?: Date;
  last_active?: Date;
  disabled?: boolean;
  support_enabled?: boolean;
  '2fa_enabled'?: boolean;
  '2fa_secret'?: string;
  '2fa_backup_code'?: string;
  default_account?: string;
  facebook_id?: string;
  twitter_id?: string;
  ip?: string;
  user_agent?: string;
  twitter?: {
    accessToken?: string;
    refreshToken?: string;
    userId?: string;
    userName?: string;
    codeVerifier?: string;
    state?: string;
    name?: string;
    dateCreated?: Date;
    expiresIn?: Date;
  };
  google?: {
    accessToken?: string;
    refreshToken?: string;
    userId?: string;
    email?: string;
    dateCreated?: Date;
  };
  has_password?: boolean;
  onboarded?: boolean;
  permission?: string;
  generalSettings?: {
    type: Schema.Types.ObjectId;
    ref: 'generalSettings';
  };
}

interface SocialProvider {
  provider: string;
  id: string;
}

interface UpdateUserData {
  onboarded?: boolean;
  permission?: string;
}

export { User, SocialProvider, UpdateUserData };
