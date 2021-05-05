# Database API

This is database API wrapping a PostgreSQL database for CRUD operations.

The architecture consists of 2 components:

- Database API application

- PostgreSQL as database itself

## Installation

As a first step of the installation, the environment file should be setup. The necessary fields are defining the variables to build the database connections string.

```conf
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=some_database_user
DATABASE_PASS=some_database_password
DATABASE_NAME=this_is_the_database_name
```

## Running (manually)

TBA

## Running (as docker container)

In order to run the application as docker container, `docker-compose` file can be used with the following commands:

1. `docker-compose up --build -d`

The command above will start 2 docker containers, one for app and one for the database.

## TODO List for future works

- TBA