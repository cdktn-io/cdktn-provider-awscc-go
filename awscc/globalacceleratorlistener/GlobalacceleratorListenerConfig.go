// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package globalacceleratorlistener

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlobalacceleratorListenerConfig struct {
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
	// The Amazon Resource Name (ARN) of the accelerator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/globalaccelerator_listener#accelerator_arn GlobalacceleratorListener#accelerator_arn}
	AcceleratorArn *string `field:"required" json:"acceleratorArn" yaml:"acceleratorArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/globalaccelerator_listener#port_ranges GlobalacceleratorListener#port_ranges}.
	PortRanges interface{} `field:"required" json:"portRanges" yaml:"portRanges"`
	// Client affinity lets you direct all requests from a user to the same endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/globalaccelerator_listener#client_affinity GlobalacceleratorListener#client_affinity}
	ClientAffinity *string `field:"optional" json:"clientAffinity" yaml:"clientAffinity"`
	// The protocol for the listener.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/globalaccelerator_listener#protocol GlobalacceleratorListener#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

