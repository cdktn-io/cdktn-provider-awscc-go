// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorsourcelocation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediatailorSourceLocationConfig struct {
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
	// <p>The HTTP configuration for the source location.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediatailor_source_location#http_configuration MediatailorSourceLocation#http_configuration}
	HttpConfiguration *MediatailorSourceLocationHttpConfiguration `field:"required" json:"httpConfiguration" yaml:"httpConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediatailor_source_location#source_location_name MediatailorSourceLocation#source_location_name}.
	SourceLocationName *string `field:"required" json:"sourceLocationName" yaml:"sourceLocationName"`
	// <p>Access configuration parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediatailor_source_location#access_configuration MediatailorSourceLocation#access_configuration}
	AccessConfiguration *MediatailorSourceLocationAccessConfiguration `field:"optional" json:"accessConfiguration" yaml:"accessConfiguration"`
	// <p>The optional configuration for a server that serves segments.
	//
	// Use this if you want the segment delivery server to be different from the source location server. For example, you can configure your source location server to be an origination server, such as MediaPackage, and the segment delivery server to be a content delivery network (CDN), such as CloudFront. If you don't specify a segment delivery server, then the source location server is used.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediatailor_source_location#default_segment_delivery_configuration MediatailorSourceLocation#default_segment_delivery_configuration}
	DefaultSegmentDeliveryConfiguration *MediatailorSourceLocationDefaultSegmentDeliveryConfiguration `field:"optional" json:"defaultSegmentDeliveryConfiguration" yaml:"defaultSegmentDeliveryConfiguration"`
	// <p>A list of the segment delivery configurations associated with this resource.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediatailor_source_location#segment_delivery_configurations MediatailorSourceLocation#segment_delivery_configurations}
	SegmentDeliveryConfigurations interface{} `field:"optional" json:"segmentDeliveryConfigurations" yaml:"segmentDeliveryConfigurations"`
	// The tags to assign to the source location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediatailor_source_location#tags MediatailorSourceLocation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

