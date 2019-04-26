docker build --no-cache -t ms-frontend:v2 .;

docker tag ms-frontend:v2 emailtovamos/ms-frontend:v2;

docker push emailtovamos/ms-frontend:v2;

kubectl delete deployment ms-frontend;

kubectl apply -f deployment.yaml;

