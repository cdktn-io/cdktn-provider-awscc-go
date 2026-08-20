// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emrstudio

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EmrStudioConfig struct {
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
	// Specifies whether the Studio authenticates users using single sign-on (SSO) or IAM.
	//
	// Amazon EMR Studio currently only supports SSO authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/emr_studio#auth_mode EmrStudio#auth_mode}
	AuthMode *string `field:"required" json:"authMode" yaml:"authMode"`
	// The default Amazon S3 location to back up EMR Studio Workspaces and notebook files.
	//
	// A Studio user can select an alternative Amazon S3 location when creating a Workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/emr_studio#default_s3_location EmrStudio#default_s3_location}
	DefaultS3Location *string `field:"required" json:"defaultS3Location" yaml:"defaultS3Location"`
	// The ID of the Amazon EMR Studio Engine security group.
	//
	// The Engine security group allows inbound network traffic from the Workspace security group, and it must be in the same VPC specified by VpcId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/emr_studio#engine_security_group_id EmrStudio#engine_security_group_id}
	EngineSecurityGroupId *string `field:"required" json:"engineSecurityGroupId" yaml:"engineSecurityGroupId"`
	// A descriptive name for the Amazon EMR Studio.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/emr_studio#name EmrStudio#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The IAM role that will be assumed by the Amazon EMR Studio.
	//
	// The service role provides a way for Amazon EMR Studio to interoperate with other AWS services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/emr_studio#service_role EmrStudio#service_role}
	ServiceRole *string `field:"required" json:"serviceRole" yaml:"serviceRole"`
	// A list of up to 5 subnet IDs to associate with the Studio.
	//
	// The subnets must belong to the VPC specified by VpcId. Studio users can create a Workspace in any of the specified subnets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/emr_studio#subnet_ids EmrStudio#subnet_ids}
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the Amazon Virtual Private Cloud (Amazon VPC) to associate with the Studio.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/emr_studio#vpc_id EmrStudio#vpc_id}
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The ID of the Amazon EMR Studio Workspace security group.
	//
	// The Workspace security group allows outbound network traffic to resources in the Engine security group, and it must be in the same VPC specified by VpcId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/emr_studio#workspace_security_group_id EmrStudio#workspace_security_group_id}
	WorkspaceSecurityGroupId *string `field:"required" json:"workspaceSecurityGroupId" yaml:"workspaceSecurityGroupId"`
	// A detailed description of the Studio.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/emr_studio#description EmrStudio#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The AWS KMS key identifier (ARN) used to encrypt AWS EMR Studio workspace and notebook files when backed up to AWS S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/emr_studio#encryption_key_arn EmrStudio#encryption_key_arn}
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// The ARN of the IAM Identity Center instance to create the Studio application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/emr_studio#idc_instance_arn EmrStudio#idc_instance_arn}
	IdcInstanceArn *string `field:"optional" json:"idcInstanceArn" yaml:"idcInstanceArn"`
	// Specifies whether IAM Identity Center user assignment is REQUIRED or OPTIONAL.
	//
	// If the value is set to REQUIRED, users must be explicitly assigned to the Studio application to access the Studio.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/emr_studio#idc_user_assignment EmrStudio#idc_user_assignment}
	IdcUserAssignment *string `field:"optional" json:"idcUserAssignment" yaml:"idcUserAssignment"`
	// Your identity provider's authentication endpoint.
	//
	// Amazon EMR Studio redirects federated users to this endpoint for authentication when logging in to a Studio with the Studio URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/emr_studio#idp_auth_url EmrStudio#idp_auth_url}
	IdpAuthUrl *string `field:"optional" json:"idpAuthUrl" yaml:"idpAuthUrl"`
	// The name of relay state parameter for external Identity Provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/emr_studio#idp_relay_state_parameter_name EmrStudio#idp_relay_state_parameter_name}
	IdpRelayStateParameterName *string `field:"optional" json:"idpRelayStateParameterName" yaml:"idpRelayStateParameterName"`
	// A list of tags to associate with the Studio.
	//
	// Tags are user-defined key-value pairs that consist of a required key string with a maximum of 128 characters, and an optional value string with a maximum of 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/emr_studio#tags EmrStudio#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A Boolean indicating whether to enable Trusted identity propagation for the Studio. The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/emr_studio#trusted_identity_propagation_enabled EmrStudio#trusted_identity_propagation_enabled}
	TrustedIdentityPropagationEnabled interface{} `field:"optional" json:"trustedIdentityPropagationEnabled" yaml:"trustedIdentityPropagationEnabled"`
	// The IAM user role that will be assumed by users and groups logged in to a Studio.
	//
	// The permissions attached to this IAM role can be scoped down for each user or group using session policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/emr_studio#user_role EmrStudio#user_role}
	UserRole *string `field:"optional" json:"userRole" yaml:"userRole"`
}

