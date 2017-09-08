package generator

import (
	"go/build"
	"io/ioutil"
	"log"
	"regexp"
	"github.com/valuppo/dotgo/util"
)

func Main(path ...string) {
	importPath := util.ImportPathInstance().Get()

	mf, err := ioutil.ReadFile(build.Default.GOPATH + "/src/github.com/valuppo/dotgo/layout/main.dotgo")
	if err != nil {
		log.Fatalln(err)
	}

	mainPath := ""
	if len(path) > 0 {
		mainPath = path[0]
	} else {
		mainPath = "."
	}

	ms := string(mf)

	reImportPath := regexp.MustCompile(`{importPath}`)
	ms = reImportPath.ReplaceAllString(ms, importPath)


	err = ioutil.WriteFile(mainPath+"/main.go", []byte(ms), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
