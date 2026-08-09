// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectsecurityprofile

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectSecurityProfileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_security_profile#instance_arn ConnectSecurityProfile#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the security profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_security_profile#security_profile_name ConnectSecurityProfile#security_profile_name}
	SecurityProfileName *string `field:"required" json:"securityProfileName" yaml:"securityProfileName"`
	// The identifier of the hierarchy group that a security profile uses to restrict access to resources in Amazon Connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_security_profile#allowed_access_control_hierarchy_group_id ConnectSecurityProfile#allowed_access_control_hierarchy_group_id}
	AllowedAccessControlHierarchyGroupId *string `field:"optional" json:"allowedAccessControlHierarchyGroupId" yaml:"allowedAccessControlHierarchyGroupId"`
	// The list of tags that a security profile uses to restrict access to resources in Amazon Connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_security_profile#allowed_access_control_tags ConnectSecurityProfile#allowed_access_control_tags}
	AllowedAccessControlTags interface{} `field:"optional" json:"allowedAccessControlTags" yaml:"allowedAccessControlTags"`
	// The list of flow-module resources to be linked to a security profile in Amazon Connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_security_profile#allowed_flow_modules ConnectSecurityProfile#allowed_flow_modules}
	AllowedFlowModules interface{} `field:"optional" json:"allowedFlowModules" yaml:"allowedFlowModules"`
	// A list of third-party applications that the security profile will give access to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_security_profile#applications ConnectSecurityProfile#applications}
	Applications interface{} `field:"optional" json:"applications" yaml:"applications"`
	// The description of the security profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_security_profile#description ConnectSecurityProfile#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_security_profile#granular_access_control_configuration ConnectSecurityProfile#granular_access_control_configuration}.
	GranularAccessControlConfiguration *ConnectSecurityProfileGranularAccessControlConfiguration `field:"optional" json:"granularAccessControlConfiguration" yaml:"granularAccessControlConfiguration"`
	// The list of resources that a security profile applies hierarchy restrictions to in Amazon Connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_security_profile#hierarchy_restricted_resources ConnectSecurityProfile#hierarchy_restricted_resources}
	HierarchyRestrictedResources *[]*string `field:"optional" json:"hierarchyRestrictedResources" yaml:"hierarchyRestrictedResources"`
	// Permissions assigned to the security profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_security_profile#permissions ConnectSecurityProfile#permissions}
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// The list of resources that a security profile applies tag restrictions to in Amazon Connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_security_profile#tag_restricted_resources ConnectSecurityProfile#tag_restricted_resources}
	TagRestrictedResources *[]*string `field:"optional" json:"tagRestrictedResources" yaml:"tagRestrictedResources"`
	// The tags used to organize, track, or control access for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_security_profile#tags ConnectSecurityProfile#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

