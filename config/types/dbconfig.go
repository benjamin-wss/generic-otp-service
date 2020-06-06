package types

type DbConfig struct {
	DbType               string
	DbHost               string
	DbPort               int
	DbName               string
	DbUser               string
	DbPassword           string
	DbPostgresSslSetting string
}
