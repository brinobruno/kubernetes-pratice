# Notes

### Pod:

The smallest object in k8s, a machine to run our container(s) with

General rule of thumb: 1 container for/per 1 pod

#### Running a pod through a file config

input:
`kubectl apply -f k8s/pod.yaml`

expected successful output:

`pod/goserver created`

checking
`kubectl get pods`
