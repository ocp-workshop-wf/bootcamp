# OpenShift 24-Hour Comprehensive Course

Welcome to your hands-on, instructor-led OpenShift course! This curriculum is designed to take you from beginner to advanced in 24 hours, with a focus on real-world use, automation, Service Mesh, and microservices.

---

## Table of Contents

- [Module 1: Introduction and Core Concepts](#module-1-introduction-and-core-concepts-2-hours)
- [Module 2: Working with OpenShift Interfaces](#module-2-working-with-openshift-interfaces-2-hours)
- [Module 3: Core OpenShift Resources](#module-3-core-openshift-resources-3-hours)
- [Module 4: Application Deployment and Management](#module-4-application-deployment-and-management-4-hours)
- [Module 5: OpenShift Service Mesh & Advanced Traffic Management](#module-5-openshift-service-mesh--advanced-traffic-management-4-hours)
- [Module 6: Reliability and Operations](#module-6-reliability-and-operations-4-hours)
- [Module 7: Observability and Monitoring](#module-7-observability-and-monitoring-3-hours)
- [Module 8: Capstone Project](#module-8-capstone-project-2-hours)
- [PNG Diagrams](#png-diagrams)

---











## Module 7: Observability and Monitoring (3 hours)

### 7.1 Metrics and Logging

- **Prometheus**: Metrics collection.
- **Logging stack**: Elasticsearch, Fluentd, Kibana.

**Diagram:**  
  ![Monitoring Architecture](./monitoring-architecture.png)

**Lab:**  
- View metrics in Prometheus/Grafana.
- Search app logs in Kibana.

**YouTube:**  
- [Monitoring OpenShift with Prometheus & Grafana](https://www.youtube.com/watch?v=QIO1WkJOGwE)
- [OpenShift Logging Overview](https://www.youtube.com/watch?v=2BstJQxHzSI)

---

### 7.2 Observability Tools and Practices

- Custom dashboards in Grafana.
- Alerts in Prometheus.

**Lab:**  
- Build a custom dashboard.
- Set a sample alert.

**YouTube:**  
- [Building Dashboards in Grafana](https://www.youtube.com/watch?v=Gr2xgKx0Mh8)

---

## Module 8: Capstone Project (2 hours)

### 8.1 End-to-End Microservices Deployment

- **Scenario**:  
  Deploy a multi-tier microservices app. Use ConfigMaps, Secrets, health probes, Service Mesh, blue/green, and observability.

**Lab:**  
1. Create all services and configs.
2. Enable Service Mesh, blue/green for a microservice.
3. Monitor with Kiali, Jaeger, Grafana.
4. Inject a fault, troubleshoot, and roll back.

**YouTube:**  
- [Real-World OpenShift Project Example](https://www.youtube.com/watch?v=_u8M1kFNHLM)
- [Bookinfo Demo for Istio/Service Mesh](https://www.youtube.com/watch?v=KpA-SyF6OMM)

---

## PNG Diagrams

- `openshift-architecture.png` - OpenShift Core Architecture
- `service-mesh-architecture.png` - OpenShift Service Mesh (Istio/Kiali/Jaeger) Architecture
- `blue-green-deployment.png` - Blue/Green Deployment Flow
- `monitoring-architecture.png` - Monitoring & Logging Stack Architecture

*Ask your instructor for these PNGs, or generate with your favorite diagram tool!*

---

## **Congratulations!**

You are ready to deploy, scale, manage, and monitor real-world apps on OpenShift!

---

*For more labs, exercises, or to request diagrams, reach out to your instructor or course designer!*

