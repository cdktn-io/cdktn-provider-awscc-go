// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprogram


type MediatailorProgramAudienceMedia struct {
	// The list of AlternateMedia defined in AudienceMedia.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#alternate_media MediatailorProgram#alternate_media}
	AlternateMedia interface{} `field:"optional" json:"alternateMedia" yaml:"alternateMedia"`
	// The Audience defined in AudienceMedia.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#audience MediatailorProgram#audience}
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
}

