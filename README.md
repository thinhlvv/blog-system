# blog-system

[![CircleCI](https://circleci.com/gh/thinhlvv/blog-system/tree/master.svg?style=svg&circle-token=47214c8eef0e763d2f4d529b884f0e9cc0141217)](https://circleci.com/gh/thinhlvv/blog-system/tree/master)
[![Coverage Status](https://coveralls.io/repos/github/thinhlvv/blog-system/badge.svg?branch=master&service=github&kill_cache=1)](https://coveralls.io/github/thinhlvv/blog-system?branch=master&kill_cache=1)

Backend serves APIs for blog system.

## Commands

```bash
# Make copy of the environment for database in development.
$ cp .env.sample .env.development

# Make copy of the environment for database in staging.
$ cp .env.sample .env.staging

# Install all dependencies.
$ make install

# Start docker services.
$ docker-compose up -d

# Run migration files (locally), default is set to development.
$ make migrate

# Rollback migration version
$ make rollback

# Run migration files (on staging).
$ make migrate ENV=staging

# Run test.
$ make test

# Start the development server.
$ make start

# Cleanup local database.
$ make clean
```

## Installation

Go version : 1.12.6

1. Run database for application and testing:
```bash
docker-compose up -d
```

Make sure you got two docker database:
```bash
docker ps 
```

2. Install dependencies:
```bash
make install
```

3. Migrate database:
```bash
make migrate
```

4. Start server
```bash
make start
```
