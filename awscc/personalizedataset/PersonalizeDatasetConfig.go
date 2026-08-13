// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizedataset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PersonalizeDatasetConfig struct {
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
	// The Amazon Resource Name (ARN) of the dataset group to add the dataset to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/personalize_dataset#dataset_group_arn PersonalizeDataset#dataset_group_arn}
	DatasetGroupArn *string `field:"required" json:"datasetGroupArn" yaml:"datasetGroupArn"`
	// The type of dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/personalize_dataset#dataset_type PersonalizeDataset#dataset_type}
	DatasetType *string `field:"required" json:"datasetType" yaml:"datasetType"`
	// The name for the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/personalize_dataset#name PersonalizeDataset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the schema to associate with the dataset. The schema defines the dataset fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/personalize_dataset#schema_arn PersonalizeDataset#schema_arn}
	SchemaArn *string `field:"required" json:"schemaArn" yaml:"schemaArn"`
	// Initial DatasetImportJob for the created dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/personalize_dataset#dataset_import_job PersonalizeDataset#dataset_import_job}
	DatasetImportJob *PersonalizeDatasetDatasetImportJob `field:"optional" json:"datasetImportJob" yaml:"datasetImportJob"`
}

