#!/bin/sh
export DB_USER="wishr" \
DB_PASSWORD="wishr" \
DB_DATABASE="wishr" \
DB_HOST="localhost" \
DB_PORT="3306" \
REGISTRATION_ENABLED="TRUE" \
USE_SECURE_COOKIE="FALSE" \
REGISTRATION_CODE="register-test" \
HOST_NAME="https://your-wishr-url.com" &&
go build && ./wishr-api