## Module 5: OpenShift Service Mesh & Advanced Traffic Management (4 hours)

### 5.1 Introduction Source-to-image (S2I)
 It is a toolkit and workflow that automates the process of building container images from source code. It takes a builder image (containing necessary build tools and dependencies) and source code, then combines them to create a runnable application image. S2I is used to create reproducible container images, making it easier for developers to deploy and manage applications in various environments, including OpenShift. 

 ![Source-to-Imag](/images/s2i-concept.webp)

 > It uses `Assemble` script vs `Run` from Docker and `Run` vs `CMD` from Docker.

 - FAQ: Why do I want another way to build and run applications?so what advantage does S2I have?
 
    - The first advantage is that the developer can rely on the expertise of the S2I authors to make sure the image will be OpenShift compatible.
    - The second advantage is that developers can avoid writing and maintaining Dockerfiles and other configuration by using the S2I defaults. 




**Hands-on Walkthroughs** 
- How to build an app using S2I?
> the source of this lesson will be in you Labs directory `/s2i/ruby/`

```bash
cd ./s2i/ruby
```
```bash
oc new-app ruby~https://gitlab.com/therayy1/openshif-labs.git --context-dir s2i/ruby --as-deployment-config
```

```bash
oc logs -f bc/openshif-labs
```
> output: push successful 

```bash
oc get pods
```
> output: you will be able to see all pods related to this deployment

```bash
oc expose svc/openshif-labs
```
> output: "route.route.openshift.io/openshif-labs exposed"

```bash
oc get route
```
> output: URL under `HOST/PORT` copy it please!

```bash
curl <URL>
```
> output: "Hello world from Ruby"

- How S2I language auto-detect works?!
- How does OpenShift know when to use S2I? And specifically, how did OpenShift know how to use the Ruby S2I builder?

  - When you start a build with OpenShift, OpenShift will first look for a dockerfile. If it finds one it will build usig something called the Docker Strategy. If OpenShift does not find a dockerfile it will attempt to use the Source Strategy instead.
   
  ![S2I Strategy](/images/S2I%20-%20visual%20selection.png)

### 🔬 Hands-on Lab: 
For builds, you will make a small tweak to an application, push it to GitLab, and then run it on your OpenShift instance.

- Create a new repository under your GitLab account (https://docs.gitlab.com/ee/gitlab-basics/create-project.html)
- Add the new repository as a remote for the labs project (https://git-scm.com/book/en/v2/Git-Basics-Working-with-Remotes)
- Create a new branch in the labs repository named `builds-lab`
- Modify the Dockerfile in the `hello-world-go` directory to change the `MESSAGE` environment variable
- Commit the change and push it to GitLab
- Create a BuildConfig for this updated `hello-world-go` directory in the `builds-lab` branch
- Add a build hook that prints the value of the `MESSAGE` environment variable
- Deploy the application based on the resulting ImageStream
---

### Checklist 📋: 
- oc logs `<your BuildConfig>` contains the build hook output with your `MESSAGE` value from step 5
- `curl <your app>` shows you the `MESSAGE` value from step 5
- You can start the build using the webhook manually

---
### Quiz
> Q1: What is the command to create a new BuildConfig for a Git URL?
- [ ] `oc start-build <GIT URL>`
- [ ] `oc new-build <GIT URL>`
- [ ] `oc import-image <GIT URL>`
- [ ] `oc new-app build <GIT URL>`

<details>
  <summary> Answer </summary>

  `oc new-build <GIT URL>`

</details>
---

### 5.2 Blue/Green Deployments & Traffic Routing

- Deploy two app versions (blue/green), split or shift traffic.
- Service Mesh allows advanced traffic splitting, routing based on headers, instant rollback.

**Lab:**  
- Deploy v1/v2 of an app.
- Use Kiali to shift traffic (80/20 to 100%).

**YouTube:**  
- [OpenShift Service Mesh: Blue/Green Deployments](https://www.youtube.com/watch?v=xWBUOq_jIVg)

---

### 5.3 Observability and Tracing

- **Kiali**: Service connections and health.
- **Jaeger**: Distributed tracing of requests.
- **Grafana**: Dashboards for metrics.

**Lab:**  
- Generate test traffic, use Kiali and Jaeger for live monitoring and tracing.

**YouTube:**  
- [OpenShift Service Mesh Observability](https://www.youtube.com/watch?v=VdIewxv9CHs)

---

### 5.4 Service Mesh Use Cases

- Fault injection, retries, circuit breaking, mTLS, policy enforcement.
- Test app resilience and security without changing code.

**Lab:**  
- Inject a failure using Service Mesh and observe via Kiali.

---

