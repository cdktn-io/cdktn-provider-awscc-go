// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2domainname

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Apigatewayv2DomainNameConfig struct {
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
	// The custom domain name for your API in Amazon API Gateway.
	//
	// Uppercase letters and the underscore (``_``) character are not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigatewayv2_domain_name#domain_name Apigatewayv2DomainName#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The domain name configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigatewayv2_domain_name#domain_name_configurations Apigatewayv2DomainName#domain_name_configurations}
	DomainNameConfigurations interface{} `field:"optional" json:"domainNameConfigurations" yaml:"domainNameConfigurations"`
	// The mutual TLS authentication configuration for a custom domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigatewayv2_domain_name#mutual_tls_authentication Apigatewayv2DomainName#mutual_tls_authentication}
	MutualTlsAuthentication *Apigatewayv2DomainNameMutualTlsAuthentication `field:"optional" json:"mutualTlsAuthentication" yaml:"mutualTlsAuthentication"`
	// The routing mode API Gateway uses to route traffic to your APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigatewayv2_domain_name#routing_mode Apigatewayv2DomainName#routing_mode}
	RoutingMode *string `field:"optional" json:"routingMode" yaml:"routingMode"`
	// The collection of tags associated with a domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigatewayv2_domain_name#tags Apigatewayv2DomainName#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

