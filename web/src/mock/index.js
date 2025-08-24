// src/mock/index.js
import MockAdapter from 'axios-mock-adapter';
import axios from '../service/request';

// 创建 Mock 实例
const mock = new MockAdapter(axios, { delayResponse: 500 });

// 示例：mock 用户列表接口
mock.onGet('/user/list').reply(200, {
  code: 0,
  data: [
    { id: 1, name: '张三' },
    { id: 2, name: '李四' }
  ]
});

// mock 登录接口
mock.onPost('/login').reply(config => {
  const { username, password } = JSON.parse(config.data);
  if (username == "admin" && password=="123456") {
    return [200, {
        code: 0,
        msg: 'success',
        data: {
        token: 'mocked-jwt-token-1234567890'
        }
    }];
  }else{
    return [200, {
        code: 1,
        msg: '用户名或密码错误',
    }];
  }
});

// 你可以继续添加更多 mock 规则
export default mock;