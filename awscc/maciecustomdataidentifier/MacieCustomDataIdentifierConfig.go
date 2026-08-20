// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package maciecustomdataidentifier

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MacieCustomDataIdentifierConfig struct {
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
	// Name of custom data identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/macie_custom_data_identifier#name MacieCustomDataIdentifier#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Regular expression for custom data identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/macie_custom_data_identifier#regex MacieCustomDataIdentifier#regex}
	Regex *string `field:"required" json:"regex" yaml:"regex"`
	// Description of custom data identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/macie_custom_data_identifier#description MacieCustomDataIdentifier#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Words to be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/macie_custom_data_identifier#ignore_words MacieCustomDataIdentifier#ignore_words}
	IgnoreWords *[]*string `field:"optional" json:"ignoreWords" yaml:"ignoreWords"`
	// Keywords to be matched against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/macie_custom_data_identifier#keywords MacieCustomDataIdentifier#keywords}
	Keywords *[]*string `field:"optional" json:"keywords" yaml:"keywords"`
	// Maximum match distance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/macie_custom_data_identifier#maximum_match_distance MacieCustomDataIdentifier#maximum_match_distance}
	MaximumMatchDistance *float64 `field:"optional" json:"maximumMatchDistance" yaml:"maximumMatchDistance"`
	// A collection of tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/macie_custom_data_identifier#tags MacieCustomDataIdentifier#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

