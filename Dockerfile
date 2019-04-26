# docker build --no-cache -t ms-frontend:latest .
FROM nginx:alpine

# RUN rm /etc/nginx/conf.d/default.conf

# RUN rm /etc/nginx/conf.d/default.conf

COPY index.html /usr/share/nginx/html

COPY nginx.conf /etc/nginx/nginx.conf