// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicediscoveryservice

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServicediscoveryServiceConfig struct {
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
	// A description for the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/servicediscovery_service#description ServicediscoveryService#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// DNS-related configurations for the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/servicediscovery_service#dns_config ServicediscoveryService#dns_config}
	DnsConfig *ServicediscoveryServiceDnsConfig `field:"optional" json:"dnsConfig" yaml:"dnsConfig"`
	// Settings for health checks. Used when routing is DNS-based.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/servicediscovery_service#health_check_config ServicediscoveryService#health_check_config}
	HealthCheckConfig *ServicediscoveryServiceHealthCheckConfig `field:"optional" json:"healthCheckConfig" yaml:"healthCheckConfig"`
	// Settings for custom health checks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/servicediscovery_service#health_check_custom_config ServicediscoveryService#health_check_custom_config}
	HealthCheckCustomConfig *ServicediscoveryServiceHealthCheckCustomConfig `field:"optional" json:"healthCheckCustomConfig" yaml:"healthCheckCustomConfig"`
	// The name of the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/servicediscovery_service#name ServicediscoveryService#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID of the namespace in which the service is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/servicediscovery_service#namespace_id ServicediscoveryService#namespace_id}
	NamespaceId *string `field:"optional" json:"namespaceId" yaml:"namespaceId"`
	// A string map that contains attributes and values for the service.
	//
	// You can specify a maximum of 30 key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/servicediscovery_service#service_attributes ServicediscoveryService#service_attributes}
	ServiceAttributes *map[string]*string `field:"optional" json:"serviceAttributes" yaml:"serviceAttributes"`
	// An array of key-value pairs to associate with the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/servicediscovery_service#tags ServicediscoveryService#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The type of service. Supported values are HTTP or DNS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/servicediscovery_service#type ServicediscoveryService#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

