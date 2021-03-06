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

package alpha

import (
	"fmt"

	"github.com/ashish-amarnath/capiyaml/cmd/alpha/capi"
	"github.com/spf13/cobra"
)

// Cmd returns the set alpha of commands
func Cmd() *cobra.Command {
	alphaCmd := &cobra.Command{
		Use:   "alpha",
		Short: "capi-yaml-gen alpha command set",
		Long:  "capi-yaml-gen alpha command set",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	alphaCmd.AddCommand(getClusterYamlCommand())
	return alphaCmd
}

func getClusterYamlCommand() *cobra.Command {
	var provider, clusterName, clusterNamespace string

	cmd := &cobra.Command{
		Use:   "get-cluster-yaml",
		Short: "generate yaml for cluster",
		Long:  "generate yaml for cluster",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Genarating yaml for cluster %q for infrastructure provider %q\n", clusterName, provider)
			coreClusterYaml, err := capi.GetCoreClusterYaml(clusterName, clusterNamespace, provider)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			fmt.Printf("%s\n", coreClusterYaml)
		},
	}

	cmd.Flags().StringVarP(&provider, "provider", "p", "", "Infrastructure provider for which yaml needs to be generated")
	cmd.Flags().StringVarP(&clusterName, "name", "n", "", "Name for the cluster")
	cmd.Flags().StringVarP(&clusterNamespace, "namespace", "", "default", "Namespace where the cluster will be created")

	return cmd
}
