// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimeappinstanceuser


type ChimeAppInstanceUserExpirationSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/chime_app_instance_user#expiration_criterion ChimeAppInstanceUser#expiration_criterion}.
	ExpirationCriterion *string `field:"optional" json:"expirationCriterion" yaml:"expirationCriterion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/chime_app_instance_user#expiration_days ChimeAppInstanceUser#expiration_days}.
	ExpirationDays *float64 `field:"optional" json:"expirationDays" yaml:"expirationDays"`
}

