/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package v1 contains API Schema definitions for the optimize.stormforge.io v1 API group.
// +groupName=optimize.stormforge.io
package v1

// +k8s:deepcopy-gen=package

//go:generate go run k8s.io/code-generator/cmd/deepcopy-gen@latest --go-header-file ../../../../hack/boilerplate.go.txt --output-file zz_generated.deepcopy.go .
//go:generate go run k8s.io/code-generator/cmd/applyconfiguration-gen@latest --go-header-file ../../../../hack/boilerplate.go.txt --output-dir=../../../applyconfigurations --output-pkg=github.com/thestormforge/agent-go/pkg/applyconfigurations .
//go:generate go run k8s.io/code-generator/cmd/client-gen@latest --go-header-file ../../../../hack/boilerplate.go.txt --output-dir=../../.. --clientset-name stormforge --output-pkg=github.com/thestormforge/agent-go/pkg --input-base=github.com/thestormforge/agent-go/pkg/api --input optimize/v1 --apply-configuration-package=github.com/thestormforge/agent-go/pkg/applyconfigurations .
