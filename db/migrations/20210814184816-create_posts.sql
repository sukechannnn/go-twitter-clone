-- +migrate Up
create table if not exists posts (
  id uuid not null primary key,
  user_id uuid not null,
  text text not null,
  created_at timestamp not null,
  foreign key (user_id) references users(id) on delete cascade on update cascade
);

create index on posts (user_id);

-- +migrate Down
drop table if exists posts;
