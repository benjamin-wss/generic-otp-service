package constants

func init() {
	SupportedDbTypes = &supportedDbTypes{
		Postgres: "postgres",
	}

	GenericDbErrorMessages =
		&genericDbErrorMessages{
			ConflictingEntryDetected: `conflictig db row located`,
		}
}
