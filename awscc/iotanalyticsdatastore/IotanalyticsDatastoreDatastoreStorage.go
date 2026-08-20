// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticsdatastore


type IotanalyticsDatastoreDatastoreStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotanalytics_datastore#customer_managed_s3 IotanalyticsDatastore#customer_managed_s3}.
	CustomerManagedS3 *IotanalyticsDatastoreDatastoreStorageCustomerManagedS3 `field:"optional" json:"customerManagedS3" yaml:"customerManagedS3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotanalytics_datastore#iot_site_wise_multi_layer_storage IotanalyticsDatastore#iot_site_wise_multi_layer_storage}.
	IotSiteWiseMultiLayerStorage *IotanalyticsDatastoreDatastoreStorageIotSiteWiseMultiLayerStorage `field:"optional" json:"iotSiteWiseMultiLayerStorage" yaml:"iotSiteWiseMultiLayerStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotanalytics_datastore#service_managed_s3 IotanalyticsDatastore#service_managed_s3}.
	ServiceManagedS3 *string `field:"optional" json:"serviceManagedS3" yaml:"serviceManagedS3"`
}

