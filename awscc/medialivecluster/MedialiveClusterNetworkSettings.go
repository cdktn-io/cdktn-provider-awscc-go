// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialivecluster


type MedialiveClusterNetworkSettings struct {
	// Default value if the customer does not define it in channel Output API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/medialive_cluster#default_route MedialiveCluster#default_route}
	DefaultRoute *string `field:"optional" json:"defaultRoute" yaml:"defaultRoute"`
	// Network mappings for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/medialive_cluster#interface_mappings MedialiveCluster#interface_mappings}
	InterfaceMappings interface{} `field:"optional" json:"interfaceMappings" yaml:"interfaceMappings"`
}

