// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package configstoredquery

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigStoredQueryTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ConfigStoredQueryTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConfigStoredQueryTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConfigStoredQueryTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigStoredQueryTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigStoredQueryTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConfigStoredQueryTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConfigStoredQueryTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

