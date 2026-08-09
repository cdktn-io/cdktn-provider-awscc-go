// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbodbnetwork

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OdbOdbNetworkConfig struct {
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
	// The AWS Availability Zone (AZ) where the ODB network is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/odb_odb_network#availability_zone OdbOdbNetwork#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The AZ ID of the AZ where the ODB network is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/odb_odb_network#availability_zone_id OdbOdbNetwork#availability_zone_id}
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// The CIDR range of the backup subnet in the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/odb_odb_network#backup_subnet_cidr OdbOdbNetwork#backup_subnet_cidr}
	BackupSubnetCidr *string `field:"optional" json:"backupSubnetCidr" yaml:"backupSubnetCidr"`
	// The CIDR range of the client subnet in the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/odb_odb_network#client_subnet_cidr OdbOdbNetwork#client_subnet_cidr}
	ClientSubnetCidr *string `field:"optional" json:"clientSubnetCidr" yaml:"clientSubnetCidr"`
	// The cross-Region Amazon S3 restore sources for the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/odb_odb_network#cross_region_s3_restore_sources OdbOdbNetwork#cross_region_s3_restore_sources}
	CrossRegionS3RestoreSources *[]*string `field:"optional" json:"crossRegionS3RestoreSources" yaml:"crossRegionS3RestoreSources"`
	// The domain name to use for the resources in the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/odb_odb_network#custom_domain_name OdbOdbNetwork#custom_domain_name}
	CustomDomainName *string `field:"optional" json:"customDomainName" yaml:"customDomainName"`
	// The DNS prefix to the default DNS domain name. The default DNS domain name is oraclevcn.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/odb_odb_network#default_dns_prefix OdbOdbNetwork#default_dns_prefix}
	DefaultDnsPrefix *string `field:"optional" json:"defaultDnsPrefix" yaml:"defaultDnsPrefix"`
	// Specifies whether to delete associated OCI networking resources along with the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/odb_odb_network#delete_associated_resources OdbOdbNetwork#delete_associated_resources}
	DeleteAssociatedResources interface{} `field:"optional" json:"deleteAssociatedResources" yaml:"deleteAssociatedResources"`
	// The user-friendly name of the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/odb_odb_network#display_name OdbOdbNetwork#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The AWS Key Management Service (KMS) access configuration for the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/odb_odb_network#kms_access OdbOdbNetwork#kms_access}
	KmsAccess *string `field:"optional" json:"kmsAccess" yaml:"kmsAccess"`
	// The AWS Key Management Service (KMS) policy document that defines permissions for key usage within the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/odb_odb_network#kms_policy_document OdbOdbNetwork#kms_policy_document}
	KmsPolicyDocument *string `field:"optional" json:"kmsPolicyDocument" yaml:"kmsPolicyDocument"`
	// Specifies the configuration for Amazon S3 access from the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/odb_odb_network#s3_access OdbOdbNetwork#s3_access}
	S3Access *string `field:"optional" json:"s3Access" yaml:"s3Access"`
	// Specifies the endpoint policy for Amazon S3 access from the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/odb_odb_network#s3_policy_document OdbOdbNetwork#s3_policy_document}
	S3PolicyDocument *string `field:"optional" json:"s3PolicyDocument" yaml:"s3PolicyDocument"`
	// The AWS Security Token Service (STS) access configuration for the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/odb_odb_network#sts_access OdbOdbNetwork#sts_access}
	StsAccess *string `field:"optional" json:"stsAccess" yaml:"stsAccess"`
	// The AWS Security Token Service (STS) policy document that defines permissions for token service usage within the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/odb_odb_network#sts_policy_document OdbOdbNetwork#sts_policy_document}
	StsPolicyDocument *string `field:"optional" json:"stsPolicyDocument" yaml:"stsPolicyDocument"`
	// Tags to assign to the Odb Network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/odb_odb_network#tags OdbOdbNetwork#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the configuration for Zero-ETL access from the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/odb_odb_network#zero_etl_access OdbOdbNetwork#zero_etl_access}
	ZeroEtlAccess *string `field:"optional" json:"zeroEtlAccess" yaml:"zeroEtlAccess"`
}

