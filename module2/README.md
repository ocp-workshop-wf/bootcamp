## Module 2: Working with OpenShift Interfaces 

### 2.1 OpenShift CLI (oc) and Web Console

- The **Web Console** offers tabs for Overview, Workloads, Networking, Builds, Storage, Monitoring, etc.
- The **CLI (`oc`)** gives full control with extra features over `kubectl`.

**Lab:**  
- Login using `oc login`
- Create a project: `oc new-project demo-project`
- List pods: `oc get pods`

**YouTube:**  
- [OpenShift CLI Basics](https://www.youtube.com/watch?v=8wFJe2U1GdI)

**Lab:**  
- Log in to the OpenShift Web Console.

    > Check the status of OpenShift
    > 
    > ##### Log in, log out
    > `oc status`

    >> ##### Uses the pre-configured OpenShift cluster
    >  `oc login`

    >> ##### Allows you to log in to any OpenShift cluster
    > `oc login <cluster address>` 
    
    <br/>

    >> ##### Check which user
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

- Identify the OpenShift components.

    ##### Pod Documentation 

    >> `oc explain pod`
    > ##### Get details on the pod's spec
    >> `oc explain pod.spec`
    > ##### Get details on the pod's containers
    > `oc explain pod.spec.containers`
    
   ##### Creating Pod from files
    >> ##### Create pod on OpenShift
    > `oc create -f pods/pod.yaml`
    >> ##### Show all currently running Pods
    > `oc get pods`

   ##### Accessing the Pod
    >> ##### Access the shell inside the container 
    > `oc rsh <pod name>`
    >> ##### Request localhost
    > `wget localhost:8080`
    >> ##### Read the output file
    > `cat <file-name>`

   ##### Deleting a Pod
    >> ##### Make sure to delete the pod by specifying the correct pod name
    > `oc delete pod <pod-name>`
    >> ##### Show all currently running pods
    > `oc get pods`

   ##### Watching a Pod use 2 different terminals.
    >> ##### Run the pod watch command on Terminal 1 
    > `oc get pods --watch`

    >> ##### Create the pod again! on Termnial 2
    >> `oc create -f pods/pod.yaml`

    >> ##### Make sure to delete the pod by specifying the correct pod name on Terminal 2
    > `oc delete pod <pod-name>`
  
___
### Quiz

#### Q1: What is the command to log in to OpenShift on the command line?
- [ ] `oc log on <cluster URL>`
- [ ] `oc login <cluster URL>`
- [ ] `minishift log on <cluster URL>`
- [ ] `minishift login <cluster URL>`
<details>
  <summary> Answer </summary>

   `oc login <cluster URL>`

</details>


#### Q2: What is the command to create a project?
- [ ] `oc project <name>`
- [ ] `oc projects <name>`
- [ ] `oc new-project <name>`
- [ ] `oc new project <name>`
<details>
  <summary> Answer </summary>

   `oc new-project <name>`

</details>

#### Q3: What is the command to switch your current project to a different project?
- [ ] `oc project <name>`
- [ ] `oc projects <name>`
- [ ] `oc new-project <name>`
- [ ] `oc new project <name>`
<details>
  <summary> Answer </summary>

   `oc project <name>`

</details>

---
---

### 2.2 Creating and Managing Projects

- Projects = Namespaces + metadata/RBAC.
- Add quotas, manage access via RBAC.

**Lab:**  
- Create a project, set a resource quota.
- Assign a user with:  
  `oc adm policy add-role-to-user edit <username> -n demo-project`

**YouTube:**  
- [Getting Started with OpenShift Web Console](https://www.youtube.com/watch?v=Qzvfi6VSFoc)

---
