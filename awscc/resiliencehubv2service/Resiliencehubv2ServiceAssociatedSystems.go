// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2service


type Resiliencehubv2ServiceAssociatedSystems struct {
	// The system ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/resiliencehubv2_service#system_arn Resiliencehubv2Service#system_arn}
	SystemArn *string `field:"optional" json:"systemArn" yaml:"systemArn"`
	// User journey IDs associated with this system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/resiliencehubv2_service#user_journey_ids Resiliencehubv2Service#user_journey_ids}
	UserJourneyIds *[]*string `field:"optional" json:"userJourneyIds" yaml:"userJourneyIds"`
}

