// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package forecastdatasetgroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ForecastDatasetGroupConfig struct {
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
	// A name for the dataset group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/forecast_dataset_group#dataset_group_name ForecastDatasetGroup#dataset_group_name}
	DatasetGroupName *string `field:"required" json:"datasetGroupName" yaml:"datasetGroupName"`
	// The domain associated with the dataset group.
	//
	// When you add a dataset to a dataset group, this value and the value specified for the Domain parameter of the CreateDataset operation must match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/forecast_dataset_group#domain ForecastDatasetGroup#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// An array of Amazon Resource Names (ARNs) of the datasets that you want to include in the dataset group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/forecast_dataset_group#dataset_arns ForecastDatasetGroup#dataset_arns}
	DatasetArns *[]*string `field:"optional" json:"datasetArns" yaml:"datasetArns"`
	// The tags of Application Insights application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/forecast_dataset_group#tags ForecastDatasetGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

