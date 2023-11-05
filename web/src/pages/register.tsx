import * as Form from '@radix-ui/react-form';
import { FormField, FormLabel, FormMessage } from '@radix-ui/react-form';
import {
  Box,
  Card,
  Flex,
  Heading,
  Text,
  TextField,
  Button,
} from '@radix-ui/themes';
import React from 'react';
import { useNavigate } from 'react-router-dom';

import { api } from '../api';

const RegisterPage = () => {
  const navigate = useNavigate();

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const data = new FormData(e.currentTarget);
    const username = data.get('username') as string;
    const password = data.get('password') as string;
    try {
      await api.registerUser(username, password);
      navigate('/sign-in');
    } catch (err) {
      alert('failed to register');
      console.error(err);
    }
  };

  return (
    <Flex align={'center'} justify={'center'} height={'100%'} width={'100%'}>
      <Flex shrink='0' gap='6' direction='column' style={{ width: 416 }}>
        <Card size='4'>
          <Form.Root onSubmit={onSubmit}>
            <Heading as='h3' size='6' trim='start' mb='5'>
              Sign up
            </Heading>

            <FormField name={'username'}>
              <Box mb='5'>
                <FormLabel>
                  <Text as='div' size='2' weight='bold' mb='2'>
                    Username
                  </Text>
                </FormLabel>
                <Form.Control asChild>
                  <TextField.Input placeholder='Enter your username' required />
                </Form.Control>
                <FormMessage match={'valueMissing'}>
                  <Text size={'2'} color={'crimson'}>
                    Username is required
                  </Text>
                </FormMessage>
                <FormMessage match={'badInput'}>
                  <Text size={'2'} color={'crimson'}>
                    Invalid username
                  </Text>
                </FormMessage>
              </Box>
            </FormField>

            <FormField name={'password'}>
              <Box mb='5'>
                <FormLabel>
                  <Text as='div' size='2' weight='bold' mb='2'>
                    Password
                  </Text>
                </FormLabel>
                <Form.Control asChild>
                  <TextField.Input
                    type={'password'}
                    placeholder='Enter your password'
                    required
                  />
                </Form.Control>
                <FormMessage match={'valueMissing'}>
                  <Text size={'2'} color={'crimson'}>
                  Password is required
                  </Text>
                </FormMessage>
              </Box>
            </FormField>

            <FormField name={'password-confirm'}>
              <Box mb='5'>
                <FormLabel>
                  <Text as='div' size='2' weight='bold' mb='2'>
                    Password Confirm
                  </Text>
                </FormLabel>
                <Form.Control asChild>
                  <TextField.Input
                    type={'password'}
                    placeholder='Enter your password again'
                    required
                  />
                </Form.Control>
                <FormMessage match={((value, formData) => value.trim() != formData.get('password'))}>
                  <Text size={'2'} color={'crimson'}>
                    Password does not match
                  </Text>
                </FormMessage>
              </Box>
            </FormField>

            <Flex mt='6' justify='end' gap='3'>
              <Button>Create an account</Button>
            </Flex>
          </Form.Root>
        </Card>
      </Flex>
    </Flex>
  );
};

export default RegisterPage;
