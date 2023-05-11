import { config } from 'config';
import passport from 'passport';
import FacebookStrategy from 'passport-facebook';
import TwitterStrategy from 'passport-twitter';

const facebook = config.get('facebook');
const twitter = config.get('twitter');

passport.serializeUser((user, done) => {
  done(null, user);
});
passport.deserializeUser((obj, done) => {
  done(null, obj);
});

if (process.env.FACEBOOK_APP_ID && process.env.FACEBOOK_APP_SECRET) {
  passport.use(
    new FacebookStrategy(
      {
        clientID: process.env.FACEBOOK_APP_ID,
        clientSecret: process.env.FACEBOOK_APP_SECRET,
        callbackURL: facebook.callback_url,
        profileFields: ['id', 'name', 'email'],
        enableProof: true,
        passReqToCallback: true,
      },
      async function (req, accessToken, refreshToken, profile, done) {
        await handleCallback(req, profile, { access: accessToken, refresh: refreshToken }, done);
      },
    ),
  );
}

if (process.env.TWITTER_API_KEY && process.env.TWITTER_API_SECRET) {
  passport.use(
    new TwitterStrategy(
      {
        consumerKey: process.env.TWITTER_API_KEY,
        consumerSecret: process.env.TWITTER_API_SECRET,
        callbackURL: twitter.callback_url,
        userProfileURL: 'https://api.twitter.com/1.1/account/verify_credentials.json?include_email=true',
        passReqToCallback: true,
      },
      async function (req, accessToken, refreshToken, profile, done) {
        await handleCallback(req, profile, { access: accessToken, refresh: refreshToken }, done);
      },
    ),
  );
}

async function handleCallback(req: Request, profile: any, tokens: { access: string; refresh: string }, done: (error: any, user?: any) => void) {
  if (!profile) return done({ message: 'Error getting profile.' });

  const provider = profile.provider;
  const email = profile.emails[0]?.value;
  const data = {
    name: profile?.name?.givenName || profile.username,
    email: email,
    ...(provider === 'facebook' && { facebook_id: profile.id }),
    ...(provider === 'twitter' && { twitter_id: profile.id }),
  };

  let userData = await userModel.get(null, email, null, {
    provider: provider,
    id: profile.id,
  });

  if (req.session.signup && !req.session.invite && userData) {
    userData.accounts = await userModel.account(userData.id);

    if (!userData.accounts.find(x => x.permission === 'owner')) {
      const accountData = await accountModel.create();
      await userModel.account.add(userData.id, accountData.id, 'owner');
      await userModel.update(userData.id, accountData.id, {
        default_account: accountData.id,
      });
    }
  }

  if (req.session.invite) {
    const inviteData = await inviteModel.get(req.session.invite);
    if (!inviteData)
      return done({
        message: 'Invalid invite. Please contact the account holder',
      });

    if (userData) {
      const social = { default_account: inviteData.account_id };
      social[`${provider}_id`] = profile.id;
      await userModel.update(userData.id, userData.account_id, social);
    } else {
      userData = await userModel.create(data, inviteData.account_id);
    }

    await userModel.account.add(userData.id, inviteData.account_id, inviteData.permission);
    await inviteModel.update(req.session.invite, { used: true });
  } else if (userData) {
    if (!userData[`${provider}_id`]) {
      const social = {};
      social[`${provider}_id`] = profile.id;
      await userModel.update(userData.id, userData.account_id, social);
    }

    await TokenModel.save(provider, tokens, userData.id);
    return done(null, profile);
  } else {
    // const accountData = await accountModel.create();
    userData = await userModel.create(data, accountData.id);
    await UserModel.account.add(userData.id, accountData.id, 'owner');
  }

  await TokenModel.save(provider, tokens, userData.id);
  done(null, profile);
}
