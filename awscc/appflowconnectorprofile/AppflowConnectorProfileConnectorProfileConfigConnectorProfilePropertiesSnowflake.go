// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflake struct {
	// The name of the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appflow_connector_profile#account_name AppflowConnectorProfile#account_name}
	AccountName *string `field:"optional" json:"accountName" yaml:"accountName"`
	// The name of the Amazon S3 bucket associated with Snow?ake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appflow_connector_profile#bucket_name AppflowConnectorProfile#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The bucket prefix that refers to the Amazon S3 bucket associated with Snow?ake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appflow_connector_profile#bucket_prefix AppflowConnectorProfile#bucket_prefix}
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// The Snow?ake Private Link service name to be used for private data transfers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appflow_connector_profile#private_link_service_name AppflowConnectorProfile#private_link_service_name}
	PrivateLinkServiceName *string `field:"optional" json:"privateLinkServiceName" yaml:"privateLinkServiceName"`
	// The region of the Snow?ake account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appflow_connector_profile#region AppflowConnectorProfile#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The name of the Amazon S3 stage that was created while setting up an Amazon S3 stage in the Snow?ake account. This is written in the following format: < Database>< Schema><Stage Name>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appflow_connector_profile#stage AppflowConnectorProfile#stage}
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
	// The name of the Snow?ake warehouse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appflow_connector_profile#warehouse AppflowConnectorProfile#warehouse}
	Warehouse *string `field:"optional" json:"warehouse" yaml:"warehouse"`
}

