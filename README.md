# Getting Started
This is a very simple hello world Kubernetes operator that used to demonstrate the basic functionalities and implementation details of an operator. The implementation of this operator based on the [Operator SDK](https://github.com/operator-framework/operator-sdk/) framework, a framework that generates all the boilerplate code of a Kubernetes operator and provides functionalities to build operators.

## Prerequisites

### Run the Operator
- [kubectl v1.11.3+](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- Kubernetes v1.11.3+ cluster

### Build the Operator
- [operator-sdk](https://github.com/operator-framework/operator-sdk/blob/master/doc/user/install-operator-sdk.md)
- [dep](https://golang.github.io/dep/docs/installation.html) version v0.5.0+
- [git](https://git-scm.com/downloads)
- [go](https://golang.org/dl/) version v1.12+
- [docker](https://docs.docker.com/install/) version 17.03+
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) version v1.11.3+
- Kubernetes v1.11.3+ cluster

## Configure a Kubernetes Cluster
As described in the prerequisites section, to run or build the operator we need to configure a Kubernetes cluster. There are multiple ways to configure a K8s cluster, therefore we can configure one of the following clusters. 

- [Minikube](https://github.com/kubernetes/minikube#installation)
- [Docker for Mac](https://docs.docker.com/docker-for-mac/install/)
- [Google Kubernetes Engine](https://github.com/siddhi-io/siddhi-operator/blob/v0.2.2/docs/gke-setup.md)

## Install the HelloWorld Operator

Install the prerequisites.
```sh
$ kubectl apply -f https://raw.githubusercontent.com/BuddhiWathsala/helloworld-k8s-operator/v0.3.0/deploy/crds/helloworld.io_helloworlds_crd.yaml
$ kubectl apply -f https://raw.githubusercontent.com/BuddhiWathsala/helloworld-k8s-operator/v0.3.0/deploy/service_account.yaml
$ kubectl apply -f https://raw.githubusercontent.com/BuddhiWathsala/helloworld-k8s-operator/v0.3.0/deploy/role.yaml
$ kubectl apply -f https://raw.githubusercontent.com/BuddhiWathsala/helloworld-k8s-operator/v0.3.0/deploy/role_binding.yaml
```
Install the operator deployment.

```sh
$ kubectl apply -f https://raw.githubusercontent.com/BuddhiWathsala/helloworld-k8s-operator/v0.3.0/deploy/operator.yaml
```

## Deploy an Example

Here, we are going to deploy a HelloWorld example app that created 
```yaml
apiVersion: helloworld.io/v1alpha1
kind: HelloWorld
metadata:
  name: example-helloworld
spec:
  size: 1
```

Apply the HelloWorld example using the following command.
```sh
$ kubectl apply -f https://raw.githubusercontent.com/BuddhiWathsala/helloworld-k8s-operator/v0.3.0/deploy/crds/helloworld.io_v1alpha1_helloworld_cr.yaml
```

These commands will create the following Kubernetes artifacts in your cluster.

```sh
➜  ~ kubectl get pods
NAME                                      READY     STATUS    RESTARTS   AGE
example-helloworld-849865f49c-5cr2n       1/1       Running   0          84s
helloworld-k8s-operator-b46f79fb5-6kvq5   1/1       Running   0          112s

➜  ~ kubectl get svc
NAME                              TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)             AGE
example-helloworld                ClusterIP   10.106.176.32    <none>        8080/TCP            97s
helloworld-k8s-operator-metrics   ClusterIP   10.108.109.142   <none>        8383/TCP,8686/TCP   108s
kubernetes                        ClusterIP   10.96.0.1        <none>        443/TCP             4d19h
```

## Send Requests

The Go application that we just deployed using this `helloworld-k8s-operator`, received HTTP GET requests and send the received payload back to the caller. However, we cannot send the requests straight away since we do not expose the deployed service outside. To expose the HelloWorld service we have to use Kubernetes port forward as follows.
```sh
➜  ~ kubectl port-forward svc/example-helloworld 8081:8080
Forwarding from 127.0.0.1:8081 -> 8080
Forwarding from [::1]:8081 -> 8080
```

Now, we can send requests to the deployed HelloWorld service from our local machine.

```sh
curl localhost:8081/HelloWorld
```

The deployed HelloWorld pod will log as follows.

```sh
➜  ~ kubectl logs -f `kubectl get pods | grep ^example-helloworld | awk '{print $1}'`
2020/05/07 11:07:37 START THE HELLO WORLD SERVER..!!
2020/05/07 11:09:26 Received message: HelloWorld
```

## References

1. https://github.com/operator-framework/operator-sdk/blob/v0.17.0/website/content/en/docs/golang/quickstart.md
1. https://github.com/siddhi-io/siddhi-operator