import { Config } from '@/interfaces/config.interface';

const config: Config = {
  domain: 'http://localhost:3000',
  'log.dir': 'log',
  stripe: {
    plans: [
      {
        id: 'free',
        name: 'Free',
        price: 0,
        interval: 'month',
        currency: {
          name: 'usd',
          symbol: '$',
        },
        features: [
          {
            name: 'Awesome feature',
            checked: true,
          },
          {
            name: 'Another amazing feature',
            checked: true,
          },
          {
            name: 'The best feature ever',
            checked: true,
          },
        ],
      },
      // additional plans...
    ],
  },
  mailgun: {
    domain: 'mail.hubhub.app',
    host: 'api.eu.mailgun.net',
    sender: 'johannes@mail.hubhub.app',
    base_url: 'https://api.eu.mailgun.net/v3',
  },
  facebook: {
    scope: ['email'],
    callback_url: 'http://localhost:8080/auth/facebook/callback',
  },
  twitter: {
    scope: [],
    callback_url: 'http://localhost:8080/auth/twitter/callback',
  },
  permissions: {
    master: {
      master: true,
      owner: true,
      admin: true,
      user: true,
      developer: true,
    },
    owner: {
      master: false,
      owner: true,
      admin: true,
      user: true,
      developer: true,
    },
    admin: {
      master: false,
      owner: false,
      admin: true,
      user: true,
      developer: false,
    },
    user: {
      master: false,
      owner: false,
      admin: false,
      user: true,
      developer: false,
    },
    developer: {
      master: false,
      owner: false,
      admin: false,
      user: true,
      developer: true,
    },
  },
  api_scopes: [
    'account.read',
    'account.update',
    'account.delete',
    'billing.read',
    'billing.update',
    'invite.create',
    'invite.read',
    'invite.delete',
    'key.create',
    'key.read',
    'key.update',
    'key.delete',
    'user.read',
    'user.update',
    'user.delete',
    'event.create',
  ],
  token: {
    duration: 604800,
  },
  throttle: {
    api: {
      max: 1000,
      windowMs: 3600000,
      headers: true,
      message: 'Too many API calls from this IP, please try again soon.',
    },
    signup: {
      max: 5,
      windowMs: 3600000,
      headers: true,
      message: 'You have created too many accounts.',
    },
    signin: {
      max: 5,
      windowMs: 300000,
      headers: true,
      message: 'Too many sign in attempts, please try again in a few minutes.',
    },
    password_reset: {
      max: 3,
      windowMs: 300000,
      headers: true,
      message: 'You have reached the limit of password reset requests.',
    },
  },
};

export default config;
