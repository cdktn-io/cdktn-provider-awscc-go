// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package eventsendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventsEndpointEventBusesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EventsEndpointEventBusesList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EventsEndpointEventBusesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EventsEndpointEventBusesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EventsEndpointEventBusesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EventsEndpointEventBusesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EventsEndpointEventBusesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEventsEndpointEventBusesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

