package conf

import (
	"flag"
	"log"
)

type Conf struct {
	path     string
	logLevel bool
}

func LoadFlags() string {
	path := flag.String("file", "", "File to execute.")
	logLevel := flag.String("logLevel", "ERROR", "Set logger to specific level, only are availables theres values: DEBUG, INFO, ERROR. By default it will be configured to ERROR level.")

	flag.Parse()

	log.Println(*logLevel)
	log.Println(*path)
	return *path
}
