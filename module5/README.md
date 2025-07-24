## Module 5: Advanced Deployment Options

### 5.1 Source-to-image (S2I)
 It is a toolkit and workflow that automates the process of building container images from source code. It takes a builder image (containing necessary build tools and dependencies) and source code, then combines them to create a runnable application image. S2I is used to create reproducible container images, making it easier for developers to deploy and manage applications in various environments, including OpenShift. 

 ![Source-to-Imag](/images/s2i-concept.webp)

 > It uses `Assemble` script vs `Run` from Docker and `Run` vs `CMD` from Docker.

 - FAQ: Why do I want another way to build and run applications?so what advantage does S2I have?
 
    - The first advantage is that the developer can rely on the expertise of the S2I authors to make sure the image will be OpenShift compatible.
    - The second advantage is that developers can avoid writing and maintaining Dockerfiles and other configuration by using the S2I defaults. 

**Hands-on Walkthroughs** 
- How to build an app using S2I?
> the source of this lesson will be in you Labs directory `/s2i/ruby/`

```bash
cd ./s2i/ruby
```
```bash
oc new-app ruby~https://gitlab.com/therayy1/openshif-labs.git --context-dir s2i/ruby --as-deployment-config
```

```bash
oc logs -f bc/openshif-labs
```
> output: push successful 

```bash
oc get pods
```
> output: you will be able to see all pods related to this deployment

```bash
oc expose svc/openshif-labs
```
> output: "route.route.openshift.io/openshif-labs exposed"

```bash
oc get route
```
> output: URL under `HOST/PORT` copy it please!

```bash
curl <URL>
```
> output: "Hello world from Ruby"

- How S2I language auto-detect works?!
- How does OpenShift know when to use S2I? And specifically, how did OpenShift know how to use the Ruby S2I builder?

  - When you start a build with OpenShift, OpenShift will first look for a dockerfile. If it finds one it will build usig something called the Docker Strategy. If OpenShift does not find a dockerfile it will attempt to use the Source Strategy instead.
   
    ![S2I Strategy](/images/S2I-visualselection.png)


### 🔬 Hands-on Lab: 
For S2I, you'll deploy a Python application without a Dockerfile and override an S2I script.

- If you haven't pushed your own version of the labs project to your GitLab account, you should follow the steps in the Builds lab to push your own version of the labs repository
- Create a new OpenShift project called `s2i-labs`
- Deploy an application using `oc new-app` with the python builder image for the `s2i/python` directory from the labs project for this course
- Create a route for your application
- Add an S2I override for the `run` script that prints a message and then calls the default `run` script
- Start a new build for the application

---

### Checklist 📋: 
- `curl <your route>` returns the Python hello world message
- `oc status` has a Route, Service, DeploymentConfig, and BuildConfig
- `oc logs <your Pod>` output contains the message you put in the override in step 5

---
### Quiz
> Q1: What is the syntax to specify a builder image to your oc new-app command?
- [ ] Put the image name before the Git repo URL with the % character in between them
- [ ] Put the image name before the Git repo URL with the ~ character in between them
- [ ] Put the image name before the Git repo URL with £ character in between them
- [ ] Put the image name before the Git repo URl with * in between them

<details>
  <summary> Answer </summary>

  Put the image name before the Git repo URL with the ~ character in between them

</details>

> Q2: What are the auto-detected languages for OpenShift builds?
- [ ] Python, Ruby, Java, PHP, Perl, Node
- [ ] Python, Ruby, Java, Perl, Node, Clojure
- [ ] Python, Go, Java, PHP, Perl, Node
- [ ] None of the above

<details>
  <summary> Answer </summary>

  Python, Ruby, Java, PHP, Perl, Node

</details>

---

### 5.2 OpenShift Volumes
Volumes allow you to manage mounted file systems in your pods using a variety of different suppliers.
    
    - Filesystem mounted in Pods
    - Many Suppliers
  
OpenShift can make files available to containers from many sources, such as secrets,ConfigMaps, attached hard disk storage, cloud storage, such as S3 and Google Compute files,and many other types of suppliers.

- Lets learn about emptyDir Volume type: 
Empty directory volumes always start out empty. The worker node that runs your application provides some temporary storage along with other information for the Pod. If the Pod is removed from a node, the EmptyDirectory contents will be deleted. 
    - Most common problems is when you run rolling out a new version, updating, configuration or changing a Pod in any way.


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

### 🔬 Hands-on Lab: 
For volumes, you'll mount a secret as a volume. 

- Create a new project named `volumes-lab`
- Create a new opaque secret based on the files in the `pods` directory from the labs repository
- Deploy the Hello World application
- Mount this secret as a volume to the path /secret-volume

---

### Checklist 📋: 

- After you shell into the pod, `ls /secret-volume` shows you the two pod files.
- You can get the original contents of the pod files with `cat`

---
### Quiz
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

# 5.3 Scaling and Debuging Your Application 
In OpenShift, scaling refers to the process of dynamically adjusting the resources allocated to an application or the overall cluster to meet changing demands. This can involve increasing or decreasing the number of pods running an application (horizontal scaling), or adjusting the resources (CPU, memory) allocated to individual pods (vertical scaling).

![scaling](/images/scaling.png)

- How does the Auto-Scale works?
The **Horizontal Pod Autoscaler (HPA)** automatically adjusts the number of pods in your application based on CPU usage to handle changing workloads. It **scales up** when resource usage is high and **scales down** when demand is low, helping optimize resource consumption.

HPA uses a formula that considers:

* **Current number of pods**
* **Current CPU usage** (in millicores)
* **Desired usage**, calculated from the target CPU utilization (percentage) and requested CPU (defined in the pod spec)

The core idea:

* If actual usage is higher than desired, HPA increases pods.
* If usage is lower, it reduces pods.
* New pod count = current pods × (current usage ÷ desired usage)


**Hands-on Walkthroughs** 

- How to manually scale your application?
  - When your application starts, the initial number of pods is controlled by the number replicas property in the deployment config spec. The default is one for oc new-app applications. You can of course edit this by hand if you have your application in a template or YAML files.

    ```bash
    oc new-app quay.io/practicalopenshift/hello-world --as-deployment-config
    ```
    ```bash
    oc describe dc/hello-world
    ```
    > output: 
    ```yml
    .....
    Deployment #1 (latest):
        Name:           hello-world-1
        Created:        40 seconds ago
        Status:         Complete
        Replicas:       1 current / 1 desired
    .....
    ```
    ```bash
    oc scale dc/hello-world --replicas=3
    ```
    ```bash
    oc describe dc/hello-world
    ```
    > output: 
    ```yaml
    .....
    Deployment #1 (latest):
        Name:           hello-world-1
        Created:        3 minutes ago
        Status:         Complete
        Replicas:       3 current / 3 desired
    .....
    ```
- How to create a HPA:
  ```bash
  oc autoscale dc/hello-world \
  --min 1 \
  --max 10 \
  --cpu-percent=80
  ```
  > output: horizontalpodautoscaler.autoscaling/hello-world autoscaled

  ```bash
  oc get hpa
  ```
> output: includes all details about the HPA specially the Targets.
  
  ```bash
  oc describe hpa/hello-world
  ```  
  ```bash
  oc get -o yaml hpa/hello-world
  ```

### Quiz
> Q1: You must have a HorizontalPodAutoscaler in order to scale up your application.
- [ ] True 
- [ ] False

<details>
  <summary> Answer </summary>

    False, You can also do it manually
  
</details>

> Q2: What command can you use to create a Horizontal Pod Autoscaler for a DeploymentConfig?
- [ ] `oc scale--auto` 
- [ ] `oc autoscale`
- [ ] You have to create the HPA using `oc create -f`
- [ ] `oc get HPA`

<details>
  <summary> Answer </summary>

  `oc autoscale` 
  
</details>

> Q3: What property in a DeploymentConfig can you use to set the number of initial replicas for its ReplicationController?
- [ ] `initialReplicas`
- [ ] `replicas`
- [ ] `defaultReplicas`
- [ ] `HPA`

<details>
  <summary> Answer </summary>

  `replicas` 
  
</details>

--- 

# OpenShift Jobs: 


**Hands-on Walkthroughs** 

### 🔬 Hands-on Lab: 

### Checklist 📋: 

### Quiz
> Q1: 
- [ ]  
- [ ]
- [ ]
- [ ] 


<details>
  <summary> Answer </summary>

    
  

</details>