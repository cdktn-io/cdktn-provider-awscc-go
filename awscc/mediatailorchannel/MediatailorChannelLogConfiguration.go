// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorchannel


type MediatailorChannelLogConfiguration struct {
	// <p>The log types.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediatailor_channel#log_types MediatailorChannel#log_types}
	LogTypes *[]*string `field:"optional" json:"logTypes" yaml:"logTypes"`
}

