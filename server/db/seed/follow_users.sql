delete from follow_users;
insert into follow_users values (
  gen_random_uuid(),
  (select id from users where screen_id = 'screen_id1'),
  (select id from users where screen_id = 'screen_id2'),
  now()
), (
  gen_random_uuid(),
  (select id from users where screen_id = 'screen_id1'),
  (select id from users where screen_id = 'screen_id3'),
  'tweet2',
  now()
), (
  gen_random_uuid(),
  (select id from users where screen_id = 'screen_id2'),
  (select id from users where screen_id = 'screen_id3'),
  now()
), (
  gen_random_uuid(),
  (select id from users where screen_id = 'screen_id3'),
  (select id from users where screen_id = 'screen_id1'),
  now()
);
