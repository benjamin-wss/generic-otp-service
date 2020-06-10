package constants

func init() {
	SupportedDbTypes = &supportedDbTypes{
		Postgres: "postgres",
	}

	GenericDbErrorMessages =
		&genericDbErrorMessages{
			ConflictingEntryDetected: `conflicting db row located`,
		}
}
