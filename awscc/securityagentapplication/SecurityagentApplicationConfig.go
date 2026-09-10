// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentapplication

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityagentApplicationConfig struct {
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
	// Identifier of a KMS key. Can be a key ID, key ARN, alias name, or alias ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityagent_application#default_kms_key_id SecurityagentApplication#default_kms_key_id}
	DefaultKmsKeyId *string `field:"optional" json:"defaultKmsKeyId" yaml:"defaultKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityagent_application#id_c_configuration SecurityagentApplication#id_c_configuration}.
	IdCConfiguration *SecurityagentApplicationIdCConfiguration `field:"optional" json:"idCConfiguration" yaml:"idCConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityagent_application#role_arn SecurityagentApplication#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Tags for the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityagent_application#tags SecurityagentApplication#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

