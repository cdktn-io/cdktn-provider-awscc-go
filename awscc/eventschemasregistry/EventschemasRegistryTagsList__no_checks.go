// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package eventschemasregistry

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventschemasRegistryTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EventschemasRegistryTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EventschemasRegistryTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EventschemasRegistryTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EventschemasRegistryTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EventschemasRegistryTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EventschemasRegistryTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEventschemasRegistryTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

