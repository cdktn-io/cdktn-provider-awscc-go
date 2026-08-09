// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticsdatastore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotanalyticsDatastoreConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotanalytics_datastore#datastore_name IotanalyticsDatastore#datastore_name}.
	DatastoreName *string `field:"optional" json:"datastoreName" yaml:"datastoreName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotanalytics_datastore#datastore_partitions IotanalyticsDatastore#datastore_partitions}.
	DatastorePartitions *IotanalyticsDatastoreDatastorePartitions `field:"optional" json:"datastorePartitions" yaml:"datastorePartitions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotanalytics_datastore#datastore_storage IotanalyticsDatastore#datastore_storage}.
	DatastoreStorage *IotanalyticsDatastoreDatastoreStorage `field:"optional" json:"datastoreStorage" yaml:"datastoreStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotanalytics_datastore#file_format_configuration IotanalyticsDatastore#file_format_configuration}.
	FileFormatConfiguration *IotanalyticsDatastoreFileFormatConfiguration `field:"optional" json:"fileFormatConfiguration" yaml:"fileFormatConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotanalytics_datastore#retention_period IotanalyticsDatastore#retention_period}.
	RetentionPeriod *IotanalyticsDatastoreRetentionPeriod `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotanalytics_datastore#tags IotanalyticsDatastore#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

