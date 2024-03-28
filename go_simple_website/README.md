# Simple Website
A Golang web server that serves a modified variant of the "Bold & Bright" template from the [Bootstrap Studio](https://bootstrapstudio.io/) web design tool.  
It modifies the template by appending a node id (passed in an an environment variable `NODE_ID`) and the time (server time) which the request was processed to each served page.  
These modifications are useful for testing caching strategies & policies.

## Deployment
### Local Development
```shell
NODE_ID=sample PORT=8080 go run .
```
### Docker
```shell
cd .docker
docker-compose up --build -d

# Stop the server and database
docker-compose down
```
