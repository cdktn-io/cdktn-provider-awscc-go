// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilesobjecttype

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CustomerprofilesObjectTypeConfig struct {
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
	// Description of the profile object type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_object_type#description CustomerprofilesObjectType#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The unique name of the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_object_type#domain_name CustomerprofilesObjectType#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The name of the profile object type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_object_type#object_type_name CustomerprofilesObjectType#object_type_name}
	ObjectTypeName *string `field:"required" json:"objectTypeName" yaml:"objectTypeName"`
	// Indicates whether a profile should be created when data is received.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_object_type#allow_profile_creation CustomerprofilesObjectType#allow_profile_creation}
	AllowProfileCreation interface{} `field:"optional" json:"allowProfileCreation" yaml:"allowProfileCreation"`
	// The default encryption key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_object_type#encryption_key CustomerprofilesObjectType#encryption_key}
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The default number of days until the data within the domain expires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_object_type#expiration_days CustomerprofilesObjectType#expiration_days}
	ExpirationDays *float64 `field:"optional" json:"expirationDays" yaml:"expirationDays"`
	// A list of the name and ObjectType field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_object_type#fields CustomerprofilesObjectType#fields}
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// A list of unique keys that can be used to map data to the profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_object_type#keys CustomerprofilesObjectType#keys}
	Keys interface{} `field:"optional" json:"keys" yaml:"keys"`
	// The maximum number of profile objects for this object type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_object_type#max_profile_object_count CustomerprofilesObjectType#max_profile_object_count}
	MaxProfileObjectCount *float64 `field:"optional" json:"maxProfileObjectCount" yaml:"maxProfileObjectCount"`
	// The format of your sourceLastUpdatedTimestamp that was previously set up.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_object_type#source_last_updated_timestamp_format CustomerprofilesObjectType#source_last_updated_timestamp_format}
	SourceLastUpdatedTimestampFormat *string `field:"optional" json:"sourceLastUpdatedTimestampFormat" yaml:"sourceLastUpdatedTimestampFormat"`
	// Defines the priority order of object types. Lower value indicates higher priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_object_type#source_priority CustomerprofilesObjectType#source_priority}
	SourcePriority *float64 `field:"optional" json:"sourcePriority" yaml:"sourcePriority"`
	// The tags (keys and values) associated with the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_object_type#tags CustomerprofilesObjectType#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A unique identifier for the object template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_object_type#template_id CustomerprofilesObjectType#template_id}
	TemplateId *string `field:"optional" json:"templateId" yaml:"templateId"`
}

