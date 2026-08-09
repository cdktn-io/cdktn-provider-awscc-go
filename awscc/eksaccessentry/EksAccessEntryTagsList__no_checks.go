// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package eksaccessentry

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EksAccessEntryTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EksAccessEntryTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EksAccessEntryTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EksAccessEntryTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EksAccessEntryTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksAccessEntryTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EksAccessEntryTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEksAccessEntryTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

