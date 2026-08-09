// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package frauddetectorlist

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FrauddetectorListConfig struct {
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
	// The name of the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/frauddetector_list#name FrauddetectorList#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/frauddetector_list#description FrauddetectorList#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The elements in this list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/frauddetector_list#elements FrauddetectorList#elements}
	Elements *[]*string `field:"optional" json:"elements" yaml:"elements"`
	// Tags associated with this list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/frauddetector_list#tags FrauddetectorList#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The variable type of the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/frauddetector_list#variable_type FrauddetectorList#variable_type}
	VariableType *string `field:"optional" json:"variableType" yaml:"variableType"`
}

