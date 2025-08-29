import request from './request';

export const login = async (params) => {
  const res = await request.post('/login', params);
  if (res.code === 0 && res.data) {
    localStorage.setItem('token', res.data);
    return res;
  } else {
    throw new Error(res.msg || '登录失败');
  }
};

export const logout = async (params) => {
  localStorage.clear();
  return Promise.resolve();
};