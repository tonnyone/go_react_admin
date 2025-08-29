import { Button, Checkbox, Col, Form, Input, Modal, Row, Space, Table, Tag } from 'antd';
import { useState } from 'react';
const columns = [
  {
    title: '用户名',
    dataIndex: 'name',
    key: 'name',
    render: text => <a>{text}</a>,
  },
  {
    title: '手机号',
    dataIndex: 'age',
    key: 'age',
  },
  {
    title: '邮箱',
    dataIndex: 'email',
    key: 'email',
  },
  {
    title: '部门',
    dataIndex: 'department',
    key: 'department',
  },
  {
    title: '角色',
    key: 'tags',
    dataIndex: 'tags',
    render: (_, { tags }) => (
      <>
        {tags.map(tag => {
          let color = tag.length > 5 ? 'geekblue' : 'green';
          if (tag === 'loser') {
            color = 'volcano';
          }
          return (
            <Tag color={color} key={tag}>
              {tag.toUpperCase()}
            </Tag>
          );
        })}
      </>
    ),
  },
  {
    title: '操作',
    key: '操作',
    render: (_, record) => (
      <Space size="middle">
        <a>修改</a>
        <a>修改角色 - {record.name}</a>
        <a>重置密码</a>
        <a>删除</a>
      </Space>
    ),
  },
];
const data = [
  {
    key: '1',
    name: 'John Brown',
    age: 32,
    address: 'New York No. 1 Lake Park',
    tags: ['nice', 'developer'],
  },
  {
    key: '2',
    name: 'Jim Green',
    age: 42,
    address: 'London No. 1 Lake Park',
    tags: ['loser'],
  },
  {
    key: '3',
    name: 'Joe Black',
    age: 32,
    address: 'Sydney No. 1 Lake Park',
    tags: ['cool', 'teacher'],
  },
];
const User = () => {
  const [open, setOpen] = useState(false);
  const [form] = Form.useForm();

  const handleAdd = (values) => {
    // TODO: 提交新增用户逻辑
    console.log('新增用户:', values);
    setOpen(false);
    form.resetFields();
  };

  return (
    <div>
      <Row justify="space-between" align="middle" style={{ marginBottom: 16 }}>
        <Col flex="1 1 auto">
          <Input.Search
            placeholder="请输入用户名/手机号/邮箱"
            allowClear
            style={{ maxWidth: 300 }}
            onSearch={value => {
              // TODO: 实现查询逻辑
              console.log('搜索:', value);
            }}
          />
        </Col>
        <Col>
          <Button type="primary" style={{ float: 'right' }} onClick={() => setOpen(true)}>
            新增
          </Button>
        </Col>
      </Row>
      <Table columns={columns} dataSource={data} />
      <Modal
        title="新增用户"
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}
        destroyOnClose
      >
        <Form layout="vertical" form={form} onFinish={handleAdd}>
          <Form.Item label="用户名" name="name" rules={[{ required: true, message: '请输入用户名' }]}> 
            <Input />
          </Form.Item>
          <Form.Item label="手机号" name="phone" rules={[{ required: true, message: '请输入手机号' }]}> 
            <Input />
          </Form.Item>
          <Form.Item label="邮箱" name="email" rules={[{ required: true, message: '请输入邮箱' }]}> 
            <Input />
          </Form.Item>
          <Form.Item label="部门" name="department"> 
            <Input />
          </Form.Item>
          <Form.Item label="角色" name="role"> 
              <Checkbox.Group
                options={[
                  { label: '管理员', value: 'admin' },
                  { label: '用户', value: 'user' },
                  { label: '访客', value: 'guest' },
                  { label: '老师', value: 'teacher' },
                  { label: '开发者', value: 'developer' },
                ]}
                style={{ width: '100%' }}
              />
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit" style={{ width: '100%' }}>
              提交
            </Button>
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
};
export default User;