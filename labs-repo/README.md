- Hello-world-go-private 
    - Used in Image and ImageStream in Module 4
    - It has the From to take the existing hello-world image then ENV to tweak the message enviroment variable.
    -  



This is the Labs repository 

Action Items: 

Section 4 Triggers should only be 15 minutes with no lab
Section 5.1 Add reference of how S2I is like Konika (sp?); no lab necessary
Section 5.3  Expand Scaling and Debugging to include Horizontal and Kubernetes
Section 5.3 Add a Scaling and Debugging lab; if short on time demo through YouTube
Section 5  Move Deployment Strategies to Section 4.
Add lab for Rolling Deployment (may already exist, confirm first)
Include Blue Green Deployment (include lab if time)
Include Canary Deployment (Copilot might have sample lab)
Section 6 Shorten Templates section (Service Mesh not needed)
Section 6.1 Add Start Up Drive (?); Consider leaving with a Kubernetes Reference Course
Section 6.2 Add more about Observatory (monitoring events, container image, etc)
Section 6.3 Helmchart include dry run and OC Apply (?)

ALL: 
Remove reference to DeploymentConfig
Look at Docker Command
Move all lab code to ReadMe file
Clean up jobs in Labs Repo


Impact of Deployment Pattern 
Canary: 
Enables gradual exposure, reducing risk. Monitoring metrics, logs, and user feedback during this phase are crucial for deciding whether to proceed. 
Blue-Green: 
Simplifies rollback if issues arise by allowing quick reversion to the stable environment. The rehearsal phase ensures the switchover process is smooth. 
Rolling Updates: 
If used, these need to be carefully managed so that service continuity is maintained during instance-by-instance updates. 
3.3 Deployed Stage 
Objective: 
Deploy the update to the live Production environment using strategies designed for zero-downtime and rapid rollback. 
Key Activities: 
Deployment Execution: 
Implement one of the following strategies: 
Blue-Green Deployment: Prepare a new (green) environment and switch traffic once tests pass. 
Rolling Updates: Update servers in batches while keeping portions of the environment active. 
Canary Releases: Gradually increase traffic to the new version until fully live. 

Summary of Pattern Effects 
Blue-Green Deployment 
Union: Prepares parallel environments. 
Rehearsal: Allows testing of the complete cutover process. 
Delivered: Facilitates a rapid switch-over with minimal risk. 
Rolling Updates 
Union: Can simulate incremental instance updates. 
Rehearsal: Tests the impact of progressive changes. 
Delivered: Ensures gradual, low-risk deployment across all instances. 
Canary Releases 
Union: Supports initial validation with minimal exposure. 
Rehearsal: Validates under near-production traffic conditions. 
Delivered: Gradually transitions all traffic, allowing continuous monitoring and immediate rollback if necessary