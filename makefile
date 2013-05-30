TOOLDIR=TestingTools

all:
	go build main.go student.go couglinks.go

tools: adduser getusers adddummy deleteuser login

adduser:
	go build -o $(TOOLDIR)/adduser ToolAdduser.go student.go couglinks.go 

getusers:
	go build -o $(TOOLDIR)/getusers ToolGetlist.go student.go

adddummy:
	go build -o $(TOOLDIR)/dummy ToolAdddummy.go student.go couglinks.go

deleteuser:
	go build -o $(TOOLDIR)/delete ToolDeleteuser.go student.go couglinks.go

login:
	go build -o $(TOOLDIR)/login ToolLoginuser.go couglinks.go student.go
