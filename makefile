TOOLDIR=TestingTools

all:
	go build main.go student.go couglinks.go

tools: adduser getusers adddummy

adduser:
	go build -o $(TOOLDIR)/adduser ToolAdduser.go student.go couglinks.go 

getusers:
	go build -o $(TOOLDIR)/getusers ToolGetlist.go student.go

adddummy:
	go build -o $(TOOLDIR)/dummy ToolAdddummy.go student.go couglinks.go
