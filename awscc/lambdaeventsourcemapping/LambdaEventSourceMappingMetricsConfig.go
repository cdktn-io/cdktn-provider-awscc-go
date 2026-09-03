// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdaeventsourcemapping


type LambdaEventSourceMappingMetricsConfig struct {
	// The metrics you want your event source mapping to produce, including ``EventCount``, ``ErrorCount``, ``KafkaMetrics``.
	//
	// +  ``EventCount`` to receive metrics related to the number of events processed by your event source mapping.
	//   +  ``ErrorCount`` (Amazon MSK and self-managed Apache Kafka) to receive metrics related to the number of errors in your event source mapping processing.
	//   +  ``KafkaMetrics`` (Amazon MSK and self-managed Apache Kafka) to receive metrics related to the Kafka consumers from your event source mapping.
	//
	//   For more information about these metrics, see [Event source mapping metrics](https://docs.aws.amazon.com/lambda/latest/dg/monitoring-metrics-types.html#event-source-mapping-metrics).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lambda_event_source_mapping#metrics LambdaEventSourceMapping#metrics}
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
}

