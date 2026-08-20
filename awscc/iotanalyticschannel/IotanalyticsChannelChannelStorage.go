// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticschannel


type IotanalyticsChannelChannelStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotanalytics_channel#customer_managed_s3 IotanalyticsChannel#customer_managed_s3}.
	CustomerManagedS3 *IotanalyticsChannelChannelStorageCustomerManagedS3 `field:"optional" json:"customerManagedS3" yaml:"customerManagedS3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotanalytics_channel#service_managed_s3 IotanalyticsChannel#service_managed_s3}.
	ServiceManagedS3 *string `field:"optional" json:"serviceManagedS3" yaml:"serviceManagedS3"`
}

