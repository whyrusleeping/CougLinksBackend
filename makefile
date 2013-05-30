TOOLDIR=TestingTools

all:
	go build main.go student.go couglinks.go

tools: adduser getusers adddummy deleteuser login

adduser:
	cp $(TOOLDIR)/ToolAdduser.go .
	go build -o $(TOOLDIR)/adduser ToolAdduser.go student.go couglinks.go 
	rm ToolAdduser.go

getusers:
	cp $(TOOLDIR)/ToolGetlist.go .
	go build -o $(TOOLDIR)/getusers ToolGetlist.go student.go
	rm ToolGetlist.go

adddummy:
	cp $(TOOLDIR)/ToolAdddummy.go .
	go build -o $(TOOLDIR)/dummy ToolAdddummy.go student.go couglinks.go utilities.go
	rm ToolAdddummy.go

deleteuser:
	cp $(TOOLDIR)/ToolDeleteuser.go .
	go build -o $(TOOLDIR)/delete ToolDeleteuser.go student.go couglinks.go
	rm ToolDeleteuser.go

login:
	cp $(TOOLDIR)/ToolLoginuser.go .
	go build -o $(TOOLDIR)/login ToolLoginuser.go couglinks.go student.go
	rm ToolLoginuser.go
