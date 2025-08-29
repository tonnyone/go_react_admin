import {
  MenuFoldOutlined,
  MenuUnfoldOutlined
} from '@ant-design/icons';
import { Button, Dropdown, Layout, Menu, theme } from 'antd';
import { useState } from 'react';
import { Outlet, useNavigate } from 'react-router-dom';
import logo from '../assets/react.svg';
import { siderMenuItems } from '../router/siderMenuItems.jsx';
import { logout } from '../service/user.js';

const MyLayout = () => {
const [collapsed, setCollapsed] = useState(false);
const { token: { colorBgContainer, borderRadiusLG }, } = theme.useToken();
const navigate = useNavigate();
const { Header, Sider, Content } = Layout;
const items = [
  {
    label: (
      <a> 个人信息 </a>
    ),
  },
  { type: 'divider'},
  { 
    label: (
      <a onClick={() => {
        logout().then(() => {
          navigate('/login');
        });
      }}> 登出 </a> 
    ),
  },
];
  return (
    <Layout style={{ width: '100vw', height: '100vh' }}>
      <Sider trigger={null} collapsible collapsed={collapsed}>
        <div className="demo-logo-vertical" >
            <img src={logo}></img>
        </div>
        <Menu
          theme="dark"
          mode="inline"
          defaultSelectedKeys={['1']}
          onClick={({key}) => {
            navigate(key)
          }}
          items={siderMenuItems}
        />
      </Sider>
      <Layout>
        <Header style={{ padding: 0, background: colorBgContainer }}>
          <Button
            type="text"
            icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
            onClick={() => setCollapsed(!collapsed)}
            style={{
              fontSize: '16px',
              width: 64,
              height: 64,
            }}
          />
          <span> RBAC Demo </span>
          <Dropdown menu={{ items}}>
            <img src={logo} alt="logo" style={{ cursor: 'pointer', float: 'right', margin: '20px' }} />
          </Dropdown>
        </Header>
        <Content
          style={{
            margin: '20px 16px',
            padding: 20,
            minHeight: 280,
            background: colorBgContainer,
            borderRadius: borderRadiusLG,
          }}
        >
          <Outlet />
        </Content>
      </Layout>
    </Layout>
  );
};
export default MyLayout;