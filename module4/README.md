[![Static Badge](https://img.shields.io/badge/Agenda-green?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/other/Agenda.md) [![Static Badge](https://img.shields.io/badge/CheatSheet-purple?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/other/CheatSheet.md) [![Static Badge](https://img.shields.io/badge/OCP-CLI-red?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/blob/main/other/ocpcli-cheatsheet.md)   [![Static Badge](https://img.shields.io/badge/Labs-maroon?style=flat&logoSize=auto)
](https://github.com/ocp-workshop-wf/bootcamp/tree/main/labs-repo)  [![Static Badge](https://img.shields.io/badge/RedHat-OpenShift-maroon?style=flat&logo=Redhat&logoSize=auto)
](https://docs.redhat.com/en/documentation/openshift_container_platform/4.19)   [![Static Badge](https://img.shields.io/badge/Kubernetes-black?style=flat&logo=Kubernetes&logoSize=auto)
](https://kubernetes.io/docs/home/)
## 🔹 Module 4: Application Deployment and Management 

## Table of Contents 
- [4.0 - OpenShift ConfigMaps](#40-openshift-configmaps) | [Lab](#-hands-on-lab-configmaps) | [Quiz](#quiz-configmaps)

- [4.1 - Secrets](#41-secrets) | [Lab](#-hands-on-lab-secrets) | [Quiz](#quiz-secrets)

- [4.2 - Images and Image Streams](#42-images-and-image-streams) | [Lab](#-hands-on-lab-images) | [Quiz](#quiz-images)

- [4.3 - Builds and BuildConfigs](#43-builds-and-buildconfigs) | [Lab](#-hands-on-lab-build-and-buildconfig) | [Quiz](#quiz-build---buildconfig)

- [4.4 - Source to Image](#44-deployment-strategies) | [Lab](#-hands-on-lab-deployment-strategies) | [Quiz](#quiz-deployment-strategies)


### 4.0 OpenShift ConfigMaps

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module3#-module-3-core-openshift-resources" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

- **Configmaps:** are a very useful resource type that OpenShift borrows from Kubernetes. ConfigMaps hold configuration data for pods to consume. This data is held in the ConfigMap separately from your running pod in OpenShift. Holding data for pods to consume is much less active job than some other types of resources in Kubernetes like `pods`, `deploymentconfigs` and `services` so when to use `ConfigMaps`

  | Component         | Development | Production               |
  |------------------|-------------|---------------------------|
  | REST API Server  | localhost   | example-api.com           |
  | Database         | localhost   | db-host.internal.com      |

  > A common case where ConfigMaps become useful is when you deploy your application to different environments. For Local development, you may wish to run non-application dependencies such as REST service or database on your machine as well, in order to simplify the development environment to connect to this REST service or database your application will need to use values that point to these local versions you can use ConfigMaps to get that kind of flexibility in OpenShift. 

<p align="center">
  <img src="/images/cm.png" alt="OpenShift Training" style="width:400px; display:block; margin:auto;" />
</p>

  > There is one piece of vocabulary you need to know when working with ConfigMaps. To <u>CONSUME</u> a configmap just means that you use the data inside of a ConfigMap from a Pod. Once the ConfigMap exists on OpenShift, you can consume or use the ConfigMap in a pod by referring to the ConfigMap in the pod definition. The word consume makes it seem like consuming a ConfigMap should use it up somehow,but this is not the case. You can consume ConfigMaps from many different pods. A common i.e  "database name" shared between an application defination and the database pod. This can help you to centralize you configuration in <u>ONE PLACE</u> and then you can update the ConfigMap, all of your pods will use the updated configuration.

  <p align="center">
    <img src="/images/cm2.png" alt="OpenShift Training" style="width:400px; display:block; margin:auto;" />
  </p>

- OpenShift give you several tools to create ConfigMaps: 
  - Command line
  - Files
  - Entire directories.
  > You also need to talk about sensitive data any resource in your OpenShift Project can read the data in a ConfigMap. For this reason, you should not store any sensitive data in the ConfigMap. OpenShift has another resource type called the `Secret` that is used for sensitive date. Also there is one more limitation of ConfigMaps that you should keep in mind, ConfigMaps have a `1MB` storage size limit. If your data could grow above the `1MB` limit you may not want to use Configmaps.


- ConfigMap Example:

    ```yaml
    apiVersion: v1 #kubernetes resource
    data:
    MESSAGE: Hello from ConfigMap #data 
    kind: ConfigMap
    metadata:
    name: message-map
    namespace: myproject #namespace
    resourceVersion: "2827192"
    selfLink: /api/v1/namespaces/myproject/configmaps/message-map
    ```

**Hands-on Walkthroughs**  

  - Creating ConfigMaps:

  ```bash
  oc create configmap message-map --from-literal MESSAGE="Hello from ConfigMap"
  # type - name of resource - to give some data -(Varaible) = data
  ```
  > output: "configmap/message-map created"

  ```bash
  # listing configmaps
  oc get cm
  ```

  > output:
    ```bash
    message-map         1      42s # 1 means that we have only 1 key value pair inside the configmap
    ```
  - Lets check the ConfigMap yaml from UI and CLI
  ```bash
  oc get -o yaml cm/message-map
  ```
  > output:
    ```yaml
    apiVersion: v1
    data: # instead of spec field we got data field here
      MESSAGE: Hello from ConfigMap # varaible we configured earlier 
    kind: ConfigMap
    metadata:
      creationTimestamp: "2025-07-18T01:09:39Z"
      name: message-map # configmap name 
      namespace: <your-namespace> # namespace/Project
      resourceVersion: "3298865818"
      uid: 7c32526a-8837-48ba-ab36-bade0095b35b
    ```


  - Consuming ConfigMaps: We'll test it out by making sure that the Hello World resonse changes after we consume the configmap 1st lets deploy our application.

    ```bash
    oc new-app quay.io/practicalopenshift/hello-world 
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

  - Consuming a ConfigMap to the application from the  **command line**

    ```bash
    oc set env deployment/hello-world --from cm/message-map 
    # ENV is just like in a dockerfile it sets an env variable based on the keys and pairs in the ConfigMap.
    ```

    > output: `deployment.apps.openshift.io/hello-world updated`

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
          - env: # the ENV we set 
            - name: MESSAGE # consumed from Configmap
              valueFrom:
                configMapKeyRef: # that is how it consumed from configmap
                  key: MESSAGE
                  name: message-map #configmap name
                  ........
    ```

  - Create ConfigMaps from **Files**:

    ```bash
    echo "Hello from ConfigMap file" > MESSAGE.txt
    # creating a simple txt file 
    ```

    ```bash
    cat MESSAGE.txt
    # just to verify
    ```

    > output:"Hello from ConfigMap file"

    ```bash
    oc create configmap file-map --from-file=MESSAGE.txt
    # create based on the txt file.
    ```

    > output: "configmap/file-map created"

    ```bash
    oc get -o yaml cm/file-map
    ```

    ```yaml
    apiVersion: v1
    data:
      MESSAGE.txt: | # confirms it was taken from the file
        Hello from ConfigMap file
    kind: ConfigMap
    metadata:
    .........
    ```

    > output: "data.MESSAGE.txt: this is the wrong syntax as it doesn't match the key in the Hello-world application to match it we need to run the following command"

    ```bash
    oc create configmap file-map-2 --from-file=MESSAGE=MESSAGE.txt
    # this time we specified the env variable = the file name
    ```

    > output: "configmap/file-map created"

    ```bash
    oc get -o yaml cm/file-map
    ```

    ```yaml
    apiVersion: v1
    data:
      MESSAGE: | # now its fixed 
        Hello from ConfigMap file
    kind: ConfigMap
    metadata:
    ```

    > output: Now as you see the data.MESSAGE: follows the same pattern for the Hello-world application.

    ```bash
    oc set env deployment/hello-world --from cm/file-map-2
    # lets update the deployment to point at this configmap-2 the correct one 
    ```

    > output: "deployment.apps.openshift.io/hello-world updated"

    ```bash
    curl < URL from oc status>
    ```

    > output: "Hello from ConfigMap file."


  - Create ConfigMaps from **Directories**:

    ```bash
    cd ./labs
    ```

    ```bash
    oc create configmap pods-example --from-file=pods
    # same but ending up with the name of the directory
    ```

    > output: "configmap/pods-example created!"

    ```bash
    oc get -o yaml configmap/pods-example
    ```

    > output:

    ```yaml
    apiVersion: v1
    data:
      pod.yaml: | # pod 1
        apiVersion: v1
        kind: Pod
        metadata:
          name: hello-world-pod
          labels:
            app: hello-world-pod
        spec:
          containers:
          - env: # seeing this here
            - name: MESSAGE # env varaible 
              value: Hi! I'm an environment variable
            image: quay.io/practicalopenshift/hello-world
            imagePullPolicy: Always
            name: hello-world-override
            resources: {}
      pod2.yaml: | # pod 2 
        apiVersion: v1
        kind: Pod
        metadata:
          name: hello-world-pod-2
          labels:
            app: hello-world-pod-2
        spec:
          containers:
          - env:
            - name: MESSAGE # message from pod 2 
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

### 🔬 Hands-on Lab (ConfigMap)

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
- Change the message that the `Deployment` uses to the ConfigMap value using the `oc set env` command
- Expose a route for your application.

---

### Checklist 📋

- Output from `oc get cm` contains your new ConfigMap

- Output from `oc get -o yaml dc/hello-world` contains the string `configMapKeyRef`

- When you run `curl <your route>` you get the value you put in the ConfigMap

---

### Quiz (ConfigMap)
>
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

  **How Are Builds Triggered?**

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

### 4.4 Source-to-image (S2I)

<p align="right">
  <a href="https://github.com/ocp-workshop-wf/bootcamp/tree/main/module5#-module-5-advanced-deployment-options" target="_blank">
    <img src="/images/top.png" alt="OpenShift Training" style="width:25px;" />
  </a>
</p>

 It is a toolkit and workflow that automates the process of building container images from source code. It takes a builder image (containing necessary build tools and dependencies) and source code, then combines them to create a runnable application image. S2I is used to create reproducible container images, making it easier for developers to deploy and manage applications in various environments, including OpenShift. 

<p align="center">
<img src="/images/s2i-concept.webp" alt="OpenShift Training" style="width:400px; align="center"/>
</p>

Kaniko (Buildpacks / Konica Buildpack Implementation)
Kaniko refers to a Cloud Native Buildpacks-based toolset, typically used in platforms like Heroku and Paketo. It:
  - Detects your application type automatically
  - Selects appropriate buildpacks
  - Produces OCI-compliant images without needing a Dockerfile
  - Emphasizes build phase separation (detect, build, and export)

| Feature                        | S2I (Source-to-Image)                   | Kaniko (Buildpacks)                         |
|-------------------------------|-----------------------------------------|---------------------------------------------|
| **Build Method**              | Uses builder image + app source         | Uses buildpacks (detect/build/export flow)  |
| **Dockerfile Needed**         | ❌ Not required                          | ❌ Not required                              |
| **Language Detection**        | Manual or via image tag                 | ✅ Automatic                                 |
| **Layer Caching**             | Basic layer reuse                       | ✅ Advanced layer caching & reuse            |
| **Custom Logic**              | Allows `assemble`, `run`, `save-artifacts` scripts | Builtpack hooks for full customization   |
| **OpenShift Native**          | ✅ Fully integrated                      | ⚠️ Needs Paketo/Stack-based integration     |
| **Output Image Format**       | Docker/OCI-compatible                   | OCI-compatible                              |
| **Use Case Fit**              | Great for OpenShift CI/CD pipelines     | Great for cloud-native CI systems like Tekton, GitHub Actions |


  > Both S2I and Kaniko offer Dockerfile-less image creation. While S2I is tailored for OpenShift with native support and scripting hooks, Kaniko leverages Cloud Native Buildpacks for broader platform compatibility and better caching. 

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