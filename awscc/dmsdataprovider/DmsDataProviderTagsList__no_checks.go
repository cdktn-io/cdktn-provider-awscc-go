// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dmsdataprovider

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DmsDataProviderTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DmsDataProviderTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DmsDataProviderTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DmsDataProviderTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DmsDataProviderTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DmsDataProviderTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DmsDataProviderTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDmsDataProviderTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

