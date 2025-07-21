List of prerequisites:
- [x] Docker or Podman
- [x] Git
- [x] vim or vi 
- [x] Visual Studio code / or any other IDE
- [x] OpenShift CLI 
- [x] iterm terminal preferred for multi-terminal purposes.


## Docker

### Building containers

#### Build an image based on the current directory
`docker build .`

#### Build an image based on the current directory with a tag
`docker build -t your-tag .`

### Running Containers

#### Check the image storage
`docker images`

> Start a container based on an image ID. Get the ID from docker images.

#### `control-c` will stop the container for all of these docker run commands.
`docker run -it <your-image-id>`

#### Start an image based on a tag
`docker run -it <image tag>`

#### Start hello-world server
`docker run -it quay.io/practicalopenshift/hello-world`

#### Start the Hello World app with port forwarding
`docker run -it -p 8080:8080 quay.io/practicalopenshift/hello-world`


### Stopping Containers

#### Get running images
`docker ps`

#### Stop a running image. The container ID will be in the docker ps output.
`docker kill <Container ID>`

### Common OpenShift 4 Resource Types

#### Pod:
A Pod is the basic deployment unit in OpenShift and Kubernetes. It represents a single instance of an application and runs one or more containers.
#### Service:
A Service is an abstraction that defines a logical set of Pods and a policy by which to access them.
#### Deployment:
A Deployment is used to define the desired state of an application and manage the application’s Pods and ReplicaSets.
(Note: OpenShift also supports ‘DeploymentConfig’, which provides additional features such as triggering deployments on image changes or config changes.)
#### ReplicaSet:
A ReplicaSet is used to maintain a stable set of replica Pods running at any given time.
#### Route:
A Route is a way to expose a Service externally.
(Note: Kubernetes uses ‘Ingress’ to expose services externally, while OpenShift uses ‘Route’.)
#### ConfigMap:
A ConfigMap is used to store configuration data as key-value pairs that can be consumed by Pods.
#### Secret:
A Secret is used to store sensitive data, such as passwords or keys, which can be used by Pods.
#### PersistentVolume (PV):
A PV is a piece of storage in the cluster that has been provisioned by an administrator or dynamically provisioned using Storage Classes.
#### PersistentVolumeClaim (PVC):
A PVC is a request for storage by a user that can be fulfilled by a PersistentVolume.
#### StatefulSet:
A StatefulSet is used for deploying, scaling, and managing a set of Pods with persistent storage and persistent identifiers.
#### DaemonSet:
A DaemonSet is used to ensure that all (or some) nodes run a copy of a Pod. As nodes are added or removed from the cluster, the DaemonSet will add or remove the required Pods.
#### ReplicationController:
A ReplicationController is the predecessor to ReplicaSet and ensures that a specified number of Pod replicas are running at any one time.
#### Job:
A Job is used to run a Pod to completion. It is used to run batch jobs that complete and then terminate.
#### CronJob:
A CronJob is used to run Jobs on a time-based schedule.
#### BuildConfig:
A BuildConfig is an OpenShift specific resource for defining build configurations. (OpenShift specific)
#### Build:
A Build is an OpenShift specific resource that represents a single build of an application. (OpenShift specific)
#### ImageStream:
An ImageStream is an OpenShift specific resource for managing sets of images. (OpenShift specific)
#### ImageStreamTag:
An ImageStreamTag is an OpenShift specific resource for accessing individual tags in an #### ImageStream. (OpenShift specific)
#### ImageStreamImport:
An ImageStreamImport is an OpenShift specific resource for importing images from external registries. (OpenShift specific)


