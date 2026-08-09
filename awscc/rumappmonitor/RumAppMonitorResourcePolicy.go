// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rumappmonitor


type RumAppMonitorResourcePolicy struct {
	// The JSON to use as the resource policy. The document can be up to 4 KB in size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/rum_app_monitor#policy_document RumAppMonitor#policy_document}
	PolicyDocument *string `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// A string value that you can use to conditionally update your policy.
	//
	// You can provide the revision ID of your existing policy to make mutating requests against that policy.
	//
	//  When you assign a policy revision ID, then later requests about that policy will be rejected with an InvalidPolicyRevisionIdException error if they don't provide the correct current revision ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/rum_app_monitor#policy_revision_id RumAppMonitor#policy_revision_id}
	PolicyRevisionId *string `field:"optional" json:"policyRevisionId" yaml:"policyRevisionId"`
}

