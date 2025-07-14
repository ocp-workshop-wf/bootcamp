## Docker

### Building containers

#### Build an image based on the current directory
`docker build .`

#### Build an image based on the current directory with a tag
`docker build -t your-tag .`

### Running Containers

#### Check the image storage
`docker images`

> Start a container based on an image ID. Get the ID from docker images.

#### `control-c` will stop the container for all of these docker run commands.
`docker run -it <your-image-id>`

#### Start an image based on a tag
`docker run -it <image tag>`

#### Start hello-world server
`docker run -it quay.io/practicalopenshift/hello-world`

#### Start the Hello World app with port forwarding
`docker run -it -p 8080:8080 quay.io/practicalopenshift/hello-world`


### Stopping Containers

#### Get running images
`docker ps`

#### Stop a running image. The container ID will be in the docker ps output.
`docker kill <Container ID>`



