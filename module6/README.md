## Module 6: Reliability and Operations (4 hours)

### 6.1 Debugging Applications

- View logs (`oc logs`), exec into containers (`oc rsh`), use `oc debug`.
- Diagnose failed pods/events.

**Lab:**  
- Break an app, debug via logs and shell.

**YouTube:**  
- [Debugging in OpenShift](https://www.youtube.com/watch?v=2KewDcsK1iA)

---

### 6.2 Deployment Strategies

- Rolling update, recreate, blue/green, canary.
- Rollbacks for quick recovery.

**Lab:**  
- Perform a rolling update, then rollback.

**YouTube:**  
- [OpenShift Deployment Strategies](https://www.youtube.com/watch?v=lUIzX6mrR1g)

---

### 6.3 Application Health Checks

- **Readiness Probe**: Can serve traffic.
- **Liveness Probe**: Should be restarted if failing.

**Lab:**  
- Add probes to deployments, watch for auto-recovery.

**YouTube:**  
- [Kubernetes Health Checks](https://www.youtube.com/watch?v=R5ka3rN5qpE)

---

### 6.4 Configuring Jobs in OpenShift

- **Jobs**: One-off tasks.
- **CronJobs**: Scheduled tasks.

**Lab:**  
- Create and schedule a CronJob.

**YouTube:**  
- [OpenShift Jobs and CronJobs](https://www.youtube.com/watch?v=SRgZyjGUPPA)

---
