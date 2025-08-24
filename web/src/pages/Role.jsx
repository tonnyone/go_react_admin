
import { Button, Col, Row, Space, Switch, Table } from 'antd';

const initialData = [
  { id: 1, name: '管理员', desc: '系统最高权限', enabled: true },
  { id: 2, name: '用户', desc: '普通用户', enabled: true },
  { id: 3, name: '访客', desc: '只读权限', enabled: false },
  { id: 4, name: '开发者', desc: '开发和测试权限', enabled: true },
];

const columns = [
  { title: 'ID', dataIndex: 'id', key: 'id', width: 60 },
  { title: '名称', dataIndex: 'name', key: 'name' },
  { title: '描述', dataIndex: 'desc', key: 'desc' },
  {
    title: '是否启用/操作',
    key: 'enabled_action',
    render: (record) => (
      <Space size="middle">
        <Switch checked={record.enabled} disabled />
        <a href="#" onClick={e => { e.preventDefault(); /* TODO: 打开编辑弹窗 */ }}>修改</a>
        <a href="#" style={{ color: 'red' }} onClick={e => { e.preventDefault(); /* TODO: 删除操作 */ }}>删除</a>
      </Space>
    ),
  },
];

const Role = () => {
  return (
    <div style={{ padding: 24 }}>
      <Row justify="space-between" align="middle" style={{ marginBottom: 16 }}>
        <Col><h2 style={{ margin: 0 }}>角色管理</h2></Col>
        <Col>
          <Button type="primary">新增</Button>
        </Col>
      </Row>
      <Table
        rowKey="id"
        columns={columns}
        dataSource={initialData}
        pagination={false}
        bordered
        style={{ background: '#fff' }}
      />
    </div>
  );
};

export default Role;