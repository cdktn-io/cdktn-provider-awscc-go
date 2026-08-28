// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretsmanagersecret

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecretsmanagerSecretConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The description of the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/secretsmanager_secret#description SecretsmanagerSecret#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A structure that specifies how to generate a password to encrypt and store in the secret.
	//
	// To include a specific string in the secret, use ``SecretString`` instead. If you omit both ``GenerateSecretString`` and ``SecretString``, you create an empty secret. When you make a change to this property, a new secret version is created.
	//  We recommend that you specify the maximum length and include every character type that the system you are generating a password for can support.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/secretsmanager_secret#generate_secret_string SecretsmanagerSecret#generate_secret_string}
	GenerateSecretString *SecretsmanagerSecretGenerateSecretString `field:"optional" json:"generateSecretString" yaml:"generateSecretString"`
	// The ARN, key ID, or alias of the KMS key that Secrets Manager uses to encrypt the secret value in the secret.
	//
	// An alias is always prefixed by ``alias/``, for example ``alias/aws/secretsmanager``. For more information, see [About aliases](https://docs.aws.amazon.com/kms/latest/developerguide/alias-about.html).
	//  To use a KMS key in a different account, use the key ARN or the alias ARN.
	//  If you don't specify this value, then Secrets Manager uses the key ``aws/secretsmanager``. If that key doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.
	//  If the secret is in a different AWS account from the credentials calling the API, then you can't use ``aws/secretsmanager`` to encrypt the secret, and you must create and use a customer managed KMS key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/secretsmanager_secret#kms_key_id SecretsmanagerSecret#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The name of the new secret.
	//
	// The secret name can contain ASCII letters, numbers, and the following characters: /_+=.@-
	//  Do not end your secret name with a hyphen followed by six characters. If you do so, you risk confusion and unexpected results when searching for a secret by partial ARN. Secrets Manager automatically adds a hyphen and six random characters after the secret name at the end of the ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/secretsmanager_secret#name SecretsmanagerSecret#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A custom type that specifies a ``Region`` and the ``KmsKeyId`` for a replica secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/secretsmanager_secret#replica_regions SecretsmanagerSecret#replica_regions}
	ReplicaRegions interface{} `field:"optional" json:"replicaRegions" yaml:"replicaRegions"`
	// The text to encrypt and store in the secret.
	//
	// We recommend you use a JSON structure of key/value pairs for your secret value. To generate a random password, use ``GenerateSecretString`` instead. If you omit both ``GenerateSecretString`` and ``SecretString``, you create an empty secret. When you make a change to this property, a new secret version is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/secretsmanager_secret#secret_string SecretsmanagerSecret#secret_string}
	SecretString *string `field:"optional" json:"secretString" yaml:"secretString"`
	// A list of tags to attach to the secret.
	//
	// Each tag is a key and value pair of strings in a JSON text string, for example:
	//   ``[{"Key":"CostCenter","Value":"12345"},{"Key":"environment","Value":"production"}]``
	//  Secrets Manager tag key names are case sensitive. A tag with the key "ABC" is a different tag from one with key "abc".
	//  Stack-level tags, tags you apply to the CloudFormation stack, are also attached to the secret.
	//  If you check tags in permissions policies as part of your security strategy, then adding or removing a tag can change permissions. If the completion of this operation would result in you losing your permissions for this secret, then Secrets Manager blocks the operation and returns an ``Access Denied`` error. For more information, see [Control access to secrets using tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#tag-secrets-abac) and [Limit access to identities with tags that match secrets' tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#auth-and-access_tags2).
	//  For information about how to format a JSON parameter for the various command line tool environments, see [Using JSON for Parameters](https://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#cli-using-param-json). If your command-line tool or SDK requires quotation marks around the parameter, you should use single quotes to avoid confusion with the double quotes required in the JSON text.
	//  The following restrictions apply to tags:
	//   +  Maximum number of tags per secret: 50
	//   +  Maximum key length: 127 Unicode characters in UTF-8
	//   +  Maximum value length: 255 Unicode characters in UTF-8
	//   +  Tag keys and values are case sensitive.
	//   +  Do not use the ``aws:`` prefix in your tag names or values because AWS reserves it for AWS use. You can't edit or delete tag names or values with this prefix. Tags with this prefix do not count against your tags per secret limit.
	//   +  If you use your tagging schema across multiple services and resources, other services might have restrictions on allowed characters. Generally allowed characters: letters, spaces, and numbers representable in UTF-8, plus the following special characters: + - = . _ : /
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The exact string that identifies the third-party partner that holds the external secret.
	//
	// For more information, see [Managed external secret partners](https://docs.aws.amazon.com/secretsmanager/latest/userguide/mes-partners.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/secretsmanager_secret#type SecretsmanagerSecret#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

