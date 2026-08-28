// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretsmanagerrotationschedule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecretsmanagerRotationScheduleConfig struct {
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
	// The ARN or name of the secret to rotate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/secretsmanager_rotation_schedule#secret_id SecretsmanagerRotationSchedule#secret_id}
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// The list of metadata needed to successfully rotate a managed external secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/secretsmanager_rotation_schedule#external_secret_rotation_metadata SecretsmanagerRotationSchedule#external_secret_rotation_metadata}
	ExternalSecretRotationMetadata interface{} `field:"optional" json:"externalSecretRotationMetadata" yaml:"externalSecretRotationMetadata"`
	// The ARN of the IAM role that is used by Secrets Manager to rotate a managed external secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/secretsmanager_rotation_schedule#external_secret_rotation_role_arn SecretsmanagerRotationSchedule#external_secret_rotation_role_arn}
	ExternalSecretRotationRoleArn *string `field:"optional" json:"externalSecretRotationRoleArn" yaml:"externalSecretRotationRoleArn"`
	// Creates a new Lambda rotation function based on one of the Secrets Manager rotation function templates.
	//
	// To use a rotation function that already exists, specify RotationLambdaARN instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/secretsmanager_rotation_schedule#hosted_rotation_lambda SecretsmanagerRotationSchedule#hosted_rotation_lambda}
	HostedRotationLambda *SecretsmanagerRotationScheduleHostedRotationLambda `field:"optional" json:"hostedRotationLambda" yaml:"hostedRotationLambda"`
	// Specifies whether to rotate the secret immediately or wait until the next scheduled rotation window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/secretsmanager_rotation_schedule#rotate_immediately_on_update SecretsmanagerRotationSchedule#rotate_immediately_on_update}
	RotateImmediatelyOnUpdate interface{} `field:"optional" json:"rotateImmediatelyOnUpdate" yaml:"rotateImmediatelyOnUpdate"`
	// The ARN of an existing Lambda rotation function.
	//
	// To specify a rotation function that is also defined in this template, use the Ref function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/secretsmanager_rotation_schedule#rotation_lambda_arn SecretsmanagerRotationSchedule#rotation_lambda_arn}
	RotationLambdaArn *string `field:"optional" json:"rotationLambdaArn" yaml:"rotationLambdaArn"`
	// A structure that defines the rotation configuration for this secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/secretsmanager_rotation_schedule#rotation_rules SecretsmanagerRotationSchedule#rotation_rules}
	RotationRules *SecretsmanagerRotationScheduleRotationRules `field:"optional" json:"rotationRules" yaml:"rotationRules"`
}

