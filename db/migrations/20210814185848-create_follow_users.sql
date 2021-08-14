-- +migrate Up
create table if not exists follow_users (
  id uuid primary key not null,
  user_id uuid not null,
  follower_id uuid not null,
  created_at timestamp not null,
  foreign key (user_id) references users(id) on delete cascade on update cascade,
  foreign key (follower_id) references users(id) on delete cascade on update cascade
);

create index on follow_users (user_id, follower_id);

-- +migrate Down
drop table if exists follow_users;
