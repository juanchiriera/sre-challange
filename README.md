# SRE Code Challange

Hi, this is my SRE Code Challange! Just a friendly reminder: it's my first time writing code in Go!

## Build & Deploy the application
In order to build your application, you need to have docker (or podman) installed. 

Navigate to the root application directory (where this file exists). Then, you may run the following commands:

    docker build -t sre-challange-deploy -f Dockerfile.deploy .
    docker run -p 8080:8080 sre-challange-deploy
## The API
Here is a list of all the endpoints provided by the API. There is an example of their usage and the expected response.
### Add a new Item to the basket

### Create new Basket instance

### Delete basket

### Get total price for basket

## Running tests
In order to run application unit tests, you'll have to run the following command:

    go test -v

