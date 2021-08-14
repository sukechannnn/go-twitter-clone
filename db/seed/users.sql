delete from users;
insert into users values (
  gen_random_uuid(),
  'example1@email.com',
  -- Below hash means 'password'
  '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8',
  'screen_id1',
  'screen_name1',
  now(),
  now()
), (
  gen_random_uuid(),
  'example2@email.com',
  -- Below hash means 'password'
  '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8',
  'screen_id2',
  'screen_name2',
  now(),
  now()
), (
  gen_random_uuid(),
  'example3@email.com',
  -- Below hash means 'password'
  '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8',
  'screen_id3',
  'screen_name3',
  now(),
  now()
);
