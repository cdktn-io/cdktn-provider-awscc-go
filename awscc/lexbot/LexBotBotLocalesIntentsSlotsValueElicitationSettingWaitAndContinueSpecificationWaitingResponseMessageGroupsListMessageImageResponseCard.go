// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationWaitingResponseMessageGroupsListMessageImageResponseCard struct {
	// A list of buttons that should be displayed on the response card.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lex_bot#buttons LexBot#buttons}
	Buttons interface{} `field:"optional" json:"buttons" yaml:"buttons"`
	// The URL of an image to display on the response card.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lex_bot#image_url LexBot#image_url}
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// The subtitle to display on the response card.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lex_bot#subtitle LexBot#subtitle}
	Subtitle *string `field:"optional" json:"subtitle" yaml:"subtitle"`
	// The title to display on the response card.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lex_bot#title LexBot#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

