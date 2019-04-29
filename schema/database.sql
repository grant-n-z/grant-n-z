-- grant-n-z
DROP DATABASE IF EXISTS grant_n_z;
CREATE DATABASE IF NOT EXISTS grant_n_z;

-- use grant-n-z
USE grant_n_z;

-- If services exit, drop services
DROP TABLE IF EXISTS services;

-- If users exit, drop users
DROP TABLE IF EXISTS users;

-- If user_services exit, drop user_services
DROP TABLE IF EXISTS user_services;

-- If policies exit, drop policies
DROP TABLE IF EXISTS user_services;

-- If roles exit, drop roles
DROP TABLE IF EXISTS roles;

-- services
-- If register service, add row this table
CREATE TABLE services (
  id int(11) NOT NULL AUTO_INCREMENT,
  uuid varchar(128) NOT NULL,
  name varchar(128) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- users
-- The user data
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

-- roles
-- The role
CREATE TABLE roles (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(128) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- user_services
-- The service data of user
CREATE TABLE user_services (
  id int(11) NOT NULL AUTO_INCREMENT,
  user_id int(11) NOT NULL,
  service_id int(11) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  INDEX (user_id),
  INDEX (service_id),
  CONSTRAINT fk_user_id
  FOREIGN KEY (user_id)
  REFERENCES users (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_service_id
  FOREIGN KEY (service_id)
  REFERENCES services (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- policies
-- The role policy of user_service id
CREATE TABLE policies (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(128) NOT NULL,
  user_service_id int(11) NOT NULL,
  role_id int(11) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  INDEX (user_service_id),
  CONSTRAINT fk_user_service_id
  FOREIGN KEY (user_service_id)
  REFERENCES user_services (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_role_id
  FOREIGN KEY (role_id)
  REFERENCES roles (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
