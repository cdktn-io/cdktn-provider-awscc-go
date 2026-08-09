// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package eventseventbus

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventsEventBusTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EventsEventBusTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EventsEventBusTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EventsEventBusTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EventsEventBusTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EventsEventBusTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EventsEventBusTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEventsEventBusTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

