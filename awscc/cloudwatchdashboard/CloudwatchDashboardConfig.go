// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchdashboard

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudwatchDashboardConfig struct {
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
	// The detailed information about the dashboard in JSON format, including the widgets to include and their location on the dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudwatch_dashboard#dashboard_body CloudwatchDashboard#dashboard_body}
	DashboardBody *string `field:"required" json:"dashboardBody" yaml:"dashboardBody"`
	// The name of the dashboard.
	//
	// The name must be between 1 and 255 characters. If you do not specify a name, one will be generated automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudwatch_dashboard#dashboard_name CloudwatchDashboard#dashboard_name}
	DashboardName *string `field:"optional" json:"dashboardName" yaml:"dashboardName"`
	// A list of key-value pairs to associate with the cloudwatch dashboard.
	//
	// You can associate up to 50 tags with a dashboard
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudwatch_dashboard#tags CloudwatchDashboard#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

