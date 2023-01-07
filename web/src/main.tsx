import {
  RouterProvider,
  createReactRouter,
  createRouteConfig,
} from '@tanstack/react-router';
import React from 'react';
import ReactDOM from 'react-dom/client';

import App from './App';
import './index.css';

const rootRoute = createRouteConfig();
const indexRoute = rootRoute.createRoute({
  path: '/',
  component: App,
});
const routeConfig = rootRoute.addChildren([indexRoute]);
const router = createReactRouter({ routeConfig });

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);

declare module '@tanstack/react-router' {
  interface RegisterRouter {
    router: typeof router
  }
}
