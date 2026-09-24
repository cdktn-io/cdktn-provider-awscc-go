// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabricoutboundexternallink


type RtbfabricOutboundExternalLinkLinkLogSettingsApplicationLogsLinkApplicationLogSampling struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/rtbfabric_outbound_external_link#error_log RtbfabricOutboundExternalLink#error_log}.
	ErrorLog *float64 `field:"required" json:"errorLog" yaml:"errorLog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/rtbfabric_outbound_external_link#filter_log RtbfabricOutboundExternalLink#filter_log}.
	FilterLog *float64 `field:"required" json:"filterLog" yaml:"filterLog"`
}

