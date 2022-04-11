#!/bin/sh

CMD_MYSQL="mysql -u root -proot_password -D product_admin_api"

$CMD_MYSQL -e "create table if not exists shops (
  id          bigint unsigned auto_increment primary key,
  name        varchar(100) not null,
  description varchar(2000) not null,
  created_at  datetime not null default current_timestamp,
  updated_at  datetime not null default current_timestamp on update current_timestamp
);

create table if not exists products (
  id          bigint unsigned auto_increment primary key,
  title       varchar(100) not null,
  description varchar(2000) not null,
  price       int unsigned not null,
  created_at  datetime not null default current_timestamp,
  updated_at  datetime not null default current_timestamp on update current_timestamp,
  shop_id     bigint unsigned,
  foreign key (shop_id) references shops(id)
);

create table if not exists users (
  user_id     varchar(20) primary key,
  password    varchar(100) not null,
  created_at  datetime not null default current_timestamp,
  updated_at  datetime not null default current_timestamp on update current_timestamp
);

create table if not exists login_sessions (
  session_id  varchar(64) primary key,
  user_id     varchar(20),
  created_at  datetime not null default current_timestamp,
  foreign key (user_id) references users(user_id)
);"
