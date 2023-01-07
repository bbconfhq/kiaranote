import {Layout, Typography} from 'antd';
import React from 'react';
import {Outlet} from 'react-router-dom';

const AuthLayout = () => {
  const { Content } = Layout;
  const { Title } = Typography;

  const ContentStyle: React.CSSProperties = {
    height: '100vh',
    width: '100%',
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    alignSelf: 'center',
    flexDirection: 'column',
    minWidth: '320px',
    maxWidth: '400px',
  };

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Content style={{...ContentStyle}}>
        <Title level={1}>KiaraNote</Title>

        <div style={{ paddingTop: 24, paddingBottom: 24, minHeight: 360, width: '100%', }}>
          <Outlet />
        </div>
      </Content>
    </Layout>
  );
};

export default AuthLayout;
