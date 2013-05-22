TOOLDIR=TestingTools

all:
	go build main.go student.go couglinks.go

tools: adduser getusers

adduser:
	go build -o $(TOOLDIR)/adduser ToolAdduser.go student.go couglinks.go 

getusers:
	go build -o $(TOOLDIR)/getusers ToolGetlist.go student.go
