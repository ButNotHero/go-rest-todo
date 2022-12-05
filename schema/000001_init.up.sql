CREATE TABLE auth_user
(
  id            serial       NOT NULL unique,
  name          varchar(255) NOT NULL,
  username      varchar(255) NOT NULL unique,
  password_hash varchar(255) NOT NULL
);

CREATE TABLE todo_list
(
  id          serial       NOT NULL unique,
  title       varchar(255) NOT NULL,
  description varchar(255)
);

CREATE TABLE user_list
(
  id      serial                                          NOT NULL unique,
  user_id int references auth_user (id) on delete cascade     NOT NULL,
  list_id int references todo_list (id) on delete cascade NOT NULL
);

CREATE TABLE todo_item
(
  id          serial       NOT NULL unique,
  title       varchar(255) NOT NULL,
  description varchar(255),
  done        boolean      NOT NULL default false
);

CREATE TABLE list_item
(
  id      serial                                          NOT NULL unique,
  item_id int references todo_item (id) on delete cascade NOT NULL,
  list_id int references todo_list (id) on delete cascade NOT NULL
);
