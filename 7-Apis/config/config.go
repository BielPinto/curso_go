package configs

var cfg *conf

type conf struct {
	DBDriver     string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBServerPort string
	DBSecret     string
	DBExperesIn  string
}

func LoadConfig(path string) (*conf, error) {

}
