import { Button, Checkbox, Form, Input, Modal } from 'antd';

const AddUser = ({person, open }) => {
  const onFinish = (values) => {
    console.log('新增用户:', values);
  };

  return (
      <Modal
        title="新增用户"
        open={open}
        onCancel={() => setOpen(false)}
        footer={null}>
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
                style={{ width: '100%' }}/>
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit" style={{ width: '100%' }}>
              提交
            </Button>
          </Form.Item>
        </Form>
      </Modal>
  );
};

export default AddUser;