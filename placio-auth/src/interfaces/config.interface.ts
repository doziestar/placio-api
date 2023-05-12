interface Currency {
  name: string;
  symbol: string;
}

interface Feature {
  name: string;
  checked: boolean;
}

interface Plan {
  id: string;
  name: string;
  price: number;
  interval: string;
  currency: Currency;
  features: Feature[];
}

interface Stripe {
  plans: Plan[];
}

interface Mailgun {
  domain: string;
  host: string;
  sender: string;
  base_url: string;
}

interface SocialAuth {
  scope: string[];
  callback_url: string;
}

interface PermissionType {
  master: boolean;
  owner: boolean;
  admin: boolean;
  user: boolean;
  developer: boolean;
}

interface Permissions {
  master: PermissionType;
  owner: PermissionType;
  admin: PermissionType;
  user: PermissionType;
  developer: PermissionType;
}

interface Throttle {
  max: number;
  windowMs: number;
  headers: boolean;
  message: string;
}

interface Throttles {
  api: Throttle;
  signup: Throttle;
  signin: Throttle;
  password_reset: Throttle;
}

interface Config {
  domain: string;
  'log.dir': string;
  stripe: Stripe;
  mailgun: Mailgun;
  facebook: SocialAuth;
  twitter: SocialAuth;
  permissions: Permissions;
  api_scopes: string[];
  token: Record<string, number>;
  throttle: Throttles;
}

export { Config, Currency, Feature, Plan, Stripe, Mailgun, SocialAuth, PermissionType, Permissions, Throttle, Throttles };
