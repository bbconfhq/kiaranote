import { Box, Section } from '@radix-ui/themes';
import React from 'react';
import {Outlet} from 'react-router-dom';

const MainLayout = () => {
  return (
    <Box height={'100%'}>
      <Section>
        <div>{/* tree */}</div>
      </Section>
      <Section>
        <Outlet />
      </Section>
    </Box>
  );
};

export default MainLayout;
