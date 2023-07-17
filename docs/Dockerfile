# --- Build stage ---
FROM node:16-alpine AS build-stage

WORKDIR /app

# Install the Redocly CLI
RUN npm install -g @redocly/openapi-cli

# Copy documentation files
COPY ./docs ./docs

# Bundle the documentation into a static HTML file
RUN openapi bundle ./docs/app/swagger.yaml -o ./docs/redoc.html

# --- Production stage ---
FROM nginx:alpine AS production-stage

# Remove default nginx page
RUN rm -rf /usr/share/nginx/html/*

# Copy the static HTML file to the nginx html directory
COPY --from=build-stage /app/docs/redoc.html /usr/share/nginx/html/index.html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
