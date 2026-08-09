// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package kendraindex

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KendraIndexTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (k *jsiiProxy_KendraIndexTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KendraIndexTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KendraIndexTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KendraIndexTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KendraIndexTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KendraIndexTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKendraIndexTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

