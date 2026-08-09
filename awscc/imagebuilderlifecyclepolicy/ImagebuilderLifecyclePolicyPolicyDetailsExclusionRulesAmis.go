// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderlifecyclepolicy


type ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmis struct {
	// Use to apply lifecycle policy actions on whether the AMI is public.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/imagebuilder_lifecycle_policy#is_public ImagebuilderLifecyclePolicy#is_public}
	IsPublic interface{} `field:"optional" json:"isPublic" yaml:"isPublic"`
	// Use to apply lifecycle policy actions on AMIs launched before a certain time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/imagebuilder_lifecycle_policy#last_launched ImagebuilderLifecyclePolicy#last_launched}
	LastLaunched *ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisLastLaunched `field:"optional" json:"lastLaunched" yaml:"lastLaunched"`
	// Use to apply lifecycle policy actions on AMIs distributed to a set of regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/imagebuilder_lifecycle_policy#regions ImagebuilderLifecyclePolicy#regions}
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Use to apply lifecycle policy actions on AMIs shared with a set of regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/imagebuilder_lifecycle_policy#shared_accounts ImagebuilderLifecyclePolicy#shared_accounts}
	SharedAccounts *[]*string `field:"optional" json:"sharedAccounts" yaml:"sharedAccounts"`
	// The AMIs to select by tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/imagebuilder_lifecycle_policy#tag_map ImagebuilderLifecyclePolicy#tag_map}
	TagMap *map[string]*string `field:"optional" json:"tagMap" yaml:"tagMap"`
}

