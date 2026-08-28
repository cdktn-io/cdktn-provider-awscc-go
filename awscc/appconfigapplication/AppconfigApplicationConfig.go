// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appconfigapplication

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppconfigApplicationConfig struct {
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
	// A name for the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appconfig_application#name AppconfigApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appconfig_application#description AppconfigApplication#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Metadata to assign to the application.
	//
	// Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appconfig_application#tags AppconfigApplication#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

