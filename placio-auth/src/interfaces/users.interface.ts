import { Document, Schema } from 'mongoose';

interface User extends Document {
  id: string;
  fingerprint?: string;
  name?: string;
  username?: string;
  email?: string;
  password?: string;
  date_created?: Date;
  last_active?: Date;
  disabled?: boolean;
  support_enabled?: boolean;
  email_verified?: boolean;
  phone_verified?: boolean;
  '2fa_enabled'?: boolean;
  '2fa_secret'?: string;
  '2fa_backup_code'?: string;
  default_account?: string;
  facebook_id?: string;
  twitter_id?: string;
  apple_id?: string;
  ip?: string;
  account?: string[];
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
  facebook?: {
    accessToken?: string;
    refreshToken?: string;
    userId?: string;
    email?: string;
  };
  apple?: {
    accessToken?: string;
    refreshToken?: string;
    userId?: string;
    email?: string;
  };
  has_password?: boolean;
  onboarded?: boolean;
  permission?: string;
  push_token?: string[];
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
