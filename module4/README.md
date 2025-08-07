[![Static Badge](https://img.shields.io/badge/Agenda-green?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/other/Agenda.md) [![Static Badge](https://img.shields.io/badge/CheatSheet-purple?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/other/CheatSheet.md) [![Static Badge](https://img.shields.io/badge/OCP-CLI-red?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/other/ocpcli-cheatsheet.md)   [![Static Badge](https://img.shields.io/badge/Labs-maroon?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/tree/main/labs-repo)  [![Static Badge](https://img.shields.io/badge/RedHat-OpenShift-maroon?style=flat&logo=Redhat&logoSize=auto)
](https://docs.redhat.com/en/documentation/openshift_container_platform/4.19)   [![Static Badge](https://img.shields.io/badge/Kubernetes-black?style=flat&logo=Kubernetes&logoSize=auto)
](https://kubernetes.io/docs/home/)
## 🔹 Module 4: Application Deployment and Management 

## Table of Contents 
- [4.1 - Secrets](#41-secrets) | [Lab](#-hands-on-lab-secrets) | [Quiz](#quiz-secrets)

- [4.2 - Images and Image Streams](#42-images-and-image-streams) | [Lab](#-hands-on-lab-images) | [Quiz](#quiz-images)

- [4.3 - Builds and BuildConfigs](#43-builds-and-buildconfigs) | [Lab](#-hands-on-lab-build-and-buildconfig) | [Quiz](#quiz-build---buildconfig)

- [4.4 - Deployment Strategies](#44-deployment-strategies) | [Lab](#-hands-on-lab-deployment-strategies) | [Quiz](#quiz-deployment-strategies)

- [4.5 - Triggers](#45-triggers) | [Quiz](#quiz-triggers)

### 4.1 Secrets
<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module4#-module-4-application-deployment-and-management" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px; float:right; margin-left:10px;" />
  </a>
</p>

- Secrets: a Kubernetes resource designed to hold sensitive information like passwords, API keys, certificates, and other credentials.
  - Basic Auth
  - SSH Key 
  - TLS Auth 
- All the above are commonly used to authenticate to private secrets outside OpenShift, such as GitHub repository. You will need to pick the right one depending on the authentication used on the private server.


**Opaque or generic secrets** have no restrictions on their structure, meaning that they can have any key value pairs inside. The word Opaque refers to the fact that Opaque secrets do not have any sort of special integrations into other parts of OpenShift. OpenShift won't know what to do with these secrets and it's up to your applications to use them correctly. This is the type of secret that you would use for authenticating to a database or a custom application. A common **<u>non-Opaque</u>** secret type is the Kubernetes `service account token` secret type.

<p align="center">
<img src="/images/serviceaccount.png" alt="Image & Image Streams Arch" style="width:400px; align="center"/>
</p>

 **Service Accounts** are a mechanism that Kubernetes and OpenShift use in order to allow pods to authenticate against internal OpenShift or Kubernetes APIs. You can use service account tokens to allow pods to access the same API that you use with the `oc-cli` tool. 

> 💡 **NOTE** 
> OpenShift security model restricts all permissions by default, so you will need to grant the service accounts permissions before they will be able to do anything interesting in the OpenShift API.

**Hands-on Walkthroughs**  

- Create an Opaque secret.
  - It is very similar to the `configmap` 
  <p align="center">
  <img src="/images/cmvss.png" alt="Image & Image Streams Arch" style="width:400px; align="center"/>
  </p>

  > 💡 **NOTE** 
  > While configmaps take the name as the argument directly after `oc create configmap`, secrets require another argument which will be the type of secret you're creating. In this case, we're creating an opaque or generic secret, these types of secrets, use the generic argument right after `oc create secret` command, after generic you'll put in the name of the secret, `oc create secret` takes the same kinds of options for initializin the secret as `oc create configmap`.

  ```bash
  oc create secret generic message-secret --from-literal MESSAGE="secret message"
  # for this we are using Literal but we can also use file. 
  ```
  > output: "secret/message-secret created"

  - Lets list the secrets!

  ```bash
  oc get secret
  ```
  > output: "message-secret        Opaque            10s"
  Unlike other types of OpenShift resources that we've worked with so far OpenShift Projects come with several secrets already initialized including `dockercfg` for the internal docker registry and service account tokens for `builder-token` default and deployer service accounts. 

  - Lets open the yaml for that secret.

  ```bash
  oc get -o yaml secrete/message-secret
  ```
  > output:

  <p align="center">
  <img src="/images/secretyaml.png" alt="Image & Image Streams Arch" style="width:400px; align="center"/>
  </p>


  ```yaml
  apiVersion: v1
  data: 
    MESSAGE: c2VjcmV0IG1lc3NhZ2U= #base64 incoded version NOT ENCRYPTION 
  kind: Secret
  metadata:
    name: message-secret
    namespace: raafat-dev
  type: Opaque
  ```

- How to use a Secret as environment variables

  ```bash
  # If you don't have the app already deployed.
  oc new-app quay.io/practicalopenshift/hello-world 
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
  # Supply a secret 
  oc set env deployment/hello-world --from secret/message-secret
  ```
  > output: "deploymentconfig.apps.openshift.io/hello-world updated"

  ```bash
  curl <URL from oc status>
  ```
  > output: "secret message"

  ```bash
  oc get -o yaml deployment/hello-world
  ```
  ```yaml
    .....
  - env:
          - name: MESSAGE
            valueFrom: # refrence
              secretKeyRef:
                key: MESSAGE # env
                name: message-secret # secret value
          image: quay.io/practicalopenshift/hello-world@sha256:2311b7a279608de9547454d1548e2de7e37e981b6f84173f2f452854d81d1b7e
          imagePullPolicy: Always
          name: hello-world
        .......
  ```
---

### 🔬 Hands-on Lab (Secrets): 
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

### Checklist 📋 (Secrets): 
- Output from `oc get secret` contains your new Secret

- Output from `oc get -o yaml deployment/hello-world` contains the string "secretKeyRef"

- When you run `curl <your route>` you get the value you put in the Secret

---
### Quiz (Secrets)
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

### 4.2 Images and Image Streams 

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module4#-module-4-application-deployment-and-management" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

OpenShift has a few resource types dedicated to managing images, so that OpenShift can support advanced image-based workflows. The Two Primary resource types that you'll use are the:
  - ImageStream 
  - ImageStreamTag 
  > Examples: `Hello-world` or `Golang`

All images such as `golang` has a variety of tags for different distributions and versions. 

**In OpenShift** the imageStream would be `golang` and all of these tags would be `ImageStramTag` as if in DockerHub. 

The OpenShift ImageStream and ImageStreamTag resources may sound like they do something new. However, these are really just the same image concepts that you already know. Just Inside of OpenShift, ImageStreams provide similar functionality to the build-it docker registry that runs on your local machine once you have docker installed.

**Image Storage inside of OpenShift** is useful because other resource types can watch or subscribe to ImageStreams and ImageStreamTags in order to receive notifications and take action when new image become available.

> For example: you can configure a Deployment to automatically deploy the most recent version of an ImageStreamTag.


<p align="center">
<img src="/images/Imagestreamtagdeployment.png" alt="Image & Image Streams Arch" style="width:400px; align="center"/>
</p>

**Hands-on Walkthroughs**  
- How to create an ImageStream
  ```bash
  oc delete all -l app=hello-world
  # just to make sure we got a clean env
  ```
  ```bash
  oc new-app quay.io/practicalopenshift/hello-world 
  # creating a deployment
  ```
  > output: you should see the following 
    ```bash
    --> Creating resources ...
    imagestream.image.openshift.io "hello-world" created
    deployment.apps "hello-world" created
    service "hello-world" created
    ```

  ```bash
  oc get is
  ```
  > output: Looking at the Image Repository Column it starts with a long domain name, this host refers to the internal container image registry hosted versions on OpenShift

  | Name  | IMAGE REPOSITORY |  TAGS |UPDATED|
  | -------- | ------- | ------- | -----| 
  | Hello-World  | `default-route-openshift-image-registry..` | latest | 27 hours ago|          

  ```bash
  #list all the imagestreamtag
  oc get imagestreamtag 
  # or oc get istag
  ```

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
  oc new-app myproject/hello-world 
  ```
- How to add more ImageStreamTags!
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
  > output "[+] Building 16.8s (5/5) FINISHED" If you got an error make sure your in the right directory 

  ```bash 
  docker login quay.io
  ```
  - Push 

  ```bash
  docker push quay.io/$REGISTRY_USERNAME/private-repo
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


<p align="center">
<img src="/images/dockerpush.png" alt="Image & Image Streams Arch" style="width:500px; align="center"/>
</p>



  - How to run this private image to OpenShift?

  ```bash
  docker build -t quay.io/$REGISTRY_USERNAME/private-repo .
  ```
  > output: 
  ```bash
  .......
          deploymentconfig.apps.openshift.io "private-repo" created
      service "private-repo" created
  ......
  ```
  <p align="center">
  <img src="/images/dockerbuild.png" alt="Image & Image Streams Arch" style="width:500px; align="center"/>
  </p>

  - Lets create a secret using the our docker credentials

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
    oc new-app quay.io/$REGISTRY_USERNAME/private-repo 
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

### 🔬 Hands-on Lab (Images): 
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

### Checklist 📋 (Images): 
- `oc get is` returns a single ImageStream
- `oc get istag`returns 2 tags for `images-lab`
- `oc get secret` shows your authentication secret
- `curl <your route>` gives you the message you supplied in step 4.

---
### Quiz (Images)
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

### 4.3 Builds and BuildConfigs

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module4#-module-4-application-deployment-and-management" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

Builds in OpenShift are managed by a resource type called **BuildConfig**.

The `BuildConfig`:
- Defines **how** the image should be built from the source code.
- Is similar in concept to the flags and options used with `docker build`.

## How Are Builds Triggered?

When using the `oc new-app` command with a Git repository, OpenShift automatically:
- Creates the `BuildConfig`.
- Initiates a build process.
- Sets up an image stream and deployment configuration.

**Hands-on Walkthroughs**  

- lets create a new build 

  ```bash
  #oc new-build <Git URL>
  oc new-build https://gitlab.com/therayy1/hello-world
  ```

  > output: "buildconfig.build.openshift.io "hello-world" created"

  - Lets see whats inside the buildconfig

  ```bash
  oc get -o yaml buildconfig/hello-world
  ```
  > output: 

  ```yaml
  apiVersion: build.openshift.io/v1 # that an OpenShift type
  kind: BuildConfig
  metadata:
    annotations:
      openshift.io/generated-by: OpenShiftNewBuild
    generation: 2
    labels:
      build: hello-world
    name: hello-world
    namespace: raafat-dev
  spec:
    failedBuildsHistoryLimit: 5
    nodeSelector: null
    output: # tells OpenShift, What to do once it builds an image
      to:
        kind: ImageStreamTag
        name: hello-world:latest
    postCommit: {}
    resources: {}
    runPolicy: Serial
    source: # What OpenShift should use to create your build.
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
    triggers: # Triggers are already configured.
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

  - Checking build logs to dignose probelems 

    ```bash
    oc get buildconfig
    ```

    > output: Hello-world Docker Git 1 

    ```bash
    oc logs -f buildconfig/hello-world
    ```

    > output: Build clones the repo, runs the docker steps and pushes the resulting image up to the image repo hosted in OpenShift 

    <p align="center">
    <img src="/images/buildlogs.png" alt="Image & Image Streams Arch" style="width:500px; align="center"/>
    </p>

- How to start a build "use 2 terminals" this is used to start an existing buildconfig, as it creates a single build object based on an existing buildconfig.

  on Terminal 1

  ```bash
  oc get pods -w
  ```
  on Terminal 2

  ```bash
  oc start-build bc/hello-world
  ```
  > output: "build.build.openshift.io/hello-world-2 started" also on Terminal 1 you will see many events happening. 

    <p align="center">
    <img src="/images/stat-build.png" alt="Image & Image Streams Arch" style="width:500px; align="center"/>
    </p>

  ```bash
  oc describe is/hello-world
  ```

  > output: "you will find there is 2 changes happened by the time you started the build, this example we didn't push anything new - in real life example most common this would create a build based on the newest source code allowing you to test your changes."

- Cancelling Builds:

  ```bash
  oc cancel-build bc/hello-world
  ```

  > output: "build.build.openshift cancelled"


### 🔬 Hands-on Lab (Build and BuildConfig): 
  For images, you'll import your own private image and tag into OpenShift.

  - Create a new repository under your GitLab account (https://docs.gitlab.com/ee/gitlab-basics/create-project.html)
  - Add the new repository as a remote for the labs project (https://git-scm.com/book/en/v2/Git-Basics-Working-with-Remotes)
  - Push the labs repository to this new remote
  - Create a new branch in the labs repository named builds-lab
  - Modify the Dockerfile in the `build-lab` directory to change the MESSAGE environment variable
  - Commit the change and push it to GitLab
  - Create a BuildConfig for this updated `build-lab` directory in the builds-lab branch
  - Deploy the application based on the resulting ImageStream

---

### Checklist 📋 (Build - Buildconfig): 
- `oc logs <your BuildConfig>` contains your MESSAGE value 
- `curl <your app>` shows you the MESSAGE value from step 5

---
### Quiz (Build - Buildconfig)
> Q1: What is the command to create a new BuildConfig for a Git URL? 
- [ ] `oc start-build <Git URL>`
- [ ] `oc new-build <Git URL>`
- [ ] `oc import-image <Git URL>`
- [ ] `oc new-app <Git URL>`

<details>
  <summary> Answer </summary>

  `oc new-build <Git URL>`

</details>


> Q2: What is the option used for build commands to build from a subdirectory of a Git project?

- [ ] --subdirectory
- [ ] --context-directory
- [ ] --subdir
- [ ] --context-dir

<details>
  <summary> Answer </summary>

  --context-dir

</details>

---

### 4.4 Deployment Strategies 

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

---

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module5" target="_blank">
    <img src="/images/nexticon.webp" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

<p align="left">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module3" target="_blank">
    <img src="/images/backred1.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>