
import request from './request';

/**
 * 获取角色列表
 * @param {Object} [data] - 查询参数
 * @param {string} [data.name] - 角色名称（模糊查询）
 * @param {string} [data.status] - 角色状态 ('active' | 'inactive')
 * @returns {Promise<Array<{
 *   id: string,
 *   name: string,
 *   description: string,
 *   status: 'active' | 'inactive',
 *   permissions: Array<string>,
 *   createdAt: string,
 *   updatedAt: string
 * }>>} 角色列表数组
 */
export async function getAllRoles(data) {
    const result = await request.get('/role', { params: { ...data, page_size: 100, current: 1 } });
    return result.list;
}

/**
 * 创建角色
 * @param {Object} data - 角色数据
 * @param {string} data.name - 角色名称
 * @param {string} [data.description] - 角色描述
 * @param {Array<string>} [data.permissions] - 权限列表
 * @returns {Promise<{
 *   id: string,
 *   name: string,
 *   description: string,
 *   status: string,
 *   permissions: Array<string>
 * }>} 创建的角色信息
 */
export async function createRole(data) {
    return await request.post('/role', data);
}