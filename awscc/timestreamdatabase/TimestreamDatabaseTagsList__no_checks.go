// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package timestreamdatabase

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TimestreamDatabaseTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TimestreamDatabaseTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TimestreamDatabaseTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TimestreamDatabaseTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TimestreamDatabaseTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TimestreamDatabaseTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TimestreamDatabaseTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTimestreamDatabaseTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

