# docker build --no-cache -t ms-frontend:latest .
FROM nginx:alpine
COPY index.html /usr/share/nginx/html