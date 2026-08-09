// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package neptunedbparametergroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NeptuneDbParameterGroupConfig struct {
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
	// Provides the customer-specified description for this DB parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/neptune_db_parameter_group#description NeptuneDbParameterGroup#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Must be `neptune1` for engine versions prior to 1.2.0.0, or `neptune1.2` for engine version `1.2.0.0` and higher.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/neptune_db_parameter_group#family NeptuneDbParameterGroup#family}
	Family *string `field:"required" json:"family" yaml:"family"`
	// The parameters to set for this DB parameter group.
	//
	// The parameters are expressed as a JSON object consisting of key-value pairs.
	//
	// Changes to dynamic parameters are applied immediately. During an update, if you have static parameters (whether they were changed or not), it triggers AWS CloudFormation to reboot the associated DB instance without failover.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/neptune_db_parameter_group#parameters NeptuneDbParameterGroup#parameters}
	Parameters *string `field:"required" json:"parameters" yaml:"parameters"`
	// Provides the name of the DB parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/neptune_db_parameter_group#name NeptuneDbParameterGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An optional array of key-value pairs to apply to this DB parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/neptune_db_parameter_group#tags NeptuneDbParameterGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

