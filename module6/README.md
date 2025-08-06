[![Static Badge](https://img.shields.io/badge/Agenda-green?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/Agenda.md)[![Static Badge](https://img.shields.io/badge/CheatSheet-purple?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/other/CheatSheet.md) [![Static Badge](https://img.shields.io/badge/OCP-CLI-red?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/other/ocpcli-cheatsheet.md)   [![Static Badge](https://img.shields.io/badge/Labs-maroon?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/tree/main/labs-repo)  [![Static Badge](https://img.shields.io/badge/RedHat-OpenShift-maroon?style=flat&logo=Redhat&logoSize=auto)
](https://docs.redhat.com/en/documentation/openshift_container_platform/4.19)   [![Static Badge](https://img.shields.io/badge/Kubernetes-black?style=flat&logo=Kubernetes&logoSize=auto)
](https://kubernetes.io/docs/home/)
## 🔹 Module 6: Mastering OpenShift

## Table of Contents 

- [6.1 - Health Check & Observability](#61-health-check--observability) | [Lab](#-hands-on-lab-health-check) | [Quiz](#quiz-health-check)

- [6.2 - Helm Charts](#62-HelmCharts) | [Lab](#-hands-on-lab-helm) | [Quiz](#quiz-helm)

---

### 6.1 Health Check & Observability

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module6#-module-6-mastering-openshift" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>



If such a probe is configured, it disables liveness and readiness checks until it succeeds.

This type of probe is only executed at startup, unlike liveness and readiness probes, which are run periodically.

**liveness probe** in OpenShift (which leverages Kubernetes) is a mechanism used to determine if a container within a pod is still running and healthy. Its primary purpose is to detect and handle situations where an application might be running but has entered an unrecoverable state, such as a deadlock or a process hanging, and is no longer able to serve requests. It alos answer the question: 
> Should we restart this Pod?

- Default is to probe every 10 seconds
- Restarts the pod after 3 failed Liveness checks.

**Readiness probe** in OpenShift (and Kubernetes) is a mechanism used to determine if a container within a pod is ready to accept incoming network traffic. Unlike a liveness probe, which indicates whether a container is alive and should be restarted if it fails, a readiness probe focuses on whether the application inside the container is fully initialized and capable of serving requests. It also answers the question: 
> Should we send traffic to this pod?

- **Use case**: If you have been developing applications for a while, you have probably had the pleasure of dealing with an application that crashed unexpectedly after running successfully for a while. When this happens to an OpenShift Application that has `liveness checks` configured, OpenShift will automatically restart the pod for many types of workloads..

- Both Readiness and Liveness Probes have a few options available. The most common option for `REST APIs` is the `HTTP GIT probe`. This type of probe makes HTTP requests to your pod at specified intervals and reports `success` if the response code is between `200` and `399`. This is a natural fit for arrest API, and it should be your go-to Readiness and Liveness Probe solution unless you have a good reason to do otherwise. For applications that aren't serving arrest API, there's also an escape hatch present in the form of the Command execution probe. Also we got the most common check which is `TCP check`.

**Startup probe** A startup probe verifies whether the application within a container is started. This can be used to adopt liveness checks on slow starting containers, avoiding them getting killed by the kubelet before they are up and running.
  - Why use a Startup Probe? 
  > Without it, a liveness probe might kill your container before it has even fully started, especially in cases like: Application with long initialization phases, or containers waiting for external services

<p align="center">
<img src="/images/probes.gif" alt="OpenShift Training" style="width:500px; align="center"/>
</p>

  -  Deployment example using a startup probe
  
  ```yaml
  apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: hello-world-app
  spec:
    replicas: 1
    selector:
      matchLabels:
        app: hello-world-app
    template:
      metadata:
        labels:
          app: hello-world-app
      spec:
        containers:
        - name: hello-world
          image: hello-world:latest
          ports:
          - containerPort: 80
          startupProbe: # here we define the startupProbe
            httpGet:
              path: /
              port: 80
            failureThreshold: 30 # means it will try 3 times 
            periodSeconds: 10 # it will try every 10 sec.
          readinessProbe:
            httpGet:
              path: /
              port: 80
            periodSeconds: 5 
            initialDelaySeconds: 10 
          livenessProbe:
            httpGet:
              path: /
              port: 80
            periodSeconds: 10 
            initialDelaySeconds: 20
  ```
  > startupProbe: Makes an HTTP GET to / on port 80. It allows up to 30 failures (i.e., ~5 minutes) before declaring the container failed. Once the startup probe succeeds once, Kubelet stops checking startupProbe and switches to liveness and readiness checks.

  - If you get an error you should see this:

  <p align="center">
  <img src="/images/startup-error.png" alt="OpenShift Training" style="width:500px; align="center"/>
  </p>

**Resource**: [Kubernetes Documentation](https://kubernetes.io/docs/concepts/configuration/liveness-readiness-startup-probes/)

**Hands-on Walkthroughs** 

- How to configure a Liveness Probe?

```bash
oc set probe deployment/hello-world \
  --liveness \
  --open-tcp=8080
#This liveness check that we're going to configure will try to set up a tcp connection to port 8080. If it succeeds, the probe succeeds,and if it fails, then the probe will fail.
```
> output: "deploymentconfig.apps.openshift.io/hello-world probes updated"

- Lets verify the Probe is configured correctly on the DeploymentConfig

```bash
oc describe deployment/hello-world
```
> output: "    Liveness:           tcp-socket :8080 delay=0s timeout=1s period=10s #success=1 #failure=3"

<p align="center">
<img src="/images/livenessprob.png" alt="OpenShift Training" style="width:500px; align="center"/>
</p>

- What happens when Livness probes fail?
- So lets set the trigger to an incorrect port, this will trigger the failer behavior and cause OpenShift to restart the pod.
```bash
# on terminal 1
oc get pods -w
```

```bash
# on terminal 2
oc set probe deployment/hello-world \
  --liveness \
  --open-tcp=8081
```
> output: After about `30` seconds, you should get a new line at the bottom of your oc get pods output. The pod is running, but you should have a new value of `one` in the `RESTARTS` column. 30 seconds gave OpenShift enough time to send three liveness checks to the pod.Because the pod failed all three checks, OpenShift decided to restart the pod and as you wait another 30 seconds and another, you'll get two and three restarts just like this. There is one more step in the pod lifecycle that we have not covered so far in this course. This pod will restart every 30 seconds until it goes into the CrashLoopBackOff. I'm going to wait a bit for it to go into the state, and then we'll reset the probe back to its original value, and here we have the CrashLoopBackOff state. This pod will restart every 30 seconds until it goes into `CrashLoopBackOff`.

<p align="center">
<img src="/images/crashloopbackoff.png" alt="OpenShift Training" style="width:500px; align="center"/>
</p>

- Now lets fix what we did.
```bash
# on terminal 1
oc get pods -w
```

```bash
# on terminal 2
oc set probe deployment/hello-world \
  --liveness \
  --open-tcp=8080
```

> output: the Pod is no longer broken and its running successfuly

<p align="center">
<img src="/images/healthy.png" alt="OpenShift Training" style="width:500px; align="center"/>
</p>


### 🔬 Hands-on Lab (Health Check): 
For DeploymentConfigs, you will edit your readiness probe to be incorrect, then make curl requests to the route and observe the behavior.

- Deploy the hello-world application

- Create a route for the application

- Add an invalid readiness check to your application

  - Use an incorrect port, a command that always returns 1, etc.

- Quickly after adding the check, make some requests to your application. They succeed--why? What's happening to the pods in the application?


### Checklist 📋 (Health Check): 
- `oc get events` should contain several "Readiness probe failed" messages

- `curl <your route>` still gives you the hello-world message

- You understand why your application is still able to serve traffic

### Quiz (Health Check)
> Q1: Which probe can cause pod restarts?
- [ ] readiness
- [ ] liveness
- [ ] all the above
- [ ] None 


<details>
  <summary> Answer </summary>

    liveness
  

</details>

--- 

### 6.2 HelmCharts

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module6#-module-6-mastering-openshift" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

Helm Charts are packages of pre-configured Kubernetes resources, acting as a package manager for Kubernetes applications. They simplify the definition, installation, and upgrading of complex applications on a Kubernetes cluster.A Helm Chart typically consists of the following components.

***Chart.yaml**
This file defines the metadata of the chart, including its name, version, description, and any dependencies on other charts.

**values.yaml**
This file contains default values for the configurable parameters within the chart's templates. Users can override these values during installation or upgrade to customize the application's deployment.

**templates/directory**
This directory holds the Kubernetes manifest files (e.g., Deployments, Services, ConfigMaps) written using Go templating language. These templates are rendered with the values from `values.yaml` (or user-provided overrides) to generate the final Kubernetes YAML manifests for deployment.

- Helm example 
```bash
mychart/
├── Chart.yaml
├── values.yaml
├── charts/
│   └── (dependent charts go here)
└── templates/
    ├── NOTES.txt
    ├── _helpers.tpl
    ├── deployment.yaml
    ├── service.yaml
```
<p align="center">
<img src="/images/helmworkflow.webp" alt="OpenShift Training" style="width:500px; align="center"/>
</p>

---

**Hands-on Walkthroughs** 
- Access the labs directory and explore helm folder
  
```bash
cd ./labs-repo/helm/mq-helm
```
```bash
helm install my-release --dry-run . > all.yaml    
```

### 🔬 Hands-on Lab (Helm): 
For Helm you will need to be creative - either you follow the same pattern I got here into the Helm directory or you might want to create your own, here are the step by step recipe. using 
  ```bash
  helm create bootcamp
  ```

- Use the `service` / `route` of hello-world application and build up your own helm-chart.

- cd Create a personal Helm repository using a Github repo

  - Create a new github repository

  - Go to Settings > Pages tab

  - Under "Source", select the main branch and save.

  - Your repository should now have a public URL -> `https://<git-org>.github.io/<repo-name>/`
  
- Clone the git repo created above
- Create a template helm chart
  - `helm create <chart-name>`
  - `helm create mysamp`
- Lint the Chart
  - `helm lint . [[comment: need to 'cd' to repo folder ]]`
- Generate yaml from the chart 
  - `helm template .`
- Generate a `tgz` of your helm chart [[Make sure your out of the helm fold]]
  - `helm package <chart-name>`
- Generate / update `index.yaml` for your Helm repository
  - `helm repo index --url <helm-repo-url> .` # Don't forget the `.`
- Add, commit and push changes to git repo
- Add a Helm repository locally
  - `helm repo add <helm-repo-name> <helm-repo-url: not GitHub repo instead use github pages repo url> `
- Refresh contents of remote Helm repositories
  - `helm repo update`
- Search for charts available in a Helm repository
  - `helm search repo <helm-repo-name>`

### Checklist 📋 (Helm): 
- You should be able to see the tgz file on your github page.
- You should be able to run `helm install <chart-name>`
- If everything works please run `helm uninstall <chart-name>` to delete all.

### Quiz (Helm)
> Q1: How do you install a Helm chart?
- [ ]  `helm install myrelease myrepo`
- [ ] `helm install myrelease myrepo:latest`
- [ ] `helm install myrelease myrepo/mychart`
- [ ] `helm install myrelease latest`


<details>
  <summary> Answer </summary>

  `helm install myrelease myrepo/mychart`
  

</details>

> Q2: How do you upgrade a Helm release?
- [ ]  `helm upgrade myrelease myrepo/latest`
- [ ] `helm upgrade myrelease myrepo/v1`
- [ ] `helm upgrade myrelease latest`
- [ ] `helm upgrade myrelease myrepo/mychart`


<details>
  <summary> Answer </summary>

  `helm upgrade myrelease myrepo/mychart`
  

</details>

> Q3: How do you create a new Helm chart?

- [ ] `helm create mychart`
- [ ] `helm get mychart`
- [ ] `helm set mychart`
- [ ] `helm add mychart`


<details>
  <summary> Answer </summary>

  `helm create mychart`
  
</details>

---

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main" target="_blank">
    <img src="/images/complete.webp" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

<p align="left">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module5" target="_blank">
    <img src="/images/backred1.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>