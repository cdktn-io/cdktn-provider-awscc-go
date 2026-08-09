// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncchannelnamespace

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppsyncChannelNamespaceConfig struct {
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
	// AppSync Api Id that this Channel Namespace belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_channel_namespace#api_id AppsyncChannelNamespace#api_id}
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Namespace indentifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_channel_namespace#name AppsyncChannelNamespace#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// String of APPSYNC_JS code to be used by the handlers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_channel_namespace#code_handlers AppsyncChannelNamespace#code_handlers}
	CodeHandlers *string `field:"optional" json:"codeHandlers" yaml:"codeHandlers"`
	// The Amazon S3 endpoint where the code is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_channel_namespace#code_s3_location AppsyncChannelNamespace#code_s3_location}
	CodeS3Location *string `field:"optional" json:"codeS3Location" yaml:"codeS3Location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_channel_namespace#handler_configs AppsyncChannelNamespace#handler_configs}.
	HandlerConfigs *AppsyncChannelNamespaceHandlerConfigs `field:"optional" json:"handlerConfigs" yaml:"handlerConfigs"`
	// List of AuthModes supported for Publish operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_channel_namespace#publish_auth_modes AppsyncChannelNamespace#publish_auth_modes}
	PublishAuthModes interface{} `field:"optional" json:"publishAuthModes" yaml:"publishAuthModes"`
	// List of AuthModes supported for Subscribe operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_channel_namespace#subscribe_auth_modes AppsyncChannelNamespace#subscribe_auth_modes}
	SubscribeAuthModes interface{} `field:"optional" json:"subscribeAuthModes" yaml:"subscribeAuthModes"`
	// An arbitrary set of tags (key-value pairs) for this AppSync API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appsync_channel_namespace#tags AppsyncChannelNamespace#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

