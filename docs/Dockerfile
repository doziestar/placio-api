FROM node:18.0.1-alpine3.14 AS builder

WORKDIR /app

RUN npm install @redocly/cli

COPY ./docs ./docs

RUN npx redocly bundle ./docs/app/swagger.yaml --output ./docs/placio.yml

CMD ["redocly", "preview-docs", "./docs/placio.yml"]