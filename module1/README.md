## Module 1: Introduction and Core Concepts 

### 1.1 Overview of OpenShift and its Architecture

OpenShift is Red Hat's enterprise Kubernetes platform. It automates deployment, scaling, and management of containerized apps, adding rich features like a web UI, CI/CD, integrated security, and developer tools.

- **Key Features:**  
  Integrated developer tools, security, operators, Service Mesh.
- **Core Architecture:**  
  Control Plane (API server, etcd), Worker Nodes, SDN, Registry, OAuth, Router, Monitoring/Logging.
- **Diagram:**  
![OpenShift Architecture](/images/ocp-arch.png)

**YouTube:**  
- [What is OpenShift?](https://youtu.be/KTN_QBuDplo)


### 1.2 Difference between Kubernetes and OpenShift

- OpenShift **includes** Kubernetes but adds a web UI, secure-by-default settings, integrated image management, CI/CD, and Service Mesh.
- Use case: OpenShift automates what vanilla Kubernetes leaves as "do it yourself".

**YouTube:**  
- [Kubernetes vs OpenShift](https://www.youtube.com/watch?v=lwzpQK4-2H0)

**Lab:**  
- Explore and compare the UI and default tooling in both OpenShift and vanilla Kubernetes.

---