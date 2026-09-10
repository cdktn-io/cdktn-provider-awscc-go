// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mpaidentitysource


type MpaIdentitySourceIdentitySourceParametersIamIdentityCenter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mpa_identity_source#instance_arn MpaIdentitySource#instance_arn}.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mpa_identity_source#region MpaIdentitySource#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
}

