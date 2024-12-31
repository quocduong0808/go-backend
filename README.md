# go-backend
go-backend

# docker
# build
> docker build . -t go-app
## net-work
> docker network create go-network
> docker network inspect go-network
## redis
> docker run -d --name redis-stack --network go-network -p 6379:6379 -p 8001:8001 redis/redis-stack
## mysql
> docker run -d --name mydbdev --network go-network -v mydbdev:/var/lib/mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root123 mysql <br/>
> docker exec -it mydbdev mysql -u root -proot123 <br/>
> create user 'dev'@'%' identified by 'dev123'; <br/>
> create database godbdev; <br/>
> grant all privileges on godbdev.*  to 'dev'@'%'; <br/>
> flush privileges; <br/>
## app
> docker run -d --name go-app --network go-network -p 8080:8080 go-app
## start apps
> docker-compose --env-file .env_dev up -d --build
