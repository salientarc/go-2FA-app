# Golang 2FA App

### Stuff Used

- golang
- postgresql

## Setup

```sh
git clone https://github.com/salientarc/go-2fa-app.git

cd go-2fa-app

go mod tidy

go run .
```

Use `setup.sql` to create database

```sh
cd go-2fa-app

psql -U postgres

\i setup.sql # in postgresql shell
```

`.env` Example

```

DBHOST=localhost
DBUSER=
DBPASSWD=
DBNAME=g2fa
DBPORT=5432
SSLMODE=disable
TIMEZONE=Asia/Kolkata

ISSUERDOMAIN=domain.com
ACCOUNTNAME=admin@example.com

```

## [API Route Examples](./docs/Routes.md)
