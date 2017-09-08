package generator

import (
	"go/build"
	"io/ioutil"
	"log"
	"github.com/valuppo/dotgo/util"
	"regexp"
)

func Route(path ...string) {
	importPath := util.ImportPathInstance().Get()

	rf, err := ioutil.ReadFile(build.Default.GOPATH + "/src/github.com/valuppo/dotgo/layout/route/route.dotgo")
	if err != nil {
		log.Fatalln(err)
	}

	routePath := ""
	if len(path) > 0 {
		routePath = path[0] + "route/"
	} else {
		routePath = "./route/"
	}

	rs := string(rf)

	reImportPath := regexp.MustCompile(`{importPath}`)
	rs = reImportPath.ReplaceAllString(rs, importPath)

	err = ioutil.WriteFile(routePath+"route.go", []byte(rs), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
