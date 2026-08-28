// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wellarchitectedworkload


type WellarchitectedWorkloadDiscoveryConfig struct {
	// Discovery integration status in respect to Trusted Advisor for the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wellarchitected_workload#trusted_advisor_integration_status WellarchitectedWorkload#trusted_advisor_integration_status}
	TrustedAdvisorIntegrationStatus *string `field:"optional" json:"trustedAdvisorIntegrationStatus" yaml:"trustedAdvisorIntegrationStatus"`
	// The mode to use for identifying resources associated with the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wellarchitected_workload#workload_resource_definition WellarchitectedWorkload#workload_resource_definition}
	WorkloadResourceDefinition *[]*string `field:"optional" json:"workloadResourceDefinition" yaml:"workloadResourceDefinition"`
}

