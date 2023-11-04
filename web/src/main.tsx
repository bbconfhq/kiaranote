import { Theme } from '@radix-ui/themes';
import React from 'react';
import ReactDOM from 'react-dom/client';
import {
  BrowserRouter, Navigate,
  Route,
  Routes,
} from 'react-router-dom';

import './index.css';
import '@radix-ui/themes/styles.css';
import AdminLayout from './components/admin-layout';
import AuthLayout from './components/auth-layout';
import RegisterPage from './pages/register';
import SignInPage from './pages/sign-in';


ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <Theme appearance='light' accentColor='blue' grayColor='sand' id={'theme-root'}>
      <BrowserRouter>
        <Routes>
          <Route element={<AdminLayout />}>
            <Route path={'/admin'} element={<Navigate to={'/admin/users'} replace />} />
          </Route>
          <Route element={<AuthLayout />}>
            <Route path={'/sign-in'} element={<SignInPage />} />
            <Route path={'/register'} element={<RegisterPage />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </Theme>
  </React.StrictMode>
);
