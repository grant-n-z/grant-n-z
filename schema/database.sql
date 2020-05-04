-- grant-n-z
DROP DATABASE IF EXISTS grant_n_z;
CREATE DATABASE IF NOT EXISTS grant_n_z;

-- use grant-n-z
USE grant_n_z;

-- If services exit, drop services
DROP TABLE IF EXISTS services;

-- If groups exit, drop groups
DROP TABLE IF EXISTS groups;

-- If roles exit, drop roles
DROP TABLE IF EXISTS roles;

-- If permissions exit, drop permissions
DROP TABLE IF EXISTS permissions;

-- If users exit, drop users
DROP TABLE IF EXISTS users;

-- If user_services exit, drop user_services
DROP TABLE IF EXISTS user_services;

-- If user_groups exit, drop user_groups
DROP TABLE IF EXISTS user_groups;

-- If service_groups exit, drop service_groups
DROP TABLE IF EXISTS service_groups;

-- If service_roles exit, drop service_roles
DROP TABLE IF EXISTS service_groups;

-- If service_permissions exit, drop service_permissions
DROP TABLE IF EXISTS service_permissions;

-- If group_roles exit, drop group_roles
DROP TABLE IF EXISTS group_groups;

-- If group_permissions exit, drop group_permissions
DROP TABLE IF EXISTS group_permissions;

-- If operator_policies exit, drop operator_policies
DROP TABLE IF EXISTS operator_policies;

-- If policies exit, drop policies
DROP TABLE IF EXISTS policies;

-- `services`
CREATE TABLE services (
  id int(11) NOT NULL AUTO_INCREMENT,
  uuid varchar(128) NOT NULL,
  name varchar(128) NOT NULL,
  secret varchar(128) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- `users`
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
CREATE TABLE permissions (
  id int(11) NOT NULL AUTO_INCREMENT,
  uuid varchar(128) NOT NULL,
  name varchar(128) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- `groups`
CREATE TABLE groups (
  id int(11) NOT NULL AUTO_INCREMENT,
  uuid varchar(128) NOT NULL,
  name varchar(128) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- `roles`
CREATE TABLE roles (
  id int(11) NOT NULL AUTO_INCREMENT,
  uuid varchar(128) NOT NULL,
  name varchar(128) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- `user_services`
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

-- `user_groups`
CREATE TABLE user_groups (
  id int(11) NOT NULL AUTO_INCREMENT,
  user_id int(11) NOT NULL,
  group_id int(11) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  INDEX (user_id),
  INDEX (group_id),
  CONSTRAINT fk_user_groups_user_id
  FOREIGN KEY (user_id)
  REFERENCES users (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_user_groups_group_id
  FOREIGN KEY (group_id)
  REFERENCES groups (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- `service_groups`
CREATE TABLE service_groups (
  id int(11) NOT NULL AUTO_INCREMENT,
  group_id int(11) NOT NULL,
  service_id int(11) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  INDEX (group_id),
  INDEX (service_id),
  CONSTRAINT fk_service_groups_group_id
  FOREIGN KEY (group_id)
  REFERENCES groups (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_service_groups_service_id
  FOREIGN KEY (service_id)
  REFERENCES services (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- `service_roles`
CREATE TABLE service_roles (
  id int(11) NOT NULL AUTO_INCREMENT,
  role_id int(11) NOT NULL,
  service_id int(11) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  INDEX (role_id),
  INDEX (service_id),
  CONSTRAINT fk_service_roles_role_id
  FOREIGN KEY (role_id)
  REFERENCES roles (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_service_roles_service_id
  FOREIGN KEY (service_id)
  REFERENCES services (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- `service_permissions`
CREATE TABLE service_permissions (
  id int(11) NOT NULL AUTO_INCREMENT,
  permission_id int(11) NOT NULL,
  service_id int(11) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  INDEX (permission_id),
  INDEX (service_id),
  CONSTRAINT fk_service_permissions_permission_id
  FOREIGN KEY (permission_id)
  REFERENCES permissions (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_service_permissions_service_id
  FOREIGN KEY (service_id)
  REFERENCES services (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- `group_roles`
CREATE TABLE group_roles (
  id int(11) NOT NULL AUTO_INCREMENT,
  role_id int(11) NOT NULL,
  group_id int(11) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  INDEX (role_id),
  INDEX (group_id),
  CONSTRAINT fk_group_roles_role_id
  FOREIGN KEY (role_id)
  REFERENCES roles (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_group_roles_group_id
  FOREIGN KEY (group_id)
  REFERENCES groups (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- `group_permissions`
CREATE TABLE group_permissions (
  id int(11) NOT NULL AUTO_INCREMENT,
  permission_id int(11) NOT NULL,
  group_id int(11) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  INDEX (permission_id),
  INDEX (group_id),
  CONSTRAINT fk_group_permissions_permission_id
  FOREIGN KEY (permission_id)
  REFERENCES permissions (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_group_permissions_group_id
  FOREIGN KEY (group_id)
  REFERENCES groups (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- `operator_policies`
CREATE TABLE operator_policies (
  id int(11) NOT NULL AUTO_INCREMENT,
  role_id int(11) NOT NULL,
  user_id int(11) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  INDEX (user_id),
  PRIMARY KEY (id),
  CONSTRAINT fk_operator_policies_role_id
  FOREIGN KEY (role_id)
  REFERENCES roles (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_operator_policies_user_id
  FOREIGN KEY (user_id)
  REFERENCES users (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- `policies`
CREATE TABLE policies (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(128) NOT NULL,
  role_id int(11) NOT NULL,
  permission_id int(11) NOT NULL,
  service_id int(11) NOT NULL,
  user_group_id int(11) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  INDEX (role_id),
  INDEX (permission_id),
  INDEX (service_id),
  INDEX (user_group_id),
  CONSTRAINT fk_policies_role_id
  FOREIGN KEY (role_id)
  REFERENCES roles (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_policies_permission_id
  FOREIGN KEY (permission_id)
  REFERENCES permissions (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_policies_service_id
  FOREIGN KEY (service_id)
  REFERENCES services (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_policies_user_group_id
  FOREIGN KEY (user_group_id)
  REFERENCES user_groups (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
