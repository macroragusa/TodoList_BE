package config

var ConnectionString string
var DebugMode bool

func LoadConfig() {
	// todo import from conf.ini and env variables
	ConnectionString = dbConfig()
	DebugMode = true
}
