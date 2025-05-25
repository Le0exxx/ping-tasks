# Ping Tasks WebService

This project represents Ping Tasks WebService for src and Deployment Code.

# Getting Started

## Part 1 
Src code and image build

- golang 1.24
- make
- docker (or equiv -> rancher desktop)

```shell
# Build and pulish image
# Pass DOCKER_REGISTRY and TAG environment variables if needed
# Image already in https://hub.docker.com/repository/docker/le0exxx/ping-tasks-webservice/general

make build_all
```

## Part 2
Deploy ingress and app

- Minikube
- make
- Helm

```shell
# Deploy ingress and app via helm
make helm_all
```

```shell
# Open minikube tunnel
 minikube tunnel
```

```shell
# Get the Prices
 curl -H "Host: ping-tasks-webservice.local" http://127.0.0.1/rabbit
```

```shell
# Clean up the deployment
make helm_uninstall
```

## Part3

This document outlines the key areas for improvement in the application, focusing on performance, security, scalability, and maintainability.

---

### 1. Query Caching to Avoid API Quota Overuse

**Current Limitation:**  
Every request triggers an external API call, quickly consuming the API quota.

**Proposed Improvement:**
- Implement a caching layer (e.g., Redis or in-memory cache).
- Use TTL (Time-To-Live) to balance freshness and efficiency.
- Avoid redundant API calls for repeated queries.

---

### 2. Authentication Layer for Security

**Current Limitation:**  
The application lacks authentication, exposing it to unauthorized access.

**Proposed Improvement:**
- Implement JWT-based or API key-based authentication.
- Optionally integrate with OAuth2 or OIDC for user identity federation.
- Secure service endpoints with proper access control.

---

### 3. Web Application Firewall (WAF) and Rate Limiting

**Current Limitation:**  
No protection against brute-force attacks, bots, or malicious inputs.

**Proposed Improvement:**
- Use AWS WAF or ingress-nginx WAF rules to:
  - Apply rate limiting.
  - Allow/block traffic by IP or geo-location.
  - Enforce body size and header validation rules.

---

### 4. Application Autoscaling with KEDA

**Current Limitation:**  
Application pods run at fixed capacity, inefficient under varying load.

**Proposed Improvement:**
- Integrate [KEDA](https://keda.sh) for event-driven autoscaling.
- Use Prometheus or custom metrics (e.g., queue depth, HTTP requests) as scaling triggers.
- Optimize resource utilization through horizontal pod autoscaling.

---

### 5. Secure API Key Management

**Current Limitation:**  
Sensitive keys are stored in plain environment variables or source code.

**Proposed Improvement:**
- Use [External Secrets Operator](https://external-secrets.io/) with a backend like:
  - AWS Secrets Manager
  - HashiCorp Vault
- Automatically sync secrets into Kubernetes securely.

---

### 6. Automated DNS & TLS Certificate Management

**Current Limitation:**  
DNS records and SSL certificates require manual updates.

**Proposed Improvement:**
- Use [external-dns](https://github.com/kubernetes-sigs/external-dns) for automated DNS management (e.g., Route 53).
- Use [cert-manager](https://cert-manager.io) to automate TLS certificate issuance and renewal.

---

### 7. CI/CD and GitOps Deployment

**Current Limitation:**  
Deployments are manual and error-prone.

**Proposed Improvement:**
- Use GitOps principles with [ArgoCD](https://argo-cd.readthedocs.io) for declarative deployments.
- Integrate with CI/CD tools (e.g., GitHub Actions, GitLab CI):
  - Run tests.
  - Build and push Docker images.
  - Promote across environments automatically.

---

### 8. Monitoring, Health Checks, and Logging

**Current Limitation:**  
Lack of structured observability and fault detection.

**Proposed Improvement:**
- Add Kubernetes readiness and liveness probes.
- Integrate with Prometheus and Grafana for monitoring and dashboards.
- Use Fluent Bit, Loki, or the ELK stack for centralized logging and alerting.

---
