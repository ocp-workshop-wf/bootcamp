[![Static Badge](https://img.shields.io/badge/CheatSheet-purple?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/CheatSheet.md)  [![Static Badge](https://img.shields.io/badge/OCP-CLI-red?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/ocpcli-cheatsheet.md)   [![Static Badge](https://img.shields.io/badge/Labs-maroon?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/tree/main/labs-repo)  [![Static Badge](https://img.shields.io/badge/RedHat-OpenShift-maroon?style=flat&logo=Redhat&logoSize=auto)
](https://docs.redhat.com/en/documentation/openshift_container_platform/4.19)   [![Static Badge](https://img.shields.io/badge/Kubernetes-black?style=flat&logo=Kubernetes&logoSize=auto)
](https://kubernetes.io/docs/home/)
## 🔹 Module 6: Mastering OpenShift
### 6.1 Templates
A template is a reusable definition of a set of OpenShift objects (like pods, services, routes, etc.) that can be parameterized and instantiated to create those objects within a project. Essentially, it's a way to automate the creation of multiple related resources, making it easier to deploy and manage applications.
    ***Template Parts***
        - Preamble: is just a header identifying the file as a template file. i.e(basic metadata: `name`, `kind`).
        - Object (Resource) List: Define the various pieces that make up your application. Each object in the list contains a full definition for a normal OpenShift resource.
        - Parameters: Allow you to make your application configurable.Common parameters include `credentials`, `database names`, more ...

**Hands-on Walkthroughs** 
- Using Templates: How to upload Template files?

    ```bash
    cd ./labs-repo/template
    ```
    ```bash
    oc create -f hello-world-template.yaml
    ```
    > output: template.template.openshift.io/hello-world created
    
    ```bash
    oc get template
    ```
    > output: 

    | NAME | DESCRIPTION | PARAMETERS | OBJECTS |
    | ---- | ----------- | ---------- | ------- |
    | hello-world |      | 1 (all set) | 4 |

    ```bash
    oc new-app hello-world
    ```
- How to use Template Parameters?

    ```bash
    oc get -o yaml template/hello-world
    ```
    > output: 
    ```yml
    .....
    spec:
      containers:
      - env:
        - name: MESSAGE
          value: ${MESSAGE} #This is the syntax that allows you to substitute parameter values into your template.
    .......
      parameters:
      - description: Message to respond to requests with
      displayName: Message
      name: MESSAGE
      value: Hello from the default value for the template
    ```
- Clean up:
  ```bash
  oc delete all --all
  ```
- Lets set some parameter values:

    ```bash
    oc new-app hello-world \
  -p MESSAGE="Hello from parameter override."
    ```
    > output: you should see a Success message.

    ```bash
    curl <your route>
    ```
    > output: "Hello from parameter override."

- Processing Templates - How to process Templates?
  - This means you can print out the final set of resources for a template taking parameters into account. Once you have the list of objects in a file, you can inspect and modify it manually beofre using `oc create`

    ```bash
    oc process hello-world
    ```
    > output: `JSON` format - you can specify it in yaml by using `-o yaml`

    ```bash
    oc process hello-world -o yaml
    ```
    > output: no longer contains {} as it shows the actual value

    ```yml
    .....
      spec:
        containers:
        - env:
          - name: MESSAGE
            value: Hello from the default value for the template
    ......
    ```
- Lets run the `oc process` command and give a parameter value.

    ```bash
    oc process hello-world -o yaml \
  -p MESSAGE="Hello from oc process"
    ```
    > output: you should see this -->
    
    ```yml
    .....
      spec:
        containers:
        - env:
          - name: MESSAGE
            value: Hello from oc process
    .....
    ```
  - Clean up:

    ```bash
    oc delete all --all
    ```

- How to use the output from `oc process` to actually create objects on the server.
  - First lets save the processed template to a file:
  
    ```bash
    oc process hello-world -o yaml \
    -p MESSAGE="Hello from oc process" \
    > processed-objects.yaml
    ```
  - Check the file content 
    
    ```bash
    head processed-objects.yaml
    ```
    
    > output: 

    ```yml
    apiVersion: v1
    items:
    - apiVersion: apps.openshift.io/v1
      kind: DeploymentConfig
      metadata:
        annotations:
          openshift.io/generated-by: PracticalOpenshift
        labels:
          app: hello-world
        name: hello-world
    ```
  - Deploy the processed file
    ```bash
    oc create -f processed-objects.yaml
    ```
    ```bash
    oc status #grab the route url
    ```
    ```bash
    curl <URL>
    ```
    > output: "Hello from oc process"

- How create a custom Template?
  - Get YAML for existing objects on the OpenShift server.
  - The dc,is... syntax lists the types of resources that you would like to export
  - Add hpa or any other type if you need them
    
    ```bash
    oc get -o yaml dc,is,bc,svc,route --export
    ```
    ```bash
    vi test-template.yaml
    ```
    - Steps for a custom template:
      - Change the items property to objects
      - Change kind from List to Template
      - Add a name property to the metadata section
      - Remove status from each resource
      - Remove most of metadata except for name, labels, and annotations
      - Remove any automatically-assigned resources such as service Virtual IPs and Route hosts
      - (optional) Add template parameters

- How to use built-in Templates
  
    ```bash
    oc get template -n openshift
    ```

    > output: list of templates.

    
    

    


### 🔬 Hands-on Lab: 
For Templates, you will create your own parameterized template for a small REST API with a database backend. First, you will need to get the application working on OpenShift.
    
  - For the REST API, you can use the `mysql-go-reader` directory in the labs project.
  - For MySQL, you can use the `mysql` image and `oc new-app`.
  - To get them working, you'll need to supply a few environment variables to both images. Specifically, you'll need to provide `MYSQL_USER`, `MYSQL_PASSWORD`, and `MYSQL_DATABASE`. Store these values in an OpenShift resource and use it to populate both sets of environment variables.
  - Expose a route for the application
  - Create an autoscaler that will scale the application from 3 to 10 pods depending on load

### Checklist 📋: 
- `MAX_PODS`: this should control the maximum number of pods for the REST API
- `MYSQL_USER`: This variable should be used for both the mysql and mysql-go-reader images
- `MYSQL_PASSWORD`: This variable should be used for both the mysql and mysql-go-reader images
- `MYSQL_DATABASE`: This variable should be used for both the mysql and mysql-go-reader images

> Supply default values so that the template will work without specifying any parameters with oc new-app.

--- 

### Quiz
> Q1: You can only work with templates as files on your filesystem.
- [ ] True
- [ ] False

<details>
  <summary> Answer </summary>

False: You can also upload the template to OpenShift as a standard REST API resource and work with it that way!

</details>

> Q2: You can use oc new-app with a template name in order to create an application according to the template.
- [ ] True
- [ ] False

<details>
  <summary> Answer </summary>

True: `oc new-app` does support templates as arguments.

</details>

> Q3: What syntax can you use to get YAML definitions for all configmaps and secrets in your project as a List?
- [ ] `oc get cm,secret --export`
- [ ] `oc get -o yaml all --export`
- [ ] `oc get -o yaml cm,secret --exprot`
- [ ] `oc get cm -o yaml, secret -o yaml`

<details>
  <summary> Answer </summary>

`oc get -o yaml cm,secret --exprot`

</details>

> Q4: What flag is used to specify a parameter value for oc new-app or oc process commands?
- [ ] `--param <value>`
- [ ] `-p <value>`
- [ ] `--p <value>`
- [ ] `parameter <value>`

<details>
  <summary> Answer </summary>

`-p <value>`

</details>

---


### 6.2 Health Check & Observability
***liveness probe*** in OpenShift (which leverages Kubernetes) is a mechanism used to determine if a container within a pod is still running and healthy. Its primary purpose is to detect and handle situations where an application might be running but has entered an unrecoverable state, such as a deadlock or a process hanging, and is no longer able to serve requests. It alos answer the question: 
> Should we restart this Pod?
> 
- Default is to probe every 10 seconds
- Restarts the pod after 3 failed Liveness checks.

***Readiness probe*** in OpenShift (and Kubernetes) is a mechanism used to determine if a container within a pod is ready to accept incoming network traffic. Unlike a liveness probe, which indicates whether a container is alive and should be restarted if it fails, a readiness probe focuses on whether the application inside the container is fully initialized and capable of serving requests. It also answers the question: 
> Should we send traffic to this pod?

- **Use case**: If you have been developing applications for a while, you have probably had the pleasure of dealing with an application that crashed unexpectedly after running successfully for a while. When this happens to an OpenShift Application that has `liveness checks` configured, OpenShift will automatically restart the pod for many types of workloads..

- Both Readiness and Liveness Probes have a few options available. The most common option for `REST APIs` is the `HTTP GIT probe`. This type of probe makes HTTP requests to your pod at specified intervals and reports `success` if the response code is between `200` and `399`. This is a natural fit for arrest API, and it should be your go-to Readiness and Liveness Probe solution unless you have a good reason to do otherwise. For applications that aren't serving arrest API, there's also an escape hatch present in the form of the Command execution probe. Also we got the most common check which is `TCP check`.

**Hands-on Walkthroughs** 

- How to configure a Liveness Probe?

```bash
oc set probe dc/hello-world \
  --liveness \
  --open-tcp=8080
#This liveness check that we're going to configure will try to set up a tcp connection to port 8080. If it succeeds, the probe succeeds,and if it fails, then the probe will fail.
```
> output: "deploymentconfig.apps.openshift.io/hello-world probes updated"

- Lets verify the Probe is configured correctly on the DeploymentConfig

```bash
oc describe dc/hello-world
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
oc set probe dc/hello-world \
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
oc set probe dc/hello-world \
  --liveness \
  --open-tcp=8080
```

> output: the Pod is no longer broken and its running successfuly

<p align="center">
<img src="/images/healthy.png" alt="OpenShift Training" style="width:500px; align="center"/>
</p>


### 🔬 Hands-on Lab: 
For DeploymentConfigs, you will edit your readiness probe to be incorrect, then make curl requests to the route and observe the behavior.

- Deploy the hello-world application

- Create a route for the application

- Add an invalid readiness check to your application

  - Use an incorrect port, a command that always returns 1, etc.

- Quickly after adding the check, make some requests to your application. They succeed--why? What's happening to the pods in the application?



### Checklist 📋: 
- `oc get events` should contain several "Readiness probe failed" messages

- `curl <your route>` still gives you the hello-world message

- You understand why your application is still able to serve traffic

### Quiz
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
Helm Charts are packages of pre-configured Kubernetes resources, acting as a package manager for Kubernetes applications. They simplify the definition, installation, and upgrading of complex applications on a Kubernetes cluster.A Helm Chart typically consists of the following components.

***Chart.yaml**
This file defines the metadata of the chart, including its name, version, description, and any dependencies on other charts.

***values.yaml***
This file contains default values for the configurable parameters within the chart's templates. Users can override these values during installation or upgrade to customize the application's deployment.

***templates/directory***
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
cd ./labs-repo/helm
```

### 🔬 Hands-on Lab: 
For Helm you will need to be creative - either you follow the same pattern I got here into the Helm directory or you might want to create your own, here are the step by step recipe.

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

### Checklist 📋: 
- You should be able to see the tgz file on your github page.
- You should be able to run `helm install <chart-name>`
- If everything works please run `helm uninstall <chart-name>` to delete all.

### Quiz
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