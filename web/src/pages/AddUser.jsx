import { Button, Card, Form, Input } from 'antd';

const AddUser = () => {
  const onFinish = (values) => {
    // TODO: 提交新增用户逻辑
    console.log('新增用户:', values);
  };

  return (
    <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', minHeight: '100vh', background: '#f5f6fa' }}>
      <Card title="新增用户" style={{ width: 400 }}>
        <Form layout="vertical" onFinish={onFinish}>
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
            <Input />
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit" style={{ width: '100%' }}>
              提交
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </div>
  );
};

export default AddUser;
