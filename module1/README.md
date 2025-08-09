[![Static Badge](https://img.shields.io/badge/Agenda-green?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/other/Agenda.md) [![Static Badge](https://img.shields.io/badge/CheatSheet-purple?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/other/CheatSheet.md) [![Static Badge](https://img.shields.io/badge/OCP-CLI-red?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/other/ocpcli-cheatsheet.md)   [![Static Badge](https://img.shields.io/badge/Labs-maroon?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/tree/main/labs-repo)  [![Static Badge](https://img.shields.io/badge/RedHat-OpenShift-maroon?style=flat&logo=Redhat&logoSize=auto)
](https://docs.redhat.com/en/documentation/openshift_container_platform/4.19)   [![Static Badge](https://img.shields.io/badge/Kubernetes-black?style=flat&logo=Kubernetes&logoSize=auto)
](https://kubernetes.io/docs/home/)
## 🔹 Module 1: Introduction and Core Concepts 

## Table of Contents

- [1.1 - Introduction to OpenShift](#11-overview-of-openshift-and-its-architecture)
- [1.2 - Difference between Kubernetes and OpenShift](#12-difference-between-kubernetes-and-openshift)

### 1.1 Overview of OpenShift and its Architecture

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module1#-module-1-introduction-and-core-concepts" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

OpenShift is Red Hat’s enterprise-grade Kubernetes platform designed for running containerized applications and workloads. Built on top of Kubernetes, it streamlines daily tasks for both developers and operations teams by automating deployment, scaling, and management. It also enhances Kubernetes with a feature-rich web UI, integrated CI/CD pipelines, built-in security, and robust developer tools. OpenShift comes in four flavors, each tailored to different use cases:
- OpenShift origin, which is the original upstream open source project from which all other models are derived. 
- The OpenShift online is Red hat's publicly hosted version of OpenShift origin, available for application development and hosting purposes. 
- OpenShift dedicated is a managed private cluster on cloud platforms like AWS and Google. 
- OpenShift enterprise is the on premise private pass offering of OpenShift.

<p align="center">
<img src="/images/openshift-flavors.png" alt="OpenShift Training" style="width:500px; align="center"/>
</p>

- In this course, we will focus on OpenShift Origin—the open-source upstream of Red Hat OpenShift. It’s built on top of Kubernetes and previously used Docker as its container runtime (now replaced by CRI-O), and includes additional developer and operational tools that enable rapid application development, deployment, and lifecycle management.

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

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module1#-module-1-introduction-and-core-concepts" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

  <p align="center">
  <img src="/images/intro-ocp.png" alt="OpenShift Training" style="width:500px; align="center"/>
  </p>


> Diagram description: One of the big advantages is being able to take advantage of `Public` or `Private` resources for running OpenShift that includes bare metal or virtualized hardware, whether on-premises or in the cloud. On the top of that we will have the operating system which is generally `RHEL or ` CentOS`. Then the K8's layer and then OCP layer.

> On the other side of the digram: Developers they write applications and test them up then deploy them into the cluster, after the push OpenShift will be running that in the pipeline, creating a a Docker `image` and pushing the image into the `image registry`. On the other side of the pipeline once the `image` gets `built` OpenShift will take that and push it in the actual cluster taking advantage of something called `imageStream`.

> Operation Engineers: OpenShift has the web console for managing the cluster, monitoring resources, and troubleshooting issues. Next the scaling out - OpenShift takes advantage of an Ansible playbook they can go ahead spin up a new host and provision it as needed which we call VPA, or increasing the amout of pods running and we call that HPA. 

<div align="center">

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

</div>
> So first off, when it comes to **architecture**, OpenShift is more monolithic—it bundles a lot of features together out of the box, while Kubernetes is modular, meaning you build and plug in the tools you need. In terms of what they are, OpenShift is a **commercial product** developed by Red Hat, while Kubernetes is an **open-source project** maintained by a broad community.Now for the **user interface**, OpenShift gives you a polished web console with a login page and role-based access built in. Kubernetes has a dashboard, but you usually have to install and configure it yourself. When it comes to application packaging, OpenShift uses **templates** which are more project-specific, whereas Kubernetes supports **Helm charts**, which are reusable and more flexible across different environments. OpenShift includes a **built-in container image registry**, so you don’t need to set one up separately. Kubernetes doesn’t include this—you’ll have to bring your own registry. For **CI/CD**, OpenShift includes tools like Jenkins pipelines out of the box, while in Kubernetes, you’ll need to integrate external CI/CD solutions like GitHub Actions or Tekton. Support-wise, OpenShift comes with **commercial support**, ideal for enterprises needing SLAs. Kubernetes has a great open-source community, but support usually depends on your vendor or your team’s skills. Finally, for **networking**, OpenShift provides advanced networking features and policies by default, while Kubernetes offers the core networking features and leaves the rest for you to configure. 

**Resource:**  
- [Kubernetes vs OpenShift](https://www.theknowledgeacademy.com/blog/openshift-vs-kubernetes/)

**Hands-on Walkthroughs**  
- Explore and compare the UI and default tooling in both OpenShift and vanilla Kubernetes.

---


<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module2" target="_blank">
    <img src="/images/nexticon.webp" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>
<p align="left">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main" target="_blank">
    <img src="/images/backred1.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>
