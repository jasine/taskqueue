# Introduction
 TaskQueue is an implement of k8s CRD, manage k8s's Job object in a Cron Queue

# Development
## Code-Generation
To implement a k8s custumer controller, we use code-generation tools. More details about code-generation is [here](https://blog.openshift.com/kubernetes-deep-dive-code-generation-customresources/).
### Install code-generation
```shell
git clone  https://github.com/kubernetes/code-generator.git
git checkout -b remotes/origin/release-1.8 # for running in kubernetes 1.8+
go install ./cmd/{defaulter-gen,client-gen,lister-gen,informer-gen,deepcopy-gen}
```
### Init porject
```
mkdir ~/go/src/github.com/taskqueue && cd ~/go/src/github.com/taskqueue 
mkdir -p pkg/apis/taskqueue/v1alpha1
```
Create doc.go, types.go and register.go here. We create global tags in [doc.go] and local tags in types.go. In register.go,we regist k8s's scheme and api.

### Generation my code
use update-codegen.sh
```
cd hack
./update.codegen.sh
```

