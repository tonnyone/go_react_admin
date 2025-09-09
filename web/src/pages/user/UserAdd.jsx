import { Button, Checkbox, Form, Input, Modal } from 'antd';
import { useEffect, useState } from 'react';

const UserAdd = ({ form, roles, open, onSubmit, onCancel, loading }) => {
  const [isEdit, setIsEdit] = useState(false);

  // 监听 form 和 open 的变化来判断是否为编辑模式
  useEffect(() => {
    if (open && form) {
      try {
        const formValues = form.getFieldsValue();
        setIsEdit(formValues && formValues.id);
      } catch (error) {
        // 如果 form 还没准备好，默认为新增模式
        setIsEdit(false);
      }
    }
  }, [open, form]);
  
  return (
      <Modal
        title={isEdit ? "修改用户" : "新增用户"}
        open={open}
        onCancel={onCancel}
        footer={null}
      >
        <Form layout="vertical" form={form} onFinish={onSubmit}>
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
              options={roles.map(role => ({ label: role.name, value: role.id, key: role.id }))}
              style={{ width: '100%' }}
            />
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit" style={{ width: '100%' }} loading={loading}>
              {isEdit ? "更新" : "提交"}
            </Button>
          </Form.Item>
        </Form>
      </Modal>
  );
};

export default UserAdd;