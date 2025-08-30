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
  response => response.data,
  error => {
    // 可全局处理错误提示
    // if (error.response) {
    //   message.error(error.response.data.message || '请求错误');
    // }
    // console.error("error: %o", error);
    return Promise.reject(error);
  }
);

export default instance;
