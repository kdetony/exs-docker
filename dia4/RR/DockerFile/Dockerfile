FROM mhart/alpine-node:6.11.4
MAINTAINER Anthony Mogrovejo "tony001983@gmail.com"
RUN apk update && apk add curl

HEALTHCHECK --interval=5s --timeout=3s --retries=3 \
CMD curl -f http://localhost:8000/whoami || exit 1

COPY package.json /app/package.json
WORKDIR /app
RUN npm install
COPY . /app

CMD ["npm", "start"]
