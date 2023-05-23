# tempest-gateway-service  
this is a vital component within `the Tempest`  
this acts as an entry point for all client requests.  
  
# How to run  
## Build  
```bash
docker build -t .
 ```
   
 ## Run  
 ```bash
docker run -p 8080:8080 -v . -e ENV_VARIABLE=value .
 ```
   
 ## Stop the container  
 ```bash
 docker stop container-name
 ```
