-- name: create-auth-server
CREATE DATABASE auth_server;

-- name: use-auth-server
USE auth_server;

-- name: drop-tokens
DROP TABLE IF EXISTS tokens;

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

