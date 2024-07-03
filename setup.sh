echo "Kindly enter your MySQL username: "
read username

echo "Kindly enter your MySQL password: "
read -s password

echo "Kindly enter your MySQL database name: "
read database

mysql -u $username -p $password < database/database.sql

echo "Database created successfully!"

go mod vendor
go mod tidy

echo "Kindly run 'go run main.go' to start the server"


