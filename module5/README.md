## Module 5: OpenShift Service Mesh & Advanced Traffic Management (4 hours)

### 5.1 Introduction to Service Mesh

- A **Service Mesh** handles communication, security, and monitoring for microservices.
- Uses **sidecar** proxies for each app container.
- **OpenShift Service Mesh** includes Istio (traffic/security), Kiali (visualization), Jaeger (tracing), Grafana (metrics).

**Diagram:**  
  ![Service Mesh Architecture](./service-mesh-architecture.png)

**Lab:**  
- Install Service Mesh Operator, enable mesh in a namespace.

**YouTube:**  
- [OpenShift Service Mesh Overview](https://www.youtube.com/watch?v=CeLkHLvBeB4)

---

### 5.2 Blue/Green Deployments & Traffic Routing

- Deploy two app versions (blue/green), split or shift traffic.
- Service Mesh allows advanced traffic splitting, routing based on headers, instant rollback.

**Lab:**  
- Deploy v1/v2 of an app.
- Use Kiali to shift traffic (80/20 to 100%).

**YouTube:**  
- [OpenShift Service Mesh: Blue/Green Deployments](https://www.youtube.com/watch?v=xWBUOq_jIVg)

---

### 5.3 Observability and Tracing

- **Kiali**: Service connections and health.
- **Jaeger**: Distributed tracing of requests.
- **Grafana**: Dashboards for metrics.

**Lab:**  
- Generate test traffic, use Kiali and Jaeger for live monitoring and tracing.

**YouTube:**  
- [OpenShift Service Mesh Observability](https://www.youtube.com/watch?v=VdIewxv9CHs)

---

### 5.4 Service Mesh Use Cases

- Fault injection, retries, circuit breaking, mTLS, policy enforcement.
- Test app resilience and security without changing code.

**Lab:**  
- Inject a failure using Service Mesh and observe via Kiali.

---

