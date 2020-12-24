# SRE Code Challange

Hi, this is my SRE Code Challange! Just a friendly reminder: it's my first time writing code in Go!

## Build & Deploy the application Locally
In order to build your application, you need to have docker (or podman) installed. 

Navigate to the root application directory (where this file exists). Then, you may run the following commands:

    docker build -t sre-challange-deploy -f Dockerfile.deploy .
    docker run -p 8080:8080 sre-challange-deploy

## CI Deplyment
The application has a CI lifecycle via GitHub Actions. It performs the following steps: 
- Unit Tests
- Build application
- Publish application to DockerHub

The image is published as `juanchiriera/sre-challange:latest`, you can access to it executing `docker pull juanchiriera/sre-challange:latest`

The **CI yaml file** can be found [here](./.github/workflows/docker-image.yml).

## The API
Here is a list of all the endpoints provided by the API. There is an example of their usage and the expected response.
### Create new Basket instance

    curl --location --request POST 'localhost:8080/basket'

### Add a new Item to the basket

#### Pen
    curl --location --request POST 'localhost:8080/product/PEN'
#### Mug
    curl --location --request POST 'localhost:8080/product/MUG'
#### T-Shirt
    curl --location --request POST 'localhost:8080/product/TSHIRT'

### Delete basket
    curl --location --request DELETE 'localhost:8080/basket'

### Get total price for basket
    curl --location --request GET 'localhost:8080/total'

### View all the contents of the basket
    curl --location --request GET 'localhost:8080/basket'

## Running tests
In order to run application unit tests, you'll have to run the following command:

    go test -v

These tests consist of the examples shown in the excersise file provided on the challange. They are as follows:

    Items: PEN, TSHIRT, MUG
    Total: 32.50€

    Items: PEN, TSHIRT, PEN
    Total: 25.00€

    Items: TSHIRT, TSHIRT, TSHIRT, PEN, TSHIRT
    Total: 65.00€

    Items: PEN, TSHIRT, PEN, PEN, MUG, TSHIRT, TSHIRT
    Total: 62.50€

