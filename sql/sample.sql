USE auth_server;

DROP TABLE IF EXISTS tokens;
DROP TABLE IF EXISTS users;

-- users table
CREATE TABLE users (
  id int(11) NOT NULL AUTO_INCREMENT,
  uuid varchar(128) NOT NULL,
  username varchar(128) NOT NULL,
  email varchar(128) NOT NULL,
  password varchar(128) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- tokens table
CREATE TABLE tokens (
  id int(11) NOT NULL AUTO_INCREMENT,
  token_type varchar(128) NOT NULL,
  token varchar(512) NOT NULL,
  refresh_token varchar(512) NOT NULL,
  user_id int(11) NOT NULL,
  expires_at datetime,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE (token),
  FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;