import {UserOutlined, LockOutlined} from '@ant-design/icons';
import {Button, Divider, Form, Input, Space, Typography} from 'antd';
import React from 'react';
import {Link} from 'react-router-dom';

const RegisterPage = () => {
  const [form] = Form.useForm();
  const { Text, Link: AntdLink } = Typography;

  const onFinish = (values: any) => {
    console.log('Success:', values);
  };

  const onFinishFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
  };

  return (
    <Form
      form={form}
      name='login'
      onFinish={onFinish}
      onFinishFailed={onFinishFailed}
      autoComplete='off'
    >
      <Form.Item
        name={'username'}
        rules={[
          {
            required: true,
            message: 'Please Input your Username'
          }
        ]}
      >
        <Input
          prefix={<UserOutlined className={'site-form-item-icon'} />}
          placeholder={'Username'}
        />
      </Form.Item>

      <Form.Item
        name='password'
        rules={[
          {
            required: true,
            message: 'Please input your password'
          }
        ]}
      >
        <Input
          prefix={<LockOutlined className={'site-form-item-icon'} />}
          type={'password'}
          placeholder={'Password'}
        />
      </Form.Item>

      <Form.Item
        name='confirm'
        dependencies={['password']}
        hasFeedback
        rules={[
          {
            required: true,
            message: 'Please confirm your password',
          },
          ({ getFieldValue }) => ({
            validator(_, value) {
              if (!value || getFieldValue('password') === value) {
                return Promise.resolve();
              }
              return Promise.reject(new Error('The two passwords that you entered do not match'));
            },
          }),
        ]}
      >
        <Input
          prefix={<LockOutlined className={'site-form-item-icon'} />}
          type={'password'}
          placeholder={'Password Confirm'}
        />
      </Form.Item>

      <Form.Item>
        <Button
          type={'primary'}
          htmlType={'submit'}
          className={'login-form-button'}
          style={{ marginTop: '1.5rem' }}
          block
        >
          Register
        </Button>
      </Form.Item>

      <Divider/>
      <Space style={{width: '100%', justifyContent: 'center'}} size={'small'}>
        <Text>Have an account?</Text>
        <Link to={'/sign-in'}>
          <AntdLink>Sign In</AntdLink>
        </Link>
      </Space>
    </Form>
  );
};

export default RegisterPage;
