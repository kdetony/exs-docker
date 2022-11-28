# Practica

1. Crear una red tipo de driver overlay, usar el nombre: webnet
2. Crear el servicio de nombre webapp con 3 replicas asociado a la red creada en el punto1, usar como imagen lherrera/webapp:1.0
3. Crear un servicio de nombre redisdb asociado a la red creada en el punto 1, usar como imagen redis:alpine
4. Validar que los aplicativos esten operativos
5. Realizar un scale en 4 replicas para el servicio webapp
