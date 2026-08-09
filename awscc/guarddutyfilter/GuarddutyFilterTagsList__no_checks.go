// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package guarddutyfilter

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GuarddutyFilterTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GuarddutyFilterTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GuarddutyFilterTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GuarddutyFilterTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GuarddutyFilterTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GuarddutyFilterTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GuarddutyFilterTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGuarddutyFilterTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

