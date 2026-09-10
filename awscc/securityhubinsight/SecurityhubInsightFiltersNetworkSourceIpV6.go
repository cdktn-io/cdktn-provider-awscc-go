// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubinsight


type SecurityhubInsightFiltersNetworkSourceIpV6 struct {
	// A finding's CIDR value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_insight#cidr SecurityhubInsight#cidr}
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
}

