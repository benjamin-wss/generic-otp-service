package enums

type DbConfigVariableKeyEnum int

const (
	DbType               DbConfigVariableKeyEnum = 1
	DbHost               DbConfigVariableKeyEnum = 2
	DbPort               DbConfigVariableKeyEnum = 3
	DbName               DbConfigVariableKeyEnum = 4
	DbUser               DbConfigVariableKeyEnum = 5
	DbPassword           DbConfigVariableKeyEnum = 6
	DbPostgresSslSetting DbConfigVariableKeyEnum = 7
)

func (enum DbConfigVariableKeyEnum) ToString() string {
	switch enum {
	case DbType:
		return "DB_TYPE"
	case DbHost:
		return "DB_HOST"
	case DbPort:
		return "DB_PORT"
	case DbName:
		return "DB_NAME"
	case DbUser:
		return "DB_USER"
	case DbPassword:
		return "DB_PASSWORD"
	case DbPostgresSslSetting:
		return "DB_POSTGRES_SSL_SETTING"
	default:
		panic("DbConfigVariableKeyEnum : Unexpected value requested/assigned.")
	}
}
