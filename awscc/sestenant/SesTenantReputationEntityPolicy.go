// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sestenant


type SesTenantReputationEntityPolicy struct {
	// The ARN of an SES-managed reputation policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/ses_tenant#reputation_entity_policy SesTenant#reputation_entity_policy}
	ReputationEntityPolicy *string `field:"optional" json:"reputationEntityPolicy" yaml:"reputationEntityPolicy"`
}

