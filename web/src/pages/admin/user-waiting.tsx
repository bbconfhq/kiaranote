import {Space, Table, Dropdown, Typography, Button} from 'antd';
import type {ColumnsType} from 'antd/es/table';
import React from 'react';

import {UserRole} from '../../constants/user';


type DataType = {
  id: number;
  username: string;

  createdAt: string;
  role: UserRole;
};

const items = [
  {
    key: 'approval-admin',
    label: '관리자로 승인',
  },
];

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
    title: '가입 신청일',
    dataIndex: 'createdAt',
  },
  {
    title: 'Action',
    key: 'action',
    render: (_, { id }) => (
      <Space size='middle'>
        <Dropdown.Button type={'primary'} onClick={(e) => console.log(id, e)} menu={{ items, onClick: (e) => { console.log(id, e); } }}>승인</Dropdown.Button>
        <Button danger>거절</Button>
      </Space>
    ),
  },
];

const data: DataType[] = [];
for (let i = 0; i < 46; i++) {
  data.push({
    id: i,
    username: `User ${i}`,
    createdAt: Intl.DateTimeFormat().format(new Date()),
    role: i % 2 ? UserRole.USER : UserRole.ADMIN,
  });
}

const { Title } = Typography;

const AdminUserWaitingListPage: React.FC = () => {

  return (
    <Typography>
      <Title level={2}>가입 대기 목록</Title>
      <Table columns={columns} dataSource={data} />
    </Typography>
  );
};

export default AdminUserWaitingListPage;
