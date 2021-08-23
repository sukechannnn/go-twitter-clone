delete from users;
insert into users values (
  gen_random_uuid(),
  'example1@email.com',
  -- Below hash means 'password'
  '$2a$10$IhhTPP22gUVZdKqssIRjJuc8invnzK.TFW0gW8HcWygyoDexHAvcC',
  'screen_id1',
  'screen_name1',
  now(),
  now()
), (
  gen_random_uuid(),
  'example2@email.com',
  -- Below hash means 'password'
  '$2a$10$IhhTPP22gUVZdKqssIRjJuc8invnzK.TFW0gW8HcWygyoDexHAvcC',
  'screen_id2',
  'screen_name2',
  now(),
  now()
), (
  gen_random_uuid(),
  'example3@email.com',
  -- Below hash means 'password'
  '$2a$10$IhhTPP22gUVZdKqssIRjJuc8invnzK.TFW0gW8HcWygyoDexHAvcC',
  'screen_id3',
  'screen_name3',
  now(),
  now()
);
