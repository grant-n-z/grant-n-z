-- grant-n-z
DROP DATABASE IF EXISTS grant_n_z;
CREATE DATABASE IF NOT EXISTS grant_n_z;

-- use grant-n-z
USE grant_n_z;

-- If services exit, drop services
DROP TABLE IF EXISTS services;

-- If users exit, drop users
DROP TABLE IF EXISTS users;

-- If permissions exit, drop permissions
DROP TABLE IF EXISTS permissions;

-- If user_services exit, drop user_services
DROP TABLE IF EXISTS user_services;

-- If roles exit, drop roles
DROP TABLE IF EXISTS roles;

-- If operate_member_roles exit, drop roles
DROP TABLE IF EXISTS operator_member_roles;

-- If service_member_roles exit, drop roles
DROP TABLE IF EXISTS service_member_roles;

-- If policies exit, drop policies
DROP TABLE IF EXISTS policies;

-- `services`
CREATE TABLE services (
  id int(11) NOT NULL AUTO_INCREMENT,
  uuid varchar(128) NOT NULL,
  name varchar(128) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- `users`
-- The all user data.
CREATE TABLE users (
  id int(11) NOT NULL AUTO_INCREMENT,
  uuid varchar(128) NOT NULL,
  username varchar(128) NOT NULL,
  email varchar(128) NOT NULL,
  password varchar(128) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- `permissions`
-- The permission data.
CREATE TABLE permissions (
  id int(11) NOT NULL AUTO_INCREMENT,
  uuid varchar(128) NOT NULL,
  name varchar(128) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- `roles`
-- The role data.
CREATE TABLE roles (
  id int(11) NOT NULL AUTO_INCREMENT,
  uuid varchar(128) NOT NULL,
  name varchar(128) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- `operator_member_roles`
-- It can operate grant_n_z server.
CREATE TABLE operator_member_roles (
  id int(11) NOT NULL AUTO_INCREMENT,
  role_id int(11) NOT NULL,
  user_id int(11) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  INDEX (user_id),
  PRIMARY KEY (id),
  CONSTRAINT fk_operate_member_roles_role_id
  FOREIGN KEY (role_id)
  REFERENCES roles (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_operate_member_roles_user_id
  FOREIGN KEY (user_id)
  REFERENCES users (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- `user_services`
-- If we want to know how many user of service, read this table.
-- If we want to know how many service of user, read this table.
CREATE TABLE user_services (
  id int(11) NOT NULL AUTO_INCREMENT,
  user_id int(11) NOT NULL,
  service_id int(11) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  INDEX (user_id),
  INDEX (service_id),
  CONSTRAINT fk_user_services_user_id
  FOREIGN KEY (user_id)
  REFERENCES users (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_user_services_service_id
  FOREIGN KEY (service_id)
  REFERENCES services (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- `service_member_roles`
-- It can operate any service.
-- Any service member role.
CREATE TABLE service_member_roles (
  id int(11) NOT NULL AUTO_INCREMENT,
  role_id int(11) NOT NULL,
  user_service_id int(11) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  INDEX (role_id),
  INDEX (user_service_id),
  PRIMARY KEY (id),
  CONSTRAINT fk_service_member_roles_role_id
  FOREIGN KEY (role_id)
  REFERENCES roles (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_service_member_roles_user_service_id
  FOREIGN KEY (user_service_id)
  REFERENCES user_services (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- `policies`
-- Access policy.
CREATE TABLE policies (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(128) NOT NULL,
  permission_id int(11) NOT NULL,
  service_member_role_id int(11) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  INDEX (permission_id),
  INDEX (service_member_role_id),
  CONSTRAINT fk_policies_permission_id
  FOREIGN KEY (permission_id)
  REFERENCES permissions (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_policies_service_member_role_id
  FOREIGN KEY (service_member_role_id)
  REFERENCES service_member_roles (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
