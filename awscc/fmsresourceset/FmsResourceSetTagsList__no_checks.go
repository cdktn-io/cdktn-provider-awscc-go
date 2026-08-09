// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package fmsresourceset

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FmsResourceSetTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FmsResourceSetTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FmsResourceSetTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FmsResourceSetTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FmsResourceSetTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FmsResourceSetTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FmsResourceSetTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFmsResourceSetTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

