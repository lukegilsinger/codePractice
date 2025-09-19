# How to use
`make start`

## Docker Commands
`docker build -t todo-app .`
`docker run -p 8080:8080 -e JWT_SECRET=local-dev-secret todo-app`
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
- [ ] database migrations
- [ ] ci/cd
- [ ] 
- [ ] 
- [ ] 
- [ ] 
- [ ] 
- [ ] 