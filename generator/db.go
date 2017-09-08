package generator

import (
	"go/build"
	"io/ioutil"
	"log"
)

func DB(path ...string) {
	dbf, err := ioutil.ReadFile(build.Default.GOPATH + "/src/github.com/valuppo/dotgo/layout/db/db.dotgo")
	if err != nil {
		log.Fatalln(err)
	}

	udbf, err := ioutil.ReadFile(build.Default.GOPATH + "/src/github.com/valuppo/dotgo/layout/db/util.dotgo")
	if err != nil {
		log.Fatalln(err)
	}

	dbPath := ""
	if len(path) > 0 {
		dbPath = path[0] + "db/"
	} else {
		dbPath = "./db/"
	}

	err = ioutil.WriteFile(dbPath+"db.go", dbf, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	err = ioutil.WriteFile(dbPath+"util.go", udbf, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
