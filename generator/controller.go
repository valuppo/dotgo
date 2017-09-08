package generator

import (
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"github.com/valuppo/dotgo/util"
)

func Controller(name string, path ...string) {
	importPath := util.ImportPathInstance().Get()
	structName := name
	funcName := strings.Title(name)

	//open controller layout files
	//cf = ControllerFile Layout
	cf, err := ioutil.ReadFile(build.Default.GOPATH + "/src/github.com/valuppo/dotgo/layout/controller/controller.dotgo")
	if err != nil {
		log.Fatalln(err)
	}

	//cast []bytes of controller file into string
	cs := string(cf)

	reImportPath := regexp.MustCompile(`{importPath}`)
	cs = reImportPath.ReplaceAllString(cs, importPath)

	//substitute all {tableName}
	reTableName := regexp.MustCompile(`{funcName}`)
	cs = reTableName.ReplaceAllString(cs, funcName)

	//substitute all {structName}
	reStructName := regexp.MustCompile(`{structName}`)
	cs = reStructName.ReplaceAllString(cs, structName)

	controllerPath := ""
	if len(path) > 0 {
		controllerPath = path[0] + "controller/"
	} else {
		controllerPath = "./controller/"
	}

	//make controller dir if not exist
	if _, err := os.Stat(controllerPath); os.IsNotExist(err) {
		os.MkdirAll(controllerPath, 0775)
	}

	//write controller layout file
	err = ioutil.WriteFile(controllerPath+name+".go", []byte(cs), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
