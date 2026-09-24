// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialivecluster

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MedialiveClusterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The hardware type for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/medialive_cluster#cluster_type MedialiveCluster#cluster_type}
	ClusterType *string `field:"optional" json:"clusterType" yaml:"clusterType"`
	// The IAM role your nodes will use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/medialive_cluster#instance_role_arn MedialiveCluster#instance_role_arn}
	InstanceRoleArn *string `field:"optional" json:"instanceRoleArn" yaml:"instanceRoleArn"`
	// The user-specified name of the Cluster to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/medialive_cluster#name MedialiveCluster#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// On premises settings which will have the interface network mappings and default Output logical interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/medialive_cluster#network_settings MedialiveCluster#network_settings}
	NetworkSettings *MedialiveClusterNetworkSettings `field:"optional" json:"networkSettings" yaml:"networkSettings"`
	// A collection of key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/medialive_cluster#tags MedialiveCluster#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

