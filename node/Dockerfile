FROM ubuntu

LABEL version="1.0"
LABEL description="node sample"

RUN apt-get update \
    && apt-get install -y  nodejs \
    && apt-get install -y --no-install-recommends npm 

COPY . /proj

WORKDIR /proj

RUN npm install

EXPOSE 3000

ENV NODE_ENV=production 

ENTRYPOINT ["node", "./src/index.js"]
