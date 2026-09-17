// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elementalinferencefeed


type ElementalinferenceFeedOutputsOutputConfigClippingDataSourceConfiguration struct {
	// The ID of the fixture whose event data you want Elemental Inference to map onto this clipping output.
	//
	// To obtain this ID, use the SearchFixtures operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elementalinference_feed#fixture_id ElementalinferenceFeed#fixture_id}
	FixtureId *string `field:"optional" json:"fixtureId" yaml:"fixtureId"`
}

