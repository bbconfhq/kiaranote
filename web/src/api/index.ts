import axios from 'axios';

const request = axios.create({
  baseURL: 'http://localhost:8080',
  withCredentials: true,
});

export const api = {
  registerUser: async (username: string, password: string) => {
    return request.post('/v1/auth/register', {
      username,
      password,
    });
  },
  signIn: async (username: string, password: string) => {
    return request.post('/v1/auth/login', {
      username,
      password,
    });
  },
  getNotes: async () => {
    return request.get('/v1/note');
  }
};
