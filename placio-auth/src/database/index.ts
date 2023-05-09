import { DB_PASSWORD, DB_USERNAME, DB_DATABASE } from '@config';

export const dbConnection = {
  url: `mongodb+srv://${DB_USERNAME}:${DB_PASSWORD}@cluster0.jmf27wp.mongodb.net/${DB_DATABASE}/?retryWrites=true&w=majority`,
  options: {
    useNewUrlParser: true,
    useUnifiedTopology: true,
  },
};
