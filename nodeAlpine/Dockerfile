FROM node:alpine

LABEL version="1.0"
LABEL description="node alpine"

COPY . /proj

WORKDIR /proj

RUN npm install

EXPOSE 3000

ENV NODE_ENV=production 

ENTRYPOINT ["node", "./src/index.js"]
