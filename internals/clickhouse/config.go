package clickhouse

type Config struct {
	Host string
	Port int

	Database string
	Username string
	Password string
}
