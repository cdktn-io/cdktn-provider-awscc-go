// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectdatalakeassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectDataLakeAssociationConfig struct {
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
	// The identifier of the analytics data set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_data_lake_association#data_set_id ConnectDataLakeAssociation#data_set_id}
	DataSetId *string `field:"required" json:"dataSetId" yaml:"dataSetId"`
	// The identifier of the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_data_lake_association#instance_id ConnectDataLakeAssociation#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The identifier of the target account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_data_lake_association#target_account_id ConnectDataLakeAssociation#target_account_id}
	TargetAccountId *string `field:"optional" json:"targetAccountId" yaml:"targetAccountId"`
}

