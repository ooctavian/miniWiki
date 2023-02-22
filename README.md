# Requirements
- Go 1.19
- Docker

# Install dependencies
```shell
make install-deps
```

# Copy settings
All example settings are in `.env-example`. 
You should copy it to `.env` and customize it.

```shell
cp .env-example .env
```

# Start database
If you want to host a postgresql server using docker, run the following command:
```shell
make start-db
```
else edit `DATABASE_URL` in `.env` file to yours instance connection url.

# Seeding database
If you need to pollute your database with fake data for testing purpose, you can use the `polluter` command: 
```shell
make seed-db
```

# Run project 
```shell 
make run
```