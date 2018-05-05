package bean

// ServerConfig ServerConfig Struct
type ServerConfig struct {
	Port        []string
	Mode        string
	GormLogMode string
	ViewLimit   int
}

// DBConfig DBConfig Struct
type DBConfig struct {
	Dialect      string
	Database     string
	User         string
	Password     string
	Host         string
	Port         int
	Charset      string
	URL          string
	MaxIdleConns int
	MaxOpenConns int
}
