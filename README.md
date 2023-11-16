# Gin Starter

A web API boilerplate for [Gin](https://gin-gonic.com/).

## Run

### Docker

1.  Build the image.

    ```sh
    docker build  -t myapp .
    ```

2.  Start a container of the newly built image.

    ```sh
    docker run --rm -p 8000:8000 my-go-api
    ```

### Local

1.  Run the app.

    ```sh
    go run .
    ```

2. Navigate to `localhost:8000`

## Helpful Commands

*   Request the data from the `/` (Home) endpoint.  This endpoint is public to anyone and does not require a JWT token
    to be sent in the Header of the request.

    ```sh
    curl --location --request GET "http://localhost:8000" | jq
    ```

*   Get the data from the `api/v1/pokemon` endpoint.  This endpoint is private and can only be accessed with a valid
    JWT token sent in the Header of the request. Replace `<REPLACE_WITH_YOUR_ACCESS_TOKEN>` with your access token,
    which can be created for testing the workflow at `Auth0 > Applications > APIs > Test > CURL`.

    ```sh
    curl --location \
        --request GET "http://localhost:8000/" \
        --header "Authorization: Bearer <REPLACE_WITH_YOUR_ACCESS_TOKEN>" | jq
    ```

## Resouces

* jwt.io
* Auth0 docs
