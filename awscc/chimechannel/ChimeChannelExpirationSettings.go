// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimechannel


type ChimeChannelExpirationSettings struct {
	// The condition the expiration period is measured from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_channel#expiration_criterion ChimeChannel#expiration_criterion}
	ExpirationCriterion *string `field:"optional" json:"expirationCriterion" yaml:"expirationCriterion"`
	// The period in days after which the system automatically deletes the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_channel#expiration_days ChimeChannel#expiration_days}
	ExpirationDays *float64 `field:"optional" json:"expirationDays" yaml:"expirationDays"`
}

