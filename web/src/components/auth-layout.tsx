import { Box, Flex } from '@radix-ui/themes';
import React from 'react';
import {Outlet} from 'react-router-dom';

const AuthLayout = () => {
  return (
    <Box height={'100%'} width={'100%'}>
      <Flex height={'100%'} width={'100%'} justify={'center'} align={'center'} direction={'column'} >
        <Outlet />
      </Flex>
    </Box>
  );
};

export default AuthLayout;
