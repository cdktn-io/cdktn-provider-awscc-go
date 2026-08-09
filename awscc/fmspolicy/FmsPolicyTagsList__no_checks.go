// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package fmspolicy

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FmsPolicyTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FmsPolicyTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FmsPolicyTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FmsPolicyTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FmsPolicyTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FmsPolicyTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FmsPolicyTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFmsPolicyTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

