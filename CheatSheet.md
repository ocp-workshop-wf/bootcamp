### Get to the right directory
```bash
cd labs
cd hello-world-go
```
---
### Build the image
```bash
docker build .
```
### Check the image storage
```bash
docker images
```
---

### Log in, log out

### Uses the pre-configured OpenShift cluster
```bash
oc login
```
### Allows you to log in to any OpenShift cluster
```bash
oc login <cluster address>
```
### Log out
```bash
oc logout
```

---

### Project Basics

### See current project

```bash
oc project
```

### Create a new project

```bash
oc new-project demo-project
```

### List all projects

```bash
oc projects
```

### Switch projects

```bash
oc project <project name>
```
---

### Get Pod Documentation

### Get built-in documentation for Pods

```bash
oc explain pod
```

### Get details on the pod's spec

```bash
oc explain pod.spec
```

### Get details on the pod's containers

```bash
oc explain pod.spec.containers
```
---

### Creating Pods from files

### Create a Pod on OpenShift based on a file

```bash
oc create -f pods/pod.yaml
```
### Show all currently running Pods

```bash
oc get pods
```
---

### Port forwarding for Pods

### Open a local port that forwards traffic to a pod

```bash
oc port-forward <pod name> <local port>:<pod port>
```

### Example of 8080 to 8080 for hello world

```bash
oc port-forward hello-world-pod 8080:8080
```
---

### Shell into Pods

### oc rsh will work with any Pod name from oc get pods

```bash
oc rsh <pod name>
```

### In the shell, check the API on port 8080

```bash
wget localhost:8080
```

### Exit the rsh session

```bash
exit
```

### Watch live updates to pods

```bash
oc get pods --watch
```

---

### Delete (stop) Pods

### Delete any OpenShift resource

```bash
oc delete <resource type> <resource name>
```

### Delete the pod for this section

```bash
oc delete pod hello-world-pod
```

### Watch live updates to pods

```bash
oc get pods --watch
```

---

### Deploying applications as DeploymentConfigs

### Deploy an existing image based on its tag

```bash
oc new-app <image tag>
```

### Deploy the Hello World image for this course

```bash
oc new-app quay.io/practicalopenshift/hello-world
```

### Deploy from Git using oc new-app

```bash
oc new-app <git repo URL>
```

### Deploy the Hello World application from Git

```bash
oc new-app https://gitlab.com/practical-openshift/hello-world.git
```

### Follow build progress (Git only)

```bash
oc logs -f bc/hello-world
```

### Set the name for the DeploymentConfig 

```bash
oc new-app <image tag> --name <desired name>
```

### Example with a name

```bash
oc new-app quay.io/practicalopenshift/hello-world --name demo-app
```
---

### Get more information about a DeploymentConfig

### Describe the DC to get its labels

```bash
oc describe dc/hello-world
```

### Get the full YAML definition

```bash
oc get -o yaml dc/hello-world
```

---
### Deleting all oc new-app resources

### Delete all application resources using labels (get them from oc describe)

```bash
oc delete all -l app=hello-world
```
---

### Starting new versions and reverting changes

### Roll out the latest version of the application

```bash
oc rollout latest dc/hello-world
```

### Roll back to the previous version of the application

```bash
oc rollback dc/hello-world
```

### Deploy an existing image based on its tag

```bash
oc new-app <image tag>
```

### For this lesson

```bash
oc new-app quay.io/practicalopenshift/hello-world
```

### Check running resources

```bash
oc status
```

### Check pods

```bash
oc get pods
```

### Standard image deployment 

```bash
oc new-app quay.io/practicalopenshift/hello-world
```

### Describe the DC to get its labels

```bash
oc describe dc/hello-world
```

### Delete all application resources using labels

```bash
oc delete all -l app=hello-world
```

### Set the name using the --name flag

```bash
oc new-app <image tag> --name <desired name>
```

### In this lesson

```bash
oc new-app quay.io/practicalopenshift/hello-world --name demo-app
```

### The specifier will match your desired name (dc/demo-app in this case)

```bash
oc describe dc/demo-app
```

### Deploy from Git using oc new-app

```bash
oc new-app <git repo URL>
```

### For this lesson

```bash
oc new-app https://gitlab.com/practical-openshift/hello-world.git
```

### Follow build progress

```bash
oc logs -f bc/hello-world
```

### Check status and pods

```bash
oc status
oc get pods
```

### Setting up

```bash
oc get pods --watch
```

### Roll out the latest version of the application

```bash
oc rollout latest dc/hello-world
```

### Roll back to the previous version of the application

```bash
oc rollback dc/hello-world
```

### Access oc explain documentation

```bash
oc explain service
```

### Get more information about Service's spec

```bash
oc explain service.spec
```
---

### Get service documentation

### Access oc explain documentation

```bash
oc explain service
```

### Get more information about Service's spec

```bash
oc explain service.spec
```

### Get YAML definition for a service

```bash
oc get -o yaml service/hello-world
```

### Get YAML definition for a route

```bash
oc get -o yaml route/hello-world
```

---

### Creating services

### Create a service for a single pod

```bash
oc expose --port 8080 pod/hello-world-pod
```

### Create a service for a DeploymentConfig 

```bash
oc expose --port 8080 dc/hello-world
```

### Check that the service and pod are connected properly

```bash
oc status
```

---
### Using Pod environment variables to find service Virtual IPs

### Inside the pod, get all environment variables

```bash
env
```
### Use the environment variables with wget

```bash
wget -qO- $HELLO_WORLD_POD_PORT_8080_TCP_ADDR:$HELLO_WORLD_POD_PORT_8080_TCP_PORT
```
---
### Creating Routes

### Create a Route based on a Service

```bash
oc expose svc/hello-world
```

### Get the Route URL

```bash
oc status
```

### Check the route

```bash
curl <route from oc status>
```

### Create a pod based on a file (just like the Pods section)

```bash
oc create -f pods/pod.yaml
```

### Create a service for the pod

```bash
oc expose --port 8080 pod/hello-world-pod
```

### Check that the service and pod are connected properly

```bash
oc status
```

### Create another pod

```bash
oc create -f pods/pod2.yaml
```

### Shell into the second pod

```bash
oc rsh hello-world-pod-2
```

### Get the service IP and Port

```bash
oc status
```
### In the shell, you can make a request to the service (because you are inside the OpenShift cluster)

```bash
wget -qO- <service IP / Port>
```

### Create a Route based on a Service

```bash
oc expose svc/hello-world
```

### Get the Route URL

```bash
oc status
```

### Check the route

```bash
curl <route from oc status>
```

### Get YAML definition

```bash
oc get -o yaml route/hello-world
```

---
### Creating ConfigMaps

### Create a ConfigMap using literal command line arguments

```bash
oc create configmap <configmap-name> --from-literal KEY="VALUE"
```

### Create from a file

```bash
oc create configmap <configmap-name> --from-file=MESSAGE.txt
```

### Create from a file with a key override

```bash
oc create configmap <configmap-name> --from-file=MESSAGE=MESSAGE.txt
```

### Same --from-file but with a directory

```bash
oc create configmap <configmap-name> --from-file pods
```

### Verify

```bash
oc get -o yaml configmap/<configmap-name>
```
---
### Consuming ConfigMaps as Environment Variables

### Set environment variables (same for all types of ConfigMap)

```bash
oc set env dc/hello-world --from cm/<configmap-name>
```

### Create a configmap using literal command line arguments

```bash
oc create configmap message-map --from-literal MESSAGE="Hello From ConfigMap"
```
### Create a configmap using literal command line arguments

```bash
oc create configmap message-map --from-literal MESSAGE="Hello From ConfigMap"
```
--- 

### Deploy the Hello World app and expose it
```bash
oc new-app quay.io/practicalopenshift/hello-world
oc expose svc/hello-world
```

### Get the route URL
```bash
oc status
```
### First check
```bash
curl <url from route>
```
### Set environment variables
```bash
oc set env dc/hello-world --from cm/message-map
```
### Second check
```bash
curl <url from route>
```
### Get the YAML
```bash
oc get -o yaml dc/hello-world
```
### First, you'll need to create a file
```bash
echo "Hello from ConfigMap file" > MESSAGE.txt
```
### Verify the file exists
```bash
cat MESSAGE.txt
```
### Create the file
```bash
oc create configmap file-map --from-file=MESSAGE.txt
```
### Create the file with a key override
```bash
oc create configmap file-map --from-file=MESSAGE=MESSAGE.txt
```
### Consume ConfigMap (same for all ConfigMaps)
```bash
oc set env dc/hello-world --from cm/file-map-2
```
### Same --from-file but with a directory
```bash
oc create configmap pods-example --from-file pods
```
### Verify
```bash
oc get -o yaml configmap/pods-example
```