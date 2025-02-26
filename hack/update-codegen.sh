#!/usr/bin/env bash

# Copyright 2023 The flock Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

source $(dirname $0)/../vendor/knative.dev/hack/codegen-library.sh
export PATH="$GOBIN:$PATH"
K8S_CODEGEN="./vendor/k8s.io/code-generator/cmd"

echo "=== Update Codegen for ${MODULE_NAME}"

group "Kubernetes Codegen"

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.

group "ack-ecr-controller codegen"

${CODEGEN_PKG}/generate-groups.sh "client,informer,lister" \
              github.com/nubank/cicd-exp-controller/pkg/ecr-controller github.com/aws-controllers-k8s/ecr-controller/apis \
              "rollouts:v1alpha1" \
              --go-header-file ${REPO_ROOT_DIR}/hack/boilerplate/boilerplate.go.txt

# Generate deep copy functions for other packages.
go run ${K8S_CODEGEN}/deepcopy-gen/main.go \
   -O zz_generated.deepcopy \
   --go-header-file ${REPO_ROOT_DIR}/hack/boilerplate/boilerplate.go.txt \
   --input-dirs $(echo \
                      github.com/nubank/flock/pkg/apis/config \
                      | sed "s/ /,/g")


group "Update Open API schemas"
${REPO_ROOT_DIR}/hack/update-openapigen.sh

group "Update deps post-codegen"

# Make sure our dependencies are up-to-date
${REPO_ROOT_DIR}/hack/update-deps.sh