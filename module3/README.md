## 🔹 Module 3: Core OpenShift Resources

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

- DeploymentConfig VS Deployment

| Point | Deployment Config | Deployment | 
| ----- | ----------------- | ---------- |
| Design | It prefers consistency. As deployer pod goes down, its not replaced but waits for it to come up again | It prefers availability over consistency. The controller manager runs in high availability mode. Hence if one controller goes down, other leader is elected. |
| Automatic Rollbacks | Supported | Not Supported | 
| Automatic trigger on config changes | Has to be mentioned explicity in resource definition | Implicit |
| Lifecycle hooks | Supported | Not Supported |
| Custom Deployment Strategies | Supported | Not Supported| 



<p align="center">
<img src="/images/deployment.png" alt="OpenShift Training"; align="center"/>
</p>

---

**Hands-on Walkthroughs**  

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

---

**Hands-on Walkthroughs**  

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

- **Servcies** They are Kubernetes resources that expose a set of pods as a network service. They provide a stable endpoint for accessing applications running within the cluster, even as individual pods are created, destroyed, or scaled.
- **Routes** Exposes a service at a hostname so that external clients can reach it. Routes are essentially DNS entries that map a hostname to a service within the OpenShift cluster - [RedHat Documentation](https://docs.redhat.com/en/documentation/openshift_container_platform/3.11/html/developer_guide/dev-guide-routes)

![OpenShift Network](/images/network.png)

---

**Hands-on Walkthroughs**  

- As always lets read about our new resource:

 ```
 oc explain svc
 ```

- Dig into the Spec.

 ```
 oc explain svc.spec
 ```

- Lets create a manual service:

  ```
  cd <lab-directory>
  ```

  ```
  oc create -f pods/pod.yaml
  ```

  ```
  oc expose pod/hello-world-pod
  ```

    > "you should see an error as you need to spicify the port!"
  
  ```
  oc expose --port 8080 pod/hellp-world-pod
  ```

  ```
  oc status
  ```

  ```
  oc create -f pods/pod2.yaml
  ```

    > We need to open a shell in the 2nd Pod

  ```
  oc rsh hello-world-pod-2
  ```

    ```bash
    $wget -qO- <service IP / Port>
    ```

  - Accessing a Service:

    ```bash
    env
    ```

    ```bash
    $wget -qO- $HELLO_WORLD_POD_PORT_8080_TCP_ADDR:$HELLO_WORLD_POD_PORT_8080_TCP_PORT
    ```

    > You should get the same output and thats the first step of learning Bash Scripting.
  >
____

- Exposing a Route:

```
oc status
```

```
oc new-app quay.io/practicalopenshift/hello-world --as-deployment-config
```

```
oc expose svc/hello-world
```

  > You should that the expose happened

```
oc status
```

> On the first line you will fine the `http-URL` copy that!

```
curl <route from oc status>
```

> You should get a Welcome! Message.

- Dig deeper:

```
oc get -o yaml route/hello-world
```

> Looking at the host its a compination of the app-name-<project-name>.IP-address.
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

- **Configmaps:** a fundamental way to manage configuration data for applications. They are Kubernetes API objects that store configuration data as key-value pairs, allowing you to decouple configuration from your application code and keep your containers portable. This means you can change an application's behavior without rebuilding its container image.

    ![ConfigMaps Stucture](/images/configmap.png)

> - Not for sensitive data
> - 1MB limit

- ConfigMap Example:

    ```yaml
    apiVersion: v1
    data:
    MESSAGE: Hello from ConfigMap
    kind: ConfigMap
    metadata:
    creationTimestamp: 2025-06-11T11:40:41Z
    name: message-map
    namespace: myproject
    resourceVersion: "2827192"
    selfLink: /api/v1/namespaces/myproject/configmaps/message-map
    uid: 60dc0569-abd8-11ea-9133-080027c1c30a
    ```

**Hands-on Walkthroughs**  

- Creating ConfigMaps:

```bash
oc create configmap message-map --from-literal MESSAGE="Hello from ConfigMap"
```

> `configmap/message-map created`

```bash
oc get cm
```

```bash
message-map         1      42s
```

```bash
oc get -o yaml cm/message-map
```

```yaml
apiVersion: v1
data:
  MESSAGE: Hello from ConfigMap
kind: ConfigMap
metadata:
  creationTimestamp: "2025-07-18T01:09:39Z"
  name: message-map
  namespace: <your-namespace>
  resourceVersion: "3298865818"
  uid: 7c32526a-8837-48ba-ab36-bade0095b35b
```

- Consuming ConfigMaps:

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

- Consuming a ConfigMap to the application

```bash
oc set env dc/hello-world --from cm/message-map
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
      - env:
        - name: MESSAGE
          valueFrom:
            configMapKeyRef:
              key: MESSAGE
              name: message-map
              ........
```

- Create ConfigMaps from Files:

```bash
echo "Hello from ConfigMap file" > MESSAGE.txt
```

```bash
cat MESSAGE.txt
```

> output:"Hello from ConfigMap file"

```bash
oc create configmap file-map --from-file=MESSAGE.txt
```

---

> output: "configmap/file-map created"

```bash
oc get -o yaml cm/file-map
```

```yaml
apiVersion: v1
data:
  MESSAGE.txt: |
    Hello from ConfigMap file
kind: ConfigMap
metadata:
.........
```

> output: "data.MESSAGE.txt: this is the wrong syntax as it doesn't match the key in the Hello-world application"

```bash
oc create configmap file-map-2 --from-file=MESSAGE=MESSAGE.txt
```

> output: "configmap/file-map created"

```bash
oc get -o yaml cm/file-map
```

```yaml
apiVersion: v1
data:
  MESSAGE: |
    Hello from ConfigMap file
kind: ConfigMap
metadata:
```

> output: Now as you see the data.MESSAGE: follows the same pattern for the Hello-world application.

```bash
oc set env dc/hello-world --from cm/file-map-2
```

> output: "deploymentconfig.apps.openshift.io/hello-world updated"

```bash
curl < URL from oc status>
```

> output: Hello from ConfigMap file.

- Create ConfigMaps from Directories:

```bash
cd ./labs
```

```bash
oc create configmap pods-example --from-file=pods
```

> output: "configmap/pods-example created!"

```bash
oc get -o yaml configmap/pods-example
```

> output:

```yaml
apiVersion: v1
data:
  pod.yaml: |
    apiVersion: v1
    kind: Pod
    metadata:
      name: hello-world-pod
      labels:
        app: hello-world-pod
    spec:
      containers:
      - env:
        - name: MESSAGE
          value: Hi! I'm an environment variable
        image: quay.io/practicalopenshift/hello-world
        imagePullPolicy: Always
        name: hello-world-override
        resources: {}
  pod2.yaml: |
    apiVersion: v1
    kind: Pod
    metadata:
      name: hello-world-pod-2
      labels:
        app: hello-world-pod-2
    spec:
      containers:
      - env:
        - name: MESSAGE
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
