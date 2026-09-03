// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eksaccessentry


type EksAccessEntryAccessPolicies struct {
	// The access scope of the access policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/eks_access_entry#access_scope EksAccessEntry#access_scope}
	AccessScope *EksAccessEntryAccessPoliciesAccessScope `field:"optional" json:"accessScope" yaml:"accessScope"`
	// The ARN of the access policy to add to the access entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/eks_access_entry#policy_arn EksAccessEntry#policy_arn}
	PolicyArn *string `field:"optional" json:"policyArn" yaml:"policyArn"`
}

