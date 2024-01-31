
go mod init example/user/hello  

go install example/user/hello    

export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))


