// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimeappinstancebot


type ChimeAppInstanceBotConfiguration struct {
	// The configuration for an Amazon Lex V2 bot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/chime_app_instance_bot#lex ChimeAppInstanceBot#lex}
	Lex *ChimeAppInstanceBotConfigurationLex `field:"required" json:"lex" yaml:"lex"`
}

