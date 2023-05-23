import { IAccountSettings, IContentSettings, IGeneralSettings, INotificationsSettings } from '@/interfaces/settings.interface';
import mongoose, { Document, Schema } from 'mongoose';

const GeneralSettingsSchema = new Schema<IGeneralSettings>({
  createdAt: { type: Date, default: Date.now },
  updatedAt: { type: Date, default: Date.now },
  id: String,
  language: String,
  theme: String,
  userID: { type: Schema.Types.ObjectId, ref: 'User' },
  privacy: { type: String, default: 'public' },
  notifications: { type: Schema.Types.ObjectId, ref: 'NotificationsSettings' },
  content: { type: Schema.Types.ObjectId, ref: 'ContentSettings' },
});

const NotificationsSettingsSchema = new Schema<INotificationsSettings>({
  id: String,
  userID: { type: String, unique: true },
  emailNotifications: Boolean,
  pushNotifications: Boolean,
  directMessageNotifications: Boolean,
  likeNotifications: Boolean,
  commentNotifications: Boolean,
  mentionNotifications: Boolean,
  followNotifications: Boolean,
});

const AccountSettingsSchema = new Schema<IAccountSettings>({
  id: String,
  accountID: { type: String, unique: true },
  twoFactorAuthentication: Boolean,
  blockedUsers: [String],
  mutedUsers: [String],
});

interface IConnectedAccount extends Document {
  id: string;
  provider: string;
  userID: string;
}

const ConnectedAccountSchema = new Schema<IConnectedAccount>({
  id: String,
  provider: String,
  userID: String,
});

const ContentSettingsSchema = new Schema<IContentSettings>({
  id: String,
  mediaVisibility: String,
  explicitContentFilter: String,
  defaultPostPrivacy: String,
  autoplayVideos: Boolean,
  displaySensitiveMedia: Boolean,
  userID: { type: String, unique: true },
});

const GeneralSettings = mongoose.model('GeneralSettings', GeneralSettingsSchema);
const NotificationsSettings = mongoose.model('NotificationsSettings', NotificationsSettingsSchema);
const AccountSettings = mongoose.model('AccountSettings', AccountSettingsSchema);
const ConnectedAccount = mongoose.model('ConnectedAccount', ConnectedAccountSchema);
const ContentSettings = mongoose.model('ContentSettings', ContentSettingsSchema);

export { GeneralSettings, NotificationsSettings, AccountSettings, ConnectedAccount, ContentSettings };
