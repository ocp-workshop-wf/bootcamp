## Module 4: Application Deployment and Management 

### 4.1 Deploying Applications on OpenShift

- **DeploymentConfig** vs. **Deployment**: DeploymentConfig supports S2I, triggers.
- S2I builds containers directly from source code.

**Lab:**  
- Deploy a Node.js or Python app from GitHub via S2I.

**YouTube:**  
- [Deploying Apps from GitHub in OpenShift](https://www.youtube.com/watch?v=NSry4lt3puQ)

---

### 4.2 Managing Application Configurations

- **ConfigMaps**: Key-value config data.
- **Secrets**: Encrypted, for sensitive info.
- Mount as env vars or volumes.

**Lab:**  
- Create and mount ConfigMaps and Secrets in a deployment.

**YouTube:**  
- [Using ConfigMaps & Secrets in OpenShift](https://www.youtube.com/watch?v=AnvOMRFwimM)

---

### 4.3 Configuring and Scaling Applications

- Scale deployments manually or with autoscaling (HPA).
- Set autoscaling thresholds for CPU/memory.

**Lab:**  
- Apply an HPA, generate load, observe auto-scaling.

**YouTube:**  
- [OpenShift Deployments & Scaling](https://www.youtube.com/watch?v=JysYQ3a7fwQ)

---
