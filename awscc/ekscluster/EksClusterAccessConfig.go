// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterAccessConfig struct {
	// Specify the authentication mode that should be used to create your cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#authentication_mode EksCluster#authentication_mode}
	AuthenticationMode *string `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// Set this value to false to avoid creating a default cluster admin Access Entry using the IAM principal used to create the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#bootstrap_cluster_creator_admin_permissions EksCluster#bootstrap_cluster_creator_admin_permissions}
	BootstrapClusterCreatorAdminPermissions interface{} `field:"optional" json:"bootstrapClusterCreatorAdminPermissions" yaml:"bootstrapClusterCreatorAdminPermissions"`
}

