## Module 2: Working with OpenShift Interfaces 

### 2.1 OpenShift CLI (oc) and Web Console

- The **Web Console** offers tabs for Overview, Workloads, Networking, Builds, Storage, Monitoring, etc.
- The **CLI (`oc`)** gives full control with extra features over `kubectl`.

**Lab:**  
- Login using 
  ```
  oc login
  ```
- Create a project: 
  ```
  oc new-project demo-project
  ```
- List pods: 
  ```
  oc get pods
  ```

**YouTube:**  
- [OpenShift CLI Basics](https://docs.redhat.com/en/documentation/openshift_container_platform/4.2/html/cli_tools/openshift-cli-oc)

**Lab:**  
- Log in to the OpenShift Web Console.

    > Check the status of OpenShift
    > 
    > ##### Log in, log out
    ```
    oc status
    ```

    >> ##### Uses the pre-configured OpenShift cluster
    ```
    oc login
    ```

    >> ##### Allows you to log in to any OpenShift cluster
    ```
    oc login <cluster address>
    ```
    
    <br/>

    >> ##### Check which user
    ```
    oc whoami
    ```

    >> ##### Log out
    ```
    oc logout
    ```


    > #### Project Basics

    >> ##### See current project
    ```
    oc project
    ```

    >> ##### Create a new project
    ```
    oc new-project myproject
    ```

    >> ##### List all projects
    ```
    oc projects
    ```

    >> ##### Switch projects
    ```
    oc project <project name>
    ```
___

- Identify the OpenShift components.

    ##### Pod Documentation 

    ```
    oc explain pod
    ```
    > ##### Get details on the pod's spec
    ```
    oc explain pod.spec
    ```
    > ##### Get details on the pod's containers
    ```
    oc explain pod.spec.containers
    ```
    
   ##### Creating Pod from files
    >> ##### Create pod on OpenShift
    ```
    oc create -f pods/pod.yaml
    ```
    >> ##### Show all currently running Pods
    ```
    oc get pods
    ```

   ##### Accessing the Pod
    >> ##### Access the shell inside the container 
    ```
    oc rsh <pod name>
    ```
    >> ##### Request localhost

    ```
    wget localhost:8080
    ```
    >> ##### Read the output file
     
    ```
    cat <file-name>
    ```

   ##### Deleting a Pod
    >> ##### Make sure to delete the pod by specifying the correct pod name
    ```
    oc delete pod <pod-name>
    ```
    >> ##### Show all currently running pods
    ```
    oc get pods
    ```

   ##### Watching a Pod use 2 different terminals.
    >> ##### Run the pod watch command on Terminal 1 
    ```
    oc get pods --watch
    ```

    >> ##### Create the pod again! on Termnial 2
    ```
    oc create -f pods/pod.yaml
    ```

    >> ##### Make sure to delete the pod by specifying the correct pod name on Terminal 2
    ```
    oc delete pod <pod-name>
    ```
  
___

### 2.2 Creating and Managing Projects (Learning YAML)

  ***YAML*** : Yaml Ain't Markup Language
  - What it is? **YAML** is a human friendly data serialization standard for all programming languages.
  - What's the difference between json and yaml? <br/>
  ***YAML*** can handle more tasks than JSON, also it supports data types like integers and floats, while JSON doesn't. In Addition JSON uses { curly brackets } to list, ***YAML*** doesn't.

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
### Arrays in Yaml vs Json

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
### [Resource: Yaml.org](https://yaml.org/)

**Lab:**  

- Use YAML to understand Pod defenations following objects such as `apiVersion, kind, spec`.

    ##### Go ahead and open `/labs-repo/pods/pod.yaml`
    
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
    >> ##### Dive deeper into pods 
      ```
      oc explain pod.spec.containers.env
      ```

### Hands-on Lab: 
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

    > #### Port forwarding for Pods
    >> ##### Open a local port that forwards traffic to a pod
    ```
    oc port-forward <pod name> <local port>:<pod port>
    ```

    > #### Example of 8080 to 8080 for hello world
      ```
      oc port-forward hello-world-pod 8080:8080
      ```
---
### Checklist: 
Once you meet all of these criteria, you have successfully completed the lab. You will need to run the commands yourself in order to grade this lab.

- Output from oc get pods contains the pod.

- Output from oc describe pod/lab-pod has the correct name and MESSAGE environment value.

- curl localhost:8080 prints the message.
---

### Quiz

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
- [ ] <[]object>
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