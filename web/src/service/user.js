import request from './request';

export const login = async (params) => {
  const res = await request.post('/login', params);
  if (res.code === 0 && res.data && res.data.token) {
    localStorage.setItem('token', res.data.token);
    return res;
  } else {
    throw new Error(res.msg || '登录失败');
  }
};