import { Strategy as FacebookStrategy } from 'passport-facebook';
import { Strategy as TwitterStrategy } from 'passport-twitter';
import { Strategy as AppleStrategy } from 'passport-apple';
import strategy from 'passport';
import { FACEBOOK_APP_ID, FACEBOOK_APP_SECRET, TWITTER_CONSUMER_KEY, TWITTER_CONSUMER_SECRET } from '@/config';

import { UserModel } from '@models/users.model';

const generateUsername = name => {
  const username = name.replace(/\s+/g, '').toLowerCase();
  const randomNumber = Math.floor(Math.random() * 10000);

  return `${username}${randomNumber}`;
};

const upsertUser = async (provider, id, email, name, accessToken, refreshToken) => {
  const existingUser = await UserModel.findOne({ email });

  if (existingUser) {
    if (existingUser[`${provider}_id`]) {
      // The user has already logged in with this provider
      return existingUser;
    } else {
      // Add the provider ID, access and refresh tokens to the existing user
      existingUser[`${provider}_id`] = id;
      existingUser[provider] = {
        accessToken,
        refreshToken,
      };
      await existingUser.save();
      return existingUser;
    }
  }

  // Create a new user if no existing user found
  let username = name;

  if (!username) {
    username = generateUsername(name);
  }

  const userData = {
    [`${provider}_id`]: id,
    email,
    name: username,
    [provider]: {
      accessToken,
      refreshToken,
    },
  };

  return await UserModel.create(userData);
};

// Facebook Strategy
strategy.use(
  new FacebookStrategy(
    {
      clientID: FACEBOOK_APP_ID,
      clientSecret: FACEBOOK_APP_SECRET,
      callbackURL: '/auth/facebook/callback',
      profileFields: ['id', 'emails', 'name'],
    },
    async (accessToken, refreshToken, profile, done) => {
      try {
        const email = profile.emails[0].value;
        const name = `${profile.name.givenName} ${profile.name.familyName}`;

        const user = await upsertUser('facebook', profile.id, email, name, accessToken, refreshToken);

        done(null, user);
      } catch (error) {
        done(error);
      }
    },
  ),
);

// Twitter Strategy
strategy.use(
  new TwitterStrategy(
    {
      consumerKey: TWITTER_CONSUMER_KEY,
      consumerSecret: TWITTER_CONSUMER_SECRET,
      callbackURL: '/auth/twitter/callback',
      includeEmail: true,
    },
    async (accessToken, refreshToken, profile, done) => {
      try {
        const email = profile.emails[0].value;
        const name = profile.displayName;

        const user = await upsertUser('twitter', profile.id, email, name, accessToken, refreshToken);

        done(null, user);
      } catch (error) {
        done(error);
      }
    },
  ),
);

// Apple Strategy
// strategy.use(
//   new AppleStrategy(
//     {
//       clientID: process.env.APPLE_CLIENT_ID,
//       teamID: process.env.APPLE_TEAM_ID,
//       keyID: process.env.APPLE_KEY_ID,
//       keyContent: process.env.APPLE_PRIVATE_KEY_CONTENT,
//       callbackURL: '/auth/apple/callback',
//       scope: ['name', 'email'],
//     },
//     async (accessToken, refreshToken, idToken, profile, done) => {
//       try {
//         const email = idToken.email;
//         const name = `${idToken.given_name} ${idToken.family_name}`;

//         const user = await upsertUser('apple', idToken.sub, email, name, accessToken, refreshToken);

//         done(null, user);
//       } catch (error) {
//         done(error);
//       }
//     },
//   ),
// );

export default strategy;
