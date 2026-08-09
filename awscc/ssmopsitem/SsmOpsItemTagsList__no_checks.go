// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ssmopsitem

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SsmOpsItemTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SsmOpsItemTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SsmOpsItemTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SsmOpsItemTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SsmOpsItemTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SsmOpsItemTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SsmOpsItemTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSsmOpsItemTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

