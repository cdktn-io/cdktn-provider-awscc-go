// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesvdmattributes


type SesVdmAttributesDashboardAttributes struct {
	// Whether emails sent from this account have engagement tracking enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_vdm_attributes#engagement_metrics SesVdmAttributes#engagement_metrics}
	EngagementMetrics *string `field:"optional" json:"engagementMetrics" yaml:"engagementMetrics"`
}

