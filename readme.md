go get github.com/gin-gonic/gin

1:To create a branch in github:
git checkout -b develop
this command creates a develop branch in github.

2.external package to create a unique identifier.
go get github.com/rs/xid

3.post request json object

{
"name": "Homemade Pizza",
"tags" : [
"italian",
"pizza",
"dinner"
],
"ingredients": [
"1 1/2 cups (355 ml) warm water (105°F-115°F)",
"1 package (2 1/4 teaspoons) of active dry yeast",
"3 3/4 cups (490 g) bread flour",
"feta cheese, firm mozzarella cheese, grated"
],
"instructions": [
"Step 1.",
"Step 2.",
"Step 3."
]
}

4. writing an open api specification is very important.
   for that we can use go swaggertool

how to install go swagger:

first clone it:

git clone https://github.com/go-swagger/go-swagger

then enter into the cloned repo and give the command below
go install ./cmd/swagger

to generate the swagger json give the below command
swagger generate spec -o ./swagger.json

to serve the swagger server :
swagger serve -F swagger ./swagger.json

http://localhost:49566

6. running the mongodb as a docker

docker run -d --name mongodb -v /home/data:/data/db -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=password -p 27017:27017 mongo:4.4.3

the above command will create the container.
once the container is created we can check the logs with the below command:

docker logs -f container-id

to stop a cotainer
docker stop container id

to remove a container:
docker rm -f container id

7. link to download the mongodb compass
8. https://www.mongodb.com/try/download/compass?tck=docs_compass

9. compass url

mongodb://admin:password@localhost:27017/admin

to install the mongodb driver in our gocode

go get go.mongodb.org/mongo-driver/mongo

pass environment variable along with go run

10. import json collection into mongo shortcut:

mongoimport --username admin --password password --authenticationDatabase admin --db demo --collection recipes --file recipes.json --jsonArray

redis is used to cache the data for faster response time of our api.
redis used the ram and we use an in-memory database for faster performaance.

to run the redis in docker:

docker run -d -v $pwd/conf:/usr/local/etc/redis --name redis -p 6379:6379 redis:6.0

to see container logs:

docker logs -f containerid

redis policy maxmemory-policy allkeys-lru
lru = least recently used

to install redis in gocode

go get github.com/go-redis/redis/v8

to verify the data being cached by redis:

docker exec -it container-id bash

after that

redis-cli

EXISTS recipes

to check if recipes exists

exit to leave the shell.

for redis gui

docker run -d --name redisinsight --link redis -p 8001:8001 redislabs/redisinsight

localhost:8001

jwts ---(javascript object notation) json web tokens

methods to protect the end points:

1. api keys
2. basic auth
3. client sessions
4. openid connect
5. openAuthorization

to integrate jwt token:

go get github.com/dgrijalva/jwt-go

to decode the jwt token

https://jwt.io/
