package generator

import (
	"os"
	"log"
	"fmt"
	"strings"
	"github.com/valuppo/dotgo/util"
)

//Init function used for initializing dotgo project
func Init(project string, path string){
	var err error

	//if no path specified use current directory as path
	if(path == ""){
		path = "."
	}

	rootPath := path + "/" + project + "/"

	util.ImportPathInstance().Set(project)

	if _, err := os.Stat(rootPath); !os.IsNotExist(err) {
		var overwrite string
		fmt.Println("Project folder with name '" + project +"' is already exist")
		fmt.Print("Do you want to overwrite the existing folder? (y/n) ")
		fmt.Scanln(&overwrite)
		if strings.ToLower(overwrite) != "y" {
			return
		}
	}

	//make project root folder
	err = os.MkdirAll(rootPath, 0775)
	if err != nil {
		log.Fatalln(err)
	}
	//make sample main go
	Main(rootPath)

	//make project route folder
	err = os.MkdirAll(rootPath + "route/", 0755)
	if err != nil{
		log.Fatalln(err)
	}
	//make sample route file
	Route(rootPath)

	//make project model folder
	err = os.MkdirAll(rootPath + "model/" , 0755)
	if err != nil{
		log.Fatalln(err)
	}
	//make sample model file
	Model("welcome", rootPath)

	//make project controller folder
	err = os.MkdirAll(rootPath + "controller/", 0755)
	if err != nil{
		log.Fatalln(err)
	}
	//make sample controller file
	Controller("welcome", rootPath)

	//make project db folder
	err = os.MkdirAll(rootPath + "db/", 0755)
	if err != nil{
		log.Fatalln(err)
	}
	//make sample controller file
	DB(rootPath)

	fmt.Println("Success generating new Dot Go project")
	fmt.Println("How to use : ")
	fmt.Println("cd " + project)
	fmt.Println("go run main.go")

}