import { LockOutlined, MailOutlined, MobileOutlined } from '@ant-design/icons';
import { Button, Form, Input, Tabs, message } from "antd";
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { login } from '../service/user';


const Login = () => {
  const navigate = useNavigate();
  const [messageApi, contextHolder] = message.useMessage();
  const [loginType, setLoginType] = useState('mobile'); // 'mobile' or 'email'

  const onFinishFailed = errorInfo => {
    console.error('Failed:', errorInfo);
  };

  const onFinish = (data) => {
    login(data)
      .then(res => {
        navigate('/user/list');
      })
      .catch(err => {
        if (messageApi) {
          messageApi.error(err.message || '登录失败，请检查用户名和密码');
        }
      });
  };

  // 使用 items 属性来定义标签页
  const tabItems = [
    {
      key: 'mobile',
      label: '手机号登录',
    },
    {
      key: 'email',
      label: '邮箱登录',
    },
  ];

  return (
    <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh', background: '#f0f2f5' }}>
      {contextHolder}
      <div style={{ width: 400, margin: '0px 0px 100px 0px', padding: '40px', boxShadow: '0 4px 8px rgba(0,0,0,0.1)', borderRadius: '8px', background: 'white' }}>
        <h2 style={{ textAlign: 'center', marginBottom: '30px' }}>系统登录</h2>
        <Tabs 
          activeKey={loginType} 
          onChange={setLoginType} 
          centered 
          items={tabItems} // 使用 items 属性
        />
        <Form
          name="basic"
          initialValues={{ remember: true }}
          onFinish={onFinish}
          onFinishFailed={onFinishFailed}
          autoComplete="off"
          size="large"
        >
          {loginType === 'mobile' ? (
            <Form.Item
              name="account"
              rules={[
                { required: true, message: "请输入手机号!" },
                { pattern: /^1\d{10}$/, message: "请输入有效的11位手机号!" }
              ]}
            >
              <Input prefix={<MobileOutlined />} placeholder="手机号" />
            </Form.Item>
          ) : (
            <Form.Item
              name="account"
              rules={[
                { required: true, message: "请输入邮箱地址!" },
                { type: 'email', message: '请输入有效的邮箱地址!' }
              ]}
            >
              <Input prefix={<MailOutlined />} placeholder="邮箱" />
            </Form.Item>
          )}

          <Form.Item
            name="password"
            rules={[{ required: true, message: "请输入密码!" }]}
          >
            <Input.Password prefix={<LockOutlined />} placeholder="密码" />
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit" style={{ width: '100%' }}>
              登 录
            </Button>
          </Form.Item>
        </Form>
      </div>
    </div>
  );
};

export default Login;
