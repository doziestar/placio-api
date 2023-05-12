import { v4 as uuidv4 } from 'uuid';
import mongoose, { Document, Schema } from 'mongoose';
import { mask } from './utility.model';

interface IKey extends Document {
  id: string;
  name?: string;
  key: string;
  scope: string[];
  date_created: Date;
  active: boolean;
  account_id: string;
  full_key: string;
}

interface KeyData {
  name?: string;
  key: string;
  scope: string[];
}

const KeySchema = new Schema<IKey>({
  id: { type: String, required: true, unique: true },
  name: { type: String },
  key: { type: String, required: true, unique: true },
  scope: [{ type: String, required: true }],
  date_created: { type: Date, required: true },
  full_key: { type: String, required: true },
  active: { type: Boolean, required: true },
  account_id: { type: String, required: true },
});

const Key = mongoose.model<IKey>('Key', KeySchema, 'key');

class KeyController {
  async create(data: KeyData, account: string) {
    const newData = {
      ...data,
      id: uuidv4(),
      active: true,
      account_id: account,
      date_created: new Date(),
    };

    const newKey = new Key(newData);
    await newKey.save();

    return {
      ...newData,
      full_key: newData.key,
      key: mask(newData.key),
    };
  }

  async get(id?: string, name?: string, account?: string) {
    const data = await Key.find({
      ...(id && { id: id }),
      ...(name && { name: name }),
      account_id: account,
    });

    return data.map(x => {
      const document = x.toObject();
      return {
        ...document,
        key: id ? document.key : mask(document.key),
      };
    });
  }

  async unique(key: string) {
    const data = await Key.find({ key: key });
    return !data.length;
  }

  async verify(key: string): Promise<IKey> {
    const data = await Key.findOne({ key: key, active: true }).select({ scope: 1, account_id: 1 });
    return data;
  }

  async update(id: string, data: Partial<IKey>, account: string) {
    return await Key.updateOne({ id: id, account_id: account }, data);
  }

  async delete(id: string, account: string) {
    return await Key.deleteOne({ id: id, account_id: account });
  }
}

export { Key, KeyController, IKey, KeyData };
