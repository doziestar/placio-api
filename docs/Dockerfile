# Pull the pre-built Redoc Docker image
FROM redocly/redoc AS redoc-stage

# The port on which Redoc will run
EXPOSE 80

# Set the environment variable for your OpenAPI definition URL
ENV SPEC_URL=https://docs.palnight.com/openapi.json

# --- Production stage ---
FROM nginx:alpine AS production-stage

# Copy Redoc files from the Redoc stage to the Nginx HTML directory
COPY --from=redoc-stage /usr/share/nginx/html /usr/share/nginx/html

# Remove default Nginx page
RUN rm -rf /usr/share/nginx/html/*

# Expose the Nginx port
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
