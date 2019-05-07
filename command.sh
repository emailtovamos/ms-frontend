docker build --no-cache -t ms-frontend:v3 .;

docker tag ms-frontend:v3 emailtovamos/ms-frontend:v3;

docker push emailtovamos/ms-frontend:v3;

kubectl delete deployment ms-frontend;

kubectl apply -f deployment.yaml;

