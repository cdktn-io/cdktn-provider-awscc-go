// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticsdatastore


type IotanalyticsDatastoreDatastorePartitionsPartitions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotanalytics_datastore#partition IotanalyticsDatastore#partition}.
	Partition *IotanalyticsDatastoreDatastorePartitionsPartitionsPartition `field:"optional" json:"partition" yaml:"partition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotanalytics_datastore#timestamp_partition IotanalyticsDatastore#timestamp_partition}.
	TimestampPartition *IotanalyticsDatastoreDatastorePartitionsPartitionsTimestampPartition `field:"optional" json:"timestampPartition" yaml:"timestampPartition"`
}

