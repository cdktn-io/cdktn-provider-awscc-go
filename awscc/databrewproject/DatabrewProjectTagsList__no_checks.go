// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package databrewproject

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabrewProjectTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatabrewProjectTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatabrewProjectTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatabrewProjectTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatabrewProjectTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatabrewProjectTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatabrewProjectTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatabrewProjectTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

