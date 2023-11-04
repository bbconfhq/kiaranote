import { api } from './instance';

export const getNotes = async () => {
  return api.get('/v1/note');
};
