// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databrewdataset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatabrewDatasetConfig struct {
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
	// Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/databrew_dataset#input DatabrewDataset#input}
	Input *DatabrewDatasetInput `field:"required" json:"input" yaml:"input"`
	// Dataset name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/databrew_dataset#name DatabrewDataset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Dataset format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/databrew_dataset#format DatabrewDataset#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Format options for dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/databrew_dataset#format_options DatabrewDataset#format_options}
	FormatOptions *DatabrewDatasetFormatOptions `field:"optional" json:"formatOptions" yaml:"formatOptions"`
	// PathOptions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/databrew_dataset#path_options DatabrewDataset#path_options}
	PathOptions *DatabrewDatasetPathOptions `field:"optional" json:"pathOptions" yaml:"pathOptions"`
	// Source type of the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/databrew_dataset#source DatabrewDataset#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/databrew_dataset#tags DatabrewDataset#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

