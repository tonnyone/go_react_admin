import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  UploadOutlined,
  UserOutlined,
  VideoCameraOutlined
} from '@ant-design/icons';
import { Button, Dropdown, Layout, Menu, theme } from 'antd';
import { useState } from 'react';
import { Outlet, useNavigate } from 'react-router-dom';
import logo from '../assets/react.svg';
const { Header, Sider, Content } = Layout;
const items = [
  {
    key: '/user/profile',
    label: (
      <a> 个人信息 </a>
    ),
  },
  { type: 'divider'},
  { key: '/login', label: '登出' },
];

const MyLayout = () => {
const [collapsed, setCollapsed] = useState(false);
const { token: { colorBgContainer, borderRadiusLG }, } = theme.useToken();
const navigate = useNavigate();
const DropdownClick = ({key}) => {
  console.log(key)
  navigate(key)
}
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
          items={[
            {
              key: '/user',
              icon: <UserOutlined />,
              label: '用户权限管理',
              children: [
                {
                    key: '/user/list',
                    icon: <VideoCameraOutlined />,
                    label: '用户管理',
                },
                {
                    key: '/user/role',
                    icon: <VideoCameraOutlined />,
                    label: '角色管理',
                },
                {
                    key: '/user/menu',
                    icon: <UploadOutlined />,
                    label: '菜单管理',
                },
                {
                    key: '/user/resource',
                    icon: <VideoCameraOutlined />,
                    label: '资源管理',
                }
            ]},
            {
                key: '/about',
                icon: <UploadOutlined />,
                label: '关于我们',
            }
          ]}
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
          <Dropdown menu={{ items, onClick: DropdownClick}}>
            <img src={logo} alt="logo" style={{cursor: 'pointer',float: 'right', margin: '20px'}} />
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