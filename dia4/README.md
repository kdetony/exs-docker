# Deploy app with Docker Swarm

this example two apis with proxy reverse (traefik)

## Create docker-compose

Remember,docker-compose first install.
```
docker stack deploy -c docker-compose.yml api
```

## List networks for Swarm
```
docker network ls 
```

## List services
```
docker stack ps api
```

## Validate proxy reverse

Open Browser and execute:
```
http://IP_SERVER_MANAGER:8080
```

now, execute:
```
http://IP_SERVER_MANAGER/books
```

```
`http://IP_SERVER_MANAGER/movies
```

## Scale service

Scaling apis: books and movies
```
docker service scale api_books=3

docker service scale api_movies=5
```

