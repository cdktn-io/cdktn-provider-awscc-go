// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerreportgenerator


type LicensemanagerReportGeneratorReportContext struct {
	// Amazon Resource Names (ARNs) of the license asset groups to include in the report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/licensemanager_report_generator#license_asset_group_arns LicensemanagerReportGenerator#license_asset_group_arns}
	LicenseAssetGroupArns *[]*string `field:"optional" json:"licenseAssetGroupArns" yaml:"licenseAssetGroupArns"`
	// Amazon Resource Names (ARNs) of the license configurations that this generator reports on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/licensemanager_report_generator#license_configuration_arns LicensemanagerReportGenerator#license_configuration_arns}
	LicenseConfigurationArns *[]*string `field:"optional" json:"licenseConfigurationArns" yaml:"licenseConfigurationArns"`
	// End date for the report data collection period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/licensemanager_report_generator#report_end_date LicensemanagerReportGenerator#report_end_date}
	ReportEndDate *string `field:"optional" json:"reportEndDate" yaml:"reportEndDate"`
	// Start date for the report data collection period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/licensemanager_report_generator#report_start_date LicensemanagerReportGenerator#report_start_date}
	ReportStartDate *string `field:"optional" json:"reportStartDate" yaml:"reportStartDate"`
}

