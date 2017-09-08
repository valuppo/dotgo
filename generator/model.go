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

func Model(name string, path ...string) {
	importPath := util.ImportPathInstance().Get()
	packageName := name
	tableName := name + "s"
	structName := name

	builderName := strings.Title(name)

	//open model layout files
	//sf = StructFile Layout
	sf, err := ioutil.ReadFile(build.Default.GOPATH + "/src/github.com/valuppo/dotgo/layout/model/struct.dotgo")
	if err != nil {
		log.Fatalln(err)
	}

	//bf = BuilderFile Layout
	bf, err := ioutil.ReadFile(build.Default.GOPATH + "/src/github.com/valuppo/dotgo/layout/model/builder.dotgo")
	if err != nil {
		log.Fatalln(err)
	}

	//cast []bytes of struct layout file & builder layout file into string
	ss := string(sf)
	bs := string(bf)

	//substitute all {packageName}
	rePackageName := regexp.MustCompile(`{packageName}`)
	ss = rePackageName.ReplaceAllString(ss, packageName)
	bs = rePackageName.ReplaceAllString(bs, packageName)

	reImportPath := regexp.MustCompile(`{importPath}`)
	ss = reImportPath.ReplaceAllString(ss, importPath)
	bs = reImportPath.ReplaceAllString(bs, importPath)

	//substitute all {tableName}
	reTableName := regexp.MustCompile(`{tableName}`)
	ss = reTableName.ReplaceAllString(ss, tableName)

	//substitute all {structName}
	reStructName := regexp.MustCompile(`{structName}`)
	ss = reStructName.ReplaceAllString(ss, structName)
	bs = reStructName.ReplaceAllString(bs, structName)

	//substitute all {builderName}
	reBuilderName := regexp.MustCompile(`{builderName}`)
	bs = reBuilderName.ReplaceAllString(bs, builderName)

	modelPath := ""
	if len(path) > 0 {
		modelPath = path[0] + "model/"
	} else {
		modelPath = "./model/"
	}

	//make directory with specify packagename if not exist
	if _, err := os.Stat(modelPath + packageName); os.IsNotExist(err) {
		os.MkdirAll(modelPath+packageName, 0775)
	}

	//write struct layout file
	err = ioutil.WriteFile(modelPath+packageName+"/struct.go", []byte(ss), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	//write builder layout file
	err = ioutil.WriteFile(modelPath+packageName+"/builder.go", []byte(bs), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
