## Module 3: Core OpenShift Resources 

### 3.1 OpenShift Resources Overview

- **Pods:** Containers grouped for management.
- **ReplicaSets & Deployments:** Maintain desired pod count, enable updates.
- **Services:** Stable endpoints for pods.
- **Routes:** HTTP(S) ingress with DNS/SSL.

**Lab:**  
- Deploy nginx:  
  ```sh
  oc create deployment nginx --image=nginx
  oc expose deployment nginx --port=80
  oc expose svc/nginx
  ```
- Visit the exposed route.

**YouTube:**  
- [Deploying Apps in OpenShift](https://www.youtube.com/watch?v=drM-p5NH_90)

---

### 3.2 Container Lifecycle Management

- Understand pod phases and restart policies.
- Add lifecycle hooks (PostStart, PreStop).

**Lab:**  
- Edit deployment YAML to add lifecycle hooks.
- Delete pods and observe automated restart.

---

### 3.3 Image Stream Management

- **ImageStreams** let you manage and tag images, enabling automated redeployments.
- **BuildConfig/S2I** automates builds from source code.

**Lab:**  
- Create an ImageStream, BuildConfig.
- Push new code/image, watch auto-update.

**YouTube:**  
- [OpenShift Image Streams Explained](https://www.youtube.com/watch?v=jA-RH0jO-J8)

---

