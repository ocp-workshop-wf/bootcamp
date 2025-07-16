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
- ReplicationControllers: DeploymentConfigs use ReplicationControllers to run their pods.
  - Deploy the application!
    ```
    oc new-app quay.io/practicalopenshift/hello-world --as-deployment-config
    ```
  - We need to look into more options into our DeploymentConfig
    ```
    oc get -o yaml dc/hello-world
    ```
  


---

### 3.3 OpenShift Networking

- **Servcies** 
- **Routes** 

**Lab:**  
- C


---

