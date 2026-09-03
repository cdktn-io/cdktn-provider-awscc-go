// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package globalacceleratorlistener


type GlobalacceleratorListenerPortRanges struct {
	// A network port number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/globalaccelerator_listener#from_port GlobalacceleratorListener#from_port}
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// A network port number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/globalaccelerator_listener#to_port GlobalacceleratorListener#to_port}
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
}

