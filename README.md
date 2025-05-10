# Golang Load Balancer with Docker Compose

This project demonstrates a simple load balancer written in Go that distributes requests among multiple backend Go servers. The full setup is containerized using Docker and orchestrated with Docker Compose.

## Features
- **Load Balancer:** Forwards incoming requests to a pool of backend servers using round-robin and health checks.
- **Backend Servers:** Simple Go HTTP servers that respond with their port and name.
- **Dockerized:** Each component runs in its own container.
- **Docker Compose:** Easily spin up the entire stack with one command.

## Project Structure
- `server.go` — Backend server code.
- `load_balancer.go` — Load balancer code.
- `Dockerfile_server` — Dockerfile for building backend server image.
- `Dockerfile_load_balancer` — Dockerfile for building load balancer image.
- `docker-compose.yml` — Compose file to orchestrate all containers.

## Usage

### 1. Build and Start All Services

```sh
docker-compose -f  docker_compose.yml up -d --build
```
- This will build images and start 3 backend servers and 1 load balancer in the background.

### 2. Access the Load Balancer

Open your browser or use curl:
```
http://localhost:8080
```
The load balancer will forward your request to one of the backend servers.

## Happy coding