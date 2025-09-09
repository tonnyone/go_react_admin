import { Alert, Button, Col, Form, Input, message, Row, Space, Spin, Table, Tag } from 'antd';
import { useEffect, useState } from 'react';
import { getAllRoles } from '../../service/role';
import { createUser, getUserList, updateUser } from '../../service/user';
import UserAdd from './UserAdd';

// 常量定义
const FORM_RESET_DELAY = 100;
const SEARCH_MAX_WIDTH = 300;

const User = () => {
  const [isUserModalVisible, setIsUserModalVisible] = useState(false);
  const [userForm] = Form.useForm();
  const [users, setUsers] = useState([]);
  const [roles, setRoles] = useState([]);
  const [loading, setLoading] = useState({
    table: false,
    roles: false,
    submit: false
  });
  const [error, setError] = useState(null);
  
  // 将分页和可能的过滤、排序条件都放在一个 state 中管理
  const [tableParams, setTableParams] = useState({
    pagination: {
      current: 1,
      pageSize: 10,
      total: 0,
    },
    // 你可以在这里添加 filters 和 sorter
  });

  // 加载角色数据 - 用于表单选项
  useEffect(() => {
    setLoading(prev => ({ ...prev, roles: true }));
    
    getAllRoles()
      .then(roles => {
        console.log("resp roles", roles);
        setRoles(roles);
      })
      .catch(err => {
        console.error("加载角色失败:", err);
        message.warning('角色数据加载失败，新增/编辑功能可能受限');
        // 角色加载失败不设置全局错误，不影响页面其他功能
      })
      .finally(() => {
        setLoading(prev => ({ ...prev, roles: false }));
      });
  }, []);

  // 加载用户列表数据
  useEffect(() => {
    fetchUserList();
  }, []);

  // 封装获取用户列表的方法，避免代码重复
  const fetchUserList = (customParams = {}) => {
    const { pagination = tableParams.pagination, ...extraParams } = customParams;
    
    setLoading(prev => ({ ...prev, table: true }));
    setError(null);
    
    getUserList({
      current: pagination.current,
      page_size: pagination.pageSize,
      ...extraParams, // 支持传入额外的查询参数，如搜索条件
    })
    .then(response => {
      setUsers(response.list);
      console.log("resp users", response.list);
      
      // 更新分页状态，特别是 total
      setTableParams(prev => ({
        ...prev,
        pagination: {
          ...pagination,
          total: response.total,
        },
      }));
    })
    .catch(err => {
      setError(err.message || '加载数据失败');
    })
    .finally(() => {
      setLoading(prev => ({ ...prev, table: false }));
    });
  };

  // 3. 创建一个处理表格变化的函数
  // Ant Design 的 Table 组件在分页、排序、过滤变化时会调用 onChange
  const handleTableChange = (pagination, filters, sorter) => {
    // 当用户点击分页或排序时，这个函数会被调用
    // 直接使用新的分页参数获取数据
    fetchUserList({ pagination, filters, sorter });
  };

  // 处理新增/修改用户
  const handleAdd = (values) => {
    console.log("提交用户数据", values);
    setLoading(prev => ({ ...prev, submit: true }));
    
    // 判断是新增还是修改
    const isEdit = values.id;
    const apiCall = isEdit ? updateUser(values.id, values) : createUser(values);
    
    apiCall
      .then(() => {
        const successMsg = isEdit ? "用户修改成功" : "用户创建成功";
        console.log(successMsg);
        message.success(successMsg);
        
        // 操作成功后关闭 Modal 并刷新数据
        setIsUserModalVisible(false);
        handleCancel(); // 重置表单
        
        // 触发表格刷新 - 使用当前的分页参数重新获取数据
        fetchUserList();
      })
      .catch(error => {
        console.error("操作失败:", error);
        const errorMsg = error.message || '操作失败';
        setError(errorMsg);
        message.error(errorMsg);
      })
      .finally(() => {
        setLoading(prev => ({ ...prev, submit: false }));
      });
  };

  const handleCancel = () => {
    // 延迟重置表单，确保在 Modal 关闭后执行
    setTimeout(() => {
      userForm.resetFields();
    }, FORM_RESET_DELAY);
    setIsUserModalVisible(false);
  };

  // 处理新增用户
  const handleAddNew = () => {
    userForm.resetFields(); // 清空表单
    setIsUserModalVisible(true);
  };

  // 处理编辑用户
  const handleEdit = (record) => {
    console.log("编辑用户", record);
    // 将用户数据设置到表单中
    userForm.setFieldsValue({
      id: record.id,
      name: record.username,
      phone: record.phone,
      email: record.email,
      department: record.department,
      role: record.roles ? record.roles.map(role => role.id) : []
    });
    setIsUserModalVisible(true);
  };

  // 处理搜索
  const handleSearch = (value) => {
    console.log('搜索:', value);
    // 搜索时重置到第一页，因为需要传递搜索参数
    const resetPagination = { ...tableParams.pagination, current: 1 };
    fetchUserList({ 
      pagination: resetPagination,
      search: value 
    });
  };

  // 表格列定义
  const tableColumns = [
    { title: '用户名', dataIndex: 'username', key: 'username' },
    { title: '手机号', dataIndex: 'phone', key: 'phone' },
    { title: '邮箱', dataIndex: 'email', key: 'email' },
    { title: '部门', dataIndex: 'department', key: 'department' },
    { title: '更新时间', dataIndex: 'updated_at', key: 'updated_at' },
    {
      title: '角色',
      key: 'roles',
      dataIndex: 'roles',
      render: (roles) => (
        <>
          {roles && roles.map(role => (
            <Tag key={role.id}>
              {role.name}
            </Tag>
          ))}
        </>
      ),
    },
    {
      title: '操作',
      key: '操作',
      render: (_, record) => (
        <Space size="middle">
          <a onClick={() => handleEdit(record)}>修改</a>
          <a>修改角色</a>
          <a>重置密码</a>
          <a>删除</a>
        </Space>
      ),
    }
  ];

  if (error) {
    return <Alert message="错误" description={error} type="error" showIcon />;
  }

  return (
    <div>
      <UserAdd 
        open={isUserModalVisible} 
        roles={roles}
        form={userForm} 
        onSubmit={handleAdd}
        onCancel={handleCancel}  
        loading={loading.submit}
      />
      <Row justify="space-between" align="middle" style={{ marginBottom: 16 }}>
        <Col flex="1 1 auto">
          <Input.Search
            placeholder="请输入用户名/手机号/邮箱"
            allowClear
            style={{ maxWidth: SEARCH_MAX_WIDTH }}
            onSearch={handleSearch}
          />
        </Col>
        <Col>
          <Button type="primary" style={{ float: 'right' }} onClick={handleAddNew}>
            新增
          </Button>
        </Col>
      </Row>
      <Spin spinning={loading.table}>
        <Table
          dataSource={users}
          columns={tableColumns}
          rowKey="id"
          pagination={tableParams.pagination}
          onChange={handleTableChange}
        />
      </Spin>
    </div>
  );
};

export default User;