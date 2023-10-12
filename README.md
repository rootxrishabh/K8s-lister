# K8s-Lister

This is a kubernetes client-go library project which can is used to query about k8s resources such as pods or deployments.

## Steps to deploy this app inside kubernetes cluster

1) Use the compiled binary and the given Dockerfile to generate a docker image.

2) Now, push to docker image to your container registery such as docker hub.

3) Start a kubernetes cluster using any k8s single-node distribution. Eg - kind

4) Start this application as a pod.
`kubectl run clientgoapp --image=[IMAGE FROM YOUR CONTAINER REGISTERY]`. You may also run this as a deployment.

5) Assign role to the default service account in order for the apps to list resources.

`kubectl create role poddeploy --resources pods, deployments --verb list`

`kubectl create rolebinding poddeploy --role poddeploy --serviceaccount default:default`

6) `kubectl logs [POD-NAME]` to get the results from the mod.