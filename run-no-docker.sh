#!/bin/sh
export DB_USER="wishr" \
DB_PASSWORD="wishr" \
DB_DATABASE="wishr" \
DB_HOST="localhost" \
DB_PORT="3306" \
REGISTRATION_ENABLED="TRUE" \
REGISTRATION_CODE="register-test"; \
go build && ./wishr-api