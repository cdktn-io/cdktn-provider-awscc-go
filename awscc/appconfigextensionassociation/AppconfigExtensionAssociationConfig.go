// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appconfigextensionassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppconfigExtensionAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appconfig_extension_association#extension_identifier AppconfigExtensionAssociation#extension_identifier}.
	ExtensionIdentifier *string `field:"optional" json:"extensionIdentifier" yaml:"extensionIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appconfig_extension_association#extension_version_number AppconfigExtensionAssociation#extension_version_number}.
	ExtensionVersionNumber *float64 `field:"optional" json:"extensionVersionNumber" yaml:"extensionVersionNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appconfig_extension_association#parameters AppconfigExtensionAssociation#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appconfig_extension_association#resource_identifier AppconfigExtensionAssociation#resource_identifier}.
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appconfig_extension_association#tags AppconfigExtensionAssociation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

