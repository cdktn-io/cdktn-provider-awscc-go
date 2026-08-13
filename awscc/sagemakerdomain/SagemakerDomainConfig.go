// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerDomainConfig struct {
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
	// The mode of authentication that members use to access the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#auth_mode SagemakerDomain#auth_mode}
	AuthMode *string `field:"required" json:"authMode" yaml:"authMode"`
	// The default user settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#default_user_settings SagemakerDomain#default_user_settings}
	DefaultUserSettings *SagemakerDomainDefaultUserSettings `field:"required" json:"defaultUserSettings" yaml:"defaultUserSettings"`
	// A name for the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#domain_name SagemakerDomain#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Specifies the VPC used for non-EFS traffic. The default value is PublicInternetOnly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#app_network_access_type SagemakerDomain#app_network_access_type}
	AppNetworkAccessType *string `field:"optional" json:"appNetworkAccessType" yaml:"appNetworkAccessType"`
	// The entity that creates and manages the required security groups for inter-app communication in VPCOnly mode.
	//
	// Required when CreateDomain.AppNetworkAccessType is VPCOnly and DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn is provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#app_security_group_management SagemakerDomain#app_security_group_management}
	AppSecurityGroupManagement *string `field:"optional" json:"appSecurityGroupManagement" yaml:"appSecurityGroupManagement"`
	// The default space settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#default_space_settings SagemakerDomain#default_space_settings}
	DefaultSpaceSettings *SagemakerDomainDefaultSpaceSettings `field:"optional" json:"defaultSpaceSettings" yaml:"defaultSpaceSettings"`
	// A collection of Domain settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#domain_settings SagemakerDomain#domain_settings}
	DomainSettings *SagemakerDomainDomainSettings `field:"optional" json:"domainSettings" yaml:"domainSettings"`
	// Indicates whether a home EFS file system is created for the domain.
	//
	// Set to Disabled to skip EFS creation and reduce domain creation time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#home_efs_file_system_creation SagemakerDomain#home_efs_file_system_creation}
	HomeEfsFileSystemCreation *string `field:"optional" json:"homeEfsFileSystemCreation" yaml:"homeEfsFileSystemCreation"`
	// SageMaker uses AWS KMS to encrypt the EFS volume attached to the domain with an AWS managed customer master key (CMK) by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#kms_key_id SagemakerDomain#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The VPC subnets that Studio uses for communication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#subnet_ids SagemakerDomain#subnet_ids}
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// Indicates whether the tags added to Domain, User Profile and Space entity is propagated to all SageMaker resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#tag_propagation SagemakerDomain#tag_propagation}
	TagPropagation *string `field:"optional" json:"tagPropagation" yaml:"tagPropagation"`
	// A list of tags to apply to the user profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#tags SagemakerDomain#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#vpc_id SagemakerDomain#vpc_id}
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

