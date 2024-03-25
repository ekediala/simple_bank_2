package database

var DatabaseConfig = struct {
	DB_DRIVER string
	DB_SOURCE string
}{
	DB_DRIVER: "postgres",
	DB_SOURCE: "postgresql://user:password@localhost:5432/simple_bank?sslmode=disable",
}
