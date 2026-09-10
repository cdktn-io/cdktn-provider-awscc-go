// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscapability


type EksCapabilityConfigurationArgoCdNetworkAccess struct {
	// A list of VPC endpoint IDs to associate with the managed Argo CD API server endpoint.
	//
	// Each VPC endpoint provides private connectivity from a specific VPC to the Argo CD server. You can specify multiple VPC endpoint IDs to enable access from multiple VPCs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_capability#vpce_ids EksCapability#vpce_ids}
	VpceIds *[]*string `field:"optional" json:"vpceIds" yaml:"vpceIds"`
}

