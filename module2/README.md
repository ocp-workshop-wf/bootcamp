## Module 2: Working with OpenShift Interfaces 

### 2.1 OpenShift CLI (oc) and Web Console

- The **Web Console** offers tabs for Overview, Workloads, Networking, Builds, Storage, Monitoring, etc.
- The **CLI (`oc`)** gives full control with extra features over `kubectl`.

**Lab:**  
- Login using `oc login`
- Create a project: `oc new-project demo-project`
- List pods: `oc get pods`

**YouTube:**  
- [OpenShift CLI Basics](https://www.youtube.com/watch?v=8wFJe2U1GdI)

---

### 2.2 Creating and Managing Projects

- Projects = Namespaces + metadata/RBAC.
- Add quotas, manage access via RBAC.

**Lab:**  
- Create a project, set a resource quota.
- Assign a user with:  
  `oc adm policy add-role-to-user edit <username> -n demo-project`

**YouTube:**  
- [Getting Started with OpenShift Web Console](https://www.youtube.com/watch?v=Qzvfi6VSFoc)

---
