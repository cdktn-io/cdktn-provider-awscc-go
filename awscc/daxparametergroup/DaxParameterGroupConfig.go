// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package daxparametergroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DaxParameterGroupConfig struct {
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
	// A description of the parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dax_parameter_group#description DaxParameterGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dax_parameter_group#parameter_group_name DaxParameterGroup#parameter_group_name}
	ParameterGroupName *string `field:"optional" json:"parameterGroupName" yaml:"parameterGroupName"`
	// An array of name-value pairs for the parameters in the group.
	//
	// Each element in the array represents a single parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dax_parameter_group#parameter_name_values DaxParameterGroup#parameter_name_values}
	ParameterNameValues *string `field:"optional" json:"parameterNameValues" yaml:"parameterNameValues"`
}

