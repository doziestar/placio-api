# --- Build stage ---
FROM node:16-alpine AS build-stage

WORKDIR /app

# Install the Redocly CLI
RUN npm install -g @redocly/cli@latest

# Copy documentation files
COPY . .

# Split the API definition file
RUN #redocly split ./docs/app/swagger.json --outDir=./docs/split
RUN redocly split redoc.yml --outDir=./placio-api-docs/split

# Bundle the split files into a static OpenAPI file
RUN #redocly bundle ./docs/split/main.json -o ./docs/swagger.json
RUN redocly build-docs ./split/main.json -o ./docs/swagger.json

# Generate the documentation HTML file
RUN redocly build-docs ./docs/swagger.json --output ./docs/redoc-static.html

# --- Production stage ---
FROM nginx:alpine AS production-stage

# Remove default nginx page
RUN rm -rf /usr/share/nginx/html/*

# Copy the static HTML file to the nginx html directory
COPY --from=build-stage /app/docs/redoc-static.html /usr/share/nginx/html/index.html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
