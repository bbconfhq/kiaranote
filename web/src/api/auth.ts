import { api } from './instance';

export const registerUser = async (username: string, password: string) => {
  return api.post('/v1/auth/register', {
    username,
    password,
  });
};

export const signIn = async (username: string, password: string) => {
  return api.post('/v1/auth/login', {
    username,
    password,
  });
};
