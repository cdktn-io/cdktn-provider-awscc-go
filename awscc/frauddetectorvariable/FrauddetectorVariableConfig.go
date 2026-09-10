// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package frauddetectorvariable

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FrauddetectorVariableConfig struct {
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
	// The source of the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/frauddetector_variable#data_source FrauddetectorVariable#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The data type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/frauddetector_variable#data_type FrauddetectorVariable#data_type}
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// The default value for the variable when no value is received.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/frauddetector_variable#default_value FrauddetectorVariable#default_value}
	DefaultValue *string `field:"required" json:"defaultValue" yaml:"defaultValue"`
	// The name of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/frauddetector_variable#name FrauddetectorVariable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/frauddetector_variable#description FrauddetectorVariable#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags associated with this variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/frauddetector_variable#tags FrauddetectorVariable#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The variable type. For more information see https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-variable.html#variable-types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/frauddetector_variable#variable_type FrauddetectorVariable#variable_type}
	VariableType *string `field:"optional" json:"variableType" yaml:"variableType"`
}

