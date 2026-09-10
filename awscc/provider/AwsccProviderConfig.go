// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type AwsccProviderConfig struct {
	// This is the AWS access key.
	//
	// It must be provided, but it can also be sourced from the `AWS_ACCESS_KEY_ID` environment variable, or via a shared credentials file if `profile` is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs#access_key AwsccProvider#access_key}
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs#alias AwsccProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// An `assume_role` block (documented below). Only one `assume_role` block may be in the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs#assume_role AwsccProvider#assume_role}
	AssumeRole *AwsccProviderAssumeRole `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// An `assume_role_with_web_identity` block (documented below). Only one `assume_role_with_web_identity` block may be in the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs#assume_role_with_web_identity AwsccProvider#assume_role_with_web_identity}
	AssumeRoleWithWebIdentity *AwsccProviderAssumeRoleWithWebIdentity `field:"optional" json:"assumeRoleWithWebIdentity" yaml:"assumeRoleWithWebIdentity"`
	// An `endpoints` block (documented below). Only one `endpoints` block may be in the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs#endpoints AwsccProvider#endpoints}
	Endpoints *AwsccProviderEndpoints `field:"optional" json:"endpoints" yaml:"endpoints"`
	// URL of a proxy to use for HTTP requests when accessing the AWS API.
	//
	// Can also be set using the `HTTP_PROXY` or `http_proxy` environment variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs#http_proxy AwsccProvider#http_proxy}
	HttpProxy *string `field:"optional" json:"httpProxy" yaml:"httpProxy"`
	// URL of a proxy to use for HTTPS requests when accessing the AWS API.
	//
	// Can also be set using the `HTTPS_PROXY` or `https_proxy` environment variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs#https_proxy AwsccProvider#https_proxy}
	HttpsProxy *string `field:"optional" json:"httpsProxy" yaml:"httpsProxy"`
	// Explicitly allow the provider to perform "insecure" SSL requests. If not set, defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs#insecure AwsccProvider#insecure}
	Insecure interface{} `field:"optional" json:"insecure" yaml:"insecure"`
	// The maximum number of times an AWS API request is retried on failure. If not set, defaults to 25.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs#max_retries AwsccProvider#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// Comma-separated list of hosts that should not use HTTP or HTTPS proxies.
	//
	// Can also be set using the `NO_PROXY` or `no_proxy` environment variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs#no_proxy AwsccProvider#no_proxy}
	NoProxy *string `field:"optional" json:"noProxy" yaml:"noProxy"`
	// This is the AWS profile name as set in the shared credentials file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs#profile AwsccProvider#profile}
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// This is the AWS region.
	//
	// It must be provided, but it can also be sourced from the `AWS_DEFAULT_REGION` environment variables, via a shared config file, or from the EC2 Instance Metadata Service if used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs#region AwsccProvider#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Amazon Resource Name of the AWS CloudFormation service role that is used on your behalf to perform operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs#role_arn AwsccProvider#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// This is the AWS secret key.
	//
	// It must be provided, but it can also be sourced from the `AWS_SECRET_ACCESS_KEY` environment variable, or via a shared credentials file if `profile` is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs#secret_key AwsccProvider#secret_key}
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// List of paths to shared config files. If not set, defaults to `~/.aws/config`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs#shared_config_files AwsccProvider#shared_config_files}
	SharedConfigFiles *[]*string `field:"optional" json:"sharedConfigFiles" yaml:"sharedConfigFiles"`
	// List of paths to shared credentials files. If not set, defaults to `~/.aws/credentials`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs#shared_credentials_files AwsccProvider#shared_credentials_files}
	SharedCredentialsFiles *[]*string `field:"optional" json:"sharedCredentialsFiles" yaml:"sharedCredentialsFiles"`
	// Skip the AWS Metadata API check.
	//
	// Useful for AWS API implementations that do not have a metadata API endpoint.  Setting to `true` prevents Terraform from authenticating via the Metadata API. You may need to use other authentication methods like static credentials, configuration variables, or environment variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs#skip_medatadata_api_check AwsccProvider#skip_medatadata_api_check}
	SkipMedatadataApiCheck interface{} `field:"optional" json:"skipMedatadataApiCheck" yaml:"skipMedatadataApiCheck"`
	// Skip the AWS Metadata API check.
	//
	// Useful for AWS API implementations that do not have a metadata API endpoint.  Setting to `true` prevents Terraform from authenticating via the Metadata API. You may need to use other authentication methods like static credentials, configuration variables, or environment variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs#skip_metadata_api_check AwsccProvider#skip_metadata_api_check}
	SkipMetadataApiCheck interface{} `field:"optional" json:"skipMetadataApiCheck" yaml:"skipMetadataApiCheck"`
	// Session token for validating temporary credentials.
	//
	// Typically provided after successful identity federation or Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit MFA code used to get temporary credentials.  It can also be sourced from the `AWS_SESSION_TOKEN` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs#token AwsccProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Product details to append to the User-Agent string sent in all AWS API calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs#user_agent AwsccProvider#user_agent}
	UserAgent interface{} `field:"optional" json:"userAgent" yaml:"userAgent"`
}

