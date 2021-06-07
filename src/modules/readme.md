Use this link for modules\
https://golang.org/doc/tutorial/call-module-code

go mod init example.com/hello\
go mod edit -replace=example.com/greetings=../greetings\
go mod tidy\
go run .\
