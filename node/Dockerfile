FROM ubuntu

LABEL version="1.0"
LABEL description="node sample"

RUN apt-get update && apt-get install nodejs npm -y

COPY . /proj

WORKDIR /proj

RUN npm install

EXPOSE 3000

ENV NODE_ENV=production 

ENTRYPOINT ["node", "./src/index.js"]
