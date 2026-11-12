// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialiveinput


type MedialiveInputSrtSettingsSrtCallerSourcesDecryption struct {
	// The algorithm used to encrypt content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#algorithm MedialiveInput#algorithm}
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// The ARN for the secret in Secrets Manager.
	//
	// Someone in your organization must create a secret and provide you with its ARN. The secret holds the passphrase that MediaLive uses to decrypt the source content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#passphrase_secret_arn MedialiveInput#passphrase_secret_arn}
	PassphraseSecretArn *string `field:"optional" json:"passphraseSecretArn" yaml:"passphraseSecretArn"`
}

