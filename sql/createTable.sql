-- ショップ
create table if not exists shops (
  id          bigint unsigned auto_increment primary key,
  name        varchar(100) not null,
  description varchar(2000) not null,
  created_at  datetime not null default current_timestamp,
  updated_at  datetime not null default current_timestamp on update current_timestamp
)

-- 商品
create table if not exists products (
  id          bigint unsigned auto_increment primary key,
  title       varchar(100) not null,
  description varchar(2000) not null,
  price       int unsigned not null,
  created_at  datetime not null default current_timestamp,
  updated_at  datetime not null default current_timestamp on update current_timestamp,
  shop_id     bigint unsigned references shops(id)
)
