# CICD-EXP-CONTROLLER

## Setup

### Pre-requisites
* [ko](https://ko.build/install/)
    ```bash
    brew install ko
    ```

* [kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installing-with-a-package-manager)

    ```bash
    brew install kind
    kind create cluster --name cicd-exp
    ```

* [helm](https://helm.sh/docs/intro/install/)

    ```bash
    brew install helm
    ```
* [kustomize](https://kubectl.docs.kubernetes.io/installation/kustomize/)

    ```bash
    brew install kustomize
    ```

* [ack-controller](https://aws-controllers-k8s.github.io/community/docs/user-docs/install/)

    ```bash
    export SERVICE=ecr
    export RELEASE_VERSION=$(curl -sL https://api.github.com/repos/aws-controllers-k8s/${SERVICE}-controller/releases/latest | jq -r '.tag_name | ltrimstr("v")')
    export ACK_SYSTEM_NAMESPACE=ack-system
    export AWS_REGION=us-west-2

    aws ecr-public get-login-password --region us-east-1 | helm registry login --username AWS --password-stdin public.ecr.aws
    helm install --create-namespace -n $ACK_SYSTEM_NAMESPACE ack-$SERVICE-controller \
    oci://public.ecr.aws/aws-controllers-k8s/$SERVICE-chart --version=$RELEASE_VERSION --set=aws.region=$AWS_REGION
    ```

    Repository example:

    ```yaml
    apiVersion: ecr.services.k8s.aws/v1alpha1
    kind: Repository
    metadata:
      name: test-cenoura
      namespace: default
    spec:
      imageScanningConfiguration:
        scanOnPush: true
      imageTagMutability: MUTABLE
      name: cenoura
    ```
