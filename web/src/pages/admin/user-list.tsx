import {Space, Table, Tag, Typography} from 'antd';
import type {ColumnsType} from 'antd/es/table';
import React from 'react';

import {UserRole} from '../../constants/user';


type DataType = {
  id: number;
  username: string;
  role: UserRole;
};

const columns: ColumnsType<DataType> = [
  {
    title: 'ID',
    dataIndex: 'id',
  },
  {
    title: 'Username',
    dataIndex: 'username',
  },
  {
    title: 'Role',
    dataIndex: 'role',
    render: (_, { role }) => (
      <Tag color={role === UserRole.ADMIN ? 'volcano' : 'geekblue'}>
        {role.toUpperCase()}
      </Tag>
    ),
  },
  {
    title: 'Action',
    key: 'action',
    render: (_, record) => (
      <Space size='middle'>
        <a href={`#edit${record.id}`}>Edit</a>
      </Space>
    ),
  },
];

const data: DataType[] = [];
for (let i = 0; i < 46; i++) {
  data.push({
    id: i,
    username: `User ${i}`,
    role: i % 2 ? UserRole.USER : UserRole.ADMIN,
  });
}

const { Title } = Typography;

const AdminUserListPage: React.FC = () => {

  return (
    <Typography>
      <Title level={2}>사용자 목록 </Title>
      <Table columns={columns} dataSource={data} />
    </Typography>
  );
};

export default AdminUserListPage;
