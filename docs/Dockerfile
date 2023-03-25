ARG NODE_ENV=development
# ARG NODE_ENV=staging
ARG NODE_VERSION=19-alpine3.16

FROM node:${NODE_VERSION} AS node

WORKDIR /app

RUN npm install @redocly/cli

COPY ./docs ./docs

RUN npx redocly bundle ./docs/app/swagger.yaml --output ./docs/placio.yml

CMD ["redocly", "preview-docs", "./docs/placio.yml"]