// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cloudformationchangeset

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudformationChangeSetTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CloudformationChangeSetTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudformationChangeSetTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudformationChangeSetTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudformationChangeSetTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudformationChangeSetTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudformationChangeSetTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudformationChangeSetTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

