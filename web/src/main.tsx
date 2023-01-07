import React from 'react';
import ReactDOM from 'react-dom/client';
import {
  BrowserRouter, Navigate,
  Route,
  Routes,
} from 'react-router-dom';

import './index.css';
import AdminLayout from './components/admin-layout';
import AuthLayout from './components/auth-layout';
import AdminUserEditPage from './pages/admin/user-edit';
import AdminUserListPage from './pages/admin/user-list';
import AdminUserWaitingListPage from './pages/admin/user-waiting';
import RegisterPage from './pages/register';
import SignInPage from './pages/sign-in';

// <RouterProvider router={router} />

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <BrowserRouter>
      <Routes>
        <Route element={<AdminLayout />}>
          <Route path={'/admin'} element={<Navigate to={'/admin/users'} replace />} />
          <Route path={'/admin/users'} element={<AdminUserListPage />} />
          <Route path={'/admin/users/waiting'} element={<AdminUserWaitingListPage />} />
          <Route path={'/admin/users/:id'} element={<AdminUserEditPage />} />
        </Route>
        <Route element={<AuthLayout />}>
          <Route path={'/sign-in'} element={<SignInPage />} />
          <Route path={'/register'} element={<RegisterPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  </React.StrictMode>
);
