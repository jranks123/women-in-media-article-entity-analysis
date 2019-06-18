package main

import (
	"flag"
)

var (
	missing  = flag.Bool("missing", false, "export missing data")
	from     = flag.String("from", "", "lower bound of date range to get the contributions for")
	to       = flag.String("to", "", "upper bound of date range to get the contributions for")
	user     = flag.String("user", "", "database user to execute queries against")
	password = flag.String("password", "", "password for database user")
	port     = flag.Int("port", 5432, "port of payment API database - 5431 for PROD / 5432 for CODE")
	stepMode = flag.Bool("step-mode", true, "determines where user in put is required before side effecting operations")
	bucket   = flag.String("bucket", "contributions-store-export-code", "S3 bucket in membership account to upload the contributions data to")
	profile  = flag.String("profile", "membership", "AWS Profile to extract credentials from the shared credentials file.")
)

func main() {

}
