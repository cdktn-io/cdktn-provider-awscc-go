// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scndataset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ScnDatasetConfig struct {
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
	// The Amazon Web Services Supply Chain instance identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/scn_dataset#instance_id ScnDataset#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The name of the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/scn_dataset#name ScnDataset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/scn_dataset#namespace ScnDataset#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The description of the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/scn_dataset#description ScnDataset#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The partition specification of the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/scn_dataset#partition_spec ScnDataset#partition_spec}
	PartitionSpec *ScnDatasetPartitionSpec `field:"optional" json:"partitionSpec" yaml:"partitionSpec"`
	// The schema of the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/scn_dataset#schema ScnDataset#schema}
	Schema *ScnDatasetSchema `field:"optional" json:"schema" yaml:"schema"`
	// The tags for the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/scn_dataset#tags ScnDataset#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

