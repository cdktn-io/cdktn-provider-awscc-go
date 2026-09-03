// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticsdataset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotanalyticsDatasetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotanalytics_dataset#actions IotanalyticsDataset#actions}.
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotanalytics_dataset#content_delivery_rules IotanalyticsDataset#content_delivery_rules}.
	ContentDeliveryRules interface{} `field:"optional" json:"contentDeliveryRules" yaml:"contentDeliveryRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotanalytics_dataset#dataset_name IotanalyticsDataset#dataset_name}.
	DatasetName *string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotanalytics_dataset#late_data_rules IotanalyticsDataset#late_data_rules}.
	LateDataRules interface{} `field:"optional" json:"lateDataRules" yaml:"lateDataRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotanalytics_dataset#retention_period IotanalyticsDataset#retention_period}.
	RetentionPeriod *IotanalyticsDatasetRetentionPeriod `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotanalytics_dataset#tags IotanalyticsDataset#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotanalytics_dataset#triggers IotanalyticsDataset#triggers}.
	Triggers interface{} `field:"optional" json:"triggers" yaml:"triggers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotanalytics_dataset#versioning_configuration IotanalyticsDataset#versioning_configuration}.
	VersioningConfiguration *IotanalyticsDatasetVersioningConfiguration `field:"optional" json:"versioningConfiguration" yaml:"versioningConfiguration"`
}

