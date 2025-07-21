## Module 4: Application Deployment and Management 

### 4.1 OpenShift ConfigMaps

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

### 🔬 Hands-on Lab: 
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

### Checklist 📋: 
- Output from `oc get cm` contains your new ConfigMap

- Output from `oc get -o yaml dc/hello-world` contains the string `configMapKeyRef`

- When you run `curl <your route>` you get the value you put in the ConfigMap

---
### Quiz
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

### 4.2 Secrets

- Secrets: a Kubernetes resource designed to hold sensitive information like passwords, API keys, certificates, and other credentials.
  - Basic Auth
  - SSH Key 
  - TLS Auth


**Hands-on Walkthroughs**  
- Create an Opaque secret

```bash
oc create secret generic message-secret --from-literal MESSAGE="secret message"
```
> output: secret/message-secret created

- Lets find out where it is!

```bash
oc get secret
```
> output: "message-secret        Opaque            10s"

- Lets open the yaml for that secret.

```bash
oc get -o yaml secrete/message-secret
```
> output:

```yaml
apiVersion: v1
data:
  MESSAGE: c2VjcmV0IG1lc3NhZ2U=
kind: Secret
metadata:
  creationTimestamp: "2025-07-19T03:42:04Z"
  name: message-secret
  namespace: raafat-dev
  resourceVersion: "3334020540"
  uid: d1e28fce-ab5d-45a3-8e9f-526f83d20135
type: Opaque
```
- This message is base64.

- How to use a Secret as environment variables

```bash
oc new-app quay.io/practicalopenshift/hello-world --as-deployment-config
```
```bash
oc expose svc/hello-world /
oc status
```
```bash
curl <url from oc status> 
```

> output: "Welcome! You can change this message by editing the MESSAGE environment variable."

```bash
oc set env dc/hello-world --from secret/message-secret
```
> output: "deploymentconfig.apps.openshift.io/hello-world updated"

```bash
curl <URL from oc status>
```
> output: "secret message"

```bash
oc get -o yaml dc/hello-world
```
```yaml
  .....
 - env:
        - name: MESSAGE
          valueFrom:
            secretKeyRef:
              key: MESSAGE
              name: message-secret
        image: quay.io/practicalopenshift/hello-world@sha256:2311b7a279608de9547454d1548e2de7e37e981b6f84173f2f452854d81d1b7e
        imagePullPolicy: Always
        name: hello-world
      .......
```
---

### 🔬 Hands-on Lab: 
For secrets, you will update a secret value by modifying its definition. Start with the following secret:

```yaml
apiVersion: v1
data:
  MESSAGE: YmFzZTY0
kind: Secret
metadata:
  name: lab-secret
type: Opaque
```
- Create a new file and put the above YAML inside
- Update the value of MESSAGE to "Hello from Lab Secret"
- Create a secret based on this YAML using `oc create -f`
- Deploy and expose the Hello World Application
- Set the MESSAGE environment varaible using the secret.
---

### Checklist 📋: 
- Output from `oc get secret` contains your new Secret

- Output from `oc get -o yaml dc/hello-world` contains the string "secretKeyRef"

- When you run `curl <your route>` you get the value you put in the Secret

---
### Quiz
> Q1: You can use all the same inputs for ConfigMaps in order to create Secrets.
- [ ] Fales
- [ ] True

<details>
  <summary> Answer </summary>

  True

</details>

> Q2: The data stored in a secret resource is securely encrypted. Even if someone could read the secret, they wouldn't get access to the information inside.
- [ ] Fales
- [ ] True

<details>
  <summary> Answer </summary>

  False "Base64 is not a secure encryption method"

</details>

> Q3: What format do secrets use for their values?
- [ ] SHA256
- [ ] RSA
- [ ] WPA2
- [ ] Base64

<details>
  <summary> Answer </summary>

 Base64 

</details>

**Resource:**  
- [Using ConfigMaps & Secrets in OpenShift](https://www.youtube.com/watch?v=AnvOMRFwimM)

---

### 4.3 Images and Image Streams 
An image is a self-contained package containing everything needed to run an application, including the code, runtime, system tools, libraries, and settings. An image stream, on the other hand, is a way to manage and track different versions of an image within OpenShift, acting as a virtual image repository. 

![Image & Image Steams architecture](/images/image-imagestream.png)

**Hands-on Walkthroughs**  
- How to create an ImageStream

```bash
oc get is
```
> output: 

| Name  | IMAGE REPOSITORY |  TAGS |UPDATED|
| -------- | ------- | ------- | -----| 
| Hello-World  | `default-route-openshift-image-registry..` | latest | 27 hours ago|          
---
```bash
oc delete is/hello-world
```
> output: "imagestream.image.openshift.io "hello-world" deleted"

- Lets import the image now.

```bash
oc import-image --confirm quay.io/practicalopenshift/hello-world
```
> output: "imagestream.image.openshift.io/hello-world imported" In addition to image description such as: 

```bash
Name:                   hello-world
Namespace:              raafat-dev
Created:                Less than a second ago
Labels:                 <none>
Annotations:            openshift.io/image.dockerRepositoryCheck=2025-07-19T05:09:16Z
Image Repository:       default-route-openshift-image-registry.apps.rm3.7wse.p1.openshiftapps.com/raafat-dev/hello-world
Image Lookup:           local=false
Unique Images:          1
Tags:                   1

latest
  tagged from quay.io/practicalopenshift/hello-world
```
- To get more information
```bash
oc get istag
```
- Once you have the name run the command to describe it:

```md
oc describe istag/<name>:latest
```
- What can you even do with an imagestream after importing it to OpenShift?
> you can easier to deploy a new app using that image

```bash
oc new-app myproject/hello-world --as-deployment-config
```
- How to add more ImageStreamTags
```bash
oc tag quay.io/image-name:tag image-name:tag
```
- Syntax (oc tag `<orignial>` `<destination>`)
```bash
oc tag quay.io/practicalopenshift/hello-world:update-message hello-world:update-message
```
> output:"Tag hello-world:update-message set to quay.io/practicalopenshift/hello-world:update-message"

```bash
oc get is
```
> output: you should be able to see `update-message` on the tag

```bash
oc get istag
```

> output: you should see both images with the tag `latest` and `update-message`

- Lets learn how to push a private image on quay.io
- locate credentials.env file and update it with your Quay.io credentials.

```bash
source credentials.env
```
```bash
cd ./labs-repo/hello-world-go-private
```
```bash
cat Dockerfile
```
- Now we need to push the image, but we need to be very careful about the syntax for remote tags

|HOST|Repository|Image Name (:tag)|
| ----| ---- | ----|
|quay.io|/`$REGISTRY_USERNAME`$|/private-repo
---
```bash
docker build -t quay.io/$REGISTRY_USERNAME/private-repo .
```
>output "[+] Building 16.8s (5/5) FINISHED"

```bash 
docker login quay.io
```
- Push 
```docker push quay.io/$REGISTRY_USERNAME/private-repo
```
> output: Access your quay.io and look for private-repo with a red lock on it which means its private, also on your terminal you should see something like this: 

```bash
The push refers to repository [quay.io/`USERNAME`/private-repo]
7b2225181d6b: Pushed 
08684ee472f3: Pushed 
c732a2540651: Pushed 
552574c585c3: Pushed 
c8dae7ec6990: Pushed 
aad63a933944: Pushed 
0dd0829302c5: Pushed 
aaed0f9cbe2b: Pushed 
latest: digest: sha256:813277ad25de25d77aee81727dc6a27751434294682b922c8ae97966a1ac1faf size: 856
```


- How to run this private image to OpenShift?

```bash
oc new-app quay.io/$REGISTRY_USERNAME/private-repo --as-deployment-config
```
> output: 
```bash
.......
        deploymentconfig.apps.openshift.io "private-repo" created
    service "private-repo" created
......
```
- In case you get an Authentication error, you need to run the following command
```bash
oc create secret docker-registry \ demo-image-pull-secret \
--docker-server=$REGISTRY_HOST \
--docker-username=$REGISTRY_USERNAME \
--docker-password=$REGISTRY_PASSWORD \ 
--docker-email=$REGISTRY_EMAIL
```
> output: secret created

- creating a secret is not enough you need to link that secret to let openshift use it to pull the image.

```bash
oc secrets link default demo-image-pull-secret --for=pull
```
- dafault here is the service account that openshift will use, using pull because we want that secret to pull an image.
```bash
oc describe serviceaccount/default
```

> output: you should see something like this

```bash
.......
Image pull secrets:  default-dockercfg-stnvn
                     demo-image-pull-secret
.....
```
- lets confirm everything went well
```bash
oc new-app quay.io/$REGISTRY_USERNAME/private-repo --as-deployment-config
```
```bash
oc expose service/private-repo
```
```bash
oc status
```
```bash
curl <URL from oc status>
```
> output: "Hello from private image registry"

---

### 🔬 Hands-on Lab: 
For images, you'll import your own private image and tag into OpenShift.

- Create a new project called `images-lab`
- Create an image named `images-lab` based on `quay.io/practicalopenshift/hello-world`
- Push this image to your Quay.io accout. Keep it private
- Create another image with the tag `images-lab:private-tag`.
- Push this tag to Quay.io
- Create an ImageStream for your private images-lab image
- Import the tag into OpenShift as an ImageStreamTag
- Deploy an application based on the images-lab:private-tag image
---

### Checklist 📋: 
- `oc get is` returns a single ImageStream
- `oc get istag`returns 2 tags for `images-lab`
- `oc get secret` shows your authentication secret
- `curl <your route>` gives you the message you supplied in step 4.

---
### Quiz
> Q1: OpenShift has its own image registry built in.
- [ ] Fales
- [ ] True

<details>
  <summary> Answer </summary>

  True

</details>

> Q2: ImageStreams correspond to which Docker concept?
- [ ] Single Image
- [ ] Image name such as Hello-world
- [ ] Image tags such as hello-world
- [ ] Running containers

<details>
  <summary> Answer </summary>

  Image name such as Hello-world

</details>

> Q3: Which service account can you add docker authentication to in order to pull from a private repository?
- [ ] Puller
- [ ] default
- [ ] Authentication
- [ ] secrets

<details>
  <summary> Answer </summary>

  default

</details>

> Q4: What command can you use to add an ImageStreamTag for an image that you already have an ImageStream for?
- [ ] `oc add tag`
- [ ] `oc tag`
- [ ] `oc copy tag`
- [ ] `oc tag:tag`

<details>
  <summary> Answer </summary>

  `oc tag`

</details>

---

### 4.4 Builds and BuildConfigs
Builds represent the process of transforming input (like source code) into a runnable image, while BuildConfigs define the entire build process for an application. A BuildConfig acts as a template, specifying how to build an image, including the source code, build strategy, and output location. 
- A build is a specific instance of a build process triggered by a BuildConfig. 
- A BuildConfig is a resource in OpenShift that defines the build process for an application. 

**Hands-on Walkthroughs**  

- lets create a new build 

```bash
oc new-build <Git URL>
```

> output: "buildconfig.build.openshift.io "hello-world" created"

- Lets see whats inside the buildconfig

```bash
oc get -o yaml buildconfig/hello-world
```
> output: 

```yaml
apiVersion: build.openshift.io/v1
kind: BuildConfig
metadata:
  annotations:
    openshift.io/generated-by: OpenShiftNewBuild
  creationTimestamp: "2025-07-19T22:53:53Z"
  generation: 2
  labels:
    build: hello-world
  name: hello-world
  namespace: raafat-dev
  resourceVersion: "3357388343"
  uid: 0d5acd9e-e53c-4d3e-a7da-50aaeda9a832
spec:
  failedBuildsHistoryLimit: 5
  nodeSelector: null
  output:
    to:
      kind: ImageStreamTag
      name: hello-world:latest
  postCommit: {}
  resources: {}
  runPolicy: Serial
  source:
    git:
      uri: https://gitlab.com/therayy1/hello-world.git
    type: Git
  strategy:
    dockerStrategy:
      from:
        kind: ImageStreamTag
        name: golang:1.17
    type: Docker
  successfulBuildsHistoryLimit: 5
  triggers:
  - github:
      secret: uEN_9w-Ftp5rVSR3Q_35
    type: GitHub
  - generic:
      secret: 8R972EmKLVDCPF5N6y7O
    type: Generic
  - type: ConfigChange
  - imageChange: {}
    type: ImageChange
```
- Lets look at the builds

```bash
oc get build
```
> output: 

| NAME | TYPE | FROM | STATUS | STARTED | DURATION|
| ---- | ---- | ---- | ------ | ------- | ------- |
|hello-world-1| Docker | Git@9e4d905 | Complete | 39 minutes ago | 56s|'

- Checking build logs

```bash
oc get buildconfig
```
```bash
oc logs -f buildconfig/hello-world
```

> output: Docker steps and pushs the image!

- How to start a build "use 2 terminals"

on Terminal 1
```oc get pods -w
```
on Terminal 2
```bash
oc start-build bc/hello-world
```
> output: "build.build.openshift.io/hello-world-2 started" also on Terminal 1 you will see many events happening.

```bash
oc describe is/hello-world
```

> output: "you will find there is 2 changes happened by the time you started the build, this example we didn't push anything new - in real life example most common this would create a build based on the newest source code allowing you to test your changes."

- Cancelling Builds:

```bash
oc cancel-build bc/hello-world
```

> output: "build.build.openshift cancelled"

---

### 4.5 Webhooks
Is a method for one application to automatically send real-time data to another application when a specific event occurs. It's essentially an automated messaging system that allows applications to communicate with each other without needing to constantly "poll" or check for updates. And this is one of the key features that enables CICD!

| OpenShift | Git Repo |
| ---------- | ------ |
| Exposes HTTPS endpoint that starts when called | Calls endpoint when developers push code|

> The way the integration is set up is that the hook is exposed on the downstream system, OpenShift in our case, and the hook is configured on the upstream system. Your Git Repository is configured to call the endpoint that's exposed on OpenShift. To mitigate the risk of having a public endpoint on your OpenShift cluster, OpenShift requires Webhook clients to pass a special token string along with a request. This token value is automatically generated for BuildConfigs created with the `oc new-app` and the `oc new-build` command line tools. Once you have set up the Webhook correctly on the GitLab side, anytime a developer pushes an update GitLab will make an HTTPS request to the Webhook URL and pass along the token. When OpenShift gets the request it will start a new Build for the linked BuildConfig. This way, you can set up an automated pipeline allowing you to push automatically to your OpenShift cluster simply by pushing to a Git Repository.

![WebHook](/images/webhook.png)


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