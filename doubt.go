In Docker for Desktops Kubernetes, I am trying to have a simple frontend (html) running on Nginx.
This is supposed to connect to the backend service (written in Go).
I can not get it to work so far as I either get no connection error 
in case I have the proxy_pass ON for nginx, otherwise simply can not connect to backend if I dont have the proxy_pass.

Following are 
portions of my frontend code (html)
nginx configuration file
Dockerfile of the frontend which involves inclusion of the conf file of nginx
Service definition of Frontend (NodePort so that I can access it at localhost:30100)
Service definition of Backend (ClusterIP)

I can ssh into the frontend Pod and do a curl http://ms-bff/getsize and get the desired result. 
So I am confused why it doesnt work when I access the frontend at localhost:30100 



// **** Frontend code in html **** 

	var url = "http://ms-bff:8081/getsize"
	$.ajax({
		url: url,
		dataType: "jsonp",
		success: function (data) {
			console.log(data.hs)
			// Doing Something with the data
		}
	});

// *** nginx configuration file****
	upstream ms-bff {
		server ms-bff;
	}
	
	server {
		listen 80;
	
		# Proxy request to rails app
		location / {
		proxy_pass http://ms-bff;
		}
	}

// ****Dockerfile for the html ****
	FROM nginx:alpine
	COPY index.html /usr/share/nginx/html
	COPY nginx.conf /etc/nginx/conf.d/default.conf

// **Service definition of frontend***
	apiVersion: v1
	kind: Service
	metadata:
	name: &service-name ms-frontend
	spec:
	selector:
		app: *service-name
	ports:
		- name: http
		port: 80
		nodePort: 30100
	type: NodePort

// ** Service definition of backend***
	apiVersion: v1
	kind: Service
	metadata:
	name: &service-name ms-bff
	spec:
	type: ClusterIP
	selector:  
		app: *service-name
	ports:
		- name: http
		port: 8081 
		targetPort: 8081
		protocol: TCP


Backend listens on 8081 port. 