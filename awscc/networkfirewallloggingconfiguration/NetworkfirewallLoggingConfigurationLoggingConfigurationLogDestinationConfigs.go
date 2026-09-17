// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkfirewallloggingconfiguration


type NetworkfirewallLoggingConfigurationLoggingConfigurationLogDestinationConfigs struct {
	// A key-value pair to configure the logDestinations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkfirewall_logging_configuration#log_destination NetworkfirewallLoggingConfiguration#log_destination}
	LogDestination *map[string]*string `field:"required" json:"logDestination" yaml:"logDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkfirewall_logging_configuration#log_destination_type NetworkfirewallLoggingConfiguration#log_destination_type}.
	LogDestinationType *string `field:"required" json:"logDestinationType" yaml:"logDestinationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/networkfirewall_logging_configuration#log_type NetworkfirewallLoggingConfiguration#log_type}.
	LogType *string `field:"required" json:"logType" yaml:"logType"`
}

