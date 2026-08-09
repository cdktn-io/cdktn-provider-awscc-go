// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package logsdelivery

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogsDeliveryTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LogsDeliveryTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LogsDeliveryTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LogsDeliveryTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LogsDeliveryTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LogsDeliveryTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LogsDeliveryTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLogsDeliveryTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

