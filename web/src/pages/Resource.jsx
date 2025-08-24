import { Button, Form, Input, Modal, Popconfirm, Select, Space, Table, TreeSelect } from 'antd';
import { useState } from 'react';

// 示例数据
const initialResources = [
  {
    id: '1',
    name: '用户资源',
    path: '/api/user',
    method: 'GET',
    children: [
      {
        id: '1-1',
        name: '用户详情',
        path: '/api/user/detail',
        method: 'GET',
      },
      {
        id: '1-2',
        name: '用户创建',
        path: '/api/user/create',
        method: 'POST',
      },
    ],
  },
  {
    id: '2',
    name: '角色资源',
    path: '/api/role',
    method: 'GET',
    children: [
      {
        id: '2-1',
        name: '角色分配',
        path: '/api/role/assign',
        method: 'POST',
      },
    ],
  },
];

const ResourceManage = () => {
  const [resources, setResources] = useState(initialResources);
  const [expandedRowKeys, setExpandedRowKeys] = useState([]);
  const [editing, setEditing] = useState(null); // {record, visible}
  const [form] = Form.useForm();
  const [searchForm] = Form.useForm();
  const [pagination, setPagination] = useState({ current: 1, pageSize: 10 });
  const [filtered, setFiltered] = useState(initialResources);

  // 递归渲染树表格数据
  const renderResources = (data) =>
    data.map(item => ({
      ...item,
      key: item.id,
      children: item.children ? renderResources(item.children) : undefined,
    }));

  // 递归过滤树结构
  const filterTree = (data, filterFn) => {
    return data
      .map(item => {
        if (item.children) {
          const children = filterTree(item.children, filterFn);
          if (children.length > 0 || filterFn(item)) {
            return { ...item, children };
          }
        } else if (filterFn(item)) {
          return { ...item };
        }
        return null;
      })
      .filter(Boolean);
  };

  // 新增/编辑弹窗提交
  const handleOk = () => {
    form.validateFields().then(values => {
      if (editing && editing.record) {
        // 修改
        const update = (list) => list.map(i => {
          if (i.id === editing.record.id) return { ...i, ...values };
          if (i.children) return { ...i, children: update(i.children) };
          return i;
        });
        setResources(update(resources));
      } else {
        // 新增
        const addTo = (list, parentId) => {
          if (!parentId) return [...list, { ...values, id: Date.now().toString() }];
          return list.map(i => {
            if (i.id === parentId) {
              return {
                ...i,
                children: i.children ? [...i.children, { ...values, id: Date.now().toString() }] : [{ ...values, id: Date.now().toString() }],
              };
            }
            if (i.children) return { ...i, children: addTo(i.children, parentId) };
            return i;
          });
        };
        setResources(addTo(resources, values.parentId));
      }
      setEditing(null);
      form.resetFields();
    });
  };

  // 删除节点
  const handleDelete = (id) => {
    const remove = (list) => list.filter(i => {
      if (i.id === id) return false;
      if (i.children) i.children = remove(i.children);
      return true;
    });
    setResources(remove(resources));
  };

  // 表格列
  const columns = [
    { title: '名称', dataIndex: 'name', key: 'name' },
    { title: '请求路径', dataIndex: 'path', key: 'path' },
    { title: '请求方法', dataIndex: 'method', key: 'method', width: 100 },
    {
      title: '操作',
      key: 'action',
      render: (_, record) => (
        <Space>
          <a onClick={() => { setEditing({ record, visible: true }); form.setFieldsValue({ ...record, parentId: undefined }); }}>修改</a>
          <Popconfirm title="确定删除该资源?" onConfirm={() => handleDelete(record.id)}>
            <a style={{ color: 'red' }}>删除</a>
          </Popconfirm>
          <a onClick={() => { setEditing({ record: null, visible: true }); form.setFieldsValue({ parentId: record.id }); }}>新增子资源</a>
        </Space>
      ),
    },
  ];

  // 查询功能
  const handleSearch = () => {
    const values = searchForm.getFieldsValue();
    const filterFn = (item) => {
      return (
        (!values.name || item.name.includes(values.name)) &&
        (!values.path || item.path.includes(values.path)) &&
        (!values.method || item.method === values.method)
      );
    };
    setFiltered(filterTree(resources, filterFn));
    setPagination({ ...pagination, current: 1 });
  };

  // 分页数据
  const getPagedData = (data) => {
    // 展开所有节点，分页只对顶层节点做
    const start = (pagination.current - 1) * pagination.pageSize;
    const end = start + pagination.pageSize;
    return data.slice(start, end);
  };

  return (
    <div style={{ padding: 24 }}>
      <Form
        form={searchForm}
        layout="inline"
        style={{ marginBottom: 16 }}
        onFinish={handleSearch}
      >
        <Form.Item name="name" label="名称">
          <Input placeholder="资源名称" allowClear />
        </Form.Item>
        <Form.Item name="path" label="请求路径">
          <Input placeholder="请求路径" allowClear />
        </Form.Item>
        <Form.Item name="method" label="请求方法">
          <Select allowClear style={{ width: 120 }} options={['GET','POST','PUT','DELETE','PATCH'].map(m=>({label:m,value:m}))} />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit">查询</Button>
        </Form.Item>
        <Form.Item>
          <Button onClick={() => { searchForm.resetFields(); setFiltered(resources); }}>重置</Button>
        </Form.Item>
      </Form>
      <Space style={{ marginBottom: 16 }}>
        <Button type="primary" onClick={() => { setEditing({ record: null, visible: true }); form.resetFields(); }}>新增资源</Button>
      </Space>
      <Table
        rowKey="id"
        columns={columns}
        dataSource={getPagedData(renderResources(filtered.length ? filtered : resources))}
        pagination={{
          current: pagination.current,
          pageSize: pagination.pageSize,
          total: (filtered.length ? filtered : resources).length,
          showSizeChanger: true,
          onChange: (page, pageSize) => setPagination({ current: page, pageSize }),
        }}
        bordered
        expandable={{
          expandedRowKeys,
          onExpand: (expanded, record) => {
            setExpandedRowKeys(expanded ? [...expandedRowKeys, record.id] : expandedRowKeys.filter(k => k !== record.id));
          },
        }}
      />
      <Modal
        title={editing && editing.record ? '修改资源' : '新增资源'}
        open={!!editing}
        onCancel={() => { setEditing(null); form.resetFields(); }}
        onOk={handleOk}
        destroyOnClose
      >
        <Form form={form} layout="vertical">
          <Form.Item label="名称" name="name" rules={[{ required: true, message: '请输入名称' }]}> 
            <Input />
          </Form.Item>
          <Form.Item label="请求路径" name="path" rules={[{ required: true, message: '请输入请求路径' }]}> 
            <Input />
          </Form.Item>
          <Form.Item label="请求方法" name="method" rules={[{ required: true, message: '请选择请求方法' }]}> 
            <Select options={['GET','POST','PUT','DELETE','PATCH'].map(m=>({label:m,value:m}))} />
          </Form.Item>
          <Form.Item label="父资源" name="parentId">
            <TreeSelect
              allowClear
              treeData={renderResources(resources)}
              fieldNames={{ label: 'name', value: 'id', children: 'children' }}
              placeholder="不选为顶级资源"
              treeDefaultExpandAll
            />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
};

export default ResourceManage;
