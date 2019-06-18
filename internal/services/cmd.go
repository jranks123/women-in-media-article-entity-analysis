package services

type DbParameters struct {
	DbName   string
	Host     string // database endpoint
	Port     int    // database port to connect to
	User     string // user to connect to the database with
	Password string // password for User
}

// Collection of parameters that are typically required when performing tasks.
type JobParameters struct {
	From    string // start date to run the job for
	To      string // end date to run the job for
	Section string
	Db      DbParameters
}
