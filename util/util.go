package util

import (
	"os"
	"go/build"
	"log"
	"strings"
)

type importPath struct{
	path string
}

var ipInstance = importPath{}

func (ic *importPath) Set(projectName string){
	gopath := build.Default.GOPATH + "/"

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	pwd = strings.Replace(pwd, gopath, "", -1)
	pwd = strings.Replace(pwd, "src/", "", -1)

	if projectName != ""{
		pwd += "/" + projectName
	}

	ic.path = pwd
}

func (ic *importPath) Get() string{
	return ic.path
}

func ImportPathInstance() *importPath{
	return &ipInstance
}