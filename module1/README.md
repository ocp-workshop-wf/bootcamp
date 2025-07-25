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
| Hands-on Labs    | 
| Checklist |
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

OpenShift is Red Hat's enterprise Kubernetes platform. It automates deployment, scaling, and management of containerized apps, adding rich features like a web UI, CI/CD, integrated security, and developer tools. OpenShift has 4 different flavors: 
- OpenShift origin, which is the original upstream open source project from which all other models are derived. 
- The OpenShift online is Red hat's publicly hosted version of OpenShift origin, available for application development and hosting purposes. 
- OpenShift dedicated is a managed private cluster on cloud platforms like AWS and Google. 
- OpenShift enterprise is the on premise private pass offering of OpenShift.

<p align="center">
<img src="/images/openshift-flavors.png" alt="OpenShift Training" style="width:500px; align="center"/>
</p>

- In this course we will focus on OpenShift Origin, OpenShift origin is based on top of Docker containers and the Kubernetes Cluster Manager, with added developer and operational centric tools that enable rapid application development, deployment and lifecycle management.

- **Key Features:**  
  Integrated developer tools, security, operators.
- **Core Architecture:**  
  Control Plane (API server, etcd), Worker Nodes, SDN, Registry, OAuth, Router, Monitoring/Logging.

<p align="center">
<img src="/images/ocp-arch.png" alt="OpenShift Training" style="width:500px; align="center"/>
</p>


**Resource:**  
- [What is OpenShift?](https://youtu.be/KTN_QBuDplo)


### 1.2 Difference between Kubernetes and OpenShift

- OpenShift **includes** Kubernetes but adds a web UI, secure-by-default settings, integrated image management, CI/CD, and Service Mesh.
- Use case: OpenShift automates what vanilla Kubernetes leaves as "do it yourself".

**Resource:**  
- [Kubernetes vs OpenShift](XXXXX)

**Hands-on Walkthroughs**  
- Explore and compare the UI and default tooling in both OpenShift and vanilla Kubernetes.

---