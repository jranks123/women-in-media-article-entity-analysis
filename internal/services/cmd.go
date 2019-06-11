package services

type DbParameters struct {
	Host     string // database endpoint
	Port     int    // database port to connect to
	User     string // user to connect to the database with
	Password string // password for User
}

// Collection of parameters that are typically required when performing tasks.
type JobParameters struct {
	From         string // start date to run the job for
	To           string // end date to run the job for
	MissingDates bool   // run the S3 export job for missing dates; to be used instead of a date range
	StepMode     bool   // should user input be required before side-effecting code is executed?
	Email        string // email to perform a task for e.g. creating a guest account
	Db           DbParameters
}
