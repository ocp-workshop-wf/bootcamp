## Module 1: Introduction and Core Concepts 

### 1.0 What You’ll Learn:
From the course  "OpenShift 4," you will learn the following OpenShift Resources and soft skills:
||||||
|-|-|-|-|-| 
| OCP CLI | OCP UI | Projects | Pods | Yaml | 
| JSON | Deployment | DeploymentConfig | Services| Route |
| ImageStream | Labels | S2I | Build | BuildConfig| 
| Rollout | Rollback | ReplicationController| ReplicaSet| ConfigMaps |
| Secrets| Webhooks | Deployment Triggers | ConfigChange Triggers | Volumes | 
| Horizontal Scalling | Vertical Scaling | AutoScaling | Rolling Strategy | Recreate Strategy | 
| Custom Strategy | Deduging | OpenShift Jobs | Templates | Health Checks |
| liveness Probe | Readiness probe| HelmCharts | Packaging | Trouble Shooting |

Deploy existing applications to OpenShift.
Understand 15 types of OpenShift resources.
Configure applications to follow OpenShift best practices.
Develop advanced application templates.

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
<img src="/images/ocp-arch.png" alt="OpenShift Training"; align="center"/>
</p>

* **Master Node**: Manages the cluster with components like API, scheduler, and etcd (data store).
* **Worker Nodes**: Run containerized applications (pods) on RHEL.
* **Developers** use Git and CI/CD to deploy apps; **Operations** use automation tools for management.
* **Routing Layer** handles external traffic; **Service Layer** manages internal communication.
* **Persistent Storage** and **Registry** support stateful apps and image management.
* Supports **physical, virtual, private, public, and hybrid** deployments.



**Resource:**  
- [What is OpenShift?](https://youtu.be/KTN_QBuDplo)


### 1.2 Difference between Kubernetes and OpenShift

| Criteria | OpenShift | Kubernetes| 
| -------- | --------- | --------- | 
| Architecture | Monolithic | Modular|
| Product vs Project | Commerical Product | Open-Source Project |
| User interface (UI) | Web console with login page | Kubernetes dashboard | 
| Templates vs Helm | Project-specific templates | Helm charts for broader use | 
| Image Registry | Built-in registry | No built-in registry | 
| CICD | CI/CD tools integrated | Requires external tools | 
| Community and support | commercial support available | Open-Source community support|
| Networking | Rich set of networking features | Core Networking Features| 

- OpenShift **includes** Kubernetes but adds a web UI, secure-by-default settings, integrated image management, CI/CD, and Service Mesh.
- Use case: OpenShift automates what vanilla Kubernetes leaves as "do it yourself".

**Resource:**  
- [Kubernetes vs OpenShift](XXXXX)

**Hands-on Walkthroughs**  
- Explore and compare the UI and default tooling in both OpenShift and vanilla Kubernetes.

---