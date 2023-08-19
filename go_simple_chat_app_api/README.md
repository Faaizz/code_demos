# Simple Chat Application API
Simple Chat Application API in Golang.  
Developed to serve as a backend target for an AWS API Gateway Web Socket API, the API persists data to AWS DynamoDB. Also, it can optionally be configured to enable AWS X-Ray Tracing (by setting the `AWS_XRAY_DAEMON_ADDRESS` environment variable).

## Deploy Development Server (Docker)
Deploy development server and local instance of AWS DynamoDB with `docker-compose`.
```shell 
cd .docker
docker-compose up --build

# Stop the server and database
docker-compose down
```

## Usage
### Sample Client Requests
```shell
export HOST_ADDR="localhost"
export HOST_PORT="8088"
# health
curl -v http://$HOST_ADDR:$HOST_PORT/healthz --header "Content-Type:application/json"
# connect
curl -v http://$HOST_ADDR:$HOST_PORT/connect -d '{"connectionId": "91123456"}' --header "Content-Type:application/json"
# set username
curl -v http://$HOST_ADDR:$HOST_PORT/username -d '{"connectionId": "91123456", "username": "Asiat"}' --header "Content-Type:application/json"
# disconnect
curl -v http://$HOST_ADDR:$HOST_PORT/disconnect -d '{"connectionId": "91123456"}' --header "Content-Type:application/json"
# list users
curl -v http://$HOST_ADDR:$HOST_PORT/online -d '{"connectionId": "91123456"}' --header "Content-Type:application/json"
# message
curl -v http://$HOST_ADDR:$HOST_PORT/message -d '{"connectionId": "911234567", "username": "Asiat33", "message": "Hello WebSocket", "url": ""}' --header "Content-Type:application/json"
```

## References
- [https://github.com/aws-samples/simple-websockets-chat-app](https://github.com/aws-samples/simple-websockets-chat-app)
- [https://www.udemy.com/course/aws-certified-developer-associate-dva-c01/learn/lecture/26101850#overview](https://www.udemy.com/course/aws-certified-developer-associate-dva-c01/learn/lecture/26101850#overview)
- [https://gowebexamples.com/http-server/](https://gowebexamples.com/http-server/)
