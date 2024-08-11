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


