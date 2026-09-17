// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package maciemember

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MacieMemberTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MacieMemberTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MacieMemberTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MacieMemberTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MacieMemberTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MacieMemberTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MacieMemberTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMacieMemberTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

