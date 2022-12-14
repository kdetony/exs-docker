# Routing Mesh

Por defecto todas las redes en Swarm usan encryptacion.
Adicionalmente se puede agregar una capa "extra" de encriptacion para una red NUEVA del tipo OVERLAY para los nodos.

## Creamos la siguiente red:
``` 
docker network create --driver overlay --subnet 10.0.9.0/24 --opt encrypted net_over
```

## Listamos redes:
```
docker network ls
```

## Vamos a crear un servicio:
``` 
docker service create  --network net_over nginx
```

OBS.

si no creamos un servicio con un nombre, tendremo un servicio con un nombre aleatorio

## Vamos ahora a escalar el servicio creado:
```
docker service scale  jolly_yonath=3
```

donde : 

jolly_yonath es el nombre del servicio


Y ahora vamos a actualizar el servicio con un puerto para exponer el servicio de la aplicacion creado:
```
docker service update --publish-add published=8090,target=80 jolly_yonath
```

Por ultimo, vamos a actualizar la cantidad de cpu para un servicio asociado a una aplicacion:
```
docker service update --limit-cpu 0.5 jolly_yonath
```

## Link:
```
https://docs.docker.com/engine/reference/commandline/service_update/
```
```
https://gdevillele.github.io/engine/reference/commandline/service_update/
```
