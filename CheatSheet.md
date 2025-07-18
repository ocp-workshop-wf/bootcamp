# Get to the right directory
cd labs
cd hello-world-go

# Build the image
docker build .

# Check the image storage
docker images
# Log in, log out

# Uses the pre-configured OpenShift cluster
oc login

# Allows you to log in to any OpenShift cluster
oc login <cluster address>

# Log out
oc logout


# Project Basics

# See current project
oc project

# Create a new project
oc new-project demo-project

# List all projects
oc projects

# Switch projects
oc project <project name>
# Get Pod Documentation

# Get built-in documentation for Pods
oc explain pod

# Get details on the pod's spec
oc explain pod.spec

# Get details on the pod's containers
oc explain pod.spec.containers


# Creating Pods from files

# Create a Pod on OpenShift based on a file
oc create -f pods/pod.yaml

# Show all currently running Pods
oc get pods


# Port forwarding for Pods

# Open a local port that forwards traffic to a pod
oc port-forward <pod name> <local port>:<pod port>

# Example of 8080 to 8080 for hello world
oc port-forward hello-world-pod 8080:8080


# Shell into Pods

# oc rsh will work with any Pod name from oc get pods
oc rsh <pod name>

# In the shell, check the API on port 8080
wget localhost:8080

# Exit the rsh session
exit

# Watch live updates to pods
oc get pods --watch


# Delete (stop) Pods

# Delete any OpenShift resource
oc delete <resource type> <resource name>

# Delete the pod for this section
oc delete pod hello-world-pod

# Watch live updates to pods
oc get pods --watch

# Deploying applications as DeploymentConfigs

# Deploy an existing image based on its tag
oc new-app <image tag>

# Deploy the Hello World image for this course
oc new-app quay.io/practicalopenshift/hello-world

# Deploy from Git using oc new-app
oc new-app <git repo URL>

# Deploy the Hello World application from Git
oc new-app https://gitlab.com/practical-openshift/hello-world.git

# Follow build progress (Git only)
oc logs -f bc/hello-world

# Set the name for the DeploymentConfig 
oc new-app <image tag> --name <desired name>

# Example with a name
oc new-app quay.io/practicalopenshift/hello-world --name demo-app


# Get more information about a DeploymentConfig

# Describe the DC to get its labels
oc describe dc/hello-world

# Get the full YAML definition
oc get -o yaml dc/hello-world


# Deleting all oc new-app resources

# Delete all application resources using labels (get them from oc describe)
oc delete all -l app=hello-world


# Starting new versions and reverting changes

# Roll out the latest version of the application
oc rollout latest dc/hello-world

# Roll back to the previous version of the application
oc rollback dc/hello-world

# Deploy an existing image based on its tag
oc new-app <image tag>

# For this lesson
oc new-app quay.io/practicalopenshift/hello-world

# Check running resources
oc status

# Check pods
oc get pods

# Standard image deployment 
oc new-app quay.io/practicalopenshift/hello-world

# Describe the DC to get its labels
oc describe dc/hello-world

# Delete all application resources using labels
oc delete all -l app=hello-world

# Set the name using the --name flag
oc new-app <image tag> --name <desired name>

# In this lesson
oc new-app quay.io/practicalopenshift/hello-world --name demo-app

# The specifier will match your desired name (dc/demo-app in this case)
oc describe dc/demo-app

# Deploy from Git using oc new-app
oc new-app <git repo URL>

# For this lesson
oc new-app https://gitlab.com/practical-openshift/hello-world.git

# Follow build progress
oc logs -f bc/hello-world

# Check status and pods
oc status
oc get pods

# Setting up
oc get pods --watch

# Roll out the latest version of the application
oc rollout latest dc/hello-world

# Roll back to the previous version of the application
oc rollback dc/hello-world

# Access oc explain documentation
oc explain service

# Get more information about Service's spec
oc explain service.spec

# Get service documentation

# Access oc explain documentation
oc explain service

# Get more information about Service's spec
oc explain service.spec

# Get YAML definition for a service
oc get -o yaml service/hello-world

# Get YAML definition for a route
oc get -o yaml route/hello-world


# Creating services

# Create a service for a single pod
oc expose --port 8080 pod/hello-world-pod

# Create a service for a DeploymentConfig 
oc expose --port 8080 dc/hello-world

# Check that the service and pod are connected properly
oc status


# Using Pod environment variables to find service Virtual IPs

# Inside the pod, get all environment variables
env

# Use the environment variables with wget
wget -qO- $HELLO_WORLD_POD_PORT_8080_TCP_ADDR:$HELLO_WORLD_POD_PORT_8080_TCP_PORT


# Creating Routes

# Create a Route based on a Service
oc expose svc/hello-world

# Get the Route URL
oc status

# Check the route
curl <route from oc status>

# Create a pod based on a file (just like the Pods section)
oc create -f pods/pod.yaml

# Create a service for the pod
oc expose --port 8080 pod/hello-world-pod

# Check that the service and pod are connected properly
oc status

# Create another pod
oc create -f pods/pod2.yaml

# Shell into the second pod
oc rsh hello-world-pod-2

# Get the service IP and Port
oc status

# In the shell, you can make a request to the service (because you are inside the OpenShift cluster)
wget -qO- <service IP / Port>

# Create a Route based on a Service
oc expose svc/hello-world

# Get the Route URL
oc status

# Check the route
curl <route from oc status>

# Get YAML definition
oc get -o yaml route/hello-world

# Creating ConfigMaps

# Create a ConfigMap using literal command line arguments
oc create configmap <configmap-name> --from-literal KEY="VALUE"

# Create from a file
oc create configmap <configmap-name> --from-file=MESSAGE.txt

# Create from a file with a key override
oc create configmap <configmap-name> --from-file=MESSAGE=MESSAGE.txt

# Same --from-file but with a directory
oc create configmap <configmap-name> --from-file pods

# Verify
oc get -o yaml configmap/<configmap-name>


# Consuming ConfigMaps as Environment Variables

# Set environment variables (same for all types of ConfigMap)
oc set env dc/hello-world --from cm/<configmap-name>

# Create a configmap using literal command line arguments
oc create configmap message-map --from-literal MESSAGE="Hello From ConfigMap"



