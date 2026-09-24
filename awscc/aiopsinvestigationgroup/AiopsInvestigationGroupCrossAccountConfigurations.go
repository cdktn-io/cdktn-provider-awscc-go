// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aiopsinvestigationgroup


type AiopsInvestigationGroupCrossAccountConfigurations struct {
	// The Investigation Role's ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/aiops_investigation_group#source_role_arn AiopsInvestigationGroup#source_role_arn}
	SourceRoleArn *string `field:"optional" json:"sourceRoleArn" yaml:"sourceRoleArn"`
}

