# curso_go 
# Course: Pos-Go Expert - 2024

This project contains the entire history of Pos-GO’s activities and challenges.

On the 0-fullcycle_challenges page you can find all the challenges.

from folders 1- follow the flow of content and small exercises.


go mod init  github.com/BielPinto/curso_go/4-database/3-advanc-GORM # go mod init {adress} 

 for API struct folder
  https://github.com/golang-standards/project-layout



//tests

# run
go test .
 # run verbose
 go test -v
 # run with coverage
 go test -coverprofile=coverrage.out

# find point that are not covered
go tool cover -html=coverage.out

# run benchmark  
go help test
go test -bench=.
go test -bench=. -run=^#    
go test -bench=. -run=^# -benchmem
go test -bench=. -run=^# -count=3
go test -bench=. -run=^# -count=3 -benchtime=3s

# run fuzz
go test -fuzz=. -run=^#
go test -fuzz=. -fuzztime=5s -run=^#


# sqlite3
  # run server and database from  module 7-api
  Ex: go run cm/server/main 
  # acess dataase sqlite3
  sqlite3 test.db 
  .help;
  .quit
  select * from products;



# Web frameoWorks Vs framworks
 # Web frameworks
  Typically thess web framworks wor focused on http and websockets.

  - Golang Echo -  https://echo.labstack.com/
    - works focused on  http  
    - very good to work events minimalista

  - Go Fiber -  https://gofiber.io/
    - 

  - GI - https://gin-gonic.com/
    -

#  frameworks
 - go buffalo -  https://gobuffalo.io/pt/
  - similar to la varal, it creates the entire ecosystem to be faster, very little used in the market, it changes the way you develop, create a folder structure
  - it covers the entire layer of your project, backend and frontend
 
 - iris -  https://github.com/kataras/iris

 #  web tookit
  A helpful toolkit for the Go programming language that provides useful, composable packages for writing HTTP-based applications.

  - Gorilla - https://gorilla.github.io/

 #  Router 
  - gorilla/mux - https://github.com/gorilla/mux

  - oa chi -  https://github.com/go-chi/chi



# Documentation
 - Doc swag : https://github.com/swaggo/swag
 - Write in main with the documentation pattern
 - swag init -g cmd/server/main.go
 - swag fmt


# Outher

- import path golang
 export PATH=$(go env GOPATH)/bin:$PATH

 # Apache Benchmark
  - Doc : https://www.digitalocean.com/community/tutorials/how-to-use-apachebench-to-do-load-testing-on-an-ubuntu-13-10-vps


  ab -n 1000 -c 100 http://localhost:3000/

# Race condition  GO
 -  It is a Golang feature to check concurrency issues in the application.
 go run -race main.go

# Go Private
- add on golang repository private
go env | grep PRIVATE
export GOPRIVATE=github.com/devfullcycle/fcutils-secret,other_repositori_private

- Add token/login  on .git
  - on file ~/.netrc insert line down
    `
    machine github.com
    login gabrielPinto
    password {token_gerado_no_github}
  `
 - for bitbucket.com add line below
   machine api.bitbucket.org

  - login edite  local .git/config   or default    ~/.gitconfig
     [url "ssh://git@github.com/"]
        insteadOf = https://github.com/

  - comand to show packeg go on cache 
    ls ~/go/pkg/mod
  
  - Go proxy
    https://proxy.golang.org
  
  - Go vendor
    go mod vendor


  - Graphql
    - url - https://gqlgen.com
    - It is a commn client-server "cpc" call, but sent in format that the server can understand and bring only the fields that are requested
     It is widely used as a front-end back-end
