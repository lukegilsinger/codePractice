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

# TODO
- [x] logging
- [x] frontend for user auth
- [ ] validation
- [ ] error handling
- [ ] testing
- [ ] drag and drop
- [x] docker
- [x] database migrations
- [ ] postgres
- [ ] ci/cd - github actions
- [ ] deploy
- [ ] file attachments
- [ ] 
- [ ] 
- [ ] 
- [ ] 