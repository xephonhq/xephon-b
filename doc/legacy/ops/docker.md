# Docker Quick Start

- https://github.com/odewahn/docker-jumpstart
- https://github.com/wsargent/docker-cheat-sheet

Explain some basic commands of docker

- `docker ps -a` to show all containers in the system
- `docker images` to list all the images
- `docker run <image-name> <command> <arg..>`
  - `<image-name>` is like github projects, `owner/name`, ie: `docker/hello-world`
  - [ ] commands
  - `<arg...>` arguments pass to the docker
- `docker tag <image-id> <account-name>/<repository>:<version>`
- `docker rmi <image-id|image-name>` to remove image

## Mount volume

https://docs.docker.com/engine/tutorials/dockervolumes

### Mount a host directory as a data volume

`docker run -v [host/path]:[path/in/container]:[mode] [container name]`

- mount read only `docker run -d -P --name web -v /src/webapp:/webapp:ro training/webapp python app.py`

## Build and publish

- [Use GitHub and auto build](https://docs.docker.com/docker-hub/builds/)
- [Manually build an publish](https://docs.docker.com/engine/getstarted/step_six/)
  - create a repository on docker-hub
  - tag
  - push to remote
