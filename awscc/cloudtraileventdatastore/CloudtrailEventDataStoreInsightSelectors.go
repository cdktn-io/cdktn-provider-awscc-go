// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudtraileventdatastore


type CloudtrailEventDataStoreInsightSelectors struct {
	// The type of Insights to log on an event data store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudtrail_event_data_store#insight_type CloudtrailEventDataStore#insight_type}
	InsightType *string `field:"optional" json:"insightType" yaml:"insightType"`
}

