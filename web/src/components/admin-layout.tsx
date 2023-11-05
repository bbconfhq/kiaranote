import { Box, Section } from '@radix-ui/themes';
import React from 'react';
import {Outlet} from 'react-router-dom';

import { SideNavigation } from './side-navigation';


const AdminLayout = () => {
  return (
    <Box height={'100%'}>
      <Section>
        <SideNavigation routes={[
          {
            label: '사용자 관리',
            pages: [
              {
                title: '사용자 목록',
                slug: '/admin/users',
              },
              {
                title: '가입 대기 목록',
                slug: '/admin/users/waiting',
              },
            ]
          }
        ]} />
      </Section>
      <Section>
        <Outlet />
      </Section>
    </Box>
  );
};

export default AdminLayout;
