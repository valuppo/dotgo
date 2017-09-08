# Dot Go Simple MVC Generator

Simple MVC Generator

## Getting Started

### Prerequisites

- sqlq - github.com/valuppo/sqlq

```
go get github.com/valuppo/sqlq
```

### Installing

- Get this library using go get

```
go get github.com/valuppo/dotgo
```

- Install this library

```
cd $GOPATH/src/github.com/valuppo/dotgo
go build
go install
```


### How To Use
- Init new project <br />
  Required : PROJECT_NAME <br />
  Optional : PROJECT_PATH (if not specified it will used current directory '.' as path)

```
dotgo init <PROJECT_NAME> [PROJECT_PATH]
```

- Create new Model <br />
  Required : MODEL_NAME

```
cd <PROJECT_DIR>
dotgo model <MODEL_NAME>
```

- Create new Controller <br />
  Required : CONTROLLER_NAME

```
cd <PROJECT_DIR>
dotgo controller <CONTROLLER_NAME>
```

## Built With

* [Golang](https://golang.org/) - Programming Language
* [Sqlq](https://github.com/valuppo/sqlq) - SQL Query Builder

## Authors

* **Antony Gunawan** - *Initial work* - [valuppo](https://github.com/valuppo)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
