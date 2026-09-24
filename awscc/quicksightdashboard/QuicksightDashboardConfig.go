// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdashboard

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightDashboardConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_dashboard#aws_account_id QuicksightDashboard#aws_account_id}.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_dashboard#dashboard_id QuicksightDashboard#dashboard_id}.
	DashboardId *string `field:"required" json:"dashboardId" yaml:"dashboardId"`
	// <p>Dashboard source entity.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_dashboard#source_entity QuicksightDashboard#source_entity}
	SourceEntity *QuicksightDashboardSourceEntity `field:"required" json:"sourceEntity" yaml:"sourceEntity"`
	// <p>Dashboard publish options.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_dashboard#dashboard_publish_options QuicksightDashboard#dashboard_publish_options}
	DashboardPublishOptions *QuicksightDashboardDashboardPublishOptions `field:"optional" json:"dashboardPublishOptions" yaml:"dashboardPublishOptions"`
	// <p>The display name of the dashboard.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_dashboard#name QuicksightDashboard#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// <p>A list of QuickSight parameters and the list's override values.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_dashboard#parameters QuicksightDashboard#parameters}
	Parameters *QuicksightDashboardParameters `field:"optional" json:"parameters" yaml:"parameters"`
	// <p>A structure that contains the permissions of the dashboard.
	//
	// You can use this structure
	//             for granting permissions by providing a list of IAM action information for each
	//             principal ARN. </p>
	//
	//         <p>To specify no permissions, omit the permissions list.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_dashboard#permissions QuicksightDashboard#permissions}
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the             dashboard.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_dashboard#tags QuicksightDashboard#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// <p>The Amazon Resource Name (ARN) of the theme that is being used for this dashboard.
	//
	// If
	//             you add a value for this field, it overrides the value that is used in the source
	//             entity. The theme ARN must exist in the same AWS account where you create the
	//             dashboard.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_dashboard#theme_arn QuicksightDashboard#theme_arn}
	ThemeArn *string `field:"optional" json:"themeArn" yaml:"themeArn"`
	// <p>A description for the first version of the dashboard being created.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_dashboard#version_description QuicksightDashboard#version_description}
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
}

