package setting

// DatabaseSettingS is the struct to init Database
type DatabaseSettingS struct {
	DBType       string
	UserType     string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}
