// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudtraildashboard


type CloudtrailDashboardWidgets struct {
	// The placeholder keys in the QueryStatement. For example: $StartTime$, $EndTime$, $Period$.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudtrail_dashboard#query_parameters CloudtrailDashboard#query_parameters}
	QueryParameters *[]*string `field:"optional" json:"queryParameters" yaml:"queryParameters"`
	// The SQL query statement on one or more event data stores.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudtrail_dashboard#query_statement CloudtrailDashboard#query_statement}
	QueryStatement *string `field:"optional" json:"queryStatement" yaml:"queryStatement"`
	// The view properties of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudtrail_dashboard#view_properties CloudtrailDashboard#view_properties}
	ViewProperties *map[string]*string `field:"optional" json:"viewProperties" yaml:"viewProperties"`
}

