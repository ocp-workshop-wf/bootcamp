[![Static Badge](https://img.shields.io/badge/CheatSheet-purple?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/CheatSheet.md)  [![Static Badge](https://img.shields.io/badge/OCP-CLI-red?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/ocpcli-cheatsheet.md)   [![Static Badge](https://img.shields.io/badge/Labs-maroon?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/tree/main/labs-repo)  [![Static Badge](https://img.shields.io/badge/Kubernetes-black?style=flat&logo=Kubernetes&logoSize=auto)
](https://kubernetes.io/docs/home/) [![Static Badge](https://img.shields.io/badge/RedHat-OpenShift-maroon?style=flat&logo=Redhat&logoSize=auto)
](https://docs.redhat.com/en/documentation/openshift_container_platform/4.19)  

## 🔹 Module 3: Core OpenShift Resources

## Table of Contents 
- [3.1 - OpenShift Resources Overview](#31-openshift-resources-overview)
 
- [3.2 - OpenShift and how to deploy applications](#32-openshift-and-how-to-deploy-applications)

- [3.3 - OpenShift Networking](#33-openshift-networking)

- [3.4 - OpenShift ConfigMaps](#34-openshift-configmaps)

### 3.1 OpenShift Resources Overview

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
    oc new-app quay.io/practicalopenshift/hello-world --as-deployment-config
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
    oc rollout latest dc/hello-world
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
    oc rollback dc/hello-world
    ```

  > output: It is very similar process "Start the previous version, and Stops the current"
    <p align="center">
    <img src="/images/rollback0.png" alt="OpenShift Training" style="width:400px; align="center"/>
    </p>

    <p align="center">
    <img src="/images/rollback1.png" alt="OpenShift Training" style="width:400px; align="center"/>
    </p>

---

### 🔬 Hands-on Lab

In the DeploymentConfig lab, you will create a custom DeploymentConfig based on the hello-world image by changing some parameters.

- First, `use oc new-app to start an application based on quay.io/practicalopenshift/hello-world`

- Use `oc new-app` to start a second version of the application using the name `lab-dc`, a custom value for the `MESSAGE` environment variable, and the same hello-world image

  - You can specify environment variables in `oc new-app` with a flag. `oc new-app --help` can help you to find the correct one

- Forward port 8080 on your local computer to port 8080 on the second pod you created

---

### Checklist 📋

- Output from `oc get pods` contains two pods

- Output from `oc describe dc/lab-dc` has the correct name and `MESSAGE` environment value

- `curl localhost:8080` prints the message you entered in step 2

---

> 💡 Cleaning Up:
 To clean up, use a single command to delete all of the resources created in step 1. You are done when `oc get dc` just has the `lab-dc` DeploymentConfig.

---

### Quiz

> Q1: What is the command to deploy the hello-world image to OpenShift as a deployment config?

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

---

### 3.3 OpenShift Networking

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
    - Environment variables are another set of key value pairs that are available in your pods. The valuse will be a mix of present values from the operating system, valuse specified in the `ENV` instruction in your dockerfile, and valuse injected by OpenShift

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

  Route                          # Top-level object representing the route definition
  ├── metadata                   # Metadata about the route
  │   ├── name                   # Name of the route object
  │   ├── namespace             # Kubernetes/OpenShift project where the route exists
  │   └── annotations           # Optional metadata to customize behavior (e.g., timeouts, load balancing)
  ├── spec                      # Desired configuration of the route
  │   ├── host                  # Public hostname exposed by the route (e.g., myapp.example.com)
  │   ├── to                    # Specifies the target backend object the route points to
  │   │   ├── kind              # Type of target, typically "Service"
  │   │   ├── name              # Name of the service the route will send traffic to
  │   │   └── weight            # Traffic distribution weight (used in canary or blue-green deployments)
  │   ├── port                  # Port configuration for the backend service
  │   │   └── targetPort        # Port on the service to direct traffic to (e.g., 8080)
  │   └── tls                   # TLS/SSL configuration for secure routes
  │       ├── termination       # Type of TLS termination: edge, passthrough, or reencrypt
  │       └── insecureEdgeTerminationPolicy  # Policy for handling HTTP when TLS is enabled (None, Allow, Redirect)
  └── status (auto-generated)   # System-generated status information (e.g., admitted routers, conditions)

  - Additional metadata to control timeout, Load balancing type, Cookies disabling, TLS termination (Edge vs Passthrough)


---

### 🔬 Hands-on Lab

For networking, you'll need to make some modifications to get a route to load balance between two pods.

- First, use `oc create` to start `pods/pod.yaml` in the labs project
- Create a service for this pod
- Modify `pods/pod2.yaml` so that the service will also hit this pod (hint: check the labels section in the pod and selector section in the service)
- Use `oc create` to start `pods/pods.yaml` in the labs project
- Expose a route for this service

---

### Checklist 📋

Once you meet all of these criteria, you have successfully completed the lab. You will need to run the commands yourself in order to grade this lab.

- Output from `oc get pods` contains two pods
- Output from `oc status` groups the two pods under the same route
- When you run `curl <your route>` several times, it will return messages from both `pod.yaml` and `pod2.yaml`

---

### Quiz

> Q1: What mechanism do services use to figure out which pods to send traffic to?

- [ ] Develpers manually update services using `oc service add-target`
- [ ] Label selectors
- [ ] Services keep target pods in their yaml under `spec.targets`
- [ ] There is no such a thing!

<details>
  <summary> Answer </summary>

   Label selectors

</details>

### Quiz

> Q2: How do pods and other resources send traffic to a service?

- [ ] The Pod's virtual IP
- [ ] Expose a Route for the service
- [ ] Use oc port-forward from inside a Pod
- [ ] Click on the service right click and select one!

<details>
  <summary> Answer </summary>

   The Pod's virtual IP

</details>

> Q3: What is the command to create a service for a pod?

- [ ] `oc expose <pod-name>`
- [ ] `oc expose --port<port><pod-name>`
- [ ] `oc expose service --port<port><pod-name>`
- [ ] `oc expose service<pod-name>`

<details>
  <summary> Answer </summary>

   `oc expose --port<port><pod-name>`

</details>

> Q4: What is the command to create a route for a service?

- [ ] `oc expose <service-name>`
- [ ] `oc expose-service <service-name`
- [ ] `oc expose port<port>`
- [ ] `oc expose service<pod-name>`

<details>
  <summary> Answer </summary>

   `oc expose <service-name>`

</details>

> Q5: True / False: OpenShift publishes virtual IPs in environment variables inside of containers.

- [ ] True
- [ ] False
- [ ] None of the above
- [ ] All the above

<details>
  <summary> Answer </summary>

   True

</details>

---

### 3.4 OpenShift ConfigMaps

- **Configmaps:** are a very useful resource type that OpenShift borrows from Kubernetes. ConfigMaps hold configuration data for pods to consume. This data is held in the ConfigMap separately from your running pod in OpenShift. Holding data for pods to consume is much less active job than some other types of resources in Kubernetes like `pods`, `deploymentconfigs` and `services` so when to use `ConfigMaps`

  | Component         | Development | Production               |
  |------------------|-------------|---------------------------|
  | REST API Server  | localhost   | example-api.com           |
  | Database         | localhost   | db-host.internal.com      |

  > A common case where ConfigMaps become useful is when you deploy your application to different environments. For Local development, you may wish to run non-application dependencies such as REST service or database on your machine as well, in order to simplify the development environment to connect to this REST sercice or database your application will need to use valuse that point to these local versions you can use ConfigMaps to get that kind of felxibility in OpenShift. 

<p align="center">
  <img src="/images/cm.png" alt="OpenShift Training" style="width:400px; display:block; margin:auto;" />
</p>

  > There is one piece of vocabulary you need to know when working with ConfigMaps. To <u>CONSUME</u> a configmap just means that you use the data inside of a ConfigMap from a Pod. Once the ConfigMap exists on OpenShift, you can consume or use the ConfigMap in a pod by referring to the ConfigMap in the pod definition. The word consume makes it seem like consuming a ConfigMap should use it up somehow,but this is not the case. You can consume ConfigMaps from many different pods. A common i.e  "database name" shared between an application defination and the database pod. This can help you to centralize you configuration in <u>ONE PLACE</u> and then you can update the ConfigMap, all of your pods will use the updated configuration.

  <p align="center">
    <img src="/images/cm2.png" alt="OpenShift Training" style="width:400px; display:block; margin:auto;" />
  </p>

- OpenShift give you several tools to create ConfigMaps: 
  - Command line
  - Files
  - Entire directories.
  > You also need to talk about sensitive data any resource in your OpenShift Project can read the data in a ConfigMap. For this reason, you should not store any sensitive data in the ConfigMap. OpenShift has another resource type called the `Secret` that is used for sensitive date. Also there is one more limitation of ConfigMaps that you should keep in mind, ConfigMaps have a `1MB` storage size limit. If your data could grow above the `1MB` limit you may not want to use Configmaps.


- ConfigMap Example:

    ```yaml
    apiVersion: v1 #kubernetes resource
    data:
    MESSAGE: Hello from ConfigMap #data 
    kind: ConfigMap
    metadata:
    name: message-map
    namespace: myproject #namespace
    resourceVersion: "2827192"
    selfLink: /api/v1/namespaces/myproject/configmaps/message-map
    ```

**Hands-on Walkthroughs**  

- Creating ConfigMaps:

```bash
oc create configmap message-map --from-literal MESSAGE="Hello from ConfigMap"
# type - name of resource - to give some data -(Varaible) = data
```
> output: "configmap/message-map created"

```bash
# listing configmaps
oc get cm
```

> output:
  ```bash
  message-map         1      42s # 1 means that we have only 1 key value pair inside the configmap
  ```
- Lets check the ConfigMap yaml from UI and CLI
```bash
oc get -o yaml cm/message-map
```
> output:
  ```yaml
  apiVersion: v1
  data: # instead of spec field we got data field here
    MESSAGE: Hello from ConfigMap # varaible we configured earlier 
  kind: ConfigMap
  metadata:
    creationTimestamp: "2025-07-18T01:09:39Z"
    name: message-map # configmap name 
    namespace: <your-namespace> # namespace/Project
    resourceVersion: "3298865818"
    uid: 7c32526a-8837-48ba-ab36-bade0095b35b
  ```


- Consuming ConfigMaps: We'll test it out by making sure that the Hello World resonse changes after we consume the configmap 1st lets deploy our application.

  ```bash
  oc new-app quay.io/practicalopenshift/hello-world --as-deployment-config
  ```

  > Once app created go ahead and expouse the service!

  ```bash
  oc expose service/hello-world
  ```

  > Once expose was done lets run status to get the URL

  ```bash
  oc status
  ```

  > We will Curl before and after configmap being applied to the application, at first it should give us the default message, at 2nd it should give us the message from inside the configmap!

  ```bash
  curl <url from oc status>
  ```

  > output:
  "Welcome! You can change this message by editing the MESSAGE environment variable."

- Consuming a ConfigMap to the application from the  **command line**

  ```bash
  oc set env dc/hello-world --from cm/message-map 
  # ENV is just like in a dockerfile it sets an env variable based on the keys and pairs in the ConfigMap.
  ```

  > output: `deploymentconfig.apps.openshift.io/hello-world updated`

  ```bash
  curl <url from oc status>
  ```

  > output: "Hello from ConfigMap"

  ```bash
  oc get -o yaml dc/hello-world
  ```

  ```yaml
      .....
      spec:
        containers:
        - env: # the ENV we set 
          - name: MESSAGE # consumed from Configmap
            valueFrom:
              configMapKeyRef: # that is how it consumed from configmap
                key: MESSAGE
                name: message-map #configmap name
                ........
  ```

- Create ConfigMaps from **Files**:

  ```bash
  echo "Hello from ConfigMap file" > MESSAGE.txt
  # creating a simple txt file 
  ```

  ```bash
  cat MESSAGE.txt
  # just to verify
  ```

  > output:"Hello from ConfigMap file"

  ```bash
  oc create configmap file-map --from-file=MESSAGE.txt
  # create based on the txt file.
  ```

  > output: "configmap/file-map created"

  ```bash
  oc get -o yaml cm/file-map
  ```

  ```yaml
  apiVersion: v1
  data:
    MESSAGE.txt: | # confirms it was taken from the file
      Hello from ConfigMap file
  kind: ConfigMap
  metadata:
  .........
  ```

  > output: "data.MESSAGE.txt: this is the wrong syntax as it doesn't match the key in the Hello-world application to match it we need to run the following command"

  ```bash
  oc create configmap file-map-2 --from-file=MESSAGE=MESSAGE.txt
  # this time we specified the env variable = the file name
  ```

  > output: "configmap/file-map created"

  ```bash
  oc get -o yaml cm/file-map
  ```

  ```yaml
  apiVersion: v1
  data:
    MESSAGE: | # now its fixed 
      Hello from ConfigMap file
  kind: ConfigMap
  metadata:
  ```

  > output: Now as you see the data.MESSAGE: follows the same pattern for the Hello-world application.

  ```bash
  oc set env dc/hello-world --from cm/file-map-2
  # lets update the deploymentconfig to point at this configmap-2 the correct one 
  ```

  > output: "deploymentconfig.apps.openshift.io/hello-world updated"

  ```bash
  curl < URL from oc status>
  ```

  > output: "Hello from ConfigMap file."


- Create ConfigMaps from **Directories**:

  ```bash
  cd ./labs
  ```

  ```bash
  oc create configmap pods-example --from-file=pods
  # same but ending up with the name of the directory
  ```

  > output: "configmap/pods-example created!"

  ```bash
  oc get -o yaml configmap/pods-example
  ```

  > output:

  ```yaml
  apiVersion: v1
  data:
    pod.yaml: | # pod 1
      apiVersion: v1
      kind: Pod
      metadata:
        name: hello-world-pod
        labels:
          app: hello-world-pod
      spec:
        containers:
        - env: # seeing this here
          - name: MESSAGE # env varaible 
            value: Hi! I'm an environment variable
          image: quay.io/practicalopenshift/hello-world
          imagePullPolicy: Always
          name: hello-world-override
          resources: {}
    pod2.yaml: | # pod 2 
      apiVersion: v1
      kind: Pod
      metadata:
        name: hello-world-pod-2
        labels:
          app: hello-world-pod-2
      spec:
        containers:
        - env:
          - name: MESSAGE # message from pod 2 
            value: Hi! I'm an environment variable in pod 2
          image: quay.io/practicalopenshift/hello-world
          imagePullPolicy: Always
          name: hello-world-override
          resources: {}
    service.yaml: |
      apiVersion: v1
      kind: Service
      metadata:
        name: hello-world-pod-service
      spec:
        selector:
          app: hello-world-pod
        ports:
          - protocol: TCP
            port: 80
            targetPort: 8080
  kind: ConfigMap

  ```

---

### 🔬 Hands-on Lab

For ConfigMaps, you'll get some hands-on practice working with YAML. Start with the following ConfigMap definition:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: lab-map
```

- Create a new file called lab-configmap.yaml
- Copy the above YAML into the file
- Modify this YAML so that the ConfigMap will have the proper key fro the hello-world application
- Us `oc create` to create the ConfigMap from the file
- Deploy the `quay.io/practicalopenshift/hello-world` image using `oc new-app`
- Change the message that the `DeploymentConfig` uses to the ConfigMap value using the `oc set env` command
- Expose a route for your application.

---

### Checklist 📋

- Output from `oc get cm` contains your new ConfigMap

- Output from `oc get -o yaml dc/hello-world` contains the string `configMapKeyRef`

- When you run `curl <your route>` you get the value you put in the ConfigMap

---

### Quiz
>
> Q1: What is the maximum amount of data that you can store in a ConfigMap?

- [ ] 1 GB
- [ ] 1 KB
- [ ] 1 MB
- [ ] 1 TB

<details>
  <summary> Answer </summary>

   1 MB

</details>

> Q2: The data for a configmap is stored in its YAML resource definition under the "configData" field name.

- [ ] True
- [ ] False

<details>
  <summary> Answer </summary>

   Fales "the field called data"

</details>

> Q3: What is the command to create a configmap using the oc tool?

- [ ] `oc create configmap <new configmap name>`
- [ ] `oc create -f configmap <new configmap name>`
- [ ] `oc get configmap <new configmap name>`
- [ ] `oc apply -f configmap <new configmap name>`

<details>
  <summary> Answer </summary>

 `oc create configmap <new configmap name>`

</details>

> Q4: What kinds of inputs can you use to create a configmap?

- [ ] Command line arguments of files
- [ ] Command line arguments, files, directories, and custom ConfigMap YAML files
- [ ] Command line arguments, files, or directories
- [ ] Command line arguments, directories only!

<details>
  <summary> Answer </summary>

   Command line arguments, files, directories, and custom ConfigMap YAML files

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