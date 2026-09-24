// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotthingprincipalattachment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotThingPrincipalAttachmentConfig struct {
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
	// The principal, which can be a certificate ARN (as returned from the CreateCertificate operation) or an Amazon Cognito ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_thing_principal_attachment#principal IotThingPrincipalAttachment#principal}
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// The name of the AWS IoT thing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_thing_principal_attachment#thing_name IotThingPrincipalAttachment#thing_name}
	ThingName *string `field:"required" json:"thingName" yaml:"thingName"`
	// The type of the relation you want to specify when you attach a principal to a thing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_thing_principal_attachment#thing_principal_type IotThingPrincipalAttachment#thing_principal_type}
	ThingPrincipalType *string `field:"optional" json:"thingPrincipalType" yaml:"thingPrincipalType"`
}

