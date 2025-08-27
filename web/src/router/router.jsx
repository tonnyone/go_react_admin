import { createBrowserRouter } from 'react-router-dom';
import MyLayout from '../components/Layout.jsx';
import About from '../pages/About.jsx';
import Login from '../pages/Login.jsx';
import Menus from '../pages/Menus.jsx';
import NotFound from '../pages/NotFound.jsx';
import Resource from '../pages/Resource.jsx';
import Role from '../pages/Role.jsx';
import User from '../pages/User.jsx';


// 路由配置（path + Component直接引用）
export const routes = [
	{ path: '/', Component: Login },
	{ path: 'login', Component: Login },
	{ path: 'about', Component: About },
	{
		path: 'user',
		Component: MyLayout,
		children: [
			{ path: 'list', index: true, Component: User },
			{ path: 'role', Component: Role },
			{ path: 'menu', Component: Menus },
			{ path: 'resource', Component: Resource },
		],
	},
	{ path: '*', Component: NotFound },
];

export const router = createBrowserRouter(routes);
