// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncapi


type AppsyncApiEventConfig struct {
	// A list of auth providers for the AppSync API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appsync_api#auth_providers AppsyncApi#auth_providers}
	AuthProviders interface{} `field:"optional" json:"authProviders" yaml:"authProviders"`
	// A list of auth modes for the AppSync API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appsync_api#connection_auth_modes AppsyncApi#connection_auth_modes}
	ConnectionAuthModes interface{} `field:"optional" json:"connectionAuthModes" yaml:"connectionAuthModes"`
	// A list of auth modes for the AppSync API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appsync_api#default_publish_auth_modes AppsyncApi#default_publish_auth_modes}
	DefaultPublishAuthModes interface{} `field:"optional" json:"defaultPublishAuthModes" yaml:"defaultPublishAuthModes"`
	// A list of auth modes for the AppSync API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appsync_api#default_subscribe_auth_modes AppsyncApi#default_subscribe_auth_modes}
	DefaultSubscribeAuthModes interface{} `field:"optional" json:"defaultSubscribeAuthModes" yaml:"defaultSubscribeAuthModes"`
	// The log config for the AppSync API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appsync_api#log_config AppsyncApi#log_config}
	LogConfig *AppsyncApiEventConfigLogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
}

