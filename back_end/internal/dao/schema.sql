-- 创建数据库
CREATE DATABASE IF NOT EXISTS go_react_admin DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE go_react_admin;

-- 用户表
CREATE TABLE IF NOT EXISTS user (
  id CHAR(36) PRIMARY KEY,
  username VARCHAR(64) NOT NULL UNIQUE,
  password VARCHAR(128) NOT NULL,
  email VARCHAR(128),
  department VARCHAR(64),
  deleted TINYINT(1) DEFAULT 0 COMMENT '软删除标记 0-正常 1-已删除',
  disabled TINYINT(1) DEFAULT 0 COMMENT '禁用标记 0-正常 1-禁用',
  create_time BIGINT UNSIGNED,
  update_time BIGINT UNSIGNED
);

-- 角色表
CREATE TABLE IF NOT EXISTS role (
  id CHAR(36) PRIMARY KEY,
  name VARCHAR(64) NOT NULL UNIQUE,
  create_time BIGINT UNSIGNED,
  update_time BIGINT UNSIGNED
);

CREATE TABLE IF NOT EXISTS menu (
  id CHAR(36) PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  parent_id CHAR(36),
  path VARCHAR(128),
  type VARCHAR(16),
  create_time BIGINT UNSIGNED,
  update_time BIGINT UNSIGNED
);

-- 资源表
CREATE TABLE IF NOT EXISTS resource (
  id CHAR(36) PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  path VARCHAR(128),
  method VARCHAR(16),
  create_time BIGINT UNSIGNED,
  update_time BIGINT UNSIGNED
);

  user_id CHAR(36) NOT NULL,
  role_id CHAR(36) NOT NULL,
  create_time BIGINT UNSIGNED,
  update_time BIGINT UNSIGNED,
  PRIMARY KEY (user_id, role_id)
);

  role_id CHAR(36) NOT NULL,
  menu_id CHAR(36) NOT NULL,
  create_time BIGINT UNSIGNED,
  update_time BIGINT UNSIGNED,
  PRIMARY KEY (role_id, menu_id)
);

  role_id CHAR(36) NOT NULL,
  resource_id CHAR(36) NOT NULL,
  create_time BIGINT UNSIGNED,
  update_time BIGINT UNSIGNED,
  PRIMARY KEY (role_id, resource_id)
);

  user_id CHAR(36) NOT NULL,
  menu_id CHAR(36) NOT NULL,
  create_time BIGINT UNSIGNED,
  update_time BIGINT UNSIGNED,
  PRIMARY KEY (user_id, menu_id)
);

  user_id CHAR(36) NOT NULL,
  resource_id CHAR(36) NOT NULL,
  create_time BIGINT UNSIGNED,
  update_time BIGINT UNSIGNED,
  PRIMARY KEY (user_id, resource_id)
);

-- 日志表
  id CHAR(36) PRIMARY KEY,
  user_id CHAR(36) COMMENT '操作用户ID',
  action VARCHAR(32) NOT NULL COMMENT '操作类型，如create/update/delete',
  object_type VARCHAR(32) NOT NULL COMMENT '操作对象类型，如user/menu/resource/role',
  object_id CHAR(36) COMMENT '操作对象ID',
  detail TEXT COMMENT '操作详情',
  create_time BIGINT UNSIGNED,
  update_time BIGINT UNSIGNED
);
