import { Document } from 'mongoose';

interface IToken extends Document {
  id: string;
  provider: string;
  jwt?: string;
  access?: string;
  refresh?: string;
  user_id: string;
}

type TokenType = 'Basic' | 'Bearer' | 'Refresh';

export { IToken, TokenType };
