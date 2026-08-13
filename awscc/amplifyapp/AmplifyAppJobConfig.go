// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package amplifyapp


type AmplifyAppJobConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/amplify_app#build_compute_type AmplifyApp#build_compute_type}.
	BuildComputeType *string `field:"optional" json:"buildComputeType" yaml:"buildComputeType"`
}

