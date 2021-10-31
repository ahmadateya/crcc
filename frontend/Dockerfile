FROM node:lts-alpine

RUN mkdir -p /nuxt-app/
WORKDIR /nuxt-app/

COPY ./package*.json /nuxt-app/
COPY . .

RUN npm install \
    && npm build

CMD ["npm", "start"]
