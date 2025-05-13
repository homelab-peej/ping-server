# Ping Server Application

This repository contains a simple HTTP server written in Go that responds with "pong" when accessed at the `/ping` endpoint.

## Table of Contents
- [Build and Run Locally](#build-and-run-locally)
  - [Prerequisites](#prerequisites)
  - [Steps to Build and Run](#steps-to-build-and-run)
- [Run in Docker](#run-in-docker)
  - [Prerequisites](#prerequisites-1)
  - [Steps to Build and Run](#steps-to-build-and-run-1)
- [Deploy on Kubernetes](#deploy-on-kubernetes)
  - [Prerequisites](#prerequisites-2)
  - [Steps to Deploy](#steps-to-deploy)

## Build and Run Locally

### Prerequisites
- Go 1.24 or later installed on your machine.
- Ensure `go` is in your system's PATH.

### Steps to Build and Run

1. **Clone the Repository:**
   ```
   git clone https://github.com/your-username/ping-server.git
   cd ping-server
   ```

2. **Build the Application:**
   ```
   go build -o ping-server
   ```

3. **Run the Application:**
   ```
   ./ping-server
   ```

4. **Test the Endpoint:**
   Open a browser or use `curl` to test:
   ```
   curl http://localhost:8080/ping
   ```
   Expected response:
   ```
   pong
   ```

## Run in Docker

### Prerequisites
- Docker installed on your machine.

### Steps to Build and Run

1. **Clone the Repository (if not already):**
   ```
   git clone https://github.com/your-username/ping-server.git
   cd ping-server
   ```

2. **Build the Docker Image:**
   ```
   docker build -t ping-server .
   ```

3. **Run the Docker Container:**
   ```
   docker run -d -p 8080:8080 --name ping-server-container ping-server
   ```

4. **Test the Endpoint:**
   Open a browser or use `curl` to test:
   ```
   curl http://localhost:8080/ping
   ```
   Expected response:
   ```
   pong
   ```

## Deploy on Kubernetes

### Prerequisites
- Kubernetes cluster up and running.
- `kubectl` configured to access your cluster.

### Steps to Deploy

1. **Clone the Repository (if not already):**
   ```
   git clone https://github.com/your-username/ping-server.git
   cd ping-server
   ```

2. **Create a Deployment Manifest (`ping-deployment.yaml`):**

   ```
   apiVersion: apps/v1
   kind: Deployment
   metadata:
     name: ping-server
     labels:
       app.kubernetes.io/name: ping-server
   spec:
     replicas: 1
     selector:
       matchLabels:
         app.kubernetes.io/name: ping-server
     template:
       metadata:
         labels:
           app.kubernetes.io/name: ping-server
       spec:
         containers:
         - name: ping-server
           image: ghcr.io/homelab-peej/ping-server:latest
           ports:
           - containerPort: 8080
   ```

3. **Create a Service Manifest (`ping-service.yaml`):**

   ```
   apiVersion: v1
   kind: Service
   metadata:
     name: ping-server
     labels:
       app.kubernetes.io/name: ping-server
   spec:
     type: LoadBalancer
     ports:
     - port: 80
       targetPort: 8080
     selector:
       app.kubernetes.io/name: ping-server
   ```

4. **Deploy the Application to Kubernetes:**
   ```
   kubectl apply -f ping-deployment.yaml
   kubectl apply -f ping-service.yaml
   ```

5. **Verify Deployment:**
   Check the status of your deployment and service:
   ```
   kubectl get deploy
   kubectl get svc
   ```

6. **Test the Endpoint:**
   Once the LoadBalancer service is provisioned (this might take a few minutes), you can find the external IP using:
   ```
   kubectl get svc ping-server
   ```
   Use the external IP to test:
   ```
   curl http://<external-ip>/ping
   ```
   Expected response:
   ```
   pong
   ```