ARG NODE_ENV=development
# ARG NODE_ENV=staging
ARG NODE_VERSION=19-alpine3.16

FROM node:${NODE_VERSION} AS node

WORKDIR /app

RUN npm install -g @redocly/cli

COPY ./docs ./docs

RUN redocly bundle ./docs/app/swagger.yaml --output ./docs/redoc.yml

CMD redocly preview-docs --host=0.0.0.0 --port=6060 ./docs/redoc.yml