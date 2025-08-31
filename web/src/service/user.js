import request from './request';

/**
 * 登录
 * @param {object} params - 登录所需的参数，例如 { username, password }
 * @returns {Promise<any>}
 */
export const login = async (params) => {
  const res = await request.post('/login', params);
  if (res.code === 0 && res.data) {
    localStorage.setItem('token', res.data);
    return res;
  } else {
    throw new Error(res.msg || '登录失败');
  }
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
  // 使用 await 等待请求完成
  const response = await request.get('/user/list', { params });
  // 检查后端返回的业务状态
  if (response.code == 0) {
    // 成功，则返回真正的数据部分，供组件直接使用
    return response.data;
  } else {
    throw new Error(response.data.message || '获取用户列表失败');
  }
}

/**
 * 更新用户信息
 * @param {string} userId - 用户ID
 * @param {object} data - 需要更新的用户数据
 * @returns {Promise<any>}
 */
export async function updateUser(userId, data) {
  try {
    const response = await request.put(`/api/user/${userId}`, data);
    if (response.data && response.data.success) {
      return response.data.data;
    } else {
      throw new Error(response.data.message || '更新用户失败');
    }
  } catch (error) {
    console.error('Service error in updateUser:', error.message);
    throw error;
  }
}

/**
 * 删除用户
 * @param {string} userId - 用户ID
 * @returns {Promise<any>}
 */
export async function deleteUser(userId) {
  try {
    const response = await request.delete(`/api/user/${userId}`);
    if (response.data && response.data.success) {
      return response.data.data;
    } else {
      throw new Error(response.data.message || '删除用户失败');
    }
  } catch (error) {
    console.error('Service error in deleteUser:', error.message);
    throw error;
  }
}

// 你可以按照这个模式添加其他函数，例如 createUser