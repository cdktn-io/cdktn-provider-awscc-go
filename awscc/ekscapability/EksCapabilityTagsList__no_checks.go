// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ekscapability

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EksCapabilityTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EksCapabilityTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EksCapabilityTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EksCapabilityTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EksCapabilityTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksCapabilityTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EksCapabilityTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEksCapabilityTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

