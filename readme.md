# OpenDI - Example API implementation

This Go API implements the OpenDI standard [API definition](https://opendi.org/api-specification/next/api).  
Intended for local use, for demonstration purposes, this implementation will be available as a Docker image for easy local deployment.

This API allows for basic CRUD operations to be performed on a Causal Decision Model database, allowing for simple persistence in any service that uses the [CDM data schema](https://opendi.org/api-specification/next/schemas/cdm-full-schema).

This project is implemented as a simple [Go](https://go.dev/) API, using the [Gin Web Framework](https://gin-gonic.com/docs/introduction/) for API endpoint implementation, and [GORM](https://gorm.io/) for object-relational mapping from Go typedefs to SQL data schemas.

## Running the API

This API is mainly intended to be used as part of the [Containerized Authoring Demo](https://github.com/opendi-org/containerized-authoring-demo). See that repo for instructions on running the whole project.

To run the API individually,
1. Install [Docker Desktop](https://www.docker.com/products/docker-desktop/) or some other form of the [Docker Engine](https://docs.docker.com/engine/install/).
2. Download the latest image for the OpenDI Go API from DockerHub (TODO: Host image on DockerHub). If this is impossible or unavailable, see [Building the API](#building-the-api) for instructions on building your own image.
3. Run a container from the image. See step 4 for Docker Desktop instructions, or step 5 for command line instructions.
4. Docker Desktop instructions:
    1. In Docker Desktop, locate the image in `Images` and click the Play button.
    2. Expand `Optional settings` and locate the `Ports` section.
    3. Set `Host port` to `8080`. This is the port that the CDD builder frontend expects to use for API calls.
    4. Click `Run`.
5. Command line instructions:
    1. Execute the following command: `docker run -p 8080:8080 opendi-api`
    2. The `-p` flag maps the local port `8080` to the corresponding port `8080` in the container. This is the port that the CDD builder frontend expects to use for API calls.
6. Test the connection. Use a REST API testing software like [Insomnia](https://insomnia.rest/), or try a command like `curl`:
    1. `curl http://localhost:8080/v0/models` should return a JSON-formatted array of Meta objects with CDM IDs, names, etc.


## Building the API

To build an updated version of the API using local project files,
1. Install [Docker Desktop](https://www.docker.com/products/docker-desktop/) or some other form of the [Docker Engine](https://docs.docker.com/engine/install/).
2. Clone this GitHub repository.
3. Within the parent directory for the cloned repository, run the following command:
    1. `docker build -t opendi-api .`

These steps should build a new Docker image using the dockerfile in this repo. The build process has several steps, and may take 2-3 minutes.

This process will create a temporary Docker image and container to create build artifacts for the Go API program, then copy those artifacts into a final runtime image. The final image can be run using the steps in [Running the API](#running-the-api) above, starting with step 3.