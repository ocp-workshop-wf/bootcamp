## Module 6: Mastering OpenShift
### 6.1 Templates
A template is a reusable definition of a set of OpenShift objects (like pods, services, routes, etc.) that can be parameterized and instantiated to create those objects within a project. Essentially, it's a way to automate the creation of multiple related resources, making it easier to deploy and manage applications.
    ***Template Parts***
        - Preamble: is just a header identifying the file as a template file. i.e(basic metadata: `name`, `kind`).
        - Object (Resource) List: Define the various pieces that make up your application. Each object in the list contains a full definition for a normal OpenShift resource.
        - Parameters: Allow you to make your application configurable.Common parameters include `credentials`, `database names`, more ...

**Hands-on Walkthroughs** 
- Using Templates: How to upload Template files?

    ```bash
    cd ./labs-repo/template
    ```
    ```bash
    oc create -f hello-world-template.yaml
    ```
    > output: template.template.openshift.io/hello-world created
    
    ```bash
    oc get template
    ```
    > output: 

    | NAME | DESCRIPTION | PARAMETERS | OBJECTS |
    | ---- | ----------- | ---------- | ------- |
    | hello-world |      | 1 (all set) | 4 |

    ```bash
    oc new-app hello-world
    ```
- How to use Template Parameters?

    ```bash
    oc get -o yaml template/hello-world
    ```
    > output: 
    ```yml
    .....
    spec:
      containers:
      - env:
        - name: MESSAGE
          value: ${MESSAGE} #This is the syntax that allows you to substitute parameter values into your template.
    .......
      parameters:
      - description: Message to respond to requests with
      displayName: Message
      name: MESSAGE
      value: Hello from the default value for the template
    ```
- Clean up:
  ```bash
  oc delete all --all
  ```
- Lets set some parameter values:

    ```bash
    oc new-app hello-world \
  -p MESSAGE="Hello from parameter override."
    ```
    > output: you should see a Success message.

    ```bash
    curl <your route>
    ```
    > output: "Hello from parameter override."

- Processing Templates - How to process Templates?
  - This means you can print out the final set of resources for a template taking parameters into account. Once you have the list of objects in a file, you can inspect and modify it manually beofre using `oc create`

    ```bash
    oc process hello-world
    ```
    > output: `JSON` format - you can specify it in yaml by using `-o yaml`

    ```bash
    oc process hello-world -o yaml
    ```
    > output: no longer contains {} as it shows the actual value

    ```yml
    .....
      spec:
        containers:
        - env:
          - name: MESSAGE
            value: Hello from the default value for the template
    ......
    ```
- Lets run the `oc process` command and give a parameter value.

    ```bash
    oc process hello-world -o yaml \
  -p MESSAGE="Hello from oc process"
    ```
    > output: you should see this -->
    
    ```yml
    .....
      spec:
        containers:
        - env:
          - name: MESSAGE
            value: Hello from oc process
    .....
    ```
  - Clean up:

    ```bash
    oc delete all --all
    ```

- How to use the output from `oc process` to actually create objects on the server.
  - First lets save the processed template to a file:
  
    ```bash
    oc process hello-world -o yaml \
  -p MESSAGE="Hello from oc process" \
  > processed-objects.yaml
    ```
  - Check the file content 
    ```bash
    head processed-objects.yaml
    ```
    > output: 

    ```yml
    apiVersion: v1
    items:
    - apiVersion: apps.openshift.io/v1
      kind: DeploymentConfig
      metadata:
        annotations:
          openshift.io/generated-by: PracticalOpenshift
        labels:
          app: hello-world
        name: hello-world
    ```
  - Deploy the processed file
    ```bash
    oc create -f processed-objects.yaml
    ```
    ```bash
    oc status #grab the route url
    ```
    ```bash
    curl <URL>
    ```
    > output: "Hello from oc process"

- How create a custom Template?
  - Get YAML for existing objects on the OpenShift server.
  - The dc,is... syntax lists the types of resources that you would like to export
  - Add hpa or any other type if you need them
    
    ```bash
    oc get -o yaml dc,is,bc,svc,route --export
    ```
    ```bash
    vi test-template.yaml
    ```
    - Steps for a custom template:
      - Change the items property to objects
      - Change kind from List to Template
      - Add a name property to the metadata section
      - Remove status from each resource
      - Remove most of metadata except for name, labels, and annotations
      - Remove any automatically-assigned resources such as service Virtual IPs and Route hosts
      - (optional) Add template parameters

- How to use built-in Templates
    ```bash
    oc get template -n openshift
    ```
    > output: list of templates.

    
    

    


### 🔬 Hands-on Lab: 
For Templates, you will create your own parameterized template for a small REST API with a database backend. First, you will need to get the application working on OpenShift.
    - For the REST API, you can use the `mysql-go-reader` directory in the labs project.
    - For MySQL, you can use the `mysql` image and `oc new-app`.
    - To get them working, you'll need to supply a few environment variables to both images. Specifically, you'll need to provide `MYSQL_USER`, `MYSQL_PASSWORD`, and `MYSQL_DATABASE`. Store these values in an OpenShift resource and use it to populate both sets of environment variables.
    - Expose a route for the application
    - Create an autoscaler that will scale the application from 3 to 10 pods depending on load

### Checklist 📋: 
- `MAX_PODS`: this should control the maximum number of pods for the REST API
- `MYSQL_USER`: This variable should be used for both the mysql and mysql-go-reader images
- `MYSQL_PASSWORD`: This variable should be used for both the mysql and mysql-go-reader images
- `MYSQL_DATABASE`: This variable should be used for both the mysql and mysql-go-reader images

> Supply default values so that the template will work without specifying any parameters with oc new-app.

--- 

### Quiz
> Q1: You can only work with templates as files on your filesystem.
- [ ] True
- [ ] False

<details>
  <summary> Answer </summary>

False: You can also upload the template to OpenShift as a standard REST API resource and work with it that way!

</details>

> Q2: You can use oc new-app with a template name in order to create an application according to the template.
- [ ] True
- [ ] False

<details>
  <summary> Answer </summary>

True: `oc new-app` does support templates as arguments.

</details>

> Q3: What syntax can you use to get YAML definitions for all configmaps and secrets in your project as a List?
- [ ] `oc get cm,secret --export`
- [ ] `oc get -o yaml all --export`
- [ ] `oc get -o yaml cm,secret --exprot`
- [ ] `oc get cm -o yaml, secret -o yaml`

<details>
  <summary> Answer </summary>

`oc get -o yaml cm,secret --exprot`

</details>

> Q4: What flag is used to specify a parameter value for oc new-app or oc process commands?
- [ ] `--param <value>`
- [ ] `-p <value>`
- [ ] `--p <value>`
- [ ] `parameter <value>`

<details>
  <summary> Answer </summary>

`-p <value>`

</details>

---


### 6.2 Advanced DeploymentConfigs


**Hands-on Walkthroughs** 

### 🔬 Hands-on Lab: 

### Checklist 📋: 

### Quiz
> Q1: 
- [ ]  
- [ ]
- [ ]
- [ ] 


<details>
  <summary> Answer </summary>

    
  

</details>