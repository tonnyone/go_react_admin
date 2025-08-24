import { Button, Form, Input } from "antd";
import { useNavigate } from 'react-router-dom';
const Login = () => {
const navigate = useNavigate();
const onFinishFailed = errorInfo => {
  console.log('Failed:', errorInfo);
};
const onFinish = values => {
  console.log('Success:', values);
  // Navigate to the dashboard or another page on successful login
  navigate('/');
};
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
