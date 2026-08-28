// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logsintegration


type LogsIntegrationResourceConfigOpenSearchResourceConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_integration#application_arn LogsIntegration#application_arn}.
	ApplicationArn *string `field:"optional" json:"applicationArn" yaml:"applicationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_integration#dashboard_viewer_principals LogsIntegration#dashboard_viewer_principals}.
	DashboardViewerPrincipals *[]*string `field:"optional" json:"dashboardViewerPrincipals" yaml:"dashboardViewerPrincipals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_integration#data_source_role_arn LogsIntegration#data_source_role_arn}.
	DataSourceRoleArn *string `field:"optional" json:"dataSourceRoleArn" yaml:"dataSourceRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_integration#kms_key_arn LogsIntegration#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/logs_integration#retention_days LogsIntegration#retention_days}.
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
}

