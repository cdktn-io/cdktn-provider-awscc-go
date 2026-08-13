// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eksfargateprofile


type EksFargateProfileSelectorsLabels struct {
	// The key name of the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/eks_fargate_profile#key EksFargateProfile#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/eks_fargate_profile#value EksFargateProfile#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

