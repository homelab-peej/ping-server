# Ping Server Application

This repository contains a simple HTTP web server that responds with "pong" when accessed at the `/ping` endpoint.

### Steps to Deploy

1. **Create a DaemonSet Manifest (`ping-ds.yaml`):**

   ```
   apiVersion: apps/v1
   kind: DaemonSet
   metadata:
   name: ping-server
   labels:
      app.kubernetes.io/name: ping-server
   spec:
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
            image: ghcr.io/homelab-peej/ping-server-test:latest
            ports:
               - containerPort: 8080
            livenessProbe:
               httpGet:
               path: /healthz
               port: 8080
               initialDelaySeconds: 5
               periodSeconds: 10
               failureThreshold: 2
            readinessProbe:
               httpGet:
               path: /readyz
               port: 8080
               initialDelaySeconds: 5
               periodSeconds: 2
               failureThreshold: 2
            startupProbe:
               httpGet:
               path: /startz
               port: 8080
               initialDelaySeconds: 5
               periodSeconds: 2
               failureThreshold: 2
   ```

2. **Create a Service Manifest (`ping-svc.yaml`):**

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

3. **Deploy the Application to Kubernetes:**
   ```
   kubectl apply -f ping-ds.yaml
   kubectl apply -f ping-svc.yaml
   ```

4. **Verify Deployment:**
   Check the status of your daemonset and service:
   ```
   kubectl get ds
   kubectl get svc
   ```

5. **Test the Endpoint:**
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