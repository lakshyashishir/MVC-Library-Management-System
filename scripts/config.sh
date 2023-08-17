#!/bin/bash

read -p "Enter your MySQL Username: " DB_USERNAME
read -s -p "Enter your MySQL Password: " DB_PASSWORD
echo
read -p "Enter your MySQL Host: " DB_HOST
read -p "Enter your MySQL Port: " DB_PORT
read -p "Enter your Database Name: " DB_NAME

echo 
yaml_content=$(cat <<EOF
dbUsername: "$DB_USERNAME"
dbPassword: "$DB_PASSWORD"
host: "$DB_HOST"
port: "$DB_PORT"
dbName: "$DB_NAME"
EOF
)

echo "$yaml_content" > config.yaml

migrate -path database/migration/ -database "mysql://$DB_USERNAME:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$DB_NAME" force 1
migrate -path database/migration/ -database "mysql://$DB_USERNAME:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$DB_NAME" -verbose up

echo "Congratulations! The yaml config file and database setup completed successfully."