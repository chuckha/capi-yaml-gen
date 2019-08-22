/*
Copyright 2019 The Kubernetes Authors.

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

package capi

import (
	"fmt"

	"encoding/json"

	v1 "k8s.io/api/core/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha2"
)

const (
	coreAPIVersion  = "cluster.x-k8s.io/v1alpha2"
	infraAPIVersion = "infrastructure.cluster.x-k8s.io/v1alpha2"

	prettyJSONPrefix = ""
	prettyJSONIndent = "	" // tab
)

// GetCoreClusterYaml returns yaml for CAPI  cluster objects
func GetCoreClusterYaml(name, namespace, infraProvider string) (string, error) {
	coreCluster := clusterv1.Cluster{}
	coreCluster.Name = name
	coreCluster.Namespace = namespace
	coreCluster.APIVersion = coreAPIVersion

	coreCluster.Spec = clusterv1.ClusterSpec{
		InfrastructureRef: &v1.ObjectReference{
			Kind:       fmt.Sprintf("%sCluster", infraProvider),
			APIVersion: infraAPIVersion,
			Name:       name,
			Namespace:  namespace,
		},
	}

	rawBytes, err := json.MarshalIndent(coreCluster, prettyJSONPrefix, prettyJSONIndent)
	if err != nil {
		fmt.Printf("Failed to generate yaml for coreCluster %q, %v", name, err)
		return "", err
	}

	return string(rawBytes), nil
}
