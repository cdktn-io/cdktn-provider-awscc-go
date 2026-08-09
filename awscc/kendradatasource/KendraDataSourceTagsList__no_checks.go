// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package kendradatasource

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KendraDataSourceTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (k *jsiiProxy_KendraDataSourceTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KendraDataSourceTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KendraDataSourceTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KendraDataSourceTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KendraDataSourceTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KendraDataSourceTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKendraDataSourceTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

