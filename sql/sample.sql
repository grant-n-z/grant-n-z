USE revel;

DROP TABLE IF EXISTS items;

-- items table
CREATE TABLE items (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(128) NOT NULL,
  category varchar(128) NOT NULL,
  created_at datetime(6),
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- items test data
INSERT INTO items(id, name, category, created_at)
VALUES (1, 'test_title01', 'Im fine', '1970-01-01 00:00:01');

INSERT INTO items(id, name, category, created_at)
VALUES (2, 'hoge', 'hoge', '1970-01-01 00:00:01');
