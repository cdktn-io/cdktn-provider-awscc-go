// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package kinesisstreamconsumer

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KinesisStreamConsumerTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (k *jsiiProxy_KinesisStreamConsumerTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KinesisStreamConsumerTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KinesisStreamConsumerTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KinesisStreamConsumerTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KinesisStreamConsumerTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KinesisStreamConsumerTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKinesisStreamConsumerTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

