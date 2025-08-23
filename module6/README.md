[![Static Badge](https://img.shields.io/badge/Agenda-green?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/Agenda.md)[![Static Badge](https://img.shields.io/badge/CheatSheet-purple?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/other/CheatSheet.md) [![Static Badge](https://img.shields.io/badge/OCP-CLI-red?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/other/ocpcli-cheatsheet.md)   [![Static Badge](https://img.shields.io/badge/Labs-maroon?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/tree/main/labs-repo)  [![Static Badge](https://img.shields.io/badge/RedHat-OpenShift-maroon?style=flat&logo=Redhat&logoSize=auto)
](https://docs.redhat.com/en/documentation/openshift_container_platform/4.19)   [![Static Badge](https://img.shields.io/badge/Kubernetes-black?style=flat&logo=Kubernetes&logoSize=auto)
](https://kubernetes.io/docs/home/) [![Static Badge](https://img.shields.io/badge/OpenShift-Bootcamp-red?style=social)
](https://github.com/ocp-workshop-wf)
## 🔹 Module 6: Mastering OpenShift

## Table of Contents 

- [6.1 - Health Check & Observability](#61-health-check--observability) | [Hands-on-Walkthrough](#hands-on-walkthrough-health-check) | [Lab](#-hands-on-lab-health-check) | [Quiz](#quiz-health-check)

- [6.2 - Helm Charts](#62-HelmCharts) | [Hands-on-Walkthrough](#hands-on-walkthrough-helm) | [Lab](#-hands-on-lab-helm) | [Quiz](#quiz-helm)

- [6.3 - Bash Scripting](#bash-scripting) | [Hands-on-Walkthrough](#hands-on-walkthrough-bash-scripting) | [Lab](#-hands-on-lab-bash-scripting) | [Quiz](#quiz-bash-scripting)

- [6.4 - Useful Resources](#63-useful-resources)

---

### 6.1 Health Check & Observability

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module6#-module-6-mastering-openshift" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>



If such a probe is configured, it disables liveness and readiness checks until it succeeds.

This type of probe is only executed at startup, unlike liveness and readiness probes, which are run periodically.

**liveness probe** in OpenShift (which leverages Kubernetes) is a mechanism used to determine if a container within a pod is still running and healthy. Its primary purpose is to detect and handle situations where an application might be running but has entered an unrecoverable state, such as a deadlock or a process hanging, and is no longer able to serve requests. It also answer the question: 
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
  > startupProbe: Makes an HTTP GIT to / on port 80. It allows up to 30 failures (i.e., ~5 minutes) before declaring the container failed. Once the startup probe succeeds once, Kubelet stops checking startupProbe and switches to liveness and readiness checks.

  - If you get an error you should see this:

  <p align="center">
  <img src="/images/startup-error.png" alt="OpenShift Training" style="width:500px; align="center"/>
  </p>

**Resource**: [Kubernetes Documentation](https://kubernetes.io/docs/concepts/configuration/liveness-readiness-startup-probes/)

### Hands-on Walkthrough (Probes)

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

**Chart.yaml**
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

### Hands-on Walkthrough (Helm)
- Access the labs directory and explore helm folder
  
```bash
cd ./labs-repo/6.2-helm/mq-helm
```
```bash
helm install my-release --dry-run . > all.yaml    
```

---

- Create a Helm chart for your application

  ```bash
  helm create myapp
  ```
  > output: A new Helm chart named "myapp" is created with the default directory structure.

- Update the `values.yaml` file with your application-specific configurations adding an Environment variable for MESSAGE="HELLO FROM HELM".
  ```yaml
  replicaCount: 1

  image:
    repository: myapp
    tag: latest
    pullPolicy: IfNotPresent

  service:
    type: ClusterIP
    port: 8080

  ingress:
    enabled: false

  env:
    - name: MESSAGE
      value: "HELLO FROM HELM"
  ```
- Update the `templates/deployment.yaml` file to use the values from `values.yaml`.

  ```yaml
  apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: {{ .Release.Name }}
    labels:
      app: {{ .Release.Name }}
  spec:
    replicas: {{ .Values.replicaCount }}
    selector:
      matchLabels:
        app: {{ .Release.Name }}
    template:
      metadata:
        labels:
          app: {{ .Release.Name }}
      spec:
        containers:
          - name: {{ .Release.Name }}
            image: {{ .Values.image.repository }}:{{ .Values.image.tag }}
            ports:
              - containerPort: {{ .Values.service.port }}
            env:
              - name: {{ .Values.env[0].name }}
                value: {{ .Values.env[0].value }}
  ```

- Update the templates/service.yaml

  ```yaml
  apiVersion: v1
  kind: Service
  metadata:
    name: {{ .Release.Name }}
    labels:
      app: {{ .Release.Name }}
  spec:
    type: {{ .Values.service.type }}
    ports:
      - port: {{ .Values.service.port }}
        targetPort: {{ .Values.service.port }}
    selector:
      app: {{ .Release.Name }}
  ```

- Update the templates/route.yaml

  ```yaml
  apiVersion: route.openshift.io/v1
  kind: Route
  metadata:
    name: {{ .Release.Name }}
    labels:
      app: {{ .Release.Name }}
  spec:
    to:
      kind: Service
      name: {{ .Release.Name }}
    port:
      targetPort: {{ .Values.service.port }}
  ```

- Run Helm install dry-run on OpenShift

  ```bash
  helm install myapp ./myapp --dry-run
  ```

- Helm template

  ```bash
  helm template myapp ./myapp
  ```
  > output: You should see the rendered Kubernetes manifests for your application.

- Helm install

  ```bash
  helm install myapp ./myapp
  ```
  > output: You should see the Helm release created.

- Check the generated resources

  ```bash
  oc get all 
  ```
  > output: You should see the resources created by the Helm chart in your namespace.

- Verify the deployment
  ```bash
  oc get all
  ```
  > output: You should see the resources created by the Helm chart in your namespace.

- Uninstall Helm deployment 

  ```bash
  helm uninstall myapp
  ```
  > output: You should see the resources being deleted.

- Package the Helm chart
  ```bash
  helm package myapp
  ```
  > output: Successfully packaged chart and saved it to: /Users/ray/bootcamp/labs-repo/6.2-helm/myapp-0.1.0.tgz

- Create a Helm repository on GitHub
  - Create a new GitHub repository
  - Go to Settings > Pages tab
  - Under "Source", select the main branch and save.
  - Your repository should now have a public URL -> `https://<git-org>.github.io/<repo-name>/`

- Push the Helm chart to your GitHub repository
  ```bash
  git add myapp-*.tgz
  git commit -m "Add Helm chart for myapp"
  git push
  ```
  > output: You should see the Helm chart pushed to your GitHub repository.

- Update the Helm repository index
  ```bash
  helm repo index --url https://<git-org>.github.io/<repo-name>/ .
  ```
- Verify the Helm repository
  ```bash
  helm repo add <helm-repo-name> https://<git-org>.github.io/<repo-name>/
  helm repo update
  ```
  > output: You should see the Helm repository updated with the new chart.


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

### Bash scripting 
Bash scripting is a way to automate tasks on Unix-like systems. It allows users to write sequences of commands in a file and execute them as a single script. Bash (Bourne Again SHell) is the most commonly used shell in Linux. Key Concepts:

  - Shebang (#!/bin/bash)

  - Variables and parameters

  - Conditionals (if, else, elif)

  - Loops (for, while, until)

  - Functions

  - Input/output (stdin, stdout, stderr)

  - File operations and redirection

**Resources:**

  - [GNU Bash Manual](https://www.gnu.org/software/bash/manual/)

  - [ShellCheck – Linter for Bash scripts](https://www.shellcheck.net/)

  - [TLDP Bash Guide](https://tldp.org/LDP/Bash-Beginners-Guide/html/)

### Hands-on Walkthrough (Bash Scripting)
- Your first script

  ```bash
  #!/bin/bash
  echo "Hello, Bash!"
  ```
- Save this as hello.sh, then run:
  ```bash
  chmod +x hello.sh
  ./hello.sh
  ```
  > output: "Hello, Bash!"

- Variables and parameters

  ```bash
  #!/bin/bash
  NAME="World"
  echo "Hello, $NAME!"
  ```
- Save this as hello_var.sh, then run:
  ```bash
  chmod +x hello_var.sh
  ./hello_var.sh
  ```
  > output: "Hello, World!"  

- Conditionals

  ```bash
  #!/bin/bash
  if [ "$1" == "hello" ]; then
      echo "Hello, World!"
  else
      echo "Goodbye, World!"
  fi
  ``` 
- Save this as conditional.sh, then run:
  ```bash
  chmod +x conditional.sh
  ./conditional.sh hello
  ```
  > output: "Hello, World!"
  ```bash
  ./conditional.sh goodbye
  ```
  > output: "Goodbye, World!"
  
  ```bash  
  ./conditional.sh
  ```
  > output: "Goodbye, World!" (default case)

- Loops

  ```bash
  #!/bin/bash
  for i in {1..5}; do
      echo "Iteration $i"
  done
  ```
- Save this as loop.sh, then run:
  ```bash
  chmod +x loop.sh
  ./loop.sh
  ```
  > output:
  ```
  Iteration 1
  Iteration 2
  Iteration 3
  Iteration 4
  Iteration 5
  ```
- Functions

  ```bash
  #!/bin/bash
  greet() {
      echo "Hello, $1!"
  }
  greet "Bash"
  ```
- Save this as functions.sh, then run:
  ```bash
  chmod +x functions.sh
  ./functions.sh
  ```
  > output: "Hello, Bash!"

- Input/output

  ```bash
  #!/bin/bash
  echo "Enter your name:"
  read NAME
  echo "Hello, $NAME!"
  ```
- Save this as input_output.sh, then run:
  ```bash
  chmod +x input_output.sh
  ./input_output.sh
  ```
  > output: "Enter your name:" (then type your name and press Enter)
  ```
  Hello, YourName!
  ```
- File operations and redirection

  ```bash
  #!/bin/bash
  echo "This is a test file." > testfile.txt
  cat testfile.txt
  ```
- Save this as file_ops.sh, then run:
  ```bash
  chmod +x file_ops.sh
  ./file_ops.sh
  ```
  > output: "This is a test file."

- Clean up created files:
  ```bash
  rm hello.sh hello_var.sh conditional.sh loop.sh functions.sh input_output.sh file_ops.sh
  rm testfile.txt
  ```
- Resources:

  - [GNU Bash Manual](https://www.gnu.org/software/bash/manual/)

  - [ShellCheck – Linter for Bash scripts](https://www.shellcheck.net/)

  - [TLDP Bash Guide](https://tldp.org/LDP/Bash-Beginners-Guide/html/)

### 🔬 Hands-on Lab:
For Bash scripting, you will create a simple script that automates a task on your OpenShift cluster.
- Create a script that does the following:
  - Lists all pods in your project
  - Checks the status of each pod
  - If a pod is not running, it should print a message indicating which pod is not running
  - If all pods are running, it should print a success message
```bash
#!/bin/bash
pods=$(oc get pods -o jsonpath='{.items[*].status.phase}')
all_running=true
for pod in $pods; do
  if [ "$pod" != "Running" ]; then
    echo "Pod $pod is not running."
    all_running=false
  fi
done
if $all_running; then
  echo "All pods are running."
else
  echo "Some pods are not running."
fi
```
- Save this script as `check_pods.sh`, then run:
```bash
chmod +x check_pods.sh
./check_pods.sh
```
> output: Depending on the status of your pods, it will either list the pods that are
not running or print "All pods are running."
- Clean up:
```bash
rm check_pods.sh
```

### Checklist 📋
- The script lists all pods in your project
- The script checks the status of each pod
- The script prints a message for each pod that is not running
- The script prints a success message if all pods are running  

### Quiz (Bash Scripting)
> Q1: What is the purpose of the shebang (#!/bin/bash) in a Bash script?
- [ ] It defines the script's name
- [ ] It specifies the interpreter to use for executing the script
- [ ] It indicates the script's version   
- [ ] It sets the script's permissions
<details>
  <summary> Answer </summary>

  It specifies the interpreter to use for executing the script 
</details>

> Q2: How do you declare a variable in a Bash script?
- [ ] var name="value"
- [ ] name = "value"
- [ ] name="value"
- [ ] name: "value"
<details>
  <summary> Answer </summary>
  name="value"
</details>

> Q3: What is the purpose of the `read` command in a Bash script?
- [ ] To read a file
- [ ] To read user input from the terminal
- [ ] To read a variable's value
- [ ] To read a command's output
<details>
  <summary> Answer </summary>
  To read user input from the terminal
</details>

> Q4: How do you create a function in a Bash script?
- [ ] function myfunc() { ... }
- [ ] myfunc() { ... }
- [ ] create myfunc() { ... }
- [ ] def myfunc() { ... }
<details>
  <summary> Answer </summary>
  myfunc() { ... }
</details>

---

### 6.4 Useful Resources

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module6#-module-6-mastering-openshift" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

- [KodeKloud - Mohammed Mummabeth](https://www.youtube.com/@KodeKloud)

- [TechWorld with Nana](https://www.youtube.com/@TechWorldwithNana)


# CONGRATULATIONS! 🎉
## Let's connect on LinkedIn
  <p align="center">
  <img src="/images/qr.png" alt="OpenShift Training" style="width:250px; align="center"/>
  </p>


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

---