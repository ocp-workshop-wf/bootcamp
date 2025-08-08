### 4.5 Triggers

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module4#-module-4-application-deployment-and-management" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

Lets learn about basic automation for deployments, configuring the deployment process itself, and how to add health checks so that OpenShift can restart your pods if necessary.

- Deployment Triggers:
  - Image Change Trigger: This trigger watches an image stream in OpenShift. When the image stream or its image stream tag has a new image, if the deployment config is configured with an image change trigger,then the deployment config will roll out the new version of your application image automatically.
<p align="center">
<img src="/images/triggers.png" alt="Triggers" style="width:500px; align="center"/>
</p>
- ConfigChange Triggers:With the ConfigChange Trigger,any change to the POD template will trigger a new rollout. If you add a volume or change the image stream tag in a pod spec, then the ConfigChange Trigger will cause your deployment to roll out another version.
  
<p align="center">
<img src="/images/configchangetrigger.png" alt="ConfigChange Triggers" style="width:500px; align="center"/>
</p>

--- 

**Hands-on Walkthroughs** 
  
- How the ConfigChange Trigger works?
    - In this exmaple you need 2 windows terminals.

    ```bash
    oc new-app quay.io/practicalopenshift/hello-world 
    ```
    ```bash
    oc describe deployment/hello-world
    ```
    > output: look for `Triggers:       Config, Image(hello-world@latest, auto=true)` 
  -  we're going to modify the pod template, which will trigger a redeploy due to the ConfigChange trigger configured for our deployment config.
    ```bash
    # on terminal 1 
    oc get pods --watch
    ```
    ```bash
   # on terminal 2
    oc set volume deployment/hello-world \
  --add \
  --type emptyDir \
  --mount-path /config-change-demo
    ```

- How to add and Remove DeploymentConfigs triggers:
  -  You learned that both the config change trigger and the image change trigger are configured by default for applications created with OC new app.
  - Lets learn how to use the `oc set triggers` command to modify the triggers with no arguments, `oc set triggers will simply print the trigger associated with the deployment config.
    
    ```bash
    oc set triggers dc/hello-world
    ```
    > output: should contain `type` config, image to `true`

    - To remove the triggers 
    ```bash
    oc set triggers dc/hello-world \ --remove --from-config
    ```
    > output: "deploymentconfig.apps.openshift.io/hello-world triggers updated"

    ```bash
    oc set triggers dc/hello-world
    # to list the triggers
    ```
    > output: You should see `TYPE` config `VALUE` false.

    ```bash
    # to re-add the config change trigger you don't need to type `--add` once you set the trigger it Automatically adds it back
    oc set triggers dc/hello-world --from-config
    ```
    ```bash
    oc set triggers deployment/hello-world
    ```
    > output: both `VALUES` are true.

- Lets replicate that with the image change trigger:

    ```bash
    oc set triggers deployment/hello-world \
  --remove\
  --from-image hello-world:latest
    ```
    ```bash
    oc set triggers deployment/hello-world
    ```
    > output: on the list you will find only the config trigger and you will no loger find the image trigger

    - So lets reverse that!
  
    ```bash
    oc set triggers dc/hello-world --from-image hello-world:latest -c hello-world
    ```
    ```bash
    oc set triggers dc/hello-world
    ```
    > output: you should find both `TYPE` in with the `VALUE` of true.


### Quiz (Triggers)
> Q1: What command will disable dc triggers?
- [ ] `oc set triggers <dc name --remove>`
- [ ] `oc set triggers <dc name> --none`
- [ ] `oc remove triggers <dc name>`
- [ ] `oc set triggers`


<details>
  <summary> Answer </summary>

    
  `oc set triggers <dc name --remove>`

</details>

> Q2: How many deployment triggers are enabled by default for oc new-app projects?
- [ ] 1
- [ ] 2
- [ ] 3
- [ ] 4


<details>
  <summary> Answer </summary>

   2
  

</details>

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