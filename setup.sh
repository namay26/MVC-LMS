#!/bin/bash

isInstalled() {
    if command -v "$1" &> /dev/null; then
        echo "$1 is installed."
    else 
        echo "$1 is not installed. Aborting..."
        exit 1
    fi
}

echo $'Checking for dependencies...\n'

isInstalled "go"
isInstalled "mysql"

echo $'Dependencies are installed.\n'

echo "Kindly enter your MySQL username: "
read username

echo "Kindly enter your MySQL password: "
read -s password

mysql -u "$username" -p"$password" -e "CREATE DATABASE IF NOT EXISTS mvc;"

echo "Created database 'mvc' successfully!"

mysql -u "$username" -p"$password" mvc < database.sql

echo "Database setup successful!"


echo "Enter Secret key for JWT: "
read secretKey

cat << EOF > db.yaml
DB_USERNAME: "$username"
DB_PASSWORD: "$password"
DB_HOST: 127.0.0.1:3306
DB_NAME: mvc
JWTSecretKey: "$secretKey"
EOF

echo "Database configuration file created successfully!"

go mod vendor
go mod tidy

echo "Kindly run the following command to start the server: go run cmd/main.go"