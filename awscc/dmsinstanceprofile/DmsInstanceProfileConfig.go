// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsinstanceprofile

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DmsInstanceProfileConfig struct {
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
	// The property describes an availability zone of the instance profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_instance_profile#availability_zone DmsInstanceProfile#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The optional description of the instance profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_instance_profile#description DmsInstanceProfile#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The property describes an identifier for the instance profile. It is used for describing/deleting/modifying. Can be name/arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_instance_profile#instance_profile_identifier DmsInstanceProfile#instance_profile_identifier}
	InstanceProfileIdentifier *string `field:"optional" json:"instanceProfileIdentifier" yaml:"instanceProfileIdentifier"`
	// The property describes a name for the instance profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_instance_profile#instance_profile_name DmsInstanceProfile#instance_profile_name}
	InstanceProfileName *string `field:"optional" json:"instanceProfileName" yaml:"instanceProfileName"`
	// The property describes kms key arn for the instance profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_instance_profile#kms_key_arn DmsInstanceProfile#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The property describes a network type for the instance profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_instance_profile#network_type DmsInstanceProfile#network_type}
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// The property describes the publicly accessible of the instance profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_instance_profile#publicly_accessible DmsInstanceProfile#publicly_accessible}
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The property describes a subnet group identifier for the instance profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_instance_profile#subnet_group_identifier DmsInstanceProfile#subnet_group_identifier}
	SubnetGroupIdentifier *string `field:"optional" json:"subnetGroupIdentifier" yaml:"subnetGroupIdentifier"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_instance_profile#tags DmsInstanceProfile#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The property describes vps security groups for the instance profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/dms_instance_profile#vpc_security_groups DmsInstanceProfile#vpc_security_groups}
	VpcSecurityGroups *[]*string `field:"optional" json:"vpcSecurityGroups" yaml:"vpcSecurityGroups"`
}

