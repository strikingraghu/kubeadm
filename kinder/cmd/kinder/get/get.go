/*
Copyright 2018 The Kubernetes Authors.

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

// Package get implements the `get` command
package get

import (
	"github.com/spf13/cobra"

	"sigs.k8s.io/kind/cmd/kind/get/clusters"
	"sigs.k8s.io/kind/cmd/kind/get/kubeconfigpath"
	"sigs.k8s.io/kind/cmd/kind/get/nodes"

	kartifacts "k8s.io/kubeadm/kinder/cmd/kinder/get/artifacts"
)

// NewCommand returns a new cobra.Command for get
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args: cobra.NoArgs,
		// TODO(bentheelder): more detailed usage
		Use:   "get",
		Short: "Gets one of [clusters, nodes, kubeconfig-path, build-artifacts]",
		Long:  "Gets one of [clusters, nodes, kubeconfig-path, build-artifacts]",
	}

	// add kind top level subcommands re-used without changes
	cmd.AddCommand(clusters.NewCommand())
	cmd.AddCommand(nodes.NewCommand())
	cmd.AddCommand(kubeconfigpath.NewCommand())

	// add kinder only commands
	cmd.AddCommand(kartifacts.NewCommand())
	return cmd
}
