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

---

### 4.2 Secrets

- **ConfigMaps**: Key-value config data.
- **Secrets**: Encrypted, for sensitive info.
- Mount as env vars or volumes.

**Hands-on Walkthroughs**  
- Create and mount ConfigMaps and Secrets in a deployment.

**Resource:**  
- [Using ConfigMaps & Secrets in OpenShift](https://www.youtube.com/watch?v=AnvOMRFwimM)

---

### 4.3 Images and Image Streams 

- Scale deployments manually or with autoscaling (HPA).
- Set autoscaling thresholds for CPU/memory.

**Hands-on Walkthroughs**  
- Apply an HPA, generate load, observe auto-scaling.

**Resource:**  
- [OpenShift Deployments & Scaling](https://www.youtube.com/watch?v=JysYQ3a7fwQ)

---

### 4.4 Builds and BuildConfigs
