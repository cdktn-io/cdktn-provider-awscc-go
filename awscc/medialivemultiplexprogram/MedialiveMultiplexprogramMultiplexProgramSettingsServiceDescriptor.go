// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialivemultiplexprogram


type MedialiveMultiplexprogramMultiplexProgramSettingsServiceDescriptor struct {
	// Name of the provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/medialive_multiplexprogram#provider_name MedialiveMultiplexprogram#provider_name}
	ProviderName *string `field:"optional" json:"providerName" yaml:"providerName"`
	// Name of the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/medialive_multiplexprogram#service_name MedialiveMultiplexprogram#service_name}
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

