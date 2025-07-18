## Module 1: Introduction and Core Concepts 

### 1.0 Bootcamp outline:
From the course "Practical OpenShift for Developers - OpenShift 4," you will gain the following skills:

Deploy existing applications to OpenShift.
Understand 15 types of OpenShift resources.
Configure applications to follow OpenShift best practices.
Develop advanced application templates.
These skills are designed for developers new to OpenShift and those familiar with Kubernetes.

|💡 Modules Stucture  | 
| -------- |
| Resource Information  | 
| Hands-on Walkthroughs |
| Labs    | 
| Quiz |

Skills needed!
- `cd`: To navigate to a specific directory
- `ls`: To list directory components
- `cat`: To read out a specific file
- `curl`: Client for URLs to transfer data to or from server using various network protocols such as HTTP, HTTPS, FTP & more.
- `echo`: To echo a parameter given
- `vi` or `vim`: Its a built in editor used to edit files.
- `git`: To interact with your source control on github.
- Other parameters and options we will learn throughout this course. 



### 1.1 Overview of OpenShift and its Architecture

OpenShift is Red Hat's enterprise Kubernetes platform. It automates deployment, scaling, and management of containerized apps, adding rich features like a web UI, CI/CD, integrated security, and developer tools.

- **Key Features:**  
  Integrated developer tools, security, operators, Service Mesh.
- **Core Architecture:**  
  Control Plane (API server, etcd), Worker Nodes, SDN, Registry, OAuth, Router, Monitoring/Logging.
- **Diagram:**  
![OpenShift Architecture](/images/ocp-arch.png)

**Resource:**  
- [What is OpenShift?](https://youtu.be/KTN_QBuDplo)


### 1.2 Difference between Kubernetes and OpenShift

- OpenShift **includes** Kubernetes but adds a web UI, secure-by-default settings, integrated image management, CI/CD, and Service Mesh.
- Use case: OpenShift automates what vanilla Kubernetes leaves as "do it yourself".

**Resource:**  
- [Kubernetes vs OpenShift](https://www.youtube.com/watch?v=lwzpQK4-2H0)

**Hands-on Walkthroughs**  
- Explore and compare the UI and default tooling in both OpenShift and vanilla Kubernetes.

---