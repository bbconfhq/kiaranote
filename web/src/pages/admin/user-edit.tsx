import {
  Button,
  Form,
  Input,
  Select, Space,
} from 'antd';
import React from 'react';
import {Navigate, useParams} from 'react-router-dom';

import {UserRole} from '../../constants/user';

const { Option } = Select;

const AdminUserEditPage: React.FC = () => {
  const { id } = useParams();
  const [form] = Form.useForm();

  if (id == null) {
    alert(`존재하지 않는 사용자입니다. (id: ${id})`);
    return <Navigate to={'/admin/users'} replace />;
  }
  const onFinish = (values: any) => {
    console.log('Received values of form: ', values);
  };

  return (
    <Form
      form={form}
      name='register'
      onFinish={onFinish}
      scrollToFirstError
    >
      <Form.Item
        label='ID'
      >
        {id}
      </Form.Item>

      <Form.Item
        name='username'
        label='Username'
        rules={[{ required: true, message: 'Please input username', whitespace: true }]}
      >
        <Input />
      </Form.Item>

      <Form.Item
        name='role'
        label='Role'
        rules={[{ required: true, message: 'Please select role' }]}
      >
        <Select placeholder='select your gender'>
          <Option value={UserRole.USER}>사용자</Option>
          <Option value={UserRole.ADMIN}>관리자</Option>
        </Select>
      </Form.Item>

      <Form.Item>
        <Space size={'small'}>
          <Button type='primary' htmlType='submit'>
          저장
          </Button>
          <Button danger htmlType='button'>
          삭제
          </Button>
        </Space>
      </Form.Item>
    </Form>
  );
};

export default AdminUserEditPage;
