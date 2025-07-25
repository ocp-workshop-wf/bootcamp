# 🚀 OpenShift CLI (oc) Cheat Sheet – One Page

---

## 🧭 Authentication & Context
```bash
oc login https://<cluster-api> --token=<token>        # Login to cluster
oc whoami                                             # Show current user
oc config get-contexts                                # List contexts
oc config use-context <context-name>                  # Switch context
oc project                                            # Show current project
oc project <project-name>                             # Switch to project
```

## 📦 Projects & Namespaces
```bash
oc new-project <name>                                 # Create new project
oc get projects                                       # List projects
oc delete project <name>                              # Delete a project
```
## 🚀 Application Deployment
```bash
oc new-app <image|template|git-url>                   # Deploy from image or Git
oc rollout status dc/<app-name>                       # Check rollout status
oc expose svc/<service-name>                          # Create route to service
oc delete all -l app=<app-name>                       # Delete all app resources
```
## 🔁 Builds & Images
```bash
oc start-build <build-config-name>                    # Start a manual build
oc logs bc/<build-config>                             # View build logs
oc get is                                             # List image streams
```
## 📜 Resource Management
```bash
oc get all                                            # Get all resources
oc get pods,svc,dc,route                              # Get specific types
oc describe pod/<pod-name>                            # Describe pod
oc delete pod/<pod-name>                              # Delete pod
```
## 🛠️ Pod Debugging
```bash
oc logs <pod-name>                                    # View logs
oc logs -f <pod-name>                                 # Stream logs
oc rsh <pod-name>                                     # Remote shell into pod
oc exec <pod-name> -- <command>                       # Run cmd in pod
oc port-forward <pod-name> <local>:<remote>           # Port forward
```
## ⚙️ ConfigMaps & Secrets
```bash
oc get configmap|secret                               # List configs/secrets
oc create configmap <name> --from-literal=key=value   # Create configmap
oc create secret generic <name> --from-literal=key=val# Create secret
oc edit configmap/<name>                              # Edit configmap
```
## 📈 Scaling & Autoscaling
```bash
oc scale dc/<name> --replicas=3                       # Manually scale
oc autoscale dc/<name> --min=1 --max=5 --cpu-percent=80 # HPA setup
```
## 🌐 Routes & Networking
```bash
oc get route                                          # List routes
oc expose svc/<svc-name>                              # Expose service via route
oc describe route/<route-name>                        # Route details
```
## 🧾 Templates & YAML
```bash
oc process -f <template.yaml>                         # Process template
oc apply -f <file.yaml>                               # Apply YAML
oc delete -f <file.yaml>                              # Delete from YAML
```







