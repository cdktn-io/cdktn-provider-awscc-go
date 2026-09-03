// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rekognitiondataset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RekognitionDatasetConfig struct {
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
	// The type of the dataset. Specify TRAIN to create a training dataset. Specify TEST to create a test dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/rekognition_dataset#dataset_type RekognitionDataset#dataset_type}
	DatasetType *string `field:"required" json:"datasetType" yaml:"datasetType"`
	// The ARN of the project to which the dataset belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/rekognition_dataset#project_arn RekognitionDataset#project_arn}
	ProjectArn *string `field:"optional" json:"projectArn" yaml:"projectArn"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/rekognition_dataset#tags RekognitionDataset#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

