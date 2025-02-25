# Welcome Application

A Simple Golang web application with Docker, and CI/CD pipelines.

## Access

Access: 
  -  http://localhost:8000/welcome/{name}   --> Welcome {name}
  -  http://localhost:8000/welcome/         --> Anonymous

## Docker Command

1. Build the Docker Image

```bash
docker build -t testing/welcome .
```

2. Run the Container

```bash
 docker run --rm -d -p 8000:5000 --name welcome-app testing/welcome:latest
```

## CI/CD Pipeline Usage

1. Manual Trigger

```bash
Requires inputs:
  - release_version: Semantic version (e.g., v1.0.0)
  - operation:
      1. build-only
      2. deploy-only
      3. build-and-deploy
```

2. Jobs

```bash
  - Validate Version: Ensures semantic versioning

  - Build & Push: Build and push docker image to docker hub

  - Deploy: SSH to VM Server to run docker container