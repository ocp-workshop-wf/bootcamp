[![Static Badge](https://img.shields.io/badge/Agenda-green?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/other/Agenda.md) [![Static Badge](https://img.shields.io/badge/CheatSheet-purple?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/other/CheatSheet.md) [![Static Badge](https://img.shields.io/badge/OCP-CLI-red?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/other/ocpcli-cheatsheet.md)   [![Static Badge](https://img.shields.io/badge/Labs-maroon?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/tree/main/labs-repo)  [![Static Badge](https://img.shields.io/badge/Kubernetes-black?style=flat&logo=Kubernetes&logoSize=auto)
](https://kubernetes.io/docs/home/) [![Static Badge](https://img.shields.io/badge/RedHat-OpenShift-maroon?style=flat&logo=Redhat&logoSize=auto)
](https://docs.redhat.com/en/documentation/openshift_container_platform/4.19)  

## 🔹 Module 3: Core OpenShift Resources

## Table of Contents 
- [3.1 - OpenShift Resources Overview](#31-openshift-resources-overview) 
 
- [3.2 - OpenShift and how to deploy applications](#32-openshift-and-how-to-deploy-applications) | [Lab](#-hands-on-lab-deploying-application) | [Quiz](#quiz-deploying-application)

- [3.3 - Resource Quotas and Limits](#33-resource-quotas-and-limits) | [Lab](#-hands-on-lab-resource-quota) | [Quiz](#quiz-resource-quota)

- [3.4 - OpenShift Volumes](#34-openshift-volumes) | [Lab](#-hands-on-lab-volumes) | [Quiz](#quiz-volumes)

- [3.5 - OpenShift Networking](#35-openshift-networking) | [Lab](#-hands-on-lab-network) | [Quiz](#quiz-network)


### 3.1 OpenShift Resources Overview

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module3#-module-3-core-openshift-resources" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

**DeploymentConfig:** Define the template for a pod and manages deploying new images or configuration changes. The DeploymentConfig template for a pod uses the same format for its pod template. This template is found in the DeploymentConfig `spec` under the `template` property. The other important thing that DeploymentConfigs have is the `replicas` parameter in the `spec` field. This configuration option tells the DeploymentConfig how many instances of a `pod` it needs to run. If the DeploymentConfig sees that there are not enough instancs, it will start new pods according to the `template` until it reaches the number specified in the replica's field. Similarly, if you change the replica's value to a lower number, the DeploymentConfig will start deleting pods until it reaches the target number.There is a lot of other behavior that DeploymentConfigs will handle for you,such as automatically triggering new deployments, controlling details about how deployments are conducted, and adding custom behavior using lifecycle hooks.

  | In This Module | Advanced Module | 
  | -------------- | --------------- | 
  | Deploy Images | | 
  | Deploy from Git | Deploy Triggers |
  | ReplicationControllers | Lifecycle hooks  |
  | Basic Configurations| Health Check|


  <p align="center">
  <img src="/images/dc.png" alt="OpenShift Training" style="width:300px; align="center"/>
  </p>


- DeploymentConfig VS Deployment

| Point | Deployment Config | Deployment | 
| ----- | ----------------- | ---------- |
| Design | It prefers consistency. As deployer pod goes down, its not replaced but waits for it to come up again | It prefers availability over consistency. The controller manager runs in high availability mode. Hence if one controller goes down, other leader is elected. |
| Automatic Rollbacks | Supported | Not Supported | 
| Automatic trigger on config changes | Has to be mentioned explicity in resource definition | Implicit |
| Lifecycle hooks | Supported | Not Supported |
| Custom Deployment Strategies | Supported | Not Supported| 



<p align="center">
<img src="/images/deployment.png" alt="OpenShift Training" style="width:400px; align="center"/>
</p>

---

**Hands-on Walkthroughs**  

- Deploy an existing image based on its tag: `oc new-app <image tag>`
  - For this lesson lets deploy using an image from quay.io lets start deploying using Kubernetes Resource first which is Deployment
      ```bash
    oc new-app quay.io/practicalopenshift/hello-world 
    ```
    > output: check your cluster ui and find your deployment yaml

    <p align="center">
    <img src="/images/deployment-ui.png" alt="OpenShift Training" style="width:500px; align="center"/>
    </p>

  - Now lets clean all and re-deploy using OpenShift Resource DeploymentConfig.
    ```bash
    oc delete all -l app=hello-world
    ```
    > output: 
    ```bash
    pod "hello-world-5b784bc45f-9f87f" deleted
    service "hello-world" deleted
    deployment.apps "hello-world" deleted
    imagestream.image.openshift.io "hello-world" deleted
    ```
    - Re-run the same command but add `--as-deployment-config` 
      ```bash
      oc new-app quay.io/practicalopenshift/hello-world --as-deployment-config
      ```
      > output: When DeploymentConfig is deployed it deploys an `ImageStream`, `Service` and `Pods` lets checkout the YAML of DeploymentConfig through the UI

    - Check running resources:

      ```bash
      oc status
      ```
    - Check pods:
      ```bash
      oc get pods
      ```
    - Get to see the service:
      ```bash
      oc get svc
      ```
    - Get to see the deployment config:
      ```bash
      oc get dc
      ```
    - Get to see image stream:
      ```bash
      oc get istag
      ```
    - Let delete using full name of the resource:
      ```bash
      oc delete svc/hello-world
      ```
      ```bash
      oc get svc
      ```
    - Check the status again and see what was effected:
      ```bash
      oc status
      ```
    - Delete specific DeploymentConfig
      ```bash
      oc get dc
      ```
    - Copy your DeploymentConfig name you need to delete in this case its `hello-world`
      ```bash
      oc delete dc/hello-world
      ```
    - Check Status again
      ```bash
      oc status
      ```
    - Clean up if something is leftover by running 
      ```bash
      oc delete all -l app=hello-world
      ```
---

- More advance way to clean up:
  - Run the application again:

    ```bash
    oc new-app quay.io/practicalopenshift/hello-world --as-deployment-config
    ```

  - Another way to Check the detatils for that DeploymentConfig and find the `label`:

    ```bash
    oc describe dc/hello-world
    ```
    <p align="center">
    <img src="/images/dc-label.png" alt="OpenShift Training" style="width:400px; align="center"/>
    </p>

    ```yaml
    metadata:
      .......
      labels:
        app: hello-world
      ......
    spec:
    ......
    ```

  - Clean up using a label selector in this case it is `app=hello-world`

    ```bash
    oc delete all -l app=hello-world
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

  <p align="right">
    <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module3#-module-3-core-openshift-resources" target="_blank">
      <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
    </a>
  </p>

  - Direct deployment Git (for GitHub)
  - Source-to-Image (S2I) "Later in this course!"
  - OpenShift Pipelines (Tekton) "will not be covered in this course"

#### Direct deployment using Git
As long as your source code is available online, `oc new-app` can build an image based on your application in the Git repository.

---

**Hands-on Walkthroughs**  

- Deploy the app from a Git Repo:

    ```bash
    oc new-app https://gitlab.com/therayy1/hello-world.git --as-deployment-config 
    ```

  ```bash
      Tags: base rhel8

      * An image stream tag will be created as "golang:1.17" that will track the source image
      * A Docker build using source code from https://gitlab.com/therayy1/hello-world.git will be created
        * The resulting image will be pushed to image stream tag "hello-world:latest"
        * Every time "golang:1.17" changes a new build will be triggered
      * This image will be deployed in deployment config "hello-world"
      * Port 8080/tcp will be load balanced by service "hello-world"
        * Other containers can access this service through the hostname "hello-world"
      * WARNING: Image "quay.io/projectquay/golang:1.17" runs as the 'root' user which may not be permitted by your cluster administrator
  --> Creating resources ...
      imagestream.image.openshift.io "golang" created
      imagestream.image.openshift.io "hello-world" created
      buildconfig.build.openshift.io "hello-world" created
      deploymentconfig.apps.openshift.io "hello-world" created
      service "hello-world" created
  --> Success
  ```
  > output: When you run `oc new-app` with a Git repository, OpenShift will try to build the image from the Dockerfile for that reason, OpenShift will attempt to download the `golang:alpine` image that we specified in the docker from instruction. When OpenShift imports an image into its internal registry,it creates an `image stream tag` as explained in this output the 2nd line says that docker build using source code from Git repo will be created, and thats why it take longer time to get deployed! Not only can OpenShift download image stream tag from repositories on the internet, but it can also create new image stream tags as a result of `builds`. OpenShift also sets up a `trigger` for the `golang:alpine` image that will start a new build when the `golang:alpine` image stream tag changes. The next line says that this will deploy to a new deployment config called `hello-world `port `8080` will be load balanced by a service. 

  - Check the line for `bc` "BuildConfig"

    ```bash
    oc logs -f bc/hello-world
    ```
  > output: 1st OpenShift will clone the repo, after cloning it will go step by step through the same dockerfile build process, after the build command succeeds OpenShift will push the image up the `hello-world` image stream. 

    - lets check our status to list what was deployed
    ```bash
    oc status
    ```
  > output: Its the same as we saw before but with an additional `buildconfig` - `bc` and we will learn about BuildConfig and Builds later in this course.
  ```bash
    svc/hello-world - 172.30.20.13:8080
    dc/hello-world deploys istag/hello-world:latest <-
    bc/hello-world docker builds https://gitlab.com/therayy1/hello-world.git on istag/golang:1.17 
    deployment #1 deployed 14 minutes ago - 1 pod
  ```

  - Lets go delete all and check the output

    ```bash
    oc delete all -l app=hello-world
    ```

  > output: You must see that `buildconfig, build & golang imagestream` got delete along with the others.
  ```bash
  replicationcontroller "hello-world-1" deleted
  service "hello-world" deleted
  deploymentconfig.apps.openshift.io "hello-world" deleted
  buildconfig.build.openshift.io "hello-world" deleted
  imagestream.image.openshift.io "golang" deleted
  imagestream.image.openshift.io "hello-world" deleted
  ```

---

- ReplicationControllers: DeploymentConfigs use ReplicationControllers to run their pods.

  <p align="center">
  <img src="/images/replicationcontroller.png" alt="OpenShift Training" style="width:400px; align="center"/>
  </p>

  - Deploy the application and check the UI!

    ```bash
    oc new-app quay.io/practicalopenshift/hello-world --as-deployment-config
    ```

  - We need to look into more options into our DeploymentConfig yaml - either from the Terminal by running the following command or from the UI from `YAML` tab.

    ```
    oc get -o yaml dc/hello-world
    ```
  > output: scroll to the top and you should see the following, The `replicationcontrollers` only job is to run a certain number of pods according to the `template`, where as the `deploymentconfig` has a number of other jobs related to orchestrating deployments. when the `deploymentconfig` is first created, it created a `replicationcontroller` and the `replicationcontroller` will start the pods. When you deploy a new version of your application using `deploymentconfig` will create a new `replicationcontroller` with the new version of the pod. Once the `replicationController` is ready the `deploymentConfig` will switch traffic over to the new `replicationController` pods, and then delete the old `replicationController`. This is the default behavior but ofcourse its configurable.
  ```yaml
  ....
  spec:
    replicas: 1
  ....
  ```

  - Lets list the ReplicationController

    ```bash
    oc get rc
    ```
  > output: Your output should include your ReplicationController.

  <p align="center">
  <img src="/images/rc-complete.png" alt="OpenShift Training" style="width:400px; align="center"/>
  </p>

---

- Rollout and Rollback
  - You need 2 different split terminals
  - Make sure that your application is running on the OCP cluster.

    ```bash
    oc new-app quay.io/practicalopenshift/hello-world 
    ```

  - Terminal 1:

    ```bash
    oc get pods --watch 
    ```

    or

    ```bash
    oc get pods -w
    ```

  - Terminal 2:

    ```bash
    oc rollout latest deplpoyment/hello-world
    ```

    > output: The first thing OCP does is to start a new deployment `starting from, Pending - ContainerCreating - Running` once its Running the previous version `Terminating` "Start new - Stop old. This is very difficult to do manually, but OpenShift contains sensible defaults that will handle the deployment for you.

  <p align="center">
  <img src="/images/rollout.png" alt="OpenShift Training" style="width:400px; align="center"/>
  </p>
    <p align="center">
  <img src="/images/rollout2.png" alt="OpenShift Training" style="width:400px; align="center"/>
  </p>
      <p align="center">
  <img src="/images/rollout3.png" alt="OpenShift Training" style="width:400px; align="center"/>
  </p>
  

  - Lets say that you've deployed a new version of your application, but you have monitoring in place and you've detected that you application's new version has an error. In this case you need to run `rollback`

    ```bash
    oc rollback deployment/hello-world
    ```

  > output: It is very similar process "Start the previous version, and Stops the current"
    <p align="center">
    <img src="/images/rollback0.png" alt="OpenShift Training" style="width:400px; align="center"/>
    </p>

    <p align="center">
    <img src="/images/rollback1.png" alt="OpenShift Training" style="width:400px; align="center"/>
    </p>

---

### 🔬 Hands-on Lab (Deploying Application)

In the DeploymentConfig lab, you will create a custom DeploymentConfig based on the hello-world image by changing some parameters.

- First, `use oc new-app to start an application based on quay.io/practicalopenshift/hello-world`

- Use `oc new-app` to start a second version of the application using the name `lab-dc`, a custom value for the `MESSAGE` environment variable, and the same hello-world image

  - You can specify environment variables in `oc new-app` with a flag. `oc new-app --help` can help you to find the correct one

- Forward port 8080 on your local computer to port 8080 on the second pod you created

---

### Checklist 📋 Deploying Application

- Output from `oc get pods` contains two pods

- Output from `oc describe deployment/lab-dc` has the correct name and `MESSAGE` environment value

- `curl localhost:8080` prints the message you entered in step 2

---

> 💡 Cleaning Up:
 To clean up, use a single command to delete all of the resources created in step 1. You are done when `oc get dc` just has the `lab-dc` DeploymentConfig.

---

### Quiz Deploying Application

> Q1: What is the command to deploy the hello-world image to OpenShift as a deployment?

- [ ] `oc new-app quay.io/practicalopenshift/hello-world`
- [ ] `oc start-app quay.io/practicalopenshift/hello-world`
- [ ] `oc new-apps quay.io/practicalopenshift/hello-world`
- [ ] `oc create -f quay.io/practicalopenshift/hello-world`

<details>
  <summary> Answer </summary>

   `oc new-app quay.io/practicalopenshift/hello-world`

</details>

> Q2: What OpenShift resource is responsible for running the correct number of pods for a Deployment?

- [ ] DeploymentConfigPodRunner
- [ ] ReplicaSet
- [ ] ReplicationController
- [ ] DeploymentConfig

<details>
  <summary> Answer </summary>

   ReplicaSet

</details>

> Q3: What is the command to create a DeploymentConfig based on a Git repository?

- [ ] `oc new-build https://github.com/practical-openshift/hello-world`
- [ ] `oc start-build quay.io/practicalopenshift/hello-world`
- [ ] `oc new-deploy quay.io/practicalopenshift/hello-world`
- [ ] `oc new-app https://github.com/practical-openshift/hello-world --as-deployment-config`

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

> Q5: What is the command to roll out a new version of your Deployment?

- [ ] `oc update deployment/app-name`
- [ ] `oc rollout deployment/app-name`
- [ ] `oc rollout latest deployment/app-name`
- [ ] `oc update latest deployment/app-name`

<details>
  <summary> Answer </summary>

  `oc rollout latest deployment/app-name`

</details>

---

### 3.3 Resource Quotas and Limits
  <p align="right">
    <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module3#-module-3-core-openshift-resources" target="_blank">
      <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
    </a>
  </p>  

  [Resource Quotas and Limits](https://docs.redhat.com/en/documentation/openshift_container_platform/4.8/html/building_applications/quotas) are used to control the amount of resources that can be consumed by a project or namespace in OpenShift. They help ensure fair resource allocation and prevent resource exhaustion.
  - **Resource Quotas**: Set limits on the total amount of resources (CPU, memory, storage) that can be used by all pods in a project. They help prevent a single project from consuming all available resources in the cluster.
  - **Resource Limits**: Set limits on the amount of resources that can be used by individual pods. They help ensure that no single pod can consume excessive resources, which could impact the performance of other pods in the cluster. 

---

  **Hands-on Walkthroughs**

  - Check-out a Resource Quota:

  ```yaml
    kind: ResourceQuota
    apiVersion: v1
    metadata:
      name: compute-deploy
      namespace: raafatadly23-dev
    spec:
      hard:
        limits.cpu: '30' # Total CPU limit for the project
        limits.memory: 30Gi # Total memory limit for the project
        requests.cpu: '3' # Total CPU requests for the project
        requests.memory: 30Gi # Total memory requests for the project
      scopes:
        - NotTerminating
  ``` 
  > Note: The `scopes` field is used to specify the scope of the quota. In this case, it is set to `NotTerminating`, which means that the quota applies only to non-terminating resources. Which that means that inside this namespace, you can only create pods that have a total of 30 CPU and 30Gi of memory requests and limits.

  - Check the status of the quota:

    ```bash
    oc describe quota compute-deploy
    ```
    > output: "You should see the hard limits and the current usage of resources in the project"

 > output:

  |Resource	|	Hard	|	Used |
  |--------	|	----	|	---- |
  | limits.cpu |      0  |   30 |
  | limits.memory |    0     |  30Gi |
  | requests.cpu |     0     |  3 |
  | requests.memory |  0     |  30Gi |

  - Create a pod that exceeds the resource limits:

    ```bash
    oc run my-pod --image=nginx --requests=cpu=5,memory=40Gi --limits=cpu=5,memory=40Gi
    ```
    > output: "You should see an error message indicating that the pod cannot be created due to resource limits"
    
    ```bash
    Error from server (Forbidden): pods "my-pod" is forbidden: exceeded quota: compute-deploy: must specify limits.cpu, limits.memory, requests.cpu, requests.memory
    ```
  - Create a pod that is within the resource limits:

    ```bash
    oc run my-pod2 --image=quay.io/practicalopenshift/hello-world --requests=cpu=250m,memory=64Mi --limits=cpu=500m,memory=500Mi
    ```
    > output: "You should see the pod created successfully" so that essentially means is that the cpu can only request x amount and then the ram can only request x amount of the memory, so you got memory and you got cpu here! So for this pod essentially what I'm trying to say is you can only `request` this 64Mi amount from `memory` and this 250m amount of `CPU` you can't go above that for this pod.

    **💡Important** When pods are done using memory they give that memory back so other resources can use it. Whereas `CPU` is not like that, when a pod is done using CPU it doesn't give it back to the cluster, so you need to set a limit on how much CPU it can use.

    ```bash
    pod/my-pod2 created
    ``` 

  - Check the status of the pod:

    ```bash
    oc get pods
    ```

    > output: "You should see the pod running"

    ```bash
    NAME       READY   STATUS    RESTARTS   AGE
    my-pod2    1/1     Running   0          1m
    ```
  - Check the status of the quota again:

    ```bash
    oc describe quota compute-deploy
    ```
    > output: "You should see the updated usage of resources in the project"
    ```yaml

    Resource		Hard		Used
    --------		----		----
    pods			10		3
    requests.cpu		4		2
    requests.memory	8Gi		4Gi
    limits.cpu		4		2
    limits.memory	8Gi		4Gi
    ```
  - Lets try it using a deployment kubernetes resource:
  ```bash
  oc create -f labs-repo/quota-lab/deployment-quota.yml
  ``` 
  > output: "You should see the deployment created successfully"

  ```bash
  oc get deployments
  ```
  > output: "You should see the deployment in the list of deployments"

  ```bash
  oc get deployments
  ``` 
  > output: "You should see the deployment in the list of deployments"

  ```bash
  NAME                READY   UP-TO-DATE   AVAILABLE   AGE
  my-deployment-quota   1/1     1            1           1m
  ``` 

  ```bash
  oc describe deployment my-deployment-quota
  ```
  > output: "You should see the deployment details including the resource limits and requests" and lets check the Yaml file of the deployment from the UI

---

### 🔬 Hands-on Lab (Resource Quota)
In the Resource Quota lab, you will create a custom Resource Quota based on the hello-world image by changing some parameters.
- First, `use oc create quota to create a Resource Quota with the following parameters:`
  - `pods=10`
  - `requests.cpu=4`
  - `requests.memory=8Gi`
  - `limits.cpu=4`
  - `limits.memory=8Gi`
- `Use oc describe quota to check the status of the Resource Quota`
- `Use oc run to create a pod that exceeds the resource limits`
- `Use oc run to create a pod that is within the resource limits`
- `Use oc get pods to check the status of the pods`
- `Use oc describe quota to check the status of the Resource Quota again`

---

### Checklist 📋 Resource Quota
- Output from `oc describe quota my-quota` contains the correct hard limits and current usage
- Output from `oc get pods` contains the pod that is within the resource limits
- Output from `oc describe quota my-quota` shows the updated usage of resources in the project

> 💡 Cleaning Up:

  To clean up, use a single command to delete all of the resources created in step 1

  ```bash
  oc delete all --all -n my-project
  ```
---

### Quiz Resource Quota
> Q1: What is the purpose of Resource Quotas in OpenShift?
- [ ] To limit the number of pods in a project
- [ ] To limit the amount of resources that can be consumed by a project
- [ ] To limit the amount of resources that can be consumed by a pod
- [ ] To limit the number of projects in a cluster
<details>
  <summary> Answer </summary>

   To limit the amount of resources that can be consumed by a project
</details>

> Q2: What is the command to create a Resource Quota in OpenShift?
- [ ] `oc create quota my-quota --hard=pods=10,requests.cpu=4,requests.memory=8Gi,limits.cpu=4,limits.memory=8Gi`
- [ ] `oc create resource-quota my-quota --hard=pods=10,requests.cpu=4,requests.memory=8Gi,limits.cpu=4,limits.memory=8Gi`
- [ ] `oc create quota my-quota --limits=pods=10,requests.cpu=4,requests.memory=8Gi,limits.cpu=4,limits.memory=8Gi`
- [ ] `oc create resource-quota my-quota --limits=pods=10,requests.cpu=4,requests.memory=8Gi,limits.cpu=4,limits.memory=8Gi`
<details>
  <summary> Answer </summary>

   `oc create resource-quota my-quota --limits=pods=10,requests.cpu=4,requests.memory=8Gi,limits.cpu=4,limits.memory=8Gi`
</details>

> Q3: What is the command to check the status of a Resource Quota in OpenShift?
- [ ] `oc describe quota my-quota`
- [ ] `oc get quota my-quota`
- [ ] `oc status quota my-quota`
- [ ] `oc check quota my-quota`
<details>
  <summary> Answer </summary>

   `oc describe quota my-quota`
</details>  

> Q4: What is the command to create a pod that exceeds the resource limits in OpenShift?
- [ ] `oc run my-pod --image=nginx --requests=cpu=2,memory=4Gi --limits=cpu=2,memory=4Gi`
- [ ] `oc create pod my-pod --image=nginx --requests=cpu=2,memory=4Gi --limits=cpu=2,memory=4Gi`
- [ ] `oc new-pod my-pod --image=nginx --requests=cpu=2,memory=4Gi --limits=cpu=2,memory=4Gi`
- [ ] `oc deploy my-pod --image=nginx --requests=cpu=2,memory=4Gi --limits=cpu=2,memory=4Gi`
<details>
  <summary> Answer </summary>
    `oc run my-pod --image=nginx --requests=cpu=2,memory=4Gi --limits=cpu=2,memory=4Gi`
</details>  


### 3.4 OpenShift Volumes

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module5#-module-5-advanced-deployment-options" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

Volumes allow you to manage mounted file systems in your pods using a variety of different suppliers.
    
    - Filesystem mounted in Pods
    - Many Suppliers
  
OpenShift can make files available to containers from many sources, such as secrets,ConfigMaps, attached hard disk storage, cloud storage, such as S3 and Google Compute files,and many other types of suppliers.

- Lets learn about emptyDir Volume type: 
Empty directory volumes always start out empty. The worker node that runs your application provides some temporary storage along with other information for the Pod. If the Pod is removed from a node, the EmptyDirectory contents will be deleted. 
    - Most common problems is when you run rolling out a new version, updating, configuration or changing a Pod in any way.

<p align="center">
<img src="/images/volumes.png" alt="OpenShift Training" style="width:400px; align="center"/>
</p>

**Hands-on Walkthroughs** 
- Define and use an empty directory:

    ```bash
    oc new-project volumes
    ```
    ```bash
    oc new-app quay.io/practicalopenshift/hello-world
    ```
    ```bash
    oc set volume deployment/hello-world \
  --add \
  --type emptyDir \
  --mount-path /empty-dir-demo 
    ```
     > output: "Generated volume name: volume-xxxx / deployment.apps/hello-world volume updated"
    ```bash
    oc get -o yaml deployment/hello-world
    ```
    > output:
    ```yml
    ......
    # The mountPath value specifies that OpenShift should mount the volume in the empty-dir-demo directory.
        volumeMounts:
        - mountPath: /empty-dir-demo 
          name: volume-XXXX 
    ......
      volumes:
      - emptyDir: {} 
        name: volume-XXXX # we didn't define a name so it generated that one!
    status:
    ......
    ```
- How to verify the emptyDir Volume is working?
So we need to access the pod's terminal.

    ```bash
    oc get pod
    ```
    - Copy the Pod name

    ```bash
    oc rsh <pod name>
    ```
    > output: In there look for the empty-dir-demo directory and verify its empty.
- How to mount ConfigMap as a Volume.
    ```bash
    oc create configmap cm-volume \
  --from-literal file.txt="ConfigMap file contents"
    ```
    > output: "configmap/cm-volume created"
   
   ```bash
    oc set volume deployment/hello-world \
    --add \
    --configmap-name cm-volume \
    --mount-path /cm-directory
   ```
    > output: "deployment.apps/hello-world volume updated"
    ```bash
    oc get -o yaml deployment/hello-world
    ```
    > output: 
    ```yaml
    .....
      volumeMounts:
        - mountPath: /empty-dir-demo
          name: volume-cnpmf
        - mountPath: /cm-directory
          name: volume-c9bct
    ......
      volumes:
      - emptyDir: {}
        name: volume-cnpmf
      - configMap:
          defaultMode: 420
          name: cm-volume
        name: volume-c9bct
    .....
    ```

### 🔬 Hands-on Lab (Volumes): 
For volumes, you'll mount a secret as a volume. 

- Create a new project named `volumes-lab`
- Create a new opaque secret based on the files in the `pods` directory from the labs repository
- Deploy the Hello World application
- Mount this secret as a volume to the path /secret-volume

---

### Checklist 📋 (Volumes): 

- After you shell into the pod, `ls /secret-volume` shows you the two pod files.
- You can get the original contents of the pod files with `cat`

---

### Quiz (Volumes)
> Q1: The emptyDir volume type can persist its data through a Pod restart
- [ ] True 
- [ ] False


<details>
  <summary> Answer </summary>

    True, this is one type of Pod event that emptyDir can handle.
  

</details>

> Q2: What is the command to add an emptyDir volume to an existing DeploymentConfig?
- [ ] `oc set volume`
- [ ] `oc create volume`
- [ ] `oc add-volume`
- [ ] `oc set DeploymentConfig --volume`


<details>
  <summary> Answer </summary>

  `oc set volume`
  
</details>

> Q3: Where can you go to find out more about configuring other types of volumes?
- [ ] Kubernetes Documentation 
- [ ] oc explain
- [ ] all the above
- [ ] None of the above

<details>
  <summary> Answer </summary>

   All the above
  
</details>

> Q4: Secrets and ConfigMaps can be used as volumes or environment variables.
- [ ] True 
- [ ] False

<details>
  <summary> Answer </summary>

    True, They can be used as either volumes or environment variables.
  
</details>

---

### 3.5 OpenShift Networking


<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module3#-module-3-core-openshift-resources" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

- [Servcies](https://docs.redhat.com/en/documentation/openshift_container_platform/3.11/html/architecture/core-concepts#services) They are Kubernetes resources that expose a set of pods as a network service. They provide a stable endpoint for accessing applications running within the cluster, even as individual pods are created, destroyed, or scaled. OpenShift tells us about `service` its a named abstraction of <u>software service</u> , no matter whether your application runs on one pod or 100 pods, you will need just one `service` to expose all of those pods to the network. External applications don't know how many pods are running in your application. Instead, all they know is that there's a `service` that they can call in order to access it. That's the abstraction the `service` will do the hardwork of splitting up traffic among all of those pods. Also the `service` depends on a `proxy` and a `selector` by `proxy` the application referring to the internally exposed IP and port that the `services` listens on. <u>This IP is only available to routing inside the OpenShift cluster."Virtual IPs"</u> Along with the `Virtual IPs`, `services` will also expose a port for example `80:80`. Virtual IPs ad Port are great for internal use, but in order to reach a service from Outside the cluster you will need to use another OpenShift resource type called the [Route](https://docs.redhat.com/en/documentation/openshift_container_platform/3.11/html/developer_guide/dev-guide-routes).


---

**Hands-on Walkthroughs**  

- Dig into the Service `spec`.

  ```bash
  oc explain svc.spec
  ```
  > output: The first field that was given is the Cluster IP, this IP is the `virtual IP` assigned to the service by OpenShift that is only exposed "This is how other pods in OpenShift will access your `service` " also looking at the `ports` object this is the list of ports exposed by the `service`, and you should have the `selector`, which explains a bit more about how `selectors work` - `selectors` use the same labels that we learned about before for `deployment`

- Lets create a manual service:

  ```bash
  oc create -f labs-repo/pods/pod.yaml
  ```

  ```bash
  oc expose pod/hello-world-pod
  ```

  > output: "you should see an error as you need to spicify the port!" 
  ```bash
  error: couldn't find port via --port flag or introspection
  See 'oc expose -h' for help and examples
  ```
  - We will expose using the port `80:80`
    ```bash
    oc expose --port 8080 pod/hello-world-pod
    ```
    > output: "service/hello-world-pod exposed"

    ```bash
    oc status
    ```
    > output: if you see this that means you did expose it correctly
    ```bash
    svc/hello-world-pod - 172.30.71.69:8080
      pod/hello-world-pod runs quay.io/practicalopenshift/hello-world
    ```
  - Lets create another pod and make a network request from the 2nd pod to the 1st one. 
    ```bash
    oc create -f labs-repo/pods/pod2.yaml
    ```

    > output: "pod/hello-world-pod-2 created"

  - We need to open a shell in the 2nd pod
    ```bash
    oc rsh hello-world-pod-2
    ```
  > output: `$`
  - run the following to make sure your at the `go` directory. 
    ```bash
    pwd
    ```
    > output: `/go` if not please `cd` into `/go` directory

  - For out 1st request, we will use the data from `oc status` to make the network request between the two pods."in this case `172.30.71.69:8080` FYI, the hello-world pods don't have `curl` installed instead we will have to use `wget` command.
    ```bash
    # wget -qO- <service IP / Port> "you should have your own" 
    # `-qO-` option to print to standard output instead of to a file. 
    wget -qO- 172.30.71.69:8080
    ```
    > output: "Hi! I'm an environment varible"

      <p align="center">
      <img src="/images/svc-req1.png" alt="OpenShift Training" style="width:400px; align="center"/>
      </p>
    
  - Accessing a Service using environment variables:
    - Environment variables are another set of key value pairs that are available in your pods. The values will be a mix of present values from the operating system, values specified in the `ENV` instruction in your dockerfile, and values injected by OpenShift

    - List all of the values from inside the pod by running the following command: 

    ```bash
    env
    ```
    > output: Because services can expose multiple ports, each one get a seprate environment variable with a port number, we can find the `address` we used for the `wget` command. By using a couple of these environment varaibles, the first is the Port `80:80 TCP address`. "this is the same IP printed out in the `oc status`"
    "HELLO_WORLD_POD_PORT_8080_TCP_PORT=8080"
    "HELLO_WORLD_POD_PORT_8080_TCP_ADDR=172.30.71.69"

    - The advantage to using environment variables rather than the direct IP is that IP may change overtime or between clusters "so lets try!"

    ```bash
    $wget -qO- $HELLO_WORLD_POD_PORT_8080_TCP_ADDR:$HELLO_WORLD_POD_PORT_8080_TCP_PORT
    ```

    > You should get the same output and thats the first step of learning Bash Scripting.

- Clean up
  ```bash
  oc delete all -l app=hello-world
  ```
____

***Route*** Openshift external network interface, it gives an external DNS name to a `service`.

- How to create a Route:
  ```bash
  oc new-app quay.io/practicalopenshift/hello-world 
  ```
  ```bash
  oc expose svc/hello-world
  ```
  > output: "route.route.openshift.io/hello-world exposed"

  ```bash
  oc status
  ```
  > output: you should see that you got a `svc` exposed to a specific `url` lets `curl` that one!

  ```bash
  # for this example I got `svc/hello-world-pod - http://hello-world-raafat-dev.apps.rm3.7wse.p1.openshiftapps.com` that what I will be using.
  curl http://hello-world-raafat-dev.apps.rm3.7wse.p1.openshiftapps.com
  ```
  > output: "Welcome! You can change this message by editing the MESSAGE environment variable. "

- Lets dig a bit deeper in Routes: 
  **Routes** are not part of Kubernetes they are new to OpenShift.

  - Lets take a look at the `route` yaml either from the UI or cli.

    ```bash
    oc get -o yaml route/hello-world
    ```
    <p align="center">
      <img src="/images/route-yml.png" alt="OpenShift Training" style="width:400px; display:block; margin:auto;" />
    </p>


    > output: Simple resource looking at the `spec` you will see a `host` "Domain Name of your route", this is something you can control in real-life projects, a `targetport`, and a `target` specified in the `to` field, the object in the `to` field has a `kind`, `name` & `weight` for this `route` the `kind` is the `service` and the `name` in this section identifies the service that the `route` should use. And the `weight` can be used to balance against more than one backend target. For now, this `route` only uses 1 `service` so the `weight` is `100`, and all requests will go to the `service` regardless of the `weight`.

- Advanced Routing options:

  <p align="center">
  <img src="/images/advanceroute.png" alt="OpenShift Training" style="width:400px; display:block; margin:auto;" />
  </p>

- Routes connect directly to pods, they don't connect to service if they do the process will get even longer, OpenShift has a router which takes care of this essentially what you do is you connect to your OpenShift cluster's IP address. OpenShift will check the URL and redirect you to a particular port so all the URLs are essentially connection you to the same IP but using the `cookies` and other methods you get redirected to a particular port depending upon which URL you connected to. This URL your connected to can be: 
  - Insecure Route:  Which is HTTP protocol "everything is in plain text" connection between you and the cluster is insecure and between the cluster to the pod is also insecure

  -  Edge Route : Which uses `HTTPS` outside the connection between you and OpenShift cluster is Secure, but the connection within the cluster to the pod is insecure. 

  - Passthrough Route: Which is secure on the outside and using the same certificate is seccure on the inside.

  - Re-encrypt Route: Where we have a different ceritificate on the outside and a different certificate on the inside "they are really complicated and not recommended"

- Route yaml sample 
```yaml
apiVersion: route.openshift.io/v1  # API group and version specific to OpenShift Route resource
kind: Route                        # Declares this is a Route object
metadata:
  name: my-app-route              # Name of the Route (used for identification within the project)
  labels:
    app: my-app                   # Labels help in selecting, grouping, or querying related objects
spec:
  to:
    kind: Service                 # Defines the type of backend this route sends traffic to
    name: my-app-service          # Name of the target service (must match an existing Kubernetes service)
    weight: 100                   # Optional: traffic weight if multiple backends (used in A/B deployments)
  port:
    targetPort: 8080             # Target port exposed by the backend service (matches container port)
  tls:
    termination: edge            # TLS termination strategy: edge, passthrough, or re-encrypt
    insecureEdgeTerminationPolicy: Redirect
                                 # Defines behavior for insecure (HTTP) traffic. Options: None, Allow, Redirect
  wildcardPolicy: None           # Whether subdomain wildcards are allowed. Options: None, Subdomain

```

**Hands-on Walkthroughs**  

- Lets use the deployed application: 
  ```bash
  # list the pods
  oc get pods 
  ```
  ```bash
  # list the services
  oc get svc
  ```
  - Expose insecure route (HTTP - HTTP)
    ```bash
    oc expose svc hello-world 
    ```
    > output: `http://hello-world-<namespace>.apps.....openshiftapps.com`

    ```bash
    oc describe pod # grab the pod IP 
    ```
    ```bash
    oc describe route # grab the endpoints IP 
    ```
    > output: comparing both you will see its the same which means the Route connects you directly to the Pod not the service.

  - Edge route (HTTPS - HTTP)
    > If you don't provide a certificate key then OpenShift is going to generate a self signed certificate and use that. Also you can create your own and use it to do that follow along:

    - Create the Root CA Certificate and Key:
    ```bash
      # Generate the CA private key
      openssl genrsa -aes256 -out ca-key.pem 4096

      # Create the CA root certificate (self-signed)
      openssl req -x509 -new -nodes -key ca-key.pem -sha256 -days 3650 -out ca-cert.pem
    ```
    - Create the Server Private Key and Certificate Signing Request (CSR):
    ```bash
    # Generate the server private key
    openssl genrsa -out server-key.pem 2048

    # Create the server Certificate Signing Request (CSR)
    openssl req -new -key server-key.pem -out server-csr.pem
    ```
    - Sign the Server CSR with the CA:
    ```bash
    # Sign the server CSR with the CA
    openssl x509 -req -in server-csr.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -days 365 -sha256
    ```
    - Create an Edge route
      - You can lookup the edge-help to get familiar with all options
        ```bash
        oc get route edge --help
        ```
    ```bash
    oc create route edge app-route-edge --service hello-world
    ```
    > output: 
      <p align="center">
      <img src="/images/edge.png" alt="OpenShift Training" style="width:200px; display:block; margin:auto;" />
      </p>

    ```bash
    oc get route
    ```

    > output: you will see the route has an `edge` termination option.
  
  - Passthrough Route (HTTPS - HTTPS same!): 
    > The communication is secure on the outside as well as in the inside, so there are 2 ports that communicate to each other over the secure channel, to do this what you need is to understand that your application should support SSL, also you need to know where are the secrets going to be kept inside the application, for example there is a certificate and key you need to provide to your application, so you need to mount it into the Pod, we will cover that in the Volumes chapter. Steps are: 
      1 - We will mount the certificate to `volumeMounts.mountPath:` inside the `spec.containers` for example `/usr/local/etc/ssl/certs`
      2 - Create a TLS `secret` for example `todo-ssl` "we will cover secrets in this course too" and it will contain the `key` and the `certificate`
      3 - Create the route

      ```bash
      oc create route passthrough --service hello-world
      ```

---

### 🔬 Hands-on Lab (Network)

For networking, you'll need to make some modifications to get a route to load balance between two pods.

- First, use `oc create` to start `pods/pod.yaml` in the labs project
- Create a service for this pod
- Modify `pods/pod2.yaml` so that the service will also hit this pod (hint: check the labels section in the pod and selector section in the service)
- Use `oc create` to start `pods/pods.yaml` in the labs project
- Expose a route for this service

---

### Checklist 📋 (Network)

Once you meet all of these criteria, you have successfully completed the lab. You will need to run the commands yourself in order to grade this lab.

- Output from `oc get pods` contains two pods
- Output from `oc status` groups the two pods under the same route
- When you run `curl <your route>` several times, it will return messages from both `pod.yaml` and `pod2.yaml`

---

### Quiz (Network)

> Q1: What mechanism do services use to figure out which pods to send traffic to?

- [ ] Develpers manually update services using `oc service add-target`
- [ ] Label selectors
- [ ] Services keep target pods in their yaml under `spec.targets`
- [ ] There is no such a thing!

<details>
  <summary> Answer </summary>

   Label selectors

</details>

> Q2: What is the command to expose a service as a route?
- [ ] `oc expose service <service-name>`
- [ ] `oc create route <route-name> --service=<service-name>`
- [ ] `oc expose <service-name>`
- [ ] `oc create route <route-name> --service=<service-name> --port=<port>`

<details>
  <summary> Answer </summary>

   `oc create route <route-name> --service=<service-name> --port=<port>`

</details>

---

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module4" target="_blank">
    <img src="/images/nexticon.webp" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

<p align="left">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module2" target="_blank">
    <img src="/images/backred1.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>
