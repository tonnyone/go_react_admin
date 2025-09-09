import request from './request';

/**
 * 登录
 * @param {object} params - 登录所需的参数，例如 { username, password }
 * @returns {Promise<any>}
 */
export const login = async (params) => {
  const token = await request.post('/login', params);
  localStorage.setItem('token', token);
  return token;
};

/**
 * 登出
 * @returns {Promise<void>}
 */
export const logout = async () => {
  localStorage.clear();
  return Promise.resolve();
};

/**
 * 获取用户列表
 * @param {object} params - 查询参数，例如 { current, pageSize, ...filters }
 * @returns {Promise<{list: Array, total: number}>}
 */
export async function getUserList(params) {
  // 拦截器已统一处理 code 检查，直接返回数据
  return await request.get('/user', { params });
}

/**
 * 更新用户信息
 * @param {string} userId - 用户ID
 * @param {object} data - 需要更新的用户数据
 * @returns {Promise<any>}
 */
export async function updateUser(userId, data) {
  return await request.put(`/user/${userId}`, data);
}

/**
 * 删除用户
 * @param {string} userId - 用户ID
 * @returns {Promise<any>}
 */
export async function deleteUser(userId) {
  return await request.delete(`/user/${userId}`);
}

/**
 * 创建新用户
 * @param {object} data - 用户数据，例如 { name, phone, email, department, role }
 * @returns {Promise<any>}
 */
export async function createUser(data) {
  return await request.post('/user', data);
}