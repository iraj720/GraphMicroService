package configs

const (
	// DefaultValues of configs
	DefaultValues = `
server:
  http_host: "127.0.0.1"
  http_port: "8080"
database:
  postgres_host: "172.18.0.125"
  postgres_port: "5432"
  postgres_dbname: "example"
  postgres_username: "example"
  postgres_password: "example"
assets:
  static_files_path: "./assets/"
  city_codes_path: "./assets/citycodes.xlsx"`
)
