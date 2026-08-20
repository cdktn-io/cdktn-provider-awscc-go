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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53_record_set#name Route53RecordSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53_record_set#type Route53RecordSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53_record_set#alias_target Route53RecordSet#alias_target}.
	AliasTarget *Route53RecordSetAliasTarget `field:"optional" json:"aliasTarget" yaml:"aliasTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53_record_set#cidr_routing_config Route53RecordSet#cidr_routing_config}.
	CidrRoutingConfig *Route53RecordSetCidrRoutingConfig `field:"optional" json:"cidrRoutingConfig" yaml:"cidrRoutingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53_record_set#comment Route53RecordSet#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53_record_set#failover Route53RecordSet#failover}.
	Failover *string `field:"optional" json:"failover" yaml:"failover"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53_record_set#geo_location Route53RecordSet#geo_location}.
	GeoLocation *Route53RecordSetGeoLocation `field:"optional" json:"geoLocation" yaml:"geoLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53_record_set#geo_proximity_location Route53RecordSet#geo_proximity_location}.
	GeoProximityLocation *Route53RecordSetGeoProximityLocation `field:"optional" json:"geoProximityLocation" yaml:"geoProximityLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53_record_set#health_check_id Route53RecordSet#health_check_id}.
	HealthCheckId *string `field:"optional" json:"healthCheckId" yaml:"healthCheckId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53_record_set#hosted_zone_id Route53RecordSet#hosted_zone_id}.
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53_record_set#hosted_zone_name Route53RecordSet#hosted_zone_name}.
	HostedZoneName *string `field:"optional" json:"hostedZoneName" yaml:"hostedZoneName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53_record_set#multi_value_answer Route53RecordSet#multi_value_answer}.
	MultiValueAnswer interface{} `field:"optional" json:"multiValueAnswer" yaml:"multiValueAnswer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53_record_set#region Route53RecordSet#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53_record_set#resource_records Route53RecordSet#resource_records}.
	ResourceRecords *[]*string `field:"optional" json:"resourceRecords" yaml:"resourceRecords"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53_record_set#set_identifier Route53RecordSet#set_identifier}.
	SetIdentifier *string `field:"optional" json:"setIdentifier" yaml:"setIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53_record_set#ttl Route53RecordSet#ttl}.
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/route53_record_set#weight Route53RecordSet#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

