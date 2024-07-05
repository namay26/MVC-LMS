echo "Kindly enter your MySQL username: "
read username

echo "Kindly enter your MySQL password: "
read -s password

echo "Kindly enter your MySQL database name: "
read database

mysql -u $username -p$password -e "CREATE DATABASE IF NOT EXISTS $database;"

echo "Database created successfully!"


echo "Enter Secret key for JWT: "
read secretKey

cat << EOF > db.yaml
DB_USERNAME: "$username"
DB_PASSWORD: "$password"
DB_HOST: 127.0.0.1:3306
DB_NAME: "$database"
JWTSecretKey: "$secretKey"
EOF

echo "Database configuration file created successfully!"

go mod vendor
go mod tidy

echo "Kindly change your directory to /cmd and run 'go run main.go' to start the server"