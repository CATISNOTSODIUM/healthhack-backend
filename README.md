# Healthhack-backend 
- [Healthhack-backend](#healthhack-backend)
  - [1. About this project](#1-about-this-project)
    - [Roadmap](#roadmap)
  - [2. Set up](#2-set-up)
    - [Build and running app locally](#build-and-running-app-locally)
    - [Deployment](#deployment)

## 1. About this project
### Roadmap
- [x] Set up database schema
- [x] User queries, mutation, Google authentication
- [x] API adding history
- [ ] API Connect to open ai
## 2. Set up
### Build and running app locally
Clone the repository and install required repositories:
```bash
go mod download
```
Make sure to configure your `.env` file before starting the server.
```
PORT=YOUR_PORT_NUMBER # DEFAULT = 8080
DSN=[YOUR_POSTGRESQL_URL]
OPENAI_TOKEN=[YOUR_TOKEN] # for text extraction
```
Then, you can start the server.
``` bash
go run cmd/server/main.go
```
### Deployment
To be updated
