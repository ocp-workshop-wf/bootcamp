## Module 4: Application Deployment and Management 

### 4.1 OpenShift ConfigMaps

- **DeploymentConfig** vs. **Deployment**: DeploymentConfig supports S2I, triggers.
- S2I builds containers directly from source code.

**Hands-on Walkthroughs**  
- Deploy a Node.js or Python app from GitHub via S2I.



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
