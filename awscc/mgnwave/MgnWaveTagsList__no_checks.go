// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package mgnwave

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MgnWaveTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MgnWaveTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MgnWaveTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MgnWaveTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MgnWaveTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MgnWaveTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MgnWaveTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMgnWaveTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

