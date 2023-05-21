import { Document } from 'mongoose';

interface IGeneralSettings extends Document {
  createdAt: Date;
  updatedAt: Date;
  id: string;
  language: string;
  theme: string;
  userID: string;
  privacy: string;
  notifications: INotificationsSettings['_id'];
  content: IContentSettings['_id'];
}

interface INotificationsSettings extends Document {
  id: string;
  userID: string;
  emailNotifications: boolean;
  pushNotifications: boolean;
  directMessageNotifications: boolean;
  likeNotifications: boolean;
  commentNotifications: boolean;
  mentionNotifications: boolean;
  followNotifications: boolean;
}

interface IAccountSettings extends Document {
  id: string;
  accountID: string;
  twoFactorAuthentication: boolean;
  blockedUsers: string[];
  mutedUsers: string[];
}

interface IContentSettings extends Document {
  id: string;
  mediaVisibility: string;
  explicitContentFilter: string;
  defaultPostPrivacy: string;
  autoplayVideos: boolean;
  displaySensitiveMedia: boolean;
  userID: string;
}

export { IGeneralSettings, INotificationsSettings, IAccountSettings, IContentSettings };
