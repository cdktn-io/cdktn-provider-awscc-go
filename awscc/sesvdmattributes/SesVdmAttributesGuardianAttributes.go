// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesvdmattributes


type SesVdmAttributesGuardianAttributes struct {
	// Whether emails sent from this account have optimized delivery algorithm enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ses_vdm_attributes#optimized_shared_delivery SesVdmAttributes#optimized_shared_delivery}
	OptimizedSharedDelivery *string `field:"optional" json:"optimizedSharedDelivery" yaml:"optimizedSharedDelivery"`
}

