-- +migrate Up
create table if not exists users (
  id uuid not null primary key,
  email varchar(191) not null unique,
  encrypted_password varchar(191) not null,
  screen_id varchar(191),
  screen_name varchar(191),
  created_at timestamp not null,
  updated_at timestamp not null
);

create index on users (email);

-- +migrate Down
drop table if exists users;
