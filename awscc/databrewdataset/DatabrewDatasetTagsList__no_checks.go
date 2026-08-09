// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package databrewdataset

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabrewDatasetTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatabrewDatasetTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatabrewDatasetTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatabrewDatasetTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatabrewDatasetTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatabrewDatasetTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatabrewDatasetTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatabrewDatasetTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

