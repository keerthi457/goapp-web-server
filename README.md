# goapp-web-server
# DevOps, Microservices

The project's goal is to deploy a microservice in a highly available Kubernetes cluster.

The microservice I choose is a simple web server application  that is containerized and deployed in a Kubernetes cluster through a Jenkins pipeline. I ensured Node affinity, Pod Disruption budget, offload TLS at the load balancer/ingress are working before it is served with traffic.


## To Recreate this project

You will need

- [Jenkins](https://www.jenkins.io/doc/book/installing/)
- [Docker](https://docs.docker.com/engine/install/ubuntu/)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [aws-cli](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2-linux.html)
- [eksctl](https://docs.aws.amazon.com/eks/latest/userguide/getting-started-eksctl.html)

Installed and configured using your credentials and docker being a Sudoer, add this repository to your Jenkins pipeline and the deployment is automated in Jenkins


## Deployment strategy

I preferred blue/green deployment strategy, as it reduces downtime by running two identical environments with only one serving at a given time


These are two pipelines I have created

- [cluster formation](https://github.com/keerthi457/jenkins-to-create-eks) - creates a Kubernetes cluster using eksctl and updates the Jenkins server with the respective config file.


- [goapp-web-server](https://github.com/keerthi457/goapp-web-server)(This is the project)




This pipeline does the following steps

1. builds and pushes the docker container
2. sets labels to nodes so as to use them in node affinity
3. deploys blue and green deployments
4. installs metrics server and uses it in horizontal pod autoscaling
5. adds pod disruption budget to both blue and green deployments
6. exposes the deployments using a load balancer port 8080

## Note
To use ingress, expose the application through NodePort and Run

- Kubectl apply -f ingress.yml

I haven't used it as I do not have a registered domain name.

Also, refer to images of deployment for a better understanding


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
