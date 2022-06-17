go9p
====

Fork of the original rminnich/go9p repo with the patches of minikube



This is go9p done in a way that I can understand.

To install:

```
export GOPATH=~/go
go get -a /github.com/gbraad/go9p
go get -a /github.com/gbraad/go9p/ufs
go install -a /github.com/gbraad/go9p/ufs       <!-- which won't work as main got renamed to StartServer for minikube

~/go/bin/ufs
```