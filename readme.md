# Go SQL Boiler Example

## Setup database

```
export DB_HOST=localhost \
    DB_PORT=3306 \
    DB_NAME=go_sqlboiler \
    DB_USER=user \
    DB_PASSWORD=usersecret
```

## Setup sqlboiler

[qlboiler](https://github.com/volatiletech/sqlboiler)

```
go get github.com/volatiletech/sqlboiler/v4
go get github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
go get github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
go get github.com/volatiletech/null/v8
```

Build sqlboiler and sqlboiler-mysql
Create sqlboiler.toml with database config
Generate models

```
go build
sqlboiler mysql -c sqlboiler.toml
```
