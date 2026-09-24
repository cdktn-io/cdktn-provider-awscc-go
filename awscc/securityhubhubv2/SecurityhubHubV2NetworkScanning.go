// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubhubv2


type SecurityhubHubV2NetworkScanning struct {
	// Whether the Network Scanning feature is enabled for this account and Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_hub_v2#status SecurityhubHubV2#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

