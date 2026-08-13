// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabricinboundexternallink


type RtbfabricInboundExternalLinkLinkLogSettingsApplicationLogsLinkApplicationLogSampling struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rtbfabric_inbound_external_link#error_log RtbfabricInboundExternalLink#error_log}.
	ErrorLog *float64 `field:"required" json:"errorLog" yaml:"errorLog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rtbfabric_inbound_external_link#filter_log RtbfabricInboundExternalLink#filter_log}.
	FilterLog *float64 `field:"required" json:"filterLog" yaml:"filterLog"`
}

