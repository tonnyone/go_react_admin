import { Alert, Button, Col, Form, Input, Row, Space, Spin, Table, Tag } from 'antd';
import { useCallback, useEffect, useState } from 'react';
import { getUserList } from '../../service/user';
import AddUser from './UserAdd';

const User = () => {
  const [isUserModalVisible, setIsUserModalVisible] = useState(false);
  const [userForm] = Form.useForm();
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);
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

  // 1. 将数据获取逻辑封装在一个 useCallback 中
  // useCallback 可以缓存这个函数，避免不必要的重新创建
  const fetchData = useCallback(async () => {
    try {
      setLoading(true);
      setError(null);

      // 使用 tableParams 中的状态来调用 API
      const response = await getUserList({
        current: tableParams.pagination.current,
        page_size: tableParams.pagination.pageSize,
        // ...可以传递其他过滤和排序参数
      });
      setUsers(response.list);
      console.log("resp users", response.list)
      // 更新分页状态，特别是 total
      setTableParams(prev => ({
        ...prev,
        pagination: {
          ...prev.pagination,
          total: response.total,
        },
      }));
    } catch (err) {
      setError(err.message || '加载数据失败');
    } finally {
      setLoading(false);
    }
  }, [tableParams.pagination.current, tableParams.pagination.pageSize]); // 依赖项：当页码或页大小变化时，fetchData 会更新

  // 2. useEffect 现在依赖 fetchData 函数
  // 当 fetchData 函数因为其依赖项变化而更新时，这个 effect 就会重新运行
  useEffect(() => {
    fetchData();
  }, [fetchData]);

  // 3. 创建一个处理表格变化的函数
  // Ant Design 的 Table 组件在分页、排序、过滤变化时会调用 onChange
  const handleTableChange = (pagination, filters, sorter) => {
    // 当用户点击分页或排序时，这个函数会被调用
    // 我们用新的分页信息来更新我们的状态
    setTableParams(prev => ({
      ...prev,
      pagination,
      // ...可以更新 filters 和 sorter
    }));
  };

  const handleAdd = (values) => {
    console.log("add user", values)
    setLoading(true)
    setTimeout(() => {
      setLoading(false)
    }, 60000)
  };

  const handleCancel = () => {
    userForm.resetFields();
    setIsUserModalVisible(false);
  };

  if (error) {
    return <Alert message="错误" description={error} type="error" showIcon />;
  }

  return (
    <div>
      <AddUser isUserModalVisible={isUserModalVisible} userForm={userForm} handleAdd={handleAdd} handleCancel={handleCancel}  loading={loading} />
      <Row justify="space-between" align="middle" style={{ marginBottom: 16 }}>
        <Col flex="1 1 auto">
          <Input.Search
            placeholder="请输入用户名/手机号/邮箱"
            allowClear
            style={{ maxWidth: 300 }}
            onSearch={value => {
              // TODO: 实现查询逻辑
              console.log('搜索:', value);
            }}
          />
        </Col>
        <Col>
          <Button type="primary" style={{ float: 'right' }} onClick={() => setIsUserModalVisible(true)}>
            新增
          </Button>
        </Col>
      </Row>
      <Spin spinning={loading}>
        <Table
          dataSource={users}
          columns={[
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
                  {
                    roles && roles.map(role => {
                      return (
                        <Tag key={role.id}>
                          {role.name}
                        </Tag>
                      );
                    })
                  }
                </>
              ),
            },
            {
              title: '操作',
              key: '操作',
              render: (_, record) => (
                <Space size="middle">
                  <a>修改</a>
                  <a>修改角色</a>
                  <a>重置密码</a>
                  <a>删除</a>
                </Space>
              ),
            }
          ]}
          rowKey="id"
          // 将我们的状态和处理函数传递给 Table 组件
          pagination={tableParams.pagination}
          onChange={handleTableChange}
        />
      </Spin>
    </div>
  );
};

export default User;