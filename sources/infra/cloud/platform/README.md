# Cloud Deployment

Cloud deployment is managed using `terraform`. AWS is used as the
cloud provider. The configuration creates the following resources:

* AWS VPC: Base VPC used for AWS EKS, ECR and Amazon MQ
* AWS EKS: Kubernetes cluster for application deployment
* Amazon MQ: Used for managed RabbitMQ
* IAM Roles: IAM roles for CI, Amazon MQ access, etc.

## Usage

First, initialize the terraform modules:

```sh
terraform init
```

Validate the plan:

```sh
terraform apply
```

If everything looks good above, type `yes` at the prompt to begin
deployment. The output of the deployment will the kubernetes cluster
details which can be used to connect to the cluster.

## Connecting to Kubernetes

You can use the command below to update your `kubectl` configuration
and authenticate with EKS:

```sh
aws eks \
    --region $(terraform output -raw region) \
    update-kubeconfig \
    --name $(terraform output -raw cluster_name)
```

If the command above runs successfully, `kubectl` should be
able to connect to the cluster.
