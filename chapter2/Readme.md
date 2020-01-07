## Target

- Dev a webapp with golang
- Build a docker image

---

- Run in k8s

  - Use command to run pod

    if you type the following command `$ kubectl run gotest --image=gotest:1.1 --port=8080` it should occue error:

    ```
    Failed to pull image "gotest:1.1": rpc error: code = Unknown desc = Error response from daemon: pull access denied for gotest, repository does not exist or may require 'docker login'
    ```

    you can fix that in [https://stackoverflow.com/questions/42564058/how-to-use-local-docker-images-with-minikube](https://stackoverflow.com/questions/42564058/how-to-use-local-docker-images-with-minikube)

    Or you can push the image to the hub.docker.io; then type the following command:
  `$ kubectl run gotest --image=yfsoftcom/gotest:1.1 --port=8080`

  `$ kubectl get po -o wide`

    ```
    NAME                      READY   STATUS    RESTARTS   AGE     IP           NODE       NOMINATED NODE   READINESS GATES
    gotest-86dfc4c4c7-r5g7l   1/1     Running   0          5m55s   172.17.0.6   minikube   <none>           <none>  
    ```

    But, there is only 1 pod; we need replice it 3. Here is a question?

    `$ kubectl `

   - Expose the port use nodeport. 

   `$ kubectl expose deployment gotest --type=NodePort --name gotest-http`


   - get the NodePort

   `$ kubectl get svc`

    - get the cluster info

    `$ kubectl cluster-info`

- Questions:

  - [ ] 3 Pod ? how to extend the pod ?

    `$ kubectl scale deployment/gotest --replicas=3`

  - [ ] How to defined the port for the service ?

    `$ ??`

  - [ ] How to upgrade the image ?

    `$ kubectl set image deployment/gotest gotest=yfsoftcom/gotest:1.2`

    check the rollout status

    `$ kubectl rollout status deployment gotest`

    run a loop to curl the api
    `for (( ; ; )) do curl 192.168.39.169:32761/foo?foo=bar; echo ''; sleep 1; done;`

  - [ ] Run a db to test

     use redis to count the view count

     create the redis pod.yml
    ```yml
    apiVersion: v1
    kind: Pod
    metadata:
      name: redis
      labels:
        name: redis
    spec:
      containers:
      - name: redis
        image: redis:alpine3.11
        resources:
          limits:
            memory: "256Mi"
            cpu: "500m"
        ports:
          - containerPort: 6379

    ```
     expose the port

     `$ kubectl expose pod redis --type=NodePort --name redis-port`

     `$ docker run -it redis:alpine3.11 /bin/sh`

     `$ redis-cli -h 192.168.39.169 -p 31311`

