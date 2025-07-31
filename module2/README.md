[![Static Badge](https://img.shields.io/badge/Agenda-green?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/other/Agenda.md) [![Static Badge](https://img.shields.io/badge/CheatSheet-purple?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/CheatSheet.md)  [![Static Badge](https://img.shields.io/badge/OCP-CLI-red?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/ocpcli-cheatsheet.md)   [![Static Badge](https://img.shields.io/badge/Labs-maroon?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/tree/main/labs-repo)  [![Static Badge](https://img.shields.io/badge/RedHat-OpenShift-maroon?style=flat&logo=Redhat&logoSize=auto)
](https://docs.redhat.com/en/documentation/openshift_container_platform/4.19)   [![Static Badge](https://img.shields.io/badge/Kubernetes-black?style=flat&logo=Kubernetes&logoSize=auto)
](https://kubernetes.io/docs/home/)
## 🔹 Module 2: Working with OpenShift Interfaces 

## Table of Contents 

- [2.0 - OpenShift Container Lifecycle](#20-openshift-container-lifecycle)
- [2.1 - OpenShift CLI and Web Consol](#21-openshift-cli-oc-and-web-console)
- [2.2 - Creating and Managing Projects (Learning YAML)](#22-creating-and-managing-projects-learning-yaml) | [Lab](#-hands-on-lab-yaml) | [Quiz](#quiz-yaml)

### 2.0 OpenShift Container Lifecycle

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module2#-module-2-working-with-openshift-interfaces" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

   <p align="center">
    <img src="/images/lifecycle.png" alt="OpenShift Training" style="width:500px; align="center"/>
    </p>

  > So this diagram shows how everything is connected in a Kubernetes or OpenShift environment. It starts with the cluster, which manages `deployments` and contains `nodes` where the actual workloads run. `Deployments` create `ReplicaSets` or `ReplicationControllers`, which make sure the right number of pods are running. Each `pod` holds one or more containers, and `services` expose these pods and route traffic to them, even providing static IPs and load balancing. Finally, `routes` are used to expose services to the outside world so users can access the apps. So continuing from there—this setup ensures high availability and scalability. For example, if a `pod` crashes, the `ReplicaSet` or `ReplicationController` automatically replaces it to keep the application running smoothly. `Services` play a critical role in making pods accessible by abstracting the backend details, so you don’t have to know the exact `pod` IP to connect to an app. And with `routes` in OpenShift, you can expose those `services` externally, enabling users to reach your apps over the internet or your corporate network. All of this happens within the `cluster`, which acts like the brain coordinating everything, from managing deployments to distributing workloads across nodes and maintaining desired state. It's a powerful system for automating and managing containerized applications.


---
### 2.1 OpenShift CLI (oc) and Web Console

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module2#-module-2-working-with-openshift-interfaces" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>


- The **Web Console** offers tabs for Overview, Workloads, Networking, Builds, Storage, Monitoring, etc.
- The **CLI (`oc`)** gives full control with extra features over `kubectl`.
- [OpenShit CLI Cheat Sheet](https://github.com/ocp-workshop-wf/bootcamp/blob/main/ocpcli-cheatsheet.md)

  **[Resource OpenShit CLI](https://www.youtube.com/watch?v=MYH8nX9J1lc)**

---

**Hands-on Walkthroughs:**  
  - Login using 

    <p align="left">
    <img src="/images/login.png" alt="OpenShift Training" style="width:200px; align="left"/>
    </p>

  ```bash
  oc login -u <username> -p <password> <url>
  ```
  - Create a [project](https://docs.redhat.com/en/documentation/openshift_container_platform/4.9/html/building_applications/projects): 

    <p align="left">
    <img src="/images/projectUI.png" alt="OpenShift Training" style="width:500px; align="left"/>
    </p>
    
  ```bash
  oc new-project demo-project
  ```
    <p align="center">
    <img src="/images/project.png" alt="OpenShift Training" style="width:500px; align="center"/>
    </p>

    > OpenShift uses project to divide up applications. Projects also allow OpenShift to guarantee resource availability for applications through the use of quotas. In this case, the quota represents an upper bound on the CPU, RAM and persistent storage that all pods in the project can use. Quotas prevent any single project from taking over all of the OpenShift's resources. While a quota does limit the amount of resources available to your application, it also guarantees the availability of those same resources.

  - List Pods?:

    ```bash
    oc get pods
    ```

  ---

  **Hands-on Walkthroughs**  
  - Log in to the OpenShift Web Console.

      - Check the status of OpenShift: Log in, log out
      
        ```bash
        oc status
        ```
      
      - Check which user
          
          ```bash
          oc whoami
          ```

      - Log out
          
          ```bash
          oc logout
          ```

  - Project Basics

    - See current project
      
      ```bash
      oc project
      ```
    - List all projects
      ```bash
      oc projects
      ```
    - Create a new project
      ```bash
      oc new-project myproject
      ```
    - Switch projects
      ```bash
      oc project <project name>
      ```
  ---

  - Identify the OpenShift components.

    - Pod Documentation 
      ```bash
      oc explain pod
      ```
    - Get details on the pod's spec
      ```bash
      oc explain pod.spec
      ```
    - Get details on the pod's containers
      ```bash
      oc explain pod.spec.containers
      ```
      
  - Creating Pod from files
    - Create pod on OpenShift
      ```bash
      oc create -f pods/pod.yaml
      ```
      <p align="center">
      <img src="/images/podcreated.png" alt="OpenShift Training" style="width:500px; align="center"/>
      </p>
        
    - Show all currently running Pods
      ```bash
      oc get pods
      ```
      > output: list pods that are deployed on your current namespace "Project"

    - Accessing the Pod
      - Access the shell inside the container 
          ```bash
          oc rsh <pod name>
          ```
        > output: `$`
      -  Request localhost

          ```bash
          wget localhost:8080
          ```
          > output: "Connecting to locathost:8080 ......" index.html saved
      - Read the output file
        
        ```bash    
        cat <file-name>
        ```
        <p align="center">
        <img src="/images/index.png" alt="OpenShift Training" style="width:500px; align="center"/>
        </p>

  - Deleting a Pod
      - Make sure to delete the pod by specifying the correct pod name
        ```bash
        oc delete pod <pod-name>
        ```
      - Show all currently running pods
        ```bash
        oc get pods
        ```

  - Watching a Pod use 2 different terminals.
    - Run the pod watch command on Terminal 1 
      ```bash
      # Terminal 1
      oc get pods --watch
      ```

    - Create the pod again! on Termnial 2
      ```bash
      # Terminal 2
      oc create -f pods/pod.yaml
      ```
        <p align="center">
        <img src="/images/pod-creation-process.png" alt="OpenShift Training" style="width:500px; align="center"/>
        </p>

  - Clean up: Make sure to delete the pod by specifying the correct pod name on Terminal 2
  ```bash
  oc delete pod <pod-name>
  ```
___

### 2.2 Creating and Managing Projects (Learning YAML)

  ***YAML*** : Yaml Ain't Markup Language
  - What it is? **YAML** is a human friendly data serialization standard for all programming languages. YAML is the defacto standard for storing Kubernetes and OpenShift resource definiations.
  
  </br>

  - What's the difference between json and yaml? 
  
<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module2#-module-2-working-with-openshift-interfaces" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

  <br/>

  ***YAML*** can handle more tasks than JSON, also it supports data types like integers and floats, while JSON doesn't. In Addition JSON uses { curly brackets } to list, ***YAML*** doesn't.

<br/>

- Lets take a look at Objects and child objects along with the `-` key value pairs in both:


| YAML    
| -------- 
  ```yaml
    outerKey:
      innerKey1: 1
      innerKey2: 2
  ```
  | JSON    
| -------- 
  ```json 
   { 
    outerKey:{
      innerKey1: 1,
      innerKey2: 2
    }
   }
  ```
#### Lists syntax in Yaml vs Json

  | YAML    
| -------- 
  ```yaml
  ingredients:
  - Tomato
  - Cilantro
  - Jalapeno
  ```
  | JSON    
| -------- 
  ```json 
  {ingredients: [
    "Tomato", "Jalapeno", "Cilantro"
    ]
  }
  ```
#### [Resource: Yaml.org](https://yaml.org/)

**Hands-on Walkthroughs**  

- Use YAML to understand Pod defenations following objects such as `apiVersion, kind, spec`.

    - Go ahead and open `/labs-repo/pods/pod.yaml`
    
      ```yaml
      apiVersion: v1
      kind: Pod 
      metadata:
        name: hello-world-pod
        labels:
          app: hello-world-pod
      spec:
        ..........
      ```
  - Adding another env-varaible to the container child object.
    <p align="center">
    <img src="/images/podyamldef.png" alt="OpenShift Training" style="width:500px; align="center"/>
    </p>

  - Dive deeper into pods 

    ```bash
    oc explain pod.spec.containers.env
    ```

### 🔬 Hands-on Lab (YAML): 
In this lab, you will create a custom Pod definition and upload it to OpenShift. This will test your skills in writing and debugging Pod YAML for the OpenShift platform.

- First, create a new file under the labs project. The path will be pods/lab-pod.yaml

- Copy the pod definition from another source into lab-pod.yaml

- Update the MESSAGE environment variable to a custom message. Update the name.

- Start a Pod based on lab-pod.yaml on the server

- Forward port 8080 from your computer to the running Pod
    > You can use to do this
    ```
    oc port-forward pod/lab-pod 8080
    ```

  - Port forwarding for Pods
  - Open a local port that forwards traffic to a pod
    ```bash
    oc port-forward <pod name> <local port>:<pod port>
    ```

  - Example of 8080 to 8080 for hello world
    ```bash
    oc port-forward hello-world-pod 8080:8080
    ```
---
#### Checklist 📋 YAML: 
Once you meet all of these criteria, you have successfully completed the lab. You will need to run the commands yourself in order to grade this lab.

- Output from oc get pods contains the pod.

- Output from oc describe pod/lab-pod has the correct name and MESSAGE environment value.

- `curl localhost:8080` prints the message.
---

### Quiz (YAML)

> Q1: What is the command to log in to OpenShift on the command line?
- [ ] `oc log on <cluster URL>`
- [ ] `oc login <cluster URL>`
- [ ] `minishift log on <cluster URL>`
- [ ] `minishift login <cluster URL>`
<details>
  <summary> Answer </summary>

   `oc login <cluster URL>`

</details>


> Q2: What is the command to create a project?
- [ ] `oc project <name>`
- [ ] `oc projects <name>`
- [ ] `oc new-project <name>`
- [ ] `oc new project <name>`
<details>
  <summary> Answer </summary>

   `oc new-project <name>`

</details>

> Q3: What is the command to switch your current project to a different project?
- [ ] `oc project <name>`
- [ ] `oc projects <name>`
- [ ] `oc new-project <name>`
- [ ] `oc new project <name>`
<details>
  <summary> Answer </summary>

   `oc project <name>`

</details>

> Q4: What is the minimum number of containers in a pod?
- [ ] 1
- [ ] 2
- [ ] 5
- [ ] 0
<details>
  <summary> Answer </summary>

   1
</details>

> Q5: What is the YAML type of pod.spec.containers
- [ ] String
- [ ] squence or list
- [ ] <[ ]object>
- [ ] object
<details>
  <summary> Answer </summary>

   squence or list

</details>

> Q6: What is the command to create a pod on OpenShift based on its YAML file?
- [ ] `oc start pod <pod-file.yaml>`
- [ ] `oc create <pod-file.yaml>`
- [ ] `oc create-f <pod-file.yaml>`
- [ ] `oc start pod -f <pod-file.yaml>`

<details>
  <summary> Answer </summary>

   `oc create-f <pod-file.yaml>`

</details>

> Q7: What is the command to start a shell in a pod on OpenShift?
- [ ] `oc shell <pod-name>`
- [ ] `oc exec <pod-name>`
- [ ] `oc sh <pod-name>`
- [ ] `oc rsh <pod-name>`

<details>
  <summary> Answer </summary>

   `oc rsh <pod-name>`

</details>

--- 

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module3" target="_blank">
    <img src="/images/nexticon.webp" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

<p align="left">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module1" target="_blank">
    <img src="/images/backred1.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>