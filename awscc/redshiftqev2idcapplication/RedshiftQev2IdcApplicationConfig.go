// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftqev2idcapplication

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RedshiftQev2IdcApplicationConfig struct {
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
	// The display name for the Amazon Redshift Query Editor (QEV2) IAM Identity Center application. It appears in the console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/redshift_qev2_idc_application#idc_display_name RedshiftQev2IdcApplication#idc_display_name}
	IdcDisplayName *string `field:"required" json:"idcDisplayName" yaml:"idcDisplayName"`
	// The Amazon Resource Name (ARN) of the IAM Identity Center instance used to create the Amazon Redshift Query Editor (QEV2) managed application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/redshift_qev2_idc_application#idc_instance_arn RedshiftQev2IdcApplication#idc_instance_arn}
	IdcInstanceArn *string `field:"required" json:"idcInstanceArn" yaml:"idcInstanceArn"`
	// The name of the Amazon Redshift Query Editor (QEV2) application in IAM Identity Center.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/redshift_qev2_idc_application#qev_2_idc_application_name RedshiftQev2IdcApplication#qev_2_idc_application_name}
	Qev2IdcApplicationName *string `field:"required" json:"qev2IdcApplicationName" yaml:"qev2IdcApplicationName"`
	// A list of tags associated with the application.
	//
	// Tags are key-value pairs that you can use to organize and identify your resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/redshift_qev2_idc_application#tags RedshiftQev2IdcApplication#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

