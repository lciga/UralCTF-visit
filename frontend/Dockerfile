# Frontend image
FROM node:18-alpine as builder
WORKDIR /app
# Copy static files
COPY . .
# Install simple server
RUN npm install -g serve
# Final
FROM alpine:latest
RUN apk add --no-cache nodejs npm \
    && npm install -g serve
WORKDIR /app
COPY --from=builder /app .
# Expose port
EXPOSE 3000
# Run static server without SPA fallback to serve individual HTML files
CMD ["serve", "-l", "3000", "."]
