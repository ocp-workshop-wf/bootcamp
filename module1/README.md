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

**Lab:**  
- Log in to the OpenShift Web Console.

    > #### Log in, log out
    >
    >> ##### Uses the pre-configured OpenShift cluster
    >  `oc login`

    >> ##### Allows you to log in to any OpenShift cluster
    > `oc login <cluster address>` 
    <br/>

    > `oc whoami`
    >> ##### Log out
    > `oc logout`


    > #### Project Basics

    >> ##### See current project
    > `oc project`

    >> ##### Create a new project
    > `oc new-project myproject`

    >> ##### List all projects
    > `oc projects`

    >> ##### Switch projects
    > `oc project <project name>`
___

- Identify the main cluster components.

___
### Quiz

#### Q1:
What is the command to log in to OpenShift on the command line?
- [ ] `oc log on <cluster URL>`
- [ ] `oc login <cluster URL>`
- [ ] `minishift log on <cluster URL>`
- [ ] `minishift login <cluster URL>`
<details>
  <summary> Answer </summary>

   `oc login <cluster URL>`

</details>


#### Q2:
What is the command to create a project?
- [ ] `oc project <name>`
- [ ] `oc projects <name>`
- [ ] `oc new-project <name>`
- [ ] `oc new project <name>`
<details>
  <summary> Answer </summary>

   `oc new-project <name>`

</details>

#### Q3:
What is the command to switch your current project to a different project?
- [ ] `oc project <name>`
- [ ] `oc projects <name>`
- [ ] `oc new-project <name>`
- [ ] `oc new project <name>`
<details>
  <summary> Answer </summary>

   `oc project <name>`

</details>

---

### 1.2 Difference between Kubernetes and OpenShift

- OpenShift **includes** Kubernetes but adds a web UI, secure-by-default settings, integrated image management, CI/CD, and Service Mesh.
- Use case: OpenShift automates what vanilla Kubernetes leaves as "do it yourself".

**YouTube:**  
- [Kubernetes vs OpenShift](https://www.youtube.com/watch?v=lwzpQK4-2H0)

**Lab:**  
- Explore and compare the UI and default tooling in both OpenShift and vanilla Kubernetes.

---