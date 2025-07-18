#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE DATABASE event_sourcing_user;
    CREATE DATABASE event_sourcing_payment;
EOSQL