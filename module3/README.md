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

### 3.2 OpenShift to deploy applications

- Direct deployment Git (for GitHub)
- Source-to-Image (S2I)
- OpenShift Pipelines (Tekton)

#### Direct deployment using Git

![Git Deployment](/images/deployfromgit.png)

**Lab:**  
- Deploy the app:
    ```
    oc new-app https://gitlab.com/therayy1/hello-world.git --as-deployment-config 
    ```


---

### 3.3 Image Stream Management

- **ImageStreams** let you manage and tag images, enabling automated redeployments.
- **BuildConfig/S2I** automates builds from source code.

**Lab:**  
- Create an ImageStream, BuildConfig.
- Push new code/image, watch auto-update.

**YouTube:**  
- [OpenShift Image Streams Explained](https://www.youtube.com/watch?v=jA-RH0jO-J8)

---

