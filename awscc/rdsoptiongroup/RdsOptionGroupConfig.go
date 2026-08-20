// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsoptiongroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RdsOptionGroupConfig struct {
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
	// Specifies the name of the engine that this option group should be associated with.
	//
	// Valid Values:
	//   +   ``mariadb``
	//   +   ``mysql``
	//   +   ``oracle-ee``
	//   +   ``oracle-ee-cdb``
	//   +   ``oracle-se2``
	//   +   ``oracle-se2-cdb``
	//   +   ``postgres``
	//   +   ``sqlserver-ee``
	//   +   ``sqlserver-se``
	//   +   ``sqlserver-ex``
	//   +   ``sqlserver-web``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_option_group#engine_name RdsOptionGroup#engine_name}
	EngineName *string `field:"required" json:"engineName" yaml:"engineName"`
	// Specifies the major version of the engine that this option group should be associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_option_group#major_engine_version RdsOptionGroup#major_engine_version}
	MajorEngineVersion *string `field:"required" json:"majorEngineVersion" yaml:"majorEngineVersion"`
	// The description of the option group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_option_group#option_group_description RdsOptionGroup#option_group_description}
	OptionGroupDescription *string `field:"required" json:"optionGroupDescription" yaml:"optionGroupDescription"`
	// A list of all available options for an option group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_option_group#option_configurations RdsOptionGroup#option_configurations}
	OptionConfigurations interface{} `field:"optional" json:"optionConfigurations" yaml:"optionConfigurations"`
	// The name of the option group to be created.
	//
	// Constraints:
	//   +  Must be 1 to 255 letters, numbers, or hyphens
	//   +  First character must be a letter
	//   +  Can't end with a hyphen or contain two consecutive hyphens
	//
	//  Example: ``myoptiongroup``
	//  If you don't specify a value for ``OptionGroupName`` property, a name is automatically created for the option group.
	//   This value is stored as a lowercase string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_option_group#option_group_name RdsOptionGroup#option_group_name}
	OptionGroupName *string `field:"optional" json:"optionGroupName" yaml:"optionGroupName"`
	// Tags to assign to the option group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_option_group#tags RdsOptionGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

