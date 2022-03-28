FROM node:16.14.0-alpine
LABEL maintainer="Memory Clutter <memclutter@gmail.com>"

ENV NUXT_HOST=0.0.0.0
ENV NUXT_PORT=3000

EXPOSE 3000

WORKDIR /var/app
COPY package.json .
COPY yarn.lock .
RUN yarn install && yarn cache clean
COPY . .
RUN yarn build

CMD ["yarn", "start"]
