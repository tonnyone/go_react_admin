import axios from 'axios';

// 创建 axios 实例
const instance = axios.create({
  baseURL: '/api', // 可根据实际情况修改
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// 请求拦截器
instance.interceptors.request.use(
  config => {
    // 除 login 接口外，自动携带 token
    if (!/\/login$/.test(config.url)) {
      const token = localStorage.getItem('token');
      if (token) {
        config.headers.Authorization = `Basic ${token}`;
      }
    }
    return config;
  },
  error => Promise.reject(error)
);

// 响应拦截器
instance.interceptors.response.use(
  response => {
    const data = response.data;
    // 统一处理业务状态码
    if (data.code !== 0) {
      // 业务错误，抛出错误让调用方处理
      const error = new Error(data.msg || data.message || '请求失败');
      error.code = data.code;
      return Promise.reject(error);
    }
    // 成功时直接返回 data 部分
    return data.data;
  },
  error => {
    // HTTP 错误处理
    if (error.response) {
      console.error('HTTP error:', error.response.status, error.response.data);
    } else {
      console.error('Network error:', error.message);
    }
    return Promise.reject(error);
  }
);

export default instance;
