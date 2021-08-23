# go-twitter-clone

## Setup

### Setup DB

```
# Setup DB
$ docker-compose up -d
$ bin/setup_db
```

### Up server

Open terminal and execute below.

```
$ cd server
$ go run server.go

# Install sql-migrate before execute below
$ sql-migrate up

# Install seed data
$ ../bin/seed_db
```

### Up client

Open another terminal and execute below.

```
$ cd client
$ npm run dev
```

### Use with test user

Access http://localhost:3000/sign_in and enter sign_in info (this is seed data),

- email: `example1@email.com`
- password: `password`

then you can see tweet.

## やり残したこと

- フォローしてるユーザーと、フォロワーの一覧表示の実装（GraphQL のクエリとしては実装済み）
- unfollow 機能の実装
- ログアウト機能の実装
- ログイン機能をセッション ID を使うようにする
- Dockerize する
- Monorepo っぽく build できるようにする（makefile + shell）
- Client のホスティング
