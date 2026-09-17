// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotTopicPublish struct {
	// The MQTT topic of the message.
	//
	// You can use a string expression that includes variables (``$variable.<variable-name>``) and input values (``$input.<input-name>.<path-to-datum>``) as the topic string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotevents_detector_model#mqtt_topic IoteventsDetectorModel#mqtt_topic}
	MqttTopic *string `field:"optional" json:"mqttTopic" yaml:"mqttTopic"`
	// You can configure the action payload when you publish a message to an IoTCore topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotevents_detector_model#payload IoteventsDetectorModel#payload}
	Payload *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotTopicPublishPayload `field:"optional" json:"payload" yaml:"payload"`
}

