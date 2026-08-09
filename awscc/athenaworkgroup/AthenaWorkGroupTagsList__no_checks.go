// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package athenaworkgroup

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AthenaWorkGroupTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AthenaWorkGroupTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AthenaWorkGroupTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AthenaWorkGroupTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AthenaWorkGroupTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AthenaWorkGroupTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AthenaWorkGroupTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAthenaWorkGroupTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

