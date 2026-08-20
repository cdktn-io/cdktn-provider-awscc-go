// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package braketspendinglimit

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BraketSpendingLimitConfig struct {
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
	// The Amazon Resource Name (ARN) of the quantum device to apply the spending limit to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/braket_spending_limit#device_arn BraketSpendingLimit#device_arn}
	DeviceArn *string `field:"required" json:"deviceArn" yaml:"deviceArn"`
	// The maximum amount that can be spent on the specified device, in USD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/braket_spending_limit#spending_limit BraketSpendingLimit#spending_limit}
	SpendingLimit *string `field:"required" json:"spendingLimit" yaml:"spendingLimit"`
	// The tags to apply to the spending limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/braket_spending_limit#tags BraketSpendingLimit#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Defines a time range for spending limits, specifying when the limit is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/braket_spending_limit#time_period BraketSpendingLimit#time_period}
	TimePeriod *BraketSpendingLimitTimePeriod `field:"optional" json:"timePeriod" yaml:"timePeriod"`
}

