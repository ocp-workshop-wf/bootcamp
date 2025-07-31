### Webhooks
Is a method for one application to automatically send real-time data to another application when a specific event occurs. It's essentially an automated messaging system that allows applications to communicate with each other without needing to constantly "poll" or check for updates. And this is one of the key features that enables CICD!

| OpenShift | Git Repo |
| ---------- | ------ |
| Exposes HTTPS endpoint that starts when called | Calls endpoint when developers push code|

> The way the integration is set up is that the hook is exposed on the downstream system, OpenShift in our case, and the hook is configured on the upstream system. Your Git Repository is configured to call the endpoint that's exposed on OpenShift. To mitigate the risk of having a public endpoint on your OpenShift cluster, OpenShift requires Webhook clients to pass a special token string along with a request. This token value is automatically generated for BuildConfigs created with the `oc new-app` and the `oc new-build` command line tools. Once you have set up the Webhook correctly on the GitLab side, anytime a developer pushes an update GitLab will make an HTTPS request to the Webhook URL and pass along the token. When OpenShift gets the request it will start a new Build for the linked BuildConfig. This way, you can set up an automated pipeline allowing you to push automatically to your OpenShift cluster simply by pushing to a Git Repository.

<p align="center">
<img src="/images/webhook.png" alt="Webhook" style="width:500px; align="center"/>
</p>


**Hands-on Walkthroughs**  

- How to call a Webhook Manually? We need 2 pieces of data!
  - Secret Token
  - Webhook URL

```bash
oc get -o yaml buildconfig/hello-world
```
> output: Look for `triggers.generic.secret` grab that copy and now you have your token.

```bash
export GENERIC_SECRET=<SECRET-FROM-BUILDCONFIG>
```
- Now we need to get the URL information 

```bash 
oc describe bc/hello-world
```
> output: look for `Webhook Generic.URL` copy and generate a variable as well.

```bash
export WEBHOOK_URL=https://api.rm3.7wse.p1.openshiftapps.com:6443/apis/build.openshift.io/v1/namespaces/raafat-dev/buildconfigs/hello-world/webhooks/$GENERIC_SECRET/generic
```
- Then we need to curl that URL

```bash
curl -X POST -k $WEBHOOK_URL
```
> output: we should see status `Success`

- Building from the update-message branch

```bash
oc new-build oc new-build https://gitlab.com/therayy1/hello-world.git#update-message
```
> output: we need to check the logs to verify we are at the correct branch

```bash
oc logs -f bc/helloworld
```
> output: on `step 3` ENV MESSAGE "Welcome! I was changed in a branch."

- Use --context-dir to build from a subdirectory to deploy a specific project in case you got many projects.

```bash
oc new-build https://gitlab.com/practical-openshift/labs.git --context-dir hello-world
```

- Configuring build hooks
  - First we need to clean our cluster from any builds and run `oc status` to make sure.

```bash
oc new-build https://gitlab.com/therayy1/hello-world.git 
```
> output: "Success"

```bash
oc set build-hook bc/hello-world \
  --post-commit \
  --script="echo Hello from build hook"
```
> output: "buildconfig.build.openshift.io/hello-world hooks updated"

```bash
oc describe bc/hello-world
```
> output: 
```yaml
.....
Strategy:               Docker
URL:                    https://gitlab.com/therayy1/hello-world.git
From Image:             ImageStreamTag golang:1.17
Output to:              ImageStreamTag hello-world:latest
Post Commit Hook:       ["/bin/sh", "-ic", "echo Hello from build hook"]
....
```
- Lets run a build and verify that this message shows up in the output

```bash
oc start-build bc/hello-world
```
> output: "build.build.openshift.io/hello-world-2 started"

- Now get the logs and check it out!

```bash
oc logs -f bc/hello-world
```
> output:

```yaml
[2/3] STEP 2/2: RUN /bin/sh -ic 'echo Hello from build hook'
sh: cannot set terminal process group (-1): Inappropriate ioctl for device
sh: no job control in this shell
Hello from build hook
```
- Clean up - remove build hook 

```bash
oc set build-hook bc/hello-world \
  --post-commit \
  --remove
```

### 🔬 Hands-on Lab: 
For builds, you will make a small tweak to an application, push it to GitLab, and then run it on your OpenShift instance.

- Create a new repository under your GitLab account (https://docs.gitlab.com/ee/gitlab-basics/create-project.html)
- Add the new repository as a remote for the labs project (https://git-scm.com/book/en/v2/Git-Basics-Working-with-Remotes)
- Create a new branch in the labs repository named `builds-lab`
- Modify the Dockerfile in the `hello-world-go` directory to change the `MESSAGE` environment variable
- Commit the change and push it to GitLab
- Create a BuildConfig for this updated `hello-world-go` directory in the `builds-lab` branch
- Add a build hook that prints the value of the `MESSAGE` environment variable
- Deploy the application based on the resulting ImageStream
---

### Checklist 📋: 
- oc logs `<your BuildConfig>` contains the build hook output with your `MESSAGE` value from step 5
- `curl <your app>` shows you the `MESSAGE` value from step 5
- You can start the build using the webhook manually

---
### Quiz
> Q1: What is the command to create a new BuildConfig for a Git URL?
- [ ] `oc start-build <GIT URL>`
- [ ] `oc new-build <GIT URL>`
- [ ] `oc import-image <GIT URL>`
- [ ] `oc new-app build <GIT URL>`

<details>
  <summary> Answer </summary>

  `oc new-build <GIT URL>`

</details>

> Q2: What is the option used for build commands to build from a subdirectory of a Git project?
- [ ] `--subdirectory`
- [ ] `context-directory`
- [ ] `subdir`
- [ ] `context-dir`

<details>
  <summary> Answer </summary>

  `context-dir`

</details>

> Q3: What is the option used to build based on a branch in a Git repository?

- [ ] Add #branch-name to the end of the URL
- [ ] Add =branch-name to the end of the URL
- [ ] Add @branch-name to the end of the URL
- [ ] Add -b branch-name to the end of the URL

<details>
  <summary> Answer </summary>

  Add #branch-name to the end of the URL

</details>

> Q4: What can you use to automatically trigger BuildConfigs to run a new build in OpenShift when events happen in a Git repository like GitLab or GitHub?

- [ ] Set cron to execute `oc new-build` every few minutes
- [ ] Webhooks
- [ ] DeploymentConfig Triggers
- [ ] There is noway we can do that!

<details>
  <summary> Answer </summary>

  WebHook

</details>

### Templates
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
  oc delete all -l app=hello-world
  ```
- Lets set some parameter values:

@@ -123,154 +123,154 @@
  - Clean up:

    ```bash
    oc delete all --all
    oc delete all -l app=hello-world
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

**Hands-on Walkthroughs** 
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

### Quiz
> Q1: What is the shebang line in a Bash script?
- [ ] `#!/bin/bash`
- [ ] `#!/usr/bin/env bash`
- [ ] `#!/bin/sh`
- [ ] `#!/bin/bash-script`
<details>
  <summary> Answer </summary>

  `#!/bin/bash` or `#!/usr/bin/env bash` are both valid shebang lines.
</details>
> Q2: How do you declare a variable in Bash?
- [ ] `var_name = value`
- [ ] `var_name:value`
- [ ] `var_name=value`
- [ ] `var_name: value`
<details>
  <summary> Answer </summary>
  `var_name=value` is the correct syntax for declaring a variable in Bash.
</details>
> Q3: How do you read user input in a Bash script?
- [ ] `read input`
- [ ] `input = read`
- [ ] `input: read`
- [ ] `input = read()`
<details>
  <summary> Answer </summary>
  `read input` is the correct way to read user input in a Bash script.
</details>
> Q4: How do you create a function in Bash?
- [ ] `function_name() { commands }`
- [ ] `function function_name { commands }`
- [ ] `function_name { commands }`
- [ ] `function_name() { commands; }`
<details>
  <summary> Answer </summary>
  `function_name() { commands }` is the correct syntax for creating a function in Bash.
</details>
### Resources:
- [GNU Bash Manual](https://www.gnu.org/software/bash/manual/)
- [ShellCheck – Linter for Bash scripts](https://www.shellcheck.net/)
- [TLDP Bash Guide](https://tldp.org/LDP/Bash-Beg
inners-Guide/html/)
- [Bash Scripting Tutorial](https://ryanstutorials.net/bash-scripting-tutorial/)
- [Advanced Bash-Scripting Guide](https://tldp.org/LDP/abs/html/)
- [Bash Cheat Sheet](https://devhints.io/bash)
- [Bash Scripting for Beginners](https://www.freecodecamp.org/news/bash-scripting-for-beginners/)
- [Bash Scripting Tutorial](https://linuxconfig.org/bash-scripting-tutorial-for
-beginners)
- [Bash Scripting Best Practices](https://www.digitalocean.com/community/tutorials
/best-practices-for-bash-scripts)
- [Bash Scripting Guide](https://tldp.org/LDP/Bash-Beginners-Guide/html/)
- [Bash Scripting Cheat Sheet](https://devhints.io/bash)
- [Bash Scripting Examples](https://www.shellscript.sh/)
- [Bash Scripting Tips](https://www.cyberciti.biz/tips/bash scripting-tips.html)
- [Bash Scripting for DevOps](https://www.udemy.com/course/bash-scripting-for-devops/)
- [Bash Scripting for System Administrators](https://www.udemy.com/course/bash-scripting-for-system-administrators/)
- [Bash Scripting for Data Scientists](https://www.udemy.com/course/bash-scripting-for-data-scientists/)
- [Bash Scripting for Web Developers](https://www.udemy.com/course/bash-scripting-for-web-developers/)
- [Bash Scripting for Network Engineers](https://www.udemy.com/course/bash-scripting-for-network-engineers/)
- [Bash Scripting for Cloud Engineers](https://www.udemy.com/course/bash-scripting-for-cloud-engineers/)
- [Bash Scripting for Security Professionals](https://www.udemy.com/course/bash-scripting-for-security-professionals/)