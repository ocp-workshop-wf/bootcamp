## Module 3: Core OpenShift Resources 

### 3.1 OpenShift Resources Overview

- **DeploymentConfig:** Define the template for a pod and manages deploying new images or configuration changes.
  - Depoly images.
  - Deploy from Git.
  - ReplicationControllers
  - Basic Configuration
  #### Advanced Deployment configs
  - Deploy Triggers
  - Lifecycle hooks
  - Health Checks
- **Replicas:** Is the number of desired replicas.
![Deployment](/images/deployment.png)

**Lab:**  
- Deploy an existing image based on its tag: `oc new-app <image tag>`
  - For this lesson
  ```
  oc new-app quay.io/practicalopenshift/hello-world --as-deployment-config
  ```
- Check running resources: 
  ```
  oc status
  ```
- Check pods: 
  ```
  oc get pods
  ```
---
- Cleaning Up after testing things out: `oc status` to make sure that your deployment is still running 
  - Get to see the service: 
    ```
    oc get svc
    ```
  - Get to see the deployment config: 
    ```
    oc get dc
    ```
  - Get to see image stream: 
    ```
    oc get istag
    ```
  - Let delete using full name of the resource: 
    ```
    oc delete svc/hello-world
    ```
    ```
    oc get svc
    ```
  - Check the status again and see what was effected: 
    ```
    oc delete dc/hello-world
    ```
    
    ```
    oc status
    ```
---
- More advance way to clean up:
  - Run the application again: 
    ```
    oc new-app quay.io/practicalopenshift/hello-world --as-deployment-config
    ```
  - Check the detatils for that DeploymentConfig: 
    ```
    oc describe dc/hello-world
    ```
  - Clean up using a label selector: 
    ```
    oc delete all -l <label-selector>
    ```
---
- Name your DeploymentConfigs:
  ```
  oc new-app quay.io/practicalopenshift/hello-world --name demo-app --as-deployment-config
  ```
  - Describe your new named DC: 
    ```
    oc describe dc/demo-app
    ```
  - Lets add another app with a different name parameter: 
    ```
    oc new-app quay.io/practicalopenshift/hello-world --name demo-app-2 --as-deployment-config
    ```
  - Run status: 
    ```
    oc status
    ```
  - Delete the first app: 
    ```
    oc delete all -l app=demo-app
    ```
  - Run status: 
    ```
    oc status
    ```
  - Delete the 2nd app: 
    ```
    oc delete all -l app=demo-app-2
    ```

---

### 3.2 OpenShift and how to deploy applications

- Direct deployment Git (for GitHub)
- Source-to-Image (S2I) "Later in this course!"
- OpenShift Pipelines (Tekton)

#### Direct deployment using Git

![Git Deployment](/images/deployfromgit.png)

**Lab:**  
- Deploy the app:
    ```
    oc new-app https://gitlab.com/therayy1/hello-world.git --as-deployment-config 
    ```
  - When you run this command, OpenShift automatically creates a BuildConfig for your application. The BuildConfig contains all necessary instructions for building the image, similar to how Docker build commands operate.
  - OpenShift will then clone the repository from the provided Git URL and proceed to build the image by executing the Dockerfile steps contained in your application. Each step results in an intermediate container.
  - Once the build successfully completes, OpenShift pushes the built image to an ImageStream, which you can utilize for deployment purposes.

- Track the progress of the build
    ```
    oc status
    ```
    - Check the line for `bc` "BuildConfig"
    ```
    oc logs -f bc/hello-world
    ```
    - Lets go delete all and check the output
    ```
    oc delete all -l app=hello-world
    ```
    - You must see that `buildconfig, build & golang imagestream` got delete along with the others.
---
- ReplicationControllers: DeploymentConfigs use ReplicationControllers to run their pods.
  - Deploy the application!
    ```
    oc new-app quay.io/practicalopenshift/hello-world --as-deployment-config
    ```
  - We need to look into more options into our DeploymentConfig
    ```
    oc get -o yaml dc/hello-world
    ```
  - Lets get the ReplicationController
    ```
    oc get rc
    ```
    - Your output should include your ReplicationController.
---
- Rollout and Rollback 
  - You need 2 different split terminals
  - Make sure that your application is running on the OCP cluster.
    ```
    oc new-app quay.io/practicalopenshift/hello-world --as-deployment-config
    ```
  - Terminal 1:
    ```
    oc get pods --watch 
    ```
    or 
    ```
    oc get pods -w
    ```
  - Terminal 2:
    ```
    oc rollout latest dc/hello-world
    ```
    - The first thing OCP does is to start a new deployment `starting from, Pending - ContainerCreating - Running` once its Running the previous version `Terminating` "Start new - Stop old"
    ```
    oc rollback dc/hello-world
    ```
    - It is very similar process "Start the previous version, and Stop current"
---

### Hands-on Lab: 
In the DeploymentConfig lab, you will create a custom DeploymentConfig based on the hello-world image by changing some parameters.

- First, `use oc new-app to start an application based on quay.io/practicalopenshift/hello-world`

- Use `oc new-app` to start a second version of the application using the name `lab-dc`, a custom value for the `MESSAGE` environment variable, and the same hello-world image

  - You can specify environment variables in `oc new-app` with a flag. `oc new-app --help` can help you to find the correct one

- Forward port 8080 on your local computer to port 8080 on the second pod you created
---
### Checklist: 
- Output from `oc get pods` contains two pods

- Output from `oc describe dc/lab-dc` has the correct name and `MESSAGE` environment value

- `curl localhost:8080` prints the message you entered in step 2

---
> Cleaning Up:
 To clean up, use a single command to delete all of the resources created in step 1. You are done when `oc get dc` just has the `lab-dc` DeploymentConfig.

---
### Quiz

>Q1: What is the command to deploy the hello-world image to OpenShift as a deployment config?
- [ ] `oc new-app quay.io/practicalopenshift/hello-world`
- [ ] `oc start-app quay.io/practicalopenshift/hello-world`
- [ ] `oc new-apps quay.io/practicalopenshift/hello-world`
- [ ] `oc create -f quay.io/practicalopenshift/hello-world`
<details>
  <summary> Answer </summary>

   `oc new-app quay.io/practicalopenshift/hello-world`

</details>

> Q2: What OpenShift resource is responsible for running the correct number of pods for a DeploymentConfig?
- [ ] DeploymentConfigPodRunner
- [ ] ReplicaSet
- [ ] ReplicationController
- [ ] DeploymentConfig
<details>
  <summary> Answer </summary>

   ReplicationController

</details>

> Q3: What is the command to create a DeploymentConfig based on a Git repository?
- [ ] `oc new-build https://github.com/practical-openshift/hello-world`
- [ ] `oc start-build quay.io/practicalopenshift/hello-world`
- [ ] `oc new-deploy quay.io/practicalopenshift/hello-world`
- [ ] `oc new-app https://github.com/practical-openshift/hello-world`
<details>
  <summary> Answer </summary>

  `oc new-app https://github.com/practical-openshift/hello-world`

</details>

> Q4: What's the easiest way to delete all the resources made from oc new-app?
- [ ] Use oc delete once for each resource you need to delete
- [ ] Use label selectors with oc delete all
- [ ] Use oc undeploy with the name in oc new-app
- [ ] There is no easy way!
<details>
  <summary> Answer </summary>

   Use label selectors with oc delete all

</details>

> Q5: What is the command to roll out a new version of your DeploymentConfig?
- [ ] `oc update dc/app-name`
- [ ] `oc rollout dc/app-name`
- [ ] `oc rollout latest dc/app-name`
- [ ] `oc update latest dc/app-name`
<details>
  <summary> Answer </summary>

  `oc rollout latest dc/app-name`

</details>
### 3.3 OpenShift Networking

- **Servcies** 
- **Routes** 

**Lab:**  
- 


---

