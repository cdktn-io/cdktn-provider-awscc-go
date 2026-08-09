// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sesconfigurationset

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SesConfigurationSetTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SesConfigurationSetTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SesConfigurationSetTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SesConfigurationSetTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SesConfigurationSetTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SesConfigurationSetTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SesConfigurationSetTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSesConfigurationSetTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

