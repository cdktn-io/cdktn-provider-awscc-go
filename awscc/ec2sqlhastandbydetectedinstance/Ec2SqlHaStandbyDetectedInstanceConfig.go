// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2sqlhastandbydetectedinstance

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2SqlHaStandbyDetectedInstanceConfig struct {
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
	// The ID of the EC2 instance to enable for SQL Server high availability standby detection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_sql_ha_standby_detected_instance#instance_id Ec2SqlHaStandbyDetectedInstance#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The ARN of the AWS Secrets Manager secret containing SQL Server access credentials to the EC2 instance.
	//
	// If not specified, AWS Systems Manager agent will use default local user credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_sql_ha_standby_detected_instance#sql_server_credentials Ec2SqlHaStandbyDetectedInstance#sql_server_credentials}
	SqlServerCredentials *string `field:"optional" json:"sqlServerCredentials" yaml:"sqlServerCredentials"`
}

