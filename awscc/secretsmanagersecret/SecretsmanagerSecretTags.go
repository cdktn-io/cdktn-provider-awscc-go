// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretsmanagersecret


type SecretsmanagerSecretTags struct {
	// The key identifier, or name, of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/secretsmanager_secret#key SecretsmanagerSecret#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The string value associated with the key of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/secretsmanager_secret#value SecretsmanagerSecret#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

