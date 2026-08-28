// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package locationapikey

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LocationApiKeyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/location_api_key#key_name LocationApiKey#key_name}.
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/location_api_key#restrictions LocationApiKey#restrictions}.
	Restrictions *LocationApiKeyRestrictions `field:"required" json:"restrictions" yaml:"restrictions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/location_api_key#description LocationApiKey#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/location_api_key#expire_time LocationApiKey#expire_time}
	ExpireTime *string `field:"optional" json:"expireTime" yaml:"expireTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/location_api_key#force_delete LocationApiKey#force_delete}.
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/location_api_key#force_update LocationApiKey#force_update}.
	ForceUpdate interface{} `field:"optional" json:"forceUpdate" yaml:"forceUpdate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/location_api_key#no_expiry LocationApiKey#no_expiry}.
	NoExpiry interface{} `field:"optional" json:"noExpiry" yaml:"noExpiry"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/location_api_key#tags LocationApiKey#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

