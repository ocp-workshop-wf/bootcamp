### Frequently Asked Questions (FAQ)
# OpenShift Bootcamp - Frequently Asked Questions (FAQ)

This FAQ section addresses common questions about OpenShift, its features, and how to use it effectively. If you have additional questions, feel free to reach out or check the official documentation.

## General

**Q: What is OpenShift?**  
A: OpenShift is Red Hat's enterprise Kubernetes platform that automates deployment, scaling, and management of containerized applications.

**Q: What is the difference between Kubernetes and OpenShift?**  
A: OpenShift is a commercial product with a built-in registry, web console, and integrated CI/CD tools, while Kubernetes is an open-source project with modular architecture and requires external tools for some features.

## CLI & Projects

**Q: How do I log in to OpenShift from the command line?**  
A: Use `oc login <cluster URL>`

**Q: How do I create a new project?**  
A: Use `oc new-project <name>`

**Q: How do I switch to a different project?**  
A: Use `oc project <name>`

## Pods & Deployments

**Q: What is the minimum number of containers in a pod?**  
A: 1

**Q: How do I create a pod from a YAML file?**  
A: Use `oc create -f <pod-file.yaml>`

**Q: How do I deploy an image as a deployment?**  
A: Use `oc new-app <image>`

## Scaling

**Q: Do I need a HorizontalPodAutoscaler to scale my app?**  
A: No, you can also scale manually.

**Q: How do I create a Horizontal Pod Autoscaler?**  
A: Use `oc autoscale <deploymentconfig>`

**Q: Which field sets the number of replicas in a DeploymentConfig?**  
A: `replicas`

## Health Checks

**Q: Which probe can cause pod restarts?**  
A: Liveness probe

**Q: What does a readiness probe do?**  
A: It determines if a container is ready to accept traffic.

## Storage & Volumes

**Q: Can emptyDir volumes persist data through pod restarts?**  
A: Yes, for pod restarts, but not node reboots.

**Q: How do I add an emptyDir volume to a DeploymentConfig?**  
A: Use `oc set volume`

## ConfigMaps & Secrets

**Q: What is the maximum data size for a ConfigMap?**  
A: 1 MB

**Q: What format do secrets use for their values?**  
A: Base64

**Q: Are secrets encrypted?**  
A: No, they are base64 encoded, not encrypted.

## Jobs

**Q: What does an OpenShift Job do?**  
A: Runs a task to completion.

**Q: Which field ensures a Job won’t restart the pod after failure?**  
A: `restartPolicy: Never`

**Q: How do you view logs of a job’s pod?**  
A: Use `oc logs -l job-name=<name>`

## Helm

**Q: How do you install a Helm chart?**  
A: `helm install myrelease myrepo/mychart`

**Q: How do you upgrade a Helm release?**  
A: `helm upgrade myrelease myrepo/mychart`

**Q: How do you create a new Helm chart?**  
A: `helm create mychart`

---
