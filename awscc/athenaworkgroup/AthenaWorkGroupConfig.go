// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AthenaWorkGroupConfig struct {
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
	// The workGroup name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/athena_work_group#name AthenaWorkGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The workgroup description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/athena_work_group#description AthenaWorkGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The option to delete the workgroup and its contents even if the workgroup contains any named queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/athena_work_group#recursive_delete_option AthenaWorkGroup#recursive_delete_option}
	RecursiveDeleteOption interface{} `field:"optional" json:"recursiveDeleteOption" yaml:"recursiveDeleteOption"`
	// The state of the workgroup: ENABLED or DISABLED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/athena_work_group#state AthenaWorkGroup#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// One or more tags, separated by commas, that you want to attach to the workgroup as you create it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/athena_work_group#tags AthenaWorkGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The workgroup configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/athena_work_group#work_group_configuration AthenaWorkGroup#work_group_configuration}
	WorkGroupConfiguration *AthenaWorkGroupWorkGroupConfiguration `field:"optional" json:"workGroupConfiguration" yaml:"workGroupConfiguration"`
	// The workgroup configuration update object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/athena_work_group#work_group_configuration_updates AthenaWorkGroup#work_group_configuration_updates}
	WorkGroupConfigurationUpdates *AthenaWorkGroupWorkGroupConfigurationUpdates `field:"optional" json:"workGroupConfigurationUpdates" yaml:"workGroupConfigurationUpdates"`
}

