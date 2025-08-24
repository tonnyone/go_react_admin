import { Button, Form, Input } from "antd";
import { useContext, useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { MessageContext } from '../App';
import { login } from '../service/user';

const Login = () => {
  const navigate = useNavigate();
  const messageApi = useContext(MessageContext);
  const [loginError, setLoginError] = useState(null);
  const onFinishFailed = errorInfo => {
    console.log('Failed:', errorInfo);
  };
  const onFinish = (data) => {
    login(data)
      .then(res => {
        console.log("success", res);
        navigate('/user/list');
        setLoginError(null);
      })
      .catch(err => {
        setLoginError('登录失败，请检查用户名和密码');
      });
  };

  useEffect(() => {
    if (loginError && messageApi) {
      messageApi.error(loginError);
    }
  }, [loginError, messageApi]);

  return (
    <div style={{justifyContent: 'center', alignItems: 'center', display: 'flex', height: '100vh'}}>
      <Form
        name="basic"
        labelCol={{ span: 8 }}
        wrapperCol={{ span: 16 }}
        style={{ maxWidth: 800}}
        initialValues={{ remember: true }}
        onFinish={onFinish}
        onFinishFailed={onFinishFailed}
        autoComplete="off"
      >
        <Form.Item
          label="Username"
          name="username"
          rules={[{ required: true, message: "Please input your username!" }]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Password"
          name="password"
          rules={[{ required: true, message: "Please input your password!" }]}
        >
          <Input.Password />
        </Form.Item>

        <Form.Item label={null}>
          <Button type="primary" htmlType="submit">
            Submit
          </Button>
          <Button type="default" htmlType="reset">
            Reset
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};

export default Login;
