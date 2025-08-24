
import { Button, Form, Input, Modal, Popconfirm, Space, Table, TreeSelect } from 'antd';
import { useState } from 'react';

// 多级嵌套菜单示例数据，含菜单和按钮类型
const initialMenus = [
	{
		key: '1',
		title: '系统管理',
		path: '/system',
		order: 1,
		type: 'menu',
		children: [
			{
				key: '1-1',
				title: '用户管理',
				path: '/system/user',
				order: 1,
				type: 'menu',
				children: [
					{
						key: '1-1-1',
						title: '用户详情',
						path: '/system/user/detail',
						order: 1,
						type: 'menu',
					},
					{
						key: '1-1-2',
						title: '用户日志',
						path: '/system/user/log',
						order: 2,
						type: 'menu',
						children: [
							{
								key: '1-1-2-1',
								title: '登录日志',
								path: '/system/user/log/login',
								order: 1,
								type: 'menu',
							},
							{
								key: '1-1-2-2',
								title: '操作日志',
								path: '/system/user/log/action',
								order: 2,
								type: 'menu',
								children: [
									{
										key: '1-1-2-2-1',
										title: '删除按钮',
										path: '',
										order: 1,
										type: 'button',
									},
								],
							},
						],
					},
				],
			},
			{
				key: '1-2',
				title: '角色管理',
				path: '/system/role',
				order: 2,
				type: 'menu',
				children: [
					{
						key: '1-2-1',
						title: '新增按钮',
						path: '',
						order: 1,
						type: 'button',
					},
				],
			},
		],
	},
	{
		key: '2',
		title: '菜单管理',
		path: '/menus',
		order: 2,
		type: 'menu',
		children: [
			{
				key: '2-1',
				title: '菜单设置',
				path: '/menus/setting',
				order: 1,
				type: 'menu',
			},
		],
	},
];

const Menus = () => {
	const [menus, setMenus] = useState(initialMenus);
	const [expandedRowKeys, setExpandedRowKeys] = useState([]);
	const [editing, setEditing] = useState(null); // {record, visible, isButton}
	const [form] = Form.useForm();

	// 递归渲染树表格数据
	const renderMenus = (data) =>
		data.map(item => ({
			...item,
			children: item.children ? renderMenus(item.children) : undefined,
		}));

	// 新增/编辑弹窗提交
	const handleOk = () => {
		form.validateFields().then(values => {
			if (editing && editing.record) {
				// 修改
				const update = (list) => list.map(i => {
					if (i.key === editing.record.key) return { ...i, ...values };
					if (i.children) return { ...i, children: update(i.children) };
					return i;
				});
				setMenus(update(menus));
			} else {
				// 新增
				const addTo = (list, parentKey) => {
					if (!parentKey) return [...list, { ...values, key: Date.now().toString() }];
					return list.map(i => {
						if (i.key === parentKey) {
							return {
								...i,
								children: i.children ? [...i.children, { ...values, key: Date.now().toString() }] : [{ ...values, key: Date.now().toString() }],
							};
						}
						if (i.children) return { ...i, children: addTo(i.children, parentKey) };
						return i;
					});
				};
				setMenus(addTo(menus, values.parentKey));
			}
			setEditing(null);
			form.resetFields();
		});
	};

	// 删除节点
	const handleDelete = (key) => {
		const remove = (list) => list.filter(i => {
			if (i.key === key) return false;
			if (i.children) i.children = remove(i.children);
			return true;
		});
		setMenus(remove(menus));
	};

	// 上移/下移操作
	const moveMenu = (key, direction) => {
		// direction: 'up' | 'down'
		const deepMove = (list) => {
			let changed = false;
			let newList = [...list];
			newList.sort((a, b) => (a.order || 0) - (b.order || 0));
			for (let i = 0; i < newList.length; i++) {
				if (newList[i].key === key) {
					if (direction === 'up' && i > 0) {
						const temp = newList[i - 1].order;
						newList[i - 1].order = newList[i].order;
						newList[i].order = temp;
						changed = true;
						break;
					}
					if (direction === 'down' && i < newList.length - 1) {
						const temp = newList[i + 1].order;
						newList[i + 1].order = newList[i].order;
						newList[i].order = temp;
						changed = true;
						break;
					}
				}
				if (newList[i].children) {
					const res = deepMove(newList[i].children);
					if (res.changed) {
						newList[i].children = res.newList;
						changed = true;
						break;
					}
				}
			}
			return { newList, changed };
		};
		const { newList } = deepMove(menus);
		setMenus(newList);
	};

	// 表格列
		const columns = [
			{ title: '名称', dataIndex: 'title', key: 'title' },
			{ title: '类型', dataIndex: 'type', key: 'type', width: 70, render: t => t === 'button' ? '按钮' : '菜单' },
			{ title: '路径', dataIndex: 'path', key: 'path' },
			{ title: '顺序', dataIndex: 'order', key: 'order', width: 70 },
			{
				title: '操作',
				key: 'action',
				render: (_, record) => (
					<Space>
						<a onClick={() => moveMenu(record.key, 'up')}>上移</a>
						<a onClick={() => moveMenu(record.key, 'down')}>下移</a>
						<a onClick={() => { setEditing({ record, visible: true }); form.setFieldsValue({ ...record, parentKey: undefined }); }}>修改</a>
						<Popconfirm title="确定删除该项?" onConfirm={() => handleDelete(record.key)}>
							<a style={{ color: 'red' }}>删除</a>
						</Popconfirm>
						{record.type !== 'button' && (
							<a onClick={() => { setEditing({ record: null, visible: true, isButton: false }); form.setFieldsValue({ parentKey: record.key, type: 'menu' }); }}>新增子菜单</a>
						)}
						{record.type !== 'button' && (
							<a onClick={() => { setEditing({ record: null, visible: true, isButton: true }); form.setFieldsValue({ parentKey: record.key, type: 'button' }); }}>新增按钮</a>
						)}
					</Space>
				),
			},
		];

	return (
		<div style={{ padding: 24 }}>
			<Space style={{ marginBottom: 16 }}>
				<Button type="primary" onClick={() => { setEditing({ record: null, visible: true }); form.resetFields(); }}>新增菜单</Button>
			</Space>
			<Table
				rowKey="key"
				columns={columns}
				dataSource={renderMenus(menus)}
				pagination={false}
				bordered
				expandable={{
					expandedRowKeys,
					onExpand: (expanded, record) => {
						setExpandedRowKeys(expanded ? [...expandedRowKeys, record.key] : expandedRowKeys.filter(k => k !== record.key));
					},
				}}
			/>
			<Modal
				title={editing && editing.record ? '修改' : (editing && editing.isButton ? '新增按钮' : '新增菜单')}
				open={!!editing}
				onCancel={() => { setEditing(null); form.resetFields(); }}
				onOk={handleOk}
				destroyOnClose
			>
				<Form form={form} layout="vertical">
					<Form.Item label={editing && editing.isButton ? '按钮名称' : '菜单名称'} name="title" rules={[{ required: true, message: '请输入名称' }]}> 
						<Input />
					</Form.Item>
					{(!editing || !editing.isButton) && (
						<Form.Item label="路径" name="path" rules={[{ required: true, message: '请输入路径' }]}> 
							<Input />
						</Form.Item>
					)}
					<Form.Item label="父菜单" name="parentKey">
						<TreeSelect
							allowClear
							treeData={renderMenus(menus).filter(i => i.type !== 'button')}
							fieldNames={{ label: 'title', value: 'key', children: 'children' }}
							placeholder="不选为顶级菜单"
							treeDefaultExpandAll
						/>
					</Form.Item>
					<Form.Item name="type" initialValue={editing && editing.isButton ? 'button' : 'menu'} hidden>
						<Input />
					</Form.Item>
				</Form>
			</Modal>
		</div>
	);
};

export default Menus;
