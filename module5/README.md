[![Static Badge](https://img.shields.io/badge/Agenda-green?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/Agenda.md) [![Static Badge](https://img.shields.io/badge/CheatSheet-purple?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/CheatSheet.md)  [![Static Badge](https://img.shields.io/badge/OCP-CLI-red?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/ocpcli-cheatsheet.md)   [![Static Badge](https://img.shields.io/badge/Labs-maroon?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/tree/main/labs-repo)  [![Static Badge](https://img.shields.io/badge/RedHat-OpenShift-maroon?style=flat&logo=Redhat&logoSize=auto)
](https://docs.redhat.com/en/documentation/openshift_container_platform/4.19)   [![Static Badge](https://img.shields.io/badge/Kubernetes-black?style=flat&logo=Kubernetes&logoSize=auto)
](https://kubernetes.io/docs/home/)
## 🔹 Module 5: Advanced Deployment Options

## Table of Contents 

- [5.1 - Source-to-image (S2I)](#51-source-to-image-s2i) | [Lab](#-hands-on-lab-s2i) | [Quiz](#quiz-s2i)

- [5.2 - OpenShift Volumes](#52-openshift-volumes) | [Lab](#-hands-on-lab-volumes) | [Quiz](#quiz-volumes)

- [5.3 - Scaling and Debuging Your Application](#53-scaling-and-debuging-your-application) | [Quiz](#quiz-scaling) 

- [5.4 - OpenShift Jobs](#openshift-jobs) 


### 5.1 Source-to-image (S2I)

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module5#-module-5-advanced-deployment-options" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

 It is a toolkit and workflow that automates the process of building container images from source code. It takes a builder image (containing necessary build tools and dependencies) and source code, then combines them to create a runnable application image. S2I is used to create reproducible container images, making it easier for developers to deploy and manage applications in various environments, including OpenShift. 

<p align="center">
<img src="/images/s2i-concept.webp" alt="OpenShift Training" style="width:400px; align="center"/>
</p>

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
oc new-app ruby~https://gitlab.com/therayy1/openshif-labs.git --context-dir s2i/ruby 
```

```bash
oc logs -f deployment/openshif-labs
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
   
<p align="center">
<img src="/images/S2I_Build_Process_Regenerated.png" alt="OpenShift Training" style="width:300px; align="center"/>
</p>



### 🔬 Hands-on Lab (S2I): 
For S2I, you'll deploy a Python application without a Dockerfile and override an S2I script.

- If you haven't pushed your own version of the labs project to your GitLab account, you should follow the steps in the Builds lab to push your own version of the labs repository
- Create a new OpenShift project called `s2i-labs`
- Deploy an application using `oc new-app` with the python builder image for the `s2i/python` directory from the labs project for this course
- Create a route for your application
- Add an S2I override for the `run` script that prints a message and then calls the default `run` script
- Start a new build for the application

---

### Checklist 📋 (S2I): 
- `curl <your route>` returns the Python hello world message
- `oc status` has a Route, Service, DeploymentConfig, and BuildConfig
- `oc logs <your Pod>` output contains the message you put in the override in step 5

---
### Quiz (S2I)
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

### 5.3 Scaling and Debuging Your Application 

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module5#-module-5-advanced-deployment-options" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

In OpenShift, scaling refers to the process of dynamically adjusting the resources allocated to an application or the overall cluster to meet changing demands. This can involve increasing or decreasing the number of pods running an application (horizontal scaling), or adjusting the resources (CPU, memory) allocated to individual pods (vertical scaling).


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

<p align="center">
<img src="/images/hpa-overview.png" alt="OpenShift Training" style="width:500px; align="center"/>
</p>


***Debuging in OpenShift*** provides a powerful way to troubleshoot and debug issues within your cluster, particularly for pods and nodes.When used with a pod, `oc debug` creates a new, temporary pod based on the existing pod's image and configuration, but with the ability to inject debugging tools or run commands within its environment. This allows you to:
- Attach to a running container: Gain a shell prompt inside a container to inspect its file system, processes, or configuration.
- Install debugging tools such as `strace`, or other utilities to analyze application behavior
- Modify container environment: Temporarily change environment variables or mount paths for testing or debugging
 
* [Debug Resource](http://redhat.com/en/blog/how-oc-debug-works#:~:text=If%20you%20have%20used%20relatively,to%20display%20its%20YAML%20output.)

**Hands-on Walkthroughs** 

- How to manually scale your application?
  - When your application starts, the initial number of pods is controlled by the number replicas property in the deployment config spec. The default is one for oc new-app applications. You can of course edit this by hand if you have your application in a template or YAML files.

    ```bash
    oc new-app quay.io/practicalopenshift/hello-world 
    ```
    ```bash
    oc describe deployment/hello-world
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
- How to create a HPA:  "Explain more about CPU add some files to run to demonstrate the HPA"

```bash
  oc autoscale deployment/hello-world \
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

- Lets debug the existing pod.
```bash
oc get pods
```
> output: copy the pod name

```bash
oc debug <pod name>
```
<p align="center">
<img src="/images/debug.png" alt="OpenShift Training"; align="center"/>
</p>


> output: your terminal will directly `rsh` in a Temporarily pod which has the same exact features and configurations of the orginial one, be aware once you exit that pod terminal it will automatically be terminated and all changes you've made will be lost.

```bash
Starting pod/hello-world-8-zbcrx-debug, command was: /bin/sh -c go run hello-world.go
Pod IP: 10.128.30.7
If you don't see a command prompt, try pressing enter.
~ $ 
```

### Quiz (Scaling)
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

> Q4: What command we use to debug a pod.
- [ ] `oc get debug`
- [ ] `oc set debug <pod name>`
- [ ] `oc debug <pod name>`
- [ ] `oc get pod debug`

<details>
  <summary> Answer </summary>

  `oc debug <pod name>`
  
</details>

--- 

### OpenShift Jobs: 

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module5#-module-5-advanced-deployment-options" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

**Jobs in OpenShift** is a Kubernetes resource used to run pods until a specified number of them successfully complete. It's designed for tasks that need to run to completion, unlike Deployments which maintain a desired state of pods. Jobs are useful for batch processing, periodic tasks, and other situations where a finite set of work needs to be done. A job, in contrast to a replication controller, runs a pod with any number of replicas to completion. 

- Creating a Job: A job configuration consists of the following key parts:
  - A pod template, which describes the application the pod will create.
  - An optional `parallelism` parameter, which specifies how many pod replicas running in parallel should execute a job. If not specified, this defaults to the value in the `completions` parameter.
  - An optional `completions` parameter, specifying how many concurrently running pods should execute a job. If not specified, this value defaults to one.
  
```yaml
apiVersion: batch/v1 # its a kubernetes resource not OpenShift
kind: Job
metadata:
  name: hello-world-job
spec:
  parallelism: 1    # Optional value for how many pod replicas a job should run in parallel; defaults to `completions`.
  completions: 1    # Optional value for how many successful pod completions are needed to mark a job completed; defaults to one.
  template:         # Template for the pod the controller creates.
    metadata:
      name: hello-world-job
    spec:
      containers:
      - name: hello-world-job
        image: <image url>
      restartPolicy: OnFailure   # The restart policy of the pod. This does not apply to the job controller.
```
- When defining a Job, you can define its maximum duration by setting the `activeDeadlineSeconds` field. It is specified in seconds and is not set by default. When not set, there is no maximum duration enforced.

```bash
  spec:
    activeDeadlineSeconds: 1800 # 30 min
```

**Hands-on Walkthroughs** 
- How to configure pre-deployment hook for `rolling strategy` you will need 2 windows terminals for this excersice.  

    ```bash
    # on terminal 1
    oc get pods --watch
    ```
    ```bash
    # on terminal 2
    oc rollout latest dc/hello-world
    ```
    > output: you should see the rolling strategy getting applied as the new version is getting deployed then switched the network to the new one and terminating the old one!

  - Lets add a deployment hook to the application, and trigger another rollout.
    
    ```bash
    oc set deployment-hook dc/hello-world --pre -c hello-world -- /bin/echo hello from pre-deploy hook
    ```
    ```bash
    oc describe dc/hello-world
    ```
    > output: "Strategy:       Rolling
  Pre-deployment hook (pod type, failure policy: Ignore):" + `Container:  hello-world` & `Command:    /bin/echo hello from pre-deploy hoo`

- Now that we-ve verified out hook is configured, let's run the `oc rollout` command once again, just as we did before.
  
    ```bash
    oc rollout latest dc/hello-world
    ```
    > output: you will see a pod ends with `pre` and `deploy` pods and again same process normal rolling deployment after the `pre` was done!

    - In case, the deployment hook pod exited very quickly, so we're going to check the OpenShift event logs in order to validate. 
      - The events are a list of various OpenShift activities that have recently occurred this includes some basic deployment hook information.
        
        ```bash
        oc get events
        ```
        > output: confirm that the `pre-hook` ran successfully.
         
- How to configure the Recreate Deployment Strategy
  - Configuring the recreated strategy is a bit different from most of the commands that you have learned so far in this course. There's no command, such as `oc set deployment strategy`. Instead you need to modify the resource definition `YAML` directly. There are a couple of ways to do this.You can download a copy of the resource with `oc get -o YAML`, make changes and `re-upload` the changed definition using `oc create`. The oc tool provides a utility that automates all of these steps called `oc edit`.
    
    ```bash
    oc edit dc/hello-world
    ```
    > output: all you have to do is remove all of the keys except for the type, replace it with `Recreate`.
    
    ```yml
      strategy:
        type: Recreate # CHANGE HERE!
        rollingParams:
          updatePeriodSeconds: 1
          intervalSeconds: 1
          timeoutSeconds: 600
          maxUnavailable: 25%
          maxSurge: 25%
          pre:
            failurePolicy: Ignore
            execNewPod:
              command:
                - /bin/echo
                - hello
                - from
                - pre-deploy
                - hook
              containerName: hello-world
        resources: {}
        activeDeadlineSeconds: 21600
    ```
<p align="center">
<img src="/images/strategy-change.png" alt="OpenShift Training" style="width:500px; align="center"/>
</p>

- Lets check the changes
  
    ```bash
    oc describe dc/hello-world
    ```
    > output: you should see `Strategy:       Recreate`

- Now that you have updated the strategy to recreate, lets watch the pods as you trigger a deployment. The deployment should follow the recreate order as described earlier.

    ```bash
    # terminal 1
    oc get pods -w
    ```
    ```bash
    oc rollout latest dc/hello-world
    ```
<p align="center">
<img src="/images/pods-recreate.png" alt="OpenShift Training" style="width:500px; align="center"/>
</p>

> output:
> Just as with the rolling strategy, OpenShift will start a deployment pod first. However, things start to change pretty quickly after that. Instead of starting the new replication controller and pods, first the old replication controller terminates and takes down its pods. Then the deployment config will schedule the new replication controller and start the new pods.


- Launch a job 

```bash
cd ./labs-repo/job
```
```bash
oc create -f hello-job.yaml
```
> output: "job.batch/hello-job created"

```bash
oc decribe job
```
> output: 
```bash
Parallelism:      1
Completions:      1
Completion Mode:  NonIndexed
Start Time:       Thu, 24 Jul 2025 20:54:52 -0700
Pods Statuses:    1 Running / 0 Succeeded / 0 Failed
```

---

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module6" target="_blank">
    <img src="/images/nexticon.webp" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

<p align="left">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module4" target="_blank">
    <img src="/images/backred1.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>