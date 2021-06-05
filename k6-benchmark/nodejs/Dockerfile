FROM node:14

WORKDIR /usr/src/app

COPY package.json index.js yarn.lock ./

RUN yarn install

CMD ["node", "index.js"]
