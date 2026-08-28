// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appconfigextension

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppconfigExtensionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appconfig_extension#actions AppconfigExtension#actions}.
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// Name of the extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appconfig_extension#name AppconfigExtension#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Description of the extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appconfig_extension#description AppconfigExtension#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appconfig_extension#latest_version_number AppconfigExtension#latest_version_number}.
	LatestVersionNumber *float64 `field:"optional" json:"latestVersionNumber" yaml:"latestVersionNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appconfig_extension#parameters AppconfigExtension#parameters}.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// An array of key-value tags to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appconfig_extension#tags AppconfigExtension#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

