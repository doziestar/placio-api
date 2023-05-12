import { promises as fs } from 'fs';
import chalk from 'chalk';
import axios from 'axios';
import FormData from 'form-data';
// import config from '../config';
import emails from '../emails/content.json';

interface EmailData {
  to: string;
  content: Record<string, string>;
  custom?: string;
  template: string;
  subject?: string;
}

interface EmailContent {
  subject?: string;
  title?: string;
  body: string;
  button: {
    url: string;
    label: string;
  };
}

// const domain: string = config.get('domain');
const domain = 'placio.io';
const settings: Record<string, string> = {
  domain: 'mail.hubhub.app',
  host: 'api.eu.mailgun.net',
  sender: 'johannes@mail.hubhub.app',
  base_url: 'https://api.eu.mailgun.net/v3',
};

export async function sendMail(data: EmailData): Promise<void> {
  console.log(chalk.blue('Sending email to:'), data.to);
  const rex =
    /^(?:[a-z0-9!#$%&amp;'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&amp;'*+/=?^_`{|}~-]+)*|'(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*')@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])$/;

  // if (rex.test(data.to)) {
  //   const content: EmailContent = emails[data.template];
  //   const html: string = await createEmail(data.custom || 'template', content, data.content);

  //   const form = new FormData();
  //   form.append('to', data.to);
  //   form.append('from', settings.sender);
  //   form.append('subject', content?.subject || data?.subject);
  //   form.append('html', html);

  //   // await axios({
  //   //   method: 'POST',
  //   //   url: `${settings.base_url}/${settings.domain}/messages`,
  //   //   headers: { 'Content-Type': `multipart/form-data; boundary=${form.getBoundary()}` },
  //   //   data: form,
  //   //   auth: {
  //   //     username: 'api',
  //   //     password: process.env.MAILGUN_API_KEY!,
  //   //   },
  //   // });

  //   console.log(chalk.green('Email sent to: ') + data.to);
  // } else {
  //   throw { message: 'Invalid email address' };
  // }
}

async function createEmail(template: string, content: EmailContent, values: Record<string, string>): Promise<string> {
  let email: string = await fs.readFile(`emails/${template}.html`, 'utf8');
  email = email.replace(/{{domain}}/g, values?.domain || domain);

  if (content) {
    content.title = content.title || content.subject;

    if (content.button.url?.includes('{{domain}}')) content.button.url = content.button.url.replace(/{{domain}}/g, values?.domain || domain);

    const lines = content.body.split('\n');

    if (content.title) lines.unshift(`Hi ${content.title},`);

    content.body = lines
      .map(
        line =>
          `<p style="color: #7e8890; font-family: 'Source Sans Pro', helvetica, sans-serif; font-size: 15px; font-weight: normal; Margin: 0; Margin-bottom: 15px; line-height: 1.6;">${line}</p>`,
      )
      .join('\n');

    email = email.replace(/{{title}}/g, content.title);
    email = email.replace('{{body}}', content.body);
    email = email.replace('{{buttonURL}}', content.button.url);
    email = email.replace('{{buttonLabel}}', content.button.label);

    if (values) {
      for (const key in values) {
        const rex = new RegExp(`{{content.${key}}}`, 'g');
        email = email.replace(rex, values[key]);
      }
    }
  }

  return email;
}
