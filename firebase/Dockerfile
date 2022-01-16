FROM node:16-buster

RUN apt-get update -y \
 && apt-get install -y openjdk-11-jre-headless \
 && npm install -g firebase-tools

WORKDIR /app

CMD ["firebase", "--version"]