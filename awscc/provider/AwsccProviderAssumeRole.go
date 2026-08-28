// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type AwsccProviderAssumeRole struct {
	// Amazon Resource Name (ARN) of the IAM Role to assume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs#role_arn AwsccProvider#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The duration, between 15 minutes and 12 hours, of the role session.
	//
	// Valid time units are ns, us (or µs), ms, s, h, or m.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs#duration AwsccProvider#duration}
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// External identifier to use when assuming the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs#external_id AwsccProvider#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// IAM policy in JSON format to use as a session policy.
	//
	// The effective permissions for the session will be the intersection between this polcy and the role's policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs#policy AwsccProvider#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Amazon Resource Names (ARNs) of IAM Policies to use as managed session policies.
	//
	// The effective permissions for the session will be the intersection between these polcy and the role's policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs#policy_arns AwsccProvider#policy_arns}
	PolicyArns *[]*string `field:"optional" json:"policyArns" yaml:"policyArns"`
	// Session name to use when assuming the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs#session_name AwsccProvider#session_name}
	SessionName *string `field:"optional" json:"sessionName" yaml:"sessionName"`
	// Map of assume role session tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs#tags AwsccProvider#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Set of assume role session tag keys to pass to any subsequent sessions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs#transitive_tag_keys AwsccProvider#transitive_tag_keys}
	TransitiveTagKeys *[]*string `field:"optional" json:"transitiveTagKeys" yaml:"transitiveTagKeys"`
}

