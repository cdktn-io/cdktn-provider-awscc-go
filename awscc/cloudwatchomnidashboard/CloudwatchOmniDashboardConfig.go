// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchomnidashboard

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudwatchOmniDashboardConfig struct {
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
	// The dashboard definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_omni_dashboard#body CloudwatchOmniDashboard#body}
	Body *string `field:"required" json:"body" yaml:"body"`
	// A name that identifies the dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_omni_dashboard#name CloudwatchOmniDashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The unique ID of the space to create the dashboard in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_omni_dashboard#space_id CloudwatchOmniDashboard#space_id}
	SpaceId *string `field:"required" json:"spaceId" yaml:"spaceId"`
	// An optional description of the dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_omni_dashboard#description CloudwatchOmniDashboard#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/cloudwatch_omni_dashboard#tags CloudwatchOmniDashboard#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

