// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudtraildashboard

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudtrailDashboardConfig struct {
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
	// The name of the dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudtrail_dashboard#name CloudtrailDashboard#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Configures the automatic refresh schedule for the dashboard.
	//
	// Includes the frequency unit (DAYS or HOURS) and value, as well as the status (ENABLED or DISABLED) of the refresh schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudtrail_dashboard#refresh_schedule CloudtrailDashboard#refresh_schedule}
	RefreshSchedule *CloudtrailDashboardRefreshSchedule `field:"optional" json:"refreshSchedule" yaml:"refreshSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudtrail_dashboard#tags CloudtrailDashboard#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Indicates whether the dashboard is protected from termination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudtrail_dashboard#termination_protection_enabled CloudtrailDashboard#termination_protection_enabled}
	TerminationProtectionEnabled interface{} `field:"optional" json:"terminationProtectionEnabled" yaml:"terminationProtectionEnabled"`
	// List of widgets on the dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/cloudtrail_dashboard#widgets CloudtrailDashboard#widgets}
	Widgets interface{} `field:"optional" json:"widgets" yaml:"widgets"`
}

