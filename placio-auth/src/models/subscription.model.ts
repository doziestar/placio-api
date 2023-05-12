import Stripe from 'stripe';
import { STRIPE_SECRET_API_KEY } from '@/config/index';

const stripe = new Stripe(STRIPE_SECRET_API_KEY, {} as Stripe.StripeConfig);

export const subscription = {
  async retrieve(id: string) {
    return await stripe.subscriptions.retrieve(id, {
      expand: ['latest_invoice.payment_intent'],
    });
  },
  async update(subscription: Stripe.Subscription, plan: string) {
    return await stripe.subscriptions.update(subscription.id, {
      items: [{ id: subscription.items.data[0].id, plan }],
    });
  },
  async delete(id: string) {
    return await stripe.subscriptions.del(id);
  },
};

export const customer = {
  async retrieve(id: string) {
    return stripe.customers.retrieve(id);
  },
  async create(email: string, token: string) {
    return await stripe.customers.create({
      email: email,
      source: token,
    });
  },
  async update(id: string, token: string) {
    return await stripe.customers.update(id, {
      source: token,
    });
  },
  async invoices(id: string, limit: number) {
    return await stripe.invoices.list({
      customer: id,
      limit: limit,
    });
  },
  async subscribe(id: string, plan: string) {
    const subscription = await stripe.subscriptions.create({
      customer: id,
      items: [{ plan: plan }],
      //   enable_incomplete_payments: true,
      expand: ['latest_invoice.payment_intent'],
    });

    // add the price
    // subscription.price = ('settings.currencySymbol' as string) + (subscription.items.data[0].plan.amount / 100).toFixed(2);

    return subscription;
  },
  async delete(customerId: string) {
    return await stripe.customers.del(customerId);
  },
};
