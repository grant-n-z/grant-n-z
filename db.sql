-- name: create-grant-n-z
CREATE DATABASE grant_n_z;

-- name: use-grant-n-z
USE grant_n_z;

-- name: drop-tokens
DROP TABLE IF EXISTS tokens;

-- name: drop-roles
DROP TABLE IF EXISTS roles;

-- name: drop-users
DROP TABLE IF EXISTS users;

-- name: create-users
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

-- name: create-tokens
CREATE TABLE tokens (
  id int(11) NOT NULL AUTO_INCREMENT,
  token_type varchar(128) NOT NULL,
  token varchar(512) NOT NULL,
  refresh_token varchar(512) NOT NULL,
  user_uuid varchar(128) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE (token)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- name: create-roles
CREATE TABLE roles (
  id int(11) NOT NULL AUTO_INCREMENT,
  type varchar(128) NOT NULL,
  user_uuid varchar(128) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

