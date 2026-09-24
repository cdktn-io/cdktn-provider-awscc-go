// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewisedataset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotsitewiseDatasetConfig struct {
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
	// The name of the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_dataset#dataset_name IotsitewiseDataset#dataset_name}
	DatasetName *string `field:"required" json:"datasetName" yaml:"datasetName"`
	// The configuration for the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_dataset#dataset_config IotsitewiseDataset#dataset_config}
	DatasetConfig *IotsitewiseDatasetDatasetConfig `field:"optional" json:"datasetConfig" yaml:"datasetConfig"`
	// A description about the dataset, and its functionality.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_dataset#dataset_description IotsitewiseDataset#dataset_description}
	DatasetDescription *string `field:"optional" json:"datasetDescription" yaml:"datasetDescription"`
	// The data source for the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_dataset#dataset_source IotsitewiseDataset#dataset_source}
	DatasetSource *IotsitewiseDatasetDatasetSource `field:"optional" json:"datasetSource" yaml:"datasetSource"`
	// The type of the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_dataset#dataset_type IotsitewiseDataset#dataset_type}
	DatasetType *string `field:"optional" json:"datasetType" yaml:"datasetType"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_dataset#tags IotsitewiseDataset#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The name of the workspace associated with the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_dataset#workspace_name IotsitewiseDataset#workspace_name}
	WorkspaceName *string `field:"optional" json:"workspaceName" yaml:"workspaceName"`
}

