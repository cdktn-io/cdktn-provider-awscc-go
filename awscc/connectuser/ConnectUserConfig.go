// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectuser

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectUserConfig struct {
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
	// The identifier of the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_user#instance_arn ConnectUser#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The identifier of the routing profile for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_user#routing_profile_arn ConnectUser#routing_profile_arn}
	RoutingProfileArn *string `field:"required" json:"routingProfileArn" yaml:"routingProfileArn"`
	// One or more security profile arns for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_user#security_profile_arns ConnectUser#security_profile_arns}
	SecurityProfileArns *[]*string `field:"required" json:"securityProfileArns" yaml:"securityProfileArns"`
	// The user name for the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_user#username ConnectUser#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// After Contact Work configurations of a user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_user#after_contact_work_configs ConnectUser#after_contact_work_configs}
	AfterContactWorkConfigs interface{} `field:"optional" json:"afterContactWorkConfigs" yaml:"afterContactWorkConfigs"`
	// Auto-accept configurations of a user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_user#auto_accept_configs ConnectUser#auto_accept_configs}
	AutoAcceptConfigs interface{} `field:"optional" json:"autoAcceptConfigs" yaml:"autoAcceptConfigs"`
	// The identifier of the user account in the directory used for identity management.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_user#directory_user_id ConnectUser#directory_user_id}
	DirectoryUserId *string `field:"optional" json:"directoryUserId" yaml:"directoryUserId"`
	// The identifier of the hierarchy group for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_user#hierarchy_group_arn ConnectUser#hierarchy_group_arn}
	HierarchyGroupArn *string `field:"optional" json:"hierarchyGroupArn" yaml:"hierarchyGroupArn"`
	// The information about the identity of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_user#identity_info ConnectUser#identity_info}
	IdentityInfo *ConnectUserIdentityInfo `field:"optional" json:"identityInfo" yaml:"identityInfo"`
	// The password for the user account.
	//
	// A password is required if you are using Amazon Connect for identity management. Otherwise, it is an error to include a password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_user#password ConnectUser#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Persistent Connection configurations of a user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_user#persistent_connection_configs ConnectUser#persistent_connection_configs}
	PersistentConnectionConfigs interface{} `field:"optional" json:"persistentConnectionConfigs" yaml:"persistentConnectionConfigs"`
	// The phone settings for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_user#phone_config ConnectUser#phone_config}
	PhoneConfig *ConnectUserPhoneConfig `field:"optional" json:"phoneConfig" yaml:"phoneConfig"`
	// Phone Number configurations of a user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_user#phone_number_configs ConnectUser#phone_number_configs}
	PhoneNumberConfigs interface{} `field:"optional" json:"phoneNumberConfigs" yaml:"phoneNumberConfigs"`
	// One or more tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_user#tags ConnectUser#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// One or more predefined attributes assigned to a user, with a level that indicates how skilled they are.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_user#user_proficiencies ConnectUser#user_proficiencies}
	UserProficiencies interface{} `field:"optional" json:"userProficiencies" yaml:"userProficiencies"`
	// Voice Enhancement configurations of a user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/connect_user#voice_enhancement_configs ConnectUser#voice_enhancement_configs}
	VoiceEnhancementConfigs interface{} `field:"optional" json:"voiceEnhancementConfigs" yaml:"voiceEnhancementConfigs"`
}

