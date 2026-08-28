// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontconnectionfunction

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudfrontConnectionFunctionConfig struct {
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
	// The code for the connection function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_connection_function#connection_function_code CloudfrontConnectionFunction#connection_function_code}
	ConnectionFunctionCode *string `field:"required" json:"connectionFunctionCode" yaml:"connectionFunctionCode"`
	// Contains configuration information about a CloudFront function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_connection_function#connection_function_config CloudfrontConnectionFunction#connection_function_config}
	ConnectionFunctionConfig *CloudfrontConnectionFunctionConnectionFunctionConfig `field:"required" json:"connectionFunctionConfig" yaml:"connectionFunctionConfig"`
	// The connection function name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_connection_function#name CloudfrontConnectionFunction#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A flag that determines whether to automatically publish the function to the ``LIVE`` stage when it?s created. To automatically publish to the ``LIVE`` stage, set this property to ``true``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_connection_function#auto_publish CloudfrontConnectionFunction#auto_publish}
	AutoPublish interface{} `field:"optional" json:"autoPublish" yaml:"autoPublish"`
	// A complex type that contains zero or more ``Tag`` elements.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_connection_function#tags CloudfrontConnectionFunction#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

