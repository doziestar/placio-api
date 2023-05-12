import { addAccount, createUser, getUser, savePassword } from '@/models/users.model';
import { Request, Response } from 'express';
import chalk from 'chalk';
import { validate } from '@models/utility.model';
import { logger, log } from 'handlebars';
import { createAccount, getAccount } from '@/models/account.model';
import { sendMail } from '@/models/mail.model';
import { AuthController } from './auth.controller';

const auth = new AuthController();

class AccountController {
  async create(req: Request, res: Response): Promise<void> {
    const data = req.body;
    validate(data, ['email', 'name', 'password']);

    // confirm_password field is a dummy field to prevent bot signups
    if (data.hasOwnProperty('confirm_password') && data.confirm_password) throw { message: 'Registration denied' };

    // check if user has already registered an account
    let userData = await getUser(null, data.email);

    if (userData) {
      // user already owns an account
      if (userData.permission === 'owner')
        throw {
          inputError: 'email',
          message: 'You have already registered an account',
        };

      // flag for authController to notify onboarding ui
      // that the user's existing account was used
      req.body.duplicate_user = true;
      req.body.has_password = userData.has_password;

      // save the new password if it exists and user doesn't have one
      if (!req.body.has_password && req.body.password) await savePassword(userData.id, req.body.password, true);
    }

    console.log(chalk.cyan('Creating account for: ') + data.email);
    const accountData = await createAccount(data.plan);
    req.body.account_id = accountData.id; // pass to auth controller to select new account

    // create the user and assign to account
    userData = !userData ? await createUser(data, accountData.id) : userData;
    await addAccount(userData.id, accountData.id, 'owner');

    console.log(chalk.green('Account created for: ') + data.email);

    // send welcome email
    await sendMail({
      to: userData.email,
      template: req.body.duplicate_user && req.body.has_password ? 'duplicate-user' : 'new-account',
      content: { name: userData.name },
    });

    // authenticate the user
    return await auth.signUp(req, res);
  }

  async getAccount(req, res) {
    const data = await getAccount(req.account);
    return res.status(200).send({ data: data });
  }

  //   async subscription(req, res) {
  //     const subscription = await subs(req.account);

  //     // format the data
  //     if (subscription?.data) {
  //       const start = new Date(subscription.data.current_period_start * 1000).toISOString().split('T')[0].split('-');
  //       const end = new Date(subscription.data.current_period_end * 1000).toISOString().split('T')[0].split('-');

  //       subscription.data = {
  //         current_period_start: `${start[2]} ${convertToMonthName(start[1])} ${start[0]}`,
  //         current_period_end: `${end[2]} ${convertToMonthName(end[1])} ${end[0]}`,
  //       };
  //     }

  //     return res.status(200).send({
  //       data: {
  //         status: subscription.status,
  //         object: subscription.data,
  //       },
  //     });
  //   }

  //   async plan(req: Request, res: Response): Promise<void> {
  //     const data = req.body;
  //     const stripeData = {};

  //     validate(data, ['plan']);

  //     // check the plan exists
  //     const plan = settings.plans.find(x => x.id === data.plan);
  //     utility.assert(plan, `Plan doesn't exist`);

  //     const accountData = await account.get(req.account);
  //     utility.assert(accountData, 'No account with that ID');

  //     // process stripe subscription for non-free accounts
  //     // if a 2-factor payment hasn't occured, create the stripe subscription
  //     if (data.plan !== 'free') {
  //       if (data.stripe === undefined) {
  //         utility.assert(data.token?.id, 'Please enter your credit card details');

  //         // create a stripe customer and subscribe them to a plan
  //         stripeData.customer = await stripe.customer.create(accountData.owner_email, data.token.id);
  //         stripeData.subscription = await stripe.customer.subscribe(stripeData.customer.id, data.plan);

  //         // check for an incomplete payment that requires 2-factor authentication
  //         if (stripeData.subscription?.latest_invoice?.payment_intent?.status === 'requires_action') {
  //           logger.info(chalk.yellow('Stripe payment requires further action'));

  //           return res.status(200).send({
  //             requires_payment_action: true,
  //             customer: { id: stripeData.customer.id },
  //             subscription: {
  //               id: stripeData.subscription.id,
  //               price: stripeData.subscription.price,
  //             },
  //             client_secret: stripeData.subscription.latest_invoice.payment_intent.client_secret,
  //           });
  //         }
  //       }

  //       // stripe info hasn't been passed back as part of 2-factor
  //       if (!data.stripe) data.stripe = stripeData;
  //     } else {
  //       // nullify stripe data on free accounts
  //       data.stripe = {
  //         customer: { id: null },
  //         subscription: { id: null },
  //       };
  //     }

  //     // update the account with plan details
  //     await account.update(req.account, {
  //       plan: data.plan,
  //       stripe_customer_id: data.stripe?.customer?.id,
  //       stripe_subscription_id: data.stripe?.subscription?.id,
  //     });

  //     // send email
  //     if (data.plan !== 'free') {
  //       await mail.send({
  //         to: accountData.owner_email,
  //         template: 'new_plan',
  //         content: {
  //           name: accountData.owner_name,
  //           plan: plan.name,
  //           price: `${plan.currency.symbol}${plan.price}`,
  //         },
  //       });
  //     }

  //     console.log(chalk.green('Customer added to plan'));
  //     log.create('Customer added to plan', { plan: plan }, req);
  //     res.status(200).send({ plan: data.plan, subscription: 'active', onboarded: false });
  //   }

  //   async updatePlan(req: Request, res: Response): Promise<void> {
  //     const data = req.body;
  //     utility.validate(data, ['plan']);

  //     const accountID = req.permission === 'master' ? data.id : req.account;
  //     const plan = settings.plans.find(x => x.id === data.plan);
  //     utility.assert(plan, 'No plan with that ID');

  //     const accountData = await account.get(accountID);
  //     utility.assert(accountData, 'Account does not exist');

  //     // user is upgrading from paid to free,
  //     // direct them to the upgrade view
  //     if (accountData.plan === 'free' && plan.id !== 'free') {
  //       if (req.permission === 'master') {
  //         throw {
  //           message: 'The account holder will need to enter their card details and upgrade to a paid plan.',
  //         };
  //       } else {
  //         return res.status(402).send({ message: 'Please upgrade your account', plan: plan.id });
  //       }
  //     }

  //     if (plan.id === 'free') {
  //       // user is downgrading - cancel the stripe subscription
  //       if (accountData.stripe_subscription_id) {
  //         const subscription = await stripe.subscription(accountData.stripe_subscription_id);
  //         await account.update(req.account, {
  //           stripe_subscription_id: null,
  //           plan: plan.id,
  //         });

  //         if (subscription.status !== 'canceled') await stripe.subscription.delete(accountData.stripe_subscription_id);
  //       }
  //     } else {
  //       // user is switching to a different paid plan
  //       if (accountData.stripe_subscription_id) {
  //         // check for active subscription
  //         let subscription = await stripe.subscription(accountData.stripe_subscription_id);

  //         if (subscription.status === 'trialing' || subscription.status === 'active') {
  //           subscription = await stripe.subscription.update(subscription, plan.id);
  //           await account.update(accountData.id, { plan: plan.id });
  //         } else if (subscription.status === 'canceled') {
  //           // user previously had a subscription, but is now cancelled - create a new one
  //           await account.update(req.account, {
  //             stripe_subscription_id: null,
  //             plan: 'free',
  //           });

  //           return req.permission === 'master'
  //             ? res.status(500).send({
  //                 message: 'The account holder will need to enter their card details and upgrade to a paid plan.',
  //               })
  //             : res.status(402).send({
  //                 message: 'Your subscription was cancelled, please upgrade your account',
  //               });
  //         }
  //       }
  //     }

  //     // notify the user
  //     await mail.send({
  //       to: accountData.owner_email,
  //       template: 'plan-updated',
  //       content: {
  //         name: accountData.owner_name,
  //         plan: plan.name,
  //       },
  //     });

  //     // done
  //     return res.status(200).send({
  //       message: `Your account has been updated to the ${plan.name} plan`,
  //       data: { plan: plan.id },
  //     });
  //   }
}

export default new AccountController();
