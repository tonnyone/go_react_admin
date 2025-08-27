import { UploadOutlined, UserOutlined, VideoCameraOutlined } from '@ant-design/icons';

export const siderMenuItems = [
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
      },
    ],
  },
  {
    key: '/about',
    icon: <UploadOutlined />,
    label: '关于我们',
  },
];
