# Multibuild

Multistage/Multibuild nos va a permitir construir imagenes sobre imagenes, esto quiere decir, que podemos usar una imagen base para construir un artefacto y luego usar otra imagen para copiar el artefacto creado, y todo esto bajo un mismo Dockerfile, vamos a crear una carpeta de nombre multibuild y dentro de ella crearemos los ficheros main.go y Dockerfile, asumamos que tenemos una pequeña aplicación desarrollada en Go!.

## First Run
```
docker build -t imgo .
```

## Second Run
```
docker build -t imgo1 .
```

## Third Run
```
docker build -t img2 ..
```

The result for size image *img2: 12MB*  aprox.


