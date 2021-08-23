delete from posts;
insert into posts values (
  gen_random_uuid(),
  (select id from users where screen_id = 'screen_id1'),
  'tweet1',
  now()
), (
  gen_random_uuid(),
  (select id from users where screen_id = 'screen_id1'),
  'tweet2',
  now()
), (
  gen_random_uuid(),
  (select id from users where screen_id = 'screen_id2'),
  'tweet3',
  now()
), (
  gen_random_uuid(),
  (select id from users where screen_id = 'screen_id3'),
  'tweet4',
  now()
);
