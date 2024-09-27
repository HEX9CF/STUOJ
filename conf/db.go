package conf

import(
	"strconv"
	_ "github.com/joho/godotenv"
	"os"
)

type DatabaseConfig struct {
	Host    string
	Port    string
	Name    string
	User   string
	Pwd    string
	MaxConn int
	MaxIdle int
}

func DatabaseConfigFromEnv() DatabaseConfig {
	// err := godotenv.Load()
	// if err != nil {
	//   log.Fatal(err)
	// }
	// fmt.Println("Loaded .env file")

	MaxConn,_:=strconv.Atoi(os.Getenv("DB_MAXOPENCONNS"))
	MaxIdle,_:=strconv.Atoi(os.Getenv("DB_MAXIDLECONNS"))

	return DatabaseConfig{
		Host:os.Getenv("DB_HOST"),
		Port:os.Getenv("DB_PORT"),
		Name:os.Getenv("DB_NAME"),
		User:os.Getenv("DB_USER"),
		Pwd:os.Getenv("DB_PASSWORD"),
		MaxConn:MaxConn,
		MaxIdle:MaxIdle,
	}
}