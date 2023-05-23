import { Document, Schema } from 'mongoose';

interface IGeneralSettings extends Document {
  createdAt: Date;
  updatedAt: Date;
  id: string;
  language: string;
  theme: string;
  userID: Schema.Types.ObjectId;
  privacy: string;
  notifications: INotificationsSettings['_id'];
  content: IContentSettings['_id'];
}

interface INotificationsSettings extends Document {
  id: string;
  userID: Schema.Types.ObjectId;
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
  accountID: Schema.Types.ObjectId;
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
  userID: Schema.Types.ObjectId;
}

export { IGeneralSettings, INotificationsSettings, IAccountSettings, IContentSettings };
