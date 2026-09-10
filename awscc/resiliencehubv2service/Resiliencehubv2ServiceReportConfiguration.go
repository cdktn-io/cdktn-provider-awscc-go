// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2service


type Resiliencehubv2ServiceReportConfiguration struct {
	// Output destinations for generated reports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/resiliencehubv2_service#report_output Resiliencehubv2Service#report_output}
	ReportOutput interface{} `field:"optional" json:"reportOutput" yaml:"reportOutput"`
}

