// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package xraygroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type XrayGroupConfig struct {
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
	// The case-sensitive name of the new group. Names must be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/xray_group#group_name XrayGroup#group_name}
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// The filter expression defining criteria by which to group traces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/xray_group#filter_expression XrayGroup#filter_expression}
	FilterExpression *string `field:"optional" json:"filterExpression" yaml:"filterExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/xray_group#insights_configuration XrayGroup#insights_configuration}.
	InsightsConfiguration *XrayGroupInsightsConfiguration `field:"optional" json:"insightsConfiguration" yaml:"insightsConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/xray_group#tags XrayGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

