[![Static Badge](https://img.shields.io/badge/Agenda-green?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/other/Agenda.md) [![Static Badge](https://img.shields.io/badge/CheatSheet-purple?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/other/CheatSheet.md) [![Static Badge](https://img.shields.io/badge/OCP-CLI-red?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/other/ocpcli-cheatsheet.md)   [![Static Badge](https://img.shields.io/badge/Labs-maroon?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/tree/main/labs-repo)  [![Static Badge](https://img.shields.io/badge/RedHat-OpenShift-maroon?style=flat&logo=Redhat&logoSize=auto)
](https://docs.redhat.com/en/documentation/openshift_container_platform/4.19)   [![Static Badge](https://img.shields.io/badge/Kubernetes-black?style=flat&logo=Kubernetes&logoSize=auto)
](https://kubernetes.io/docs/home/)
## 🔹 Module 5: Advanced Deployment Options

## Table of Contents 

- [5.1 - Deployment Strategies](#51-openshift-deployment-strategies) | [Lab](#-hands-on-lab-deployment-strategies) | [Quiz](#quiz-deployment-strategies)

- [5.2 - Scaling and Debuging Your Application](#52-scaling-and-debuging-your-application) | [Quiz](#quiz-scaling) 

- [5.3 - OpenShift Jobs](#53-openshift-jobs) 

### 5.1 Deployment Strategies 

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module4#-module-4-application-deployment-and-management" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

| Stategy Type | Step 1 | Step 2 | Step 3 |
| ------------ | -------- | ----- | ------ |
| Rolling Strategy "Default" | Start new version | Switch traffic to new version | Stop old version | 
| Recreate Strategy | Stop the old version | Start new Version | Switch Traffic to new version |
| Custom Strategy | Start | Run custom deployment image | End |


***Resources***
- [12 App Factor](https://12factor.net/)
- [ Custom Strategy using custom Image ](https://docs.openshift.com/en/container-platform/3.11/dev_guide/deployments/deployment_strategies.html#custom-strategy)

<p align="center">
<img src="/images/rollingdeploymentgif.gif" alt="OpenShift Training" style="width:500px; align="center"/>
</p>

- A rolling strategy supports pre and post hooks. The pre hook runs, of course, before the deployment Config, starts a new version and the post hook runs after the deployment Config stops the old version.

<p align="center">
<img src="/images/Recreate-Deployment.jpg" alt="OpenShift Training" style="width:500px; align="center"/>
</p>

- The recreate strategy, on the other hand, supports pre and post hooks, as well as a mid hook that is executed while no Pods are running. That is, after the recreate strategy stops the old version, but before it starts the new one.

- This image defines a Strategy Hook in OpenShift using a YAML-style syntax, specifically for a pre hook during a deployment strategy. Here's a regenerated clean version of the example for reference or use in documentation:

  ```yml
  pre:
    failurePolicy: Abort
    execNewPod:
      containerName: hello-world
      command: ["/bin/echo", "Hello from pre-hook"]
      env:
        - name: DEMO_ENV_VAR
          value: DEMO_VALUE
      volumes: []
  ```

- `pre`: – Indicates the hook will run before the strategy (rolling or recreate).

- `failurePolicy`: – Controls what happens if the hook fails. Options are:

- `Abort` – stop the deployment

- `Retry` – try again

- `Ignore` – continue regardless

- `execNewPod`: – Runs the hook in a new pod with:

- `containerName`: – the target container

- `command`: – command to execute (in this case, a simple echo)

- `env`: – environment variables

- `volumes`: – volumes to mount (empty in this example)

**Blue-Green Deployment**:
  Blue-Green deployment is a strategy where two identical environments (called "blue" and "green") are maintained. The "blue" environment is the live/production system, and the "green" is the staging version running the new code. When a new version is ready:

  - It’s deployed to the green environment.

  - After testing and validation, traffic is switched from blue to green.

  - If issues arise, rollback is as simple as redirecting traffic back to blue.

  - How it works in OpenShift: 
    - Traffic switch is achieved using `Routes`, you can update the route's `spec.to.name` field to point to a different `service` (e.g from `my-blue-app` to `my-green-app`).
    
    </br>

    <p align="center">
    <img src="/images/gbdeploy.gif" alt="OpenShift Training" style="width:500px; align="center"/>
    </p>

**Hands-on Walkthroughs**  

- How to configure pre-deployment hook for `rolling strategy` you will need 2 windows terminals for this excersice.  

    ```bash
    # on terminal 1
    oc get pods --watch
    ```
    ```bash
    # on terminal 2
    oc rollout latest deployment/hello-world
    ```
    > output: you should see the rolling strategy getting applied as the new version is getting deployed then switched the network to the new one and terminating the old one!

  - Lets add a deployment hook to the application, and trigger another rollout.
    
    ```bash
    oc set deployment-hook deployment/hello-world --pre -c hello-world -- /bin/echo hello from pre-deploy hook
    ```
    ```bash
    oc describe deployment/hello-world
    ```
    > output: "Strategy:       Rolling
  Pre-deployment hook (pod type, failure policy: Ignore):" + `Container:  hello-world` & `Command:    /bin/echo hello from pre-deploy hoo`

- Now that we-ve verified out hook is configured, let's run the `oc rollout` command once again, just as we did before.
  
    ```bash
    oc rollout latest deployment/hello-world
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
    oc edit deployment/hello-world
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
    oc describe deployment/hello-world
    ```
    > output: you should see `Strategy:       Recreate`

- Now that you have updated the strategy to recreate, lets watch the pods as you trigger a deployment. The deployment should follow the recreate order as described earlier.

    ```bash
    # terminal 1
    oc get pods -w
    ```
    ```bash
    oc rollout latest deployment/hello-world
    ```
<p align="center">
<img src="/images/pods-recreate.png" alt="OpenShift Training" style="width:500px; align="center"/>
</p>

> output:
> Just as with the rolling strategy, OpenShift will start a deployment pod first. However, things start to change pretty quickly after that. Instead of starting the new replication controller and pods, first the old replication controller terminates and takes down its pods. Then the deployment config will schedule the new replication controller and start the new pods.

- Blue Green Deployment 
  ```yaml
  # Route pointing initially to blue
  apiVersion: route.openshift.io/v1
  kind: Route
  metadata:
    name: my-app
  spec:
    to:
      kind: Service
      name: my-app-blue
  ```
  > once testing on the green environment is complete, switch:

  ```yaml
  spec:
  to:
    kind: Service
    name: my-app-green
  ```
  - You can also Patch your route by running the following command
  ```bash
  oc patch route my-app -p '{"spec":{"to":{"name":"my-green-app"}}}'
  ```
  > output: Quick rollback or just re-point the route, with that your now able to test in production-like environment, and also you reduced the risk of downtime during deployment.

- Lets try to do it ourseleves:
  ```bash
  # deploy blue
  oc new-app quay.io/practicalopenshift/hello-world --name=blue-app
  ```
  ```bash
  # expose blue
  oc expose service/blue-app
  ```
  ```bash
  # deploy green
  oc new-app quay.io/practicalopenshift/hello-world --name=green-app
  ```
  ```bash
  # expose green
  oc expose service/green-app
  ```
  - Before switching :

    <p align="center">
    <img src="/images/blue.png" alt="OpenShift Training" style="width:500px; align="center"/>
    </p>

  - Lets switch route to green 

    ```bash
    oc patch route blue-app -p '{"spec":{"to":{"name":"green-app"}}}'
    ```
    > output: "route.route.openshift.io/blue-app patched"
  
  - After switching: 

    <p align="center">
    <img src="/images/afterblue.png" alt="OpenShift Training" style="width:500px; align="center"/>
    </p>

  - Rollback 
    ```bash
    oc patch route blue-app -p '{"spec":{"to":{"name":"blue-app"}}}'
    ```
    > output: it should be the same as the `Before Switching Image` 

 ***Resource*** [RedHat Documentation](https://www.redhat.com/en/topics/devops/what-is-blue-green-deployment)

 **Canary Release**: 
 Is a deployment strategy where a new version of the application is released to a small subset of users first. If no issues are detected, traffic is gradually increased to the new version until it becomes fully live.
The name comes from the "canary in the coal mine" analogy — testing in a low-risk environment before exposing to all users.

- How it works in OpenShift: 
In OpenShift, Routes can be used to direct traffic to multiple backends with weights, allowing a portion of traffic to be sent to a new version. This works well when:

  - You have two versions of the same app deployed (v1 and v2).
  - You use a single Route with traffic weights defined.

  **Compare**
  - Blue Green vs Canary
    - Blue Green: Switches traffic between two complete environments (blue and green).
    - Canary: Gradually shifts traffic to a new version while monitoring for issues.

  | Feature               | Blue Green                     | Canary                          |
  |-----------------------|--------------------------------|---------------------------------|
  | Deployment Scope      | Entire environment             | Subset of users                 |
  | Traffic Switching      | Instant switch                 | Gradual shift                   |
  | Risk Level            | Higher (all or nothing)       | Lower (monitoring during rollout) |
  | Rollback Complexity    | Simple (switch back)           | More complex (gradual rollback) |
  | User Impact            | All users at once              | Small group initially           |
  | Monitoring              | Less emphasis on monitoring     | High emphasis on monitoring      |

  > Canary is useful if new version is backwards compatible vs Blue Green is more suitable for major changes that require a full environment switch NOT backwards compatible.

**Hands-on Walkthroughs**
  - In this example we are looking at a weighted routing.
  ```yaml
    apiVersion: route.openshift.io/v1
    kind: Route
    metadata:
      name: my-app
    spec:
      to:
        kind: Service
        name: my-app-v1
        weight: 80 # app v1 
      alternateBackends:
        - kind: Service
          name: my-app-v2
          weight: 20 # app v2  
  ```
  > In this example 80% of traffic goes to V1, 20% is routed to V2. You can always adjust the weights (e.g, 60/40, 50/50, 0/100).

  - The Ideal way to use Canary:
    1. Start with 90/10 split.
    2. Monitor for errors, latency and logs.
    3. If stable, move to 70/30, then 50/50.
    4. Continue until 0/100 to fully promote.
    5. Remove the old version when your done.
    > By following these steps you will have a very low risk exposure, easy to monitor real-time behavior of the new version, and can be automated for rollout.

    <p align="center">
    <img src="/images/canarydgif.gif" alt="OpenShift Training" style="width:500px; align="center"/>
    </p>


  - Lets try it out:
      ```bash
      # deploy v1
      oc new-app quay.io/practicalopenshift/hello-world --name=app-v1
      ```
      ```bash
      oc expose service/app-v1
      ```
      ```bash
      # deploy v2
      oc new-app quay.io/practicalopenshift/hello-world --name=app-v2
      ```
      > after deploying app v2 we have to craft our route with the desired weights in this example ill do 90/10

      - The weight yaml 
      ```yaml
      apiVersion: route.openshift.io/v1
      kind: Route
      metadata:
        name: my-app
      spec:
        to:
          kind: Service
          name: app-v1
          weight: 90
        alternateBackends:
          - kind: Service
            name: app-v2
            weight: 10
      ```
      > once you have your yaml ready you will have to `Apply` it to your cluster in this case, we will save the file as `canary-route.yaml`
      ```bash
      oc apply -f canary.yaml
      ```
      > output "route.route.openshift.io/canary-route created"

    - Lets monitor the traffic now 
      ```bash
      oc logs -f deployment/app-v2
      ```
      ```bash
      # curl the url 
      curl http://canary-route-raafat-dev.apps.rm3.7wse.p1.openshiftapps.com
      ```
      > output: " Welcome! You can change this message by editing the MESSAGE environment variable."  

    - Increase traffic gradually:

      <p align="center">
      <img src="/images/canarybefore.png" alt="OpenShift Training" style="width:500px; align="center"/>
      </p>

      ```bash
      # 70/30
      oc patch route canary-route -p '{"spec":{"to":{"weight":70},"alternateBackends":[{"kind":"Service","name":"app-v2","weight":30}]}}'
      ```
      
      <p align="center">
      <img src="/images/canaryafter.png" alt="OpenShift Training" style="width:500px; align="center"/>
      </p>

    - Clean up:
      ```bash
      oc delete all -l app=hello-world
      ```
  **Resource** [Redhat Documentation](https://developers.redhat.com/articles/2024/03/26/canary-deployment-strategy-openshift-service-mesh#products_canary_deployment)
     

---

### 🔬 Hands-on Lab (Deployment strategies): 
For Deployment Hooks, you will add a mid-deployment hook for the recreate strategy

- Create a new project called "advanced-dc-labs"

- Deploy the hello-world application

- Switch your application to use the Recreate strategy

- Add a mid-deployment hook that prints out "Hello from mid-Deployment hook."

- Roll out a new version of your application

### Checklist 📋 (Deployment strategies): 

- `oc get events` output contains your message from step 4

- `oc describe deployment/hello-world` shows the Recreate strategy and hook

### Quiz (Deployment strategies)

> Q1: Which deployment strategy always has downtime during the deployment?
- [ ] Rolling
- [ ] Recreate
- [ ] Custom
- [ ] All the above


<details>
  <summary> Answer </summary>

    Recreate 
  

</details>

> Q2: How many deployment strategy hooks does the rolling strategy have?
- [ ] 1
- [ ] 2
- [ ] 3
- [ ] 4


<details>
  <summary> Answer </summary>

   2 `pre & post` deployment
  

</details>

> Q3: What is the oc command to change from rolling to recreate strategy?
- [ ] `oc set deployment-strategy <deployment name>`
- [ ] `oc set deployment-strategy --recreate <deployment name>`
- [ ] `oc set strategy <deployment name>`
- [ ] There isn't one


<details>
  <summary> Answer </summary>

   There isn't one
  

</details>

> Q4: A deployment Strategy that allows you to smoothly switch traffic 
- [ ] Blue-Green 
- [ ] Rolling Strategy
- [ ] Canary Strategy
- [ ] Recreate Strategy


<details>
  <summary> Answer </summary>

   Canary Strategy
  

</details>

> Q4: A deployment Strategy that allows you to switch traffic between to running application at the same time. 
- [ ] Blue-Green 
- [ ] Rolling Strategy
- [ ] Canary Strategy
- [ ] Recreate Strategy


<details>
  <summary> Answer </summary>

   Blue-Green 
  

</details>

--- 



### 5.2 Scaling and Debuging Your Application 

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
    oc scale deployment/hello-world --replicas=3
    ```
    ```bash
    oc describe deployment/hello-world
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
  --cpu-percent=10
```
  > output: horizontalpodautoscaler.autoscaling/hello-world autoscaled
```bash
# access the pod terminal and run this command to increase cpu 
while true; do sha1sum /dev/zero; done
```
```bash
# In terminal 2
  watch oc get hpa
```
> output: includes all details about the HPA specially the Targets.

<p align="center">
<img src="/images/hpa-scaling.png" alt="OpenShift Training" style="width:500px; align="center"/>
</p>

  
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

### 5.3 OpenShift Jobs: 

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

- While jobs don't have strategies like deployment, you can:
  - Use `initContainers` for setup before the main container runs.
  - Run cleanup via another job after completion
  ```yaml
   initContainers:
   - name: setup
     image: busybox
     command: ['sh', '-c', 'echo Setup running...']
  ```
  > That your accessing the shell and running the command echo.


**Hands-on Walkthroughs** 

- Launch a job 

```bash
cd ./labs-repo/job
```
```bash
oc apply -f hello-job.yaml
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

### Quiz (Jobs)
> Q1: What does an OpenShift Job do?
- [ ] Runs infinitely 
- [ ] Creates a pod for monitoring
- [ ] Runs a task to completion
- [ ] Deploys a web server 

<details>
  <summary> Answer </summary>

    Runs a task to completion
  
</details>

> Q2: Which field ensures a Job won’t restart the pod after failure?
- [ ] `replicas`
- [ ] `restartPolicy: Never`
- [ ] `strategy: Recreate`
- [ ] `podSelector`

<details>
  <summary> Answer </summary>

   `restartPolicy: Never`
  
</details>

> Q3: How do you view logs of a job’s pod?
- [ ] `oc describe job <name>`
- [ ] `oc logs <job-name>`
- [ ] `oc logs -l job-name=<name>`
- [ ] `oc get events`

<details>
  <summary> Answer </summary>

  `oc logs -l job-name=<name>`
  
</details>
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