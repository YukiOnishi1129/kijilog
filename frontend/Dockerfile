FROM node:16.13.0-alpine3.12

WORKDIR /frontend
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile

EXPOSE 3000