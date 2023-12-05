package cmdarguments

import (
	"flag"
	"rv4-request/io"
)

func ParseArguments() (string, string) {
	var maillistPath, campaignName string
	// Flag -ml to read from csv file with a given name. Default value is "maillist.csv"
	flag.StringVar(&maillistPath, "ml", "", "Name of the file to read from. No default value.")

	// Flag -camp to choose campaign with a given name. No default value.
	flag.StringVar(&campaignName, "camp", "", "Name of the campaign. No default value.")

	// Parse arguments.
	flag.Parse()

	// Assuming we need both arguments to not be empty.
	condition := campaignName != "" && maillistPath != ""
	if condition {
		io.ReadCsvFile(maillistPath)
	} else {
		// Stdout -h in case of error or wrong arguments.
		flag.Usage()
	}
	return maillistPath, campaignName
}
