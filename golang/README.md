
go mod init example/user/hello  

go install example/user/hello    

export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))


# APP IDEAS
- [ ] todo list
- [ ] todo list with daily, weekly, monthly, yearly, etc.
- [ ] movies and tv shows rankings, progress
- [ ] pokemon database, app, images
- [ ] rent tools app
