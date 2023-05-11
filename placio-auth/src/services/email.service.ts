// emailService.ts
import nodemailer from 'nodemailer';
import fs from 'fs';
import path from 'path';
import handlebars from 'handlebars';
import emailContent from '../emails/content.json';

const templatePath = path.join(__dirname, '../emails/template.html');
const emailTemplate = fs.readFileSync(templatePath, 'utf8');
const compiledTemplate = handlebars.compile(emailTemplate);

const transporter = nodemailer.createTransport({
  // Configure your email transport using your preferred email provider
});

async function sendEmail(to: string, subject: string, html: string) {
  const mailOptions = {
    from: process.env.EMAIL_FROM,
    to,
    subject,
    html,
  };

  await transporter.sendMail(mailOptions);
}

export async function sendWelcomeEmail(to: string) {
  const content = emailContent['new-user'];
  const html = compiledTemplate({
    title: content.subject,
    body: content.body,
    button: content.button,
    domain: process.env.DOMAIN,
  });

  await sendEmail(to, content.subject, html);
}

export async function sendEmailVerification(to: string) {
  const content = emailContent['email-verification'];
  const html = compiledTemplate({
    title: content.subject,
    body: content.body,
    button: content.button,
    domain: process.env.DOMAIN,
  });

  await sendEmail(to, content.subject, html);
}

export async function sendPasswordReset(to: string, token: string) {
  const content = emailContent['password-reset'];
  const html = compiledTemplate({
    title: content.subject,
    body: content.body,
    button: content.button,
    domain: process.env.DOMAIN,
    content: { token },
  });

  await sendEmail(to, content.subject, html);
}
