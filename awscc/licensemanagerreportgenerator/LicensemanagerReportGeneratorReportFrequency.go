// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerreportgenerator


type LicensemanagerReportGeneratorReportFrequency struct {
	// Time period between each report. The period can be daily, weekly, or monthly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/licensemanager_report_generator#period LicensemanagerReportGenerator#period}
	Period *string `field:"optional" json:"period" yaml:"period"`
	// Number of times within the frequency period that a report is generated. The only supported value is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/licensemanager_report_generator#value LicensemanagerReportGenerator#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

