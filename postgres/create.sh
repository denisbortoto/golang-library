 #!/bin/bash

set -e
set -u

function create_user_and_database() {
  local database=$1
  echo "Creating user and database '$database'"
  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE DATABASE $database;
    GRANT ALL PRIVILEGES ON DATABASE $database TO $POSTGRES_USER;
EOSQL
}

function create_schemas() {
  local database=$1
  for s in $(echo $POSTGRES_SCHEMAS | tr ',' ' '); do
    echo "Creating schema '$s'"
    psql -v --username "$POSTGRES_USER" -d $database <<-EOSQL
      CREATE SCHEMA $s;
EOSQL
  done
}

echo "database creation requested: $POSTGRES_DATABASE"
  for db in $(echo $POSTGRES_DATABASE | tr ',' ' '); do
    create_user_and_database $db
     create_schemas $db
  done
echo "database created"