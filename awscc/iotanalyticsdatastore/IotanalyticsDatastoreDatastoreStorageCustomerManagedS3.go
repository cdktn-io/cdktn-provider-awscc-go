// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticsdatastore


type IotanalyticsDatastoreDatastoreStorageCustomerManagedS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotanalytics_datastore#bucket IotanalyticsDatastore#bucket}.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotanalytics_datastore#key_prefix IotanalyticsDatastore#key_prefix}.
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotanalytics_datastore#role_arn IotanalyticsDatastore#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

