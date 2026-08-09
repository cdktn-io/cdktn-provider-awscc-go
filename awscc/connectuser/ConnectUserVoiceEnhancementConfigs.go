// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectuser


type ConnectUserVoiceEnhancementConfigs struct {
	// The channels that agents can handle in the Contact Control Panel (CCP).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_user#channel ConnectUser#channel}
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// The Voice Enhancement Mode setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_user#voice_enhancement_mode ConnectUser#voice_enhancement_mode}
	VoiceEnhancementMode *string `field:"optional" json:"voiceEnhancementMode" yaml:"voiceEnhancementMode"`
}

