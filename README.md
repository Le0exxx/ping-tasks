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

# Part 3: Ready to production
1. Query Caching to Avoid API Quota Overuse
Current Limitation: Each request to the web service triggers an external API call, consuming API quota.

Improvement:

Implement a caching layer (e.g., Redis or in-memory) to store API responses for repeated requests.

Use TTL (time-to-live) strategy to ensure data freshness while avoiding redundant calls.

2. Authentication Layer for Security
Current Limitation: The application lacks an authentication mechanism, making it publicly accessible.

Improvement:

Add JWT-based or API key-based authentication for service endpoints.

Integrate with OAuth2 or OIDC for secure user access (if applicable).

3. Web Application Firewall (WAF) and Rate Limiting
Current Limitation: The application is vulnerable to brute force, bots, and malicious payloads.

Improvement:

Use AWS WAF or ingress-nginx WAF rules to:

Rate limit requests

Block/allow by IP/geo

Enforce body size limits and header checks

4. Application Autoscaling with KEDA
Current Limitation: App runs at static capacity, under-utilizing or over-consuming resources.

Improvement:

Integrate KEDA with a metrics provider (e.g., Prometheus, custom metrics)

Enable horizontal pod autoscaling based on actual demand (e.g., queue depth, HTTP traffic, CPU/memory)

5. Secure API Key Management
Current Limitation: API keys are stored in environment variables or source code.

Improvement:

Use External Secrets Operator with a backend like AWS Secrets Manager or HashiCorp Vault.

Automate secret syncing with Kubernetes to keep secrets secure and centrally managed.

6. Automated DNS & TLS Certificate Management
Current Limitation: Manual configuration of DNS records and SSL certificates.

Improvement:

Use external-dns for dynamic DNS record management (e.g., Route53)

Use cert-manager to auto-provision and renew Let's Encrypt TLS certificates.

7. CI/CD and GitOps Deployment
Current Limitation: Manual deployments or scripts with potential human error.

Improvement:

Use GitOps (ArgoCD) for declarative deployment pipelines.

Integrate with CI/CD tools (GitHub Actions, GitLab CI, etc.) to automate testing, image builds, and environment promotion.

8. Monitoring, Health Checks, and Logging
Current Limitation: No structured monitoring or alerting in place.

Improvement:

Add readiness/liveness probes to deployments for better fault detection.

Integrate with Prometheus + Grafana for metrics monitoring.

Use a log collector (e.g., Fluent Bit + Loki or ELK Stack) to centralize application logs for observability.