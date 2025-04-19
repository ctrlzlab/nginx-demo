# 🚀 NGINX + Go Backend Demo

This project demonstrates a simple web application where:

- A **Go-based backend** server serves region-specific APIs
- An **NGINX server** acts as a reverse proxy, forwarding requests based on URL parameters (e.g., `/eu/`, `/us/`)

Both services are containerized using **Docker** and managed via **Docker Compose**.

## 🧪 Run the Project
Start the application using Docker Compose:

```bash
docker compose up --build
```

## 🔍 Test the Proxy Behavior
Use curl or your browser to test the regional routing:

```bash
curl http://localhost:9090/eu/home    # Forwards the request to the Europe backend
curl http://localhost:9090/us/home    # Forwards the request to the USA backend
```

