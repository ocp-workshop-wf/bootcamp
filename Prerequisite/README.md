List of prerequisites:
- [x] Text Editor
- [x] Visual Studio code - Preferred
- [x] OpenShift CLI 
- [x] 9 GB RAM
- [x] 10+ GB Disk Space
- [x] 4 virtual CPUs recommended 

### Common OpenShift Resource Types

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


