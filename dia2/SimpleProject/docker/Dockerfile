FROM php:7-apache
MAINTAINER Anthony Mogrovejo "tony001983@gmail.com"
RUN apt-get update -y && apt-get install -y libpng-dev curl libcurl4-openssl-dev
RUN docker-php-ext-install pdo pdo_mysql gd curl
RUN a2enmod rewrite
RUN service apache2 restart
