import {
  UserOutlined,
} from '@ant-design/icons';
import {Layout, Menu, MenuProps} from 'antd';
import React, { useEffect } from 'react';
import {Link, Outlet} from 'react-router-dom';

import { getNotes } from '../api/note';

type MenuItem = Required<MenuProps>['items'][number];

const getItem = (
  label: React.ReactNode,
  key: React.Key,
  icon?: React.ReactNode,
  children?: MenuItem[],
): MenuItem => ({
  key,
  icon,
  children,
  label,
} as MenuItem);

const items: MenuItem[] = [
  getItem('사용자 관리', 'user-mgmt', <UserOutlined />, [
    getItem(<Link to={'/admin/users'}>사용자 목록</Link>, 'user-list'),
    getItem(<Link to={'/admin/users/waiting'}>가입 대기 목록</Link>, 'user-waiting-list'),
  ]),
];
const AdminLayout = () => {
  const { Content, Sider } = Layout;
  useEffect(() => {
    getNotes().then((res) => console.log(res));
  }, []);

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Sider theme={'light'} collapsed={false}>
        <Menu style={{padding: '8px 4px'}} theme='light' defaultSelectedKeys={['user-list']} mode='inline' openKeys={['user-mgmt']} items={items} />
      </Sider>
      <Layout className='site-layout'>
        <Content style={{ margin: '0 16px' }}>
          <div style={{ paddingTop: 24, paddingBottom: 24, minHeight: 360 }}>
            <Outlet />
          </div>
        </Content>
      </Layout>
    </Layout>
  );
};

export default AdminLayout;
