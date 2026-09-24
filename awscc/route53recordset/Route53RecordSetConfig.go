// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53recordset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Route53RecordSetConfig struct {
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
	// The name of the record that you want to create, update, or delete.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_record_set#name Route53RecordSet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The DNS record type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_record_set#type Route53RecordSet#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Alias resource record sets only: Information about the AWS resource, such as a CloudFront distribution or an Amazon S3 bucket, that you want to route traffic to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_record_set#alias_target Route53RecordSet#alias_target}
	AliasTarget *Route53RecordSetAliasTarget `field:"optional" json:"aliasTarget" yaml:"aliasTarget"`
	// The object that is specified in resource record set object when you are linking a resource record set to a CIDR location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_record_set#cidr_routing_config Route53RecordSet#cidr_routing_config}
	CidrRoutingConfig *Route53RecordSetCidrRoutingConfig `field:"optional" json:"cidrRoutingConfig" yaml:"cidrRoutingConfig"`
	// Optional: Any comments you want to include about a change batch request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_record_set#comment Route53RecordSet#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// To configure failover, you add the Failover element to two resource record sets.
	//
	// For one resource record set, you specify PRIMARY as the value for Failover; for the other resource record set, you specify SECONDARY. In addition, you include the HealthCheckId element and specify the health check that you want Amazon Route 53 to perform for each resource record set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_record_set#failover Route53RecordSet#failover}
	Failover *string `field:"optional" json:"failover" yaml:"failover"`
	// A complex type that lets you control how Amazon Route 53 responds to DNS queries based on the geographic origin of the query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_record_set#geo_location Route53RecordSet#geo_location}
	GeoLocation *Route53RecordSetGeoLocation `field:"optional" json:"geoLocation" yaml:"geoLocation"`
	// If you want Amazon Route 53 to return this resource record set in response to a DNS query only when the status of a health check is healthy, include the HealthCheckId element and specify the ID of the applicable health check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_record_set#health_check_id Route53RecordSet#health_check_id}
	HealthCheckId *string `field:"optional" json:"healthCheckId" yaml:"healthCheckId"`
	// The ID of the hosted zone that you want to create records in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_record_set#hosted_zone_id Route53RecordSet#hosted_zone_id}
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
	// The name of the hosted zone that you want to create records in.
	//
	// You must include a trailing dot (for example, www.example.com.) as part of the HostedZoneName.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_record_set#hosted_zone_name Route53RecordSet#hosted_zone_name}
	HostedZoneName *string `field:"optional" json:"hostedZoneName" yaml:"hostedZoneName"`
	// To route traffic approximately randomly to multiple resources, such as web servers, create one multivalue answer record for each resource and specify true for MultiValueAnswer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_record_set#multi_value_answer Route53RecordSet#multi_value_answer}
	MultiValueAnswer interface{} `field:"optional" json:"multiValueAnswer" yaml:"multiValueAnswer"`
	// The Amazon EC2 Region where you created the resource that this resource record set refers to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_record_set#region Route53RecordSet#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// One or more values that correspond with the value that you specified for the Type property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_record_set#resource_records Route53RecordSet#resource_records}
	ResourceRecords *[]*string `field:"optional" json:"resourceRecords" yaml:"resourceRecords"`
	// An identifier that differentiates among multiple resource record sets that have the same combination of name and type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_record_set#set_identifier Route53RecordSet#set_identifier}
	SetIdentifier *string `field:"optional" json:"setIdentifier" yaml:"setIdentifier"`
	// The resource record cache time to live (TTL), in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_record_set#ttl Route53RecordSet#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
	// Among resource record sets that have the same combination of DNS name and type, a value that determines the proportion of DNS queries that Amazon Route 53 responds to using the current resource record set.
	//
	// Route 53 calculates the sum of the weights for the resource record sets that have the same combination of DNS name and type. Route 53 then responds to queries based on the ratio of a resource's weight to the total.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/route53_record_set#weight Route53RecordSet#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

