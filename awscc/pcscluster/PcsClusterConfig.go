// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcscluster

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PcsClusterConfig struct {
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
	// The networking configuration for the cluster's control plane.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcs_cluster#networking PcsCluster#networking}
	Networking *PcsClusterNetworking `field:"required" json:"networking" yaml:"networking"`
	// The cluster management and job scheduling software associated with the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcs_cluster#scheduler PcsCluster#scheduler}
	Scheduler *PcsClusterScheduler `field:"required" json:"scheduler" yaml:"scheduler"`
	// The size of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcs_cluster#size PcsCluster#size}
	Size *string `field:"required" json:"size" yaml:"size"`
	// The name that identifies the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcs_cluster#name PcsCluster#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Additional options related to the Slurm scheduler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcs_cluster#slurm_configuration PcsCluster#slurm_configuration}
	SlurmConfiguration *PcsClusterSlurmConfiguration `field:"optional" json:"slurmConfiguration" yaml:"slurmConfiguration"`
	// 1 or more tags added to the resource.
	//
	// Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcs_cluster#tags PcsCluster#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

