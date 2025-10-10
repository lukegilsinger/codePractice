# How to use
`make start`



## Docker Commands
`docker build -t todo-app .`
`docker run -p 8080:8080 -e JWT_SECRET=local-dev-secret todo-app`

`docker run --name todo_postgres -e POSTGRES_PASSWORD=password -d -p 5431:5432 postgres`

## Docker Compose
`docker-compose up --build`
`docker-compose down`

## Note
* If schema changes, then you need to remove database.db file

# BUGS
update status

# TODO
- [x] logging
- [x] frontend for user auth
- [x] docker
- [x] database migrations
- [x] postgres
- [x] priority and freq
- [ ] repeated tasks
- [ ] seperate front end
- [ ] history table
- [ ] validation
- [ ] error handling
- [ ] testing
- [ ] drag and drop
- [ ] ci/cd - github actions
- [ ] deploy
- [ ] file attachments
- [ ] images
- [ ]
- [ ] 
- [ ]
- [ ]
- [ ]
- [ ] 