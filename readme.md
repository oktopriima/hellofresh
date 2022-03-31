
# Hellofresh

Take home exercise from Hellofresh
- Language : Golang 1.17
- Framework : Chi
- ORM : sqlx
- Database : MySQL

## Run Locally

Clone the project

```bash
  git clone https://github.com/oktopriima/hellofresh
```

Go to the project directory

```bash
  cd $GOPATH/src/gitthub.com/oktopriima/hellofresh
```

Setup env
```bash
  cp .env.example .env
  # change CONFIG_PATH to your root directory of this project
  # change CONFIG_ENV to "local"

  cp app-example.yaml app-local.yaml
  # adjust the environment value with your local computer
```

Migrate database
```bash
  # install rubenv migration tools
  go get -v github.com/rubenv/sql-migrate/... 

  # only for the migration
  # because I use different tools for migration
  cp dbconfig-example.yml dbconfig.yml
  # adjust the database connection

  sql-migrate up --env=local
```

Install dependencies

```bash
  go mod tidy -compact=1.17
  go mod vendor
```

Start the server

```bash
  go build
  ./hellofresh
```

## Authors

- [@oktopriima](https://www.github.com/oktopriima)