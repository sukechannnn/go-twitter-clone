-- +migrate Up
create extension if not exists "pgcrypto";

create table if not exists users (
  id uuid primary key not null,
  email varchar(191) not null unique,
  encrypted_password varchar(191) not null,
  screen_id varchar(191) not null unique,
  screen_name varchar(191) not null,
  created_at timestamp not null,
  updated_at timestamp not null
);

create index on users (email);
create index on users (screen_id);

-- +migrate Down
drop table if exists users;
