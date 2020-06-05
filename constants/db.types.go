package constants

type supportedDbTypes struct {
	Postgres string
}

var SupportedDbTypes *supportedDbTypes

func init() {
	SupportedDbTypes = &supportedDbTypes{
		Postgres: "postgres",
	}
}
