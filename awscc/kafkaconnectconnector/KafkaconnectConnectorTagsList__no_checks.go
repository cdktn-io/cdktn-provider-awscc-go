// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package kafkaconnectconnector

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KafkaconnectConnectorTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (k *jsiiProxy_KafkaconnectConnectorTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KafkaconnectConnectorTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KafkaconnectConnectorTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KafkaconnectConnectorTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KafkaconnectConnectorTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KafkaconnectConnectorTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKafkaconnectConnectorTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

