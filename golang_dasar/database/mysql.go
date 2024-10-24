package database

var connection string

func init() {
	connection = "MySQL"
}

func GetData() string {
	return connection
}