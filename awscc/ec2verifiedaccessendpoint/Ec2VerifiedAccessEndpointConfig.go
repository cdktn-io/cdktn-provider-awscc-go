// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2verifiedaccessendpoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2VerifiedAccessEndpointConfig struct {
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
	// The type of attachment used to provide connectivity between the AWS Verified Access endpoint and the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_endpoint#attachment_type Ec2VerifiedAccessEndpoint#attachment_type}
	AttachmentType *string `field:"required" json:"attachmentType" yaml:"attachmentType"`
	// The type of AWS Verified Access endpoint.
	//
	// Incoming application requests will be sent to an IP address, load balancer or a network interface depending on the endpoint type specified.The type of AWS Verified Access endpoint. Incoming application requests will be sent to an IP address, load balancer or a network interface depending on the endpoint type specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_endpoint#endpoint_type Ec2VerifiedAccessEndpoint#endpoint_type}
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// The ID of the AWS Verified Access group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_endpoint#verified_access_group_id Ec2VerifiedAccessEndpoint#verified_access_group_id}
	VerifiedAccessGroupId *string `field:"required" json:"verifiedAccessGroupId" yaml:"verifiedAccessGroupId"`
	// The DNS name for users to reach your application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_endpoint#application_domain Ec2VerifiedAccessEndpoint#application_domain}
	ApplicationDomain *string `field:"optional" json:"applicationDomain" yaml:"applicationDomain"`
	// The options for cidr type endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_endpoint#cidr_options Ec2VerifiedAccessEndpoint#cidr_options}
	CidrOptions *Ec2VerifiedAccessEndpointCidrOptions `field:"optional" json:"cidrOptions" yaml:"cidrOptions"`
	// A description for the AWS Verified Access endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_endpoint#description Ec2VerifiedAccessEndpoint#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ARN of a public TLS/SSL certificate imported into or created with ACM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_endpoint#domain_certificate_arn Ec2VerifiedAccessEndpoint#domain_certificate_arn}
	DomainCertificateArn *string `field:"optional" json:"domainCertificateArn" yaml:"domainCertificateArn"`
	// A custom identifier that gets prepended to a DNS name that is generated for the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_endpoint#endpoint_domain_prefix Ec2VerifiedAccessEndpoint#endpoint_domain_prefix}
	EndpointDomainPrefix *string `field:"optional" json:"endpointDomainPrefix" yaml:"endpointDomainPrefix"`
	// The load balancer details if creating the AWS Verified Access endpoint as load-balancer type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_endpoint#load_balancer_options Ec2VerifiedAccessEndpoint#load_balancer_options}
	LoadBalancerOptions *Ec2VerifiedAccessEndpointLoadBalancerOptions `field:"optional" json:"loadBalancerOptions" yaml:"loadBalancerOptions"`
	// The options for network-interface type endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_endpoint#network_interface_options Ec2VerifiedAccessEndpoint#network_interface_options}
	NetworkInterfaceOptions *Ec2VerifiedAccessEndpointNetworkInterfaceOptions `field:"optional" json:"networkInterfaceOptions" yaml:"networkInterfaceOptions"`
	// The AWS Verified Access policy document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_endpoint#policy_document Ec2VerifiedAccessEndpoint#policy_document}
	PolicyDocument *string `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// The status of the Verified Access policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_endpoint#policy_enabled Ec2VerifiedAccessEndpoint#policy_enabled}
	PolicyEnabled interface{} `field:"optional" json:"policyEnabled" yaml:"policyEnabled"`
	// The options for rds type endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_endpoint#rds_options Ec2VerifiedAccessEndpoint#rds_options}
	RdsOptions *Ec2VerifiedAccessEndpointRdsOptions `field:"optional" json:"rdsOptions" yaml:"rdsOptions"`
	// The IDs of the security groups for the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_endpoint#security_group_ids Ec2VerifiedAccessEndpoint#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The configuration options for customer provided KMS encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_endpoint#sse_specification Ec2VerifiedAccessEndpoint#sse_specification}
	SseSpecification *Ec2VerifiedAccessEndpointSseSpecification `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_verified_access_endpoint#tags Ec2VerifiedAccessEndpoint#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

