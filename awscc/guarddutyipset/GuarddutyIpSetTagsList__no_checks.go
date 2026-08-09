// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package guarddutyipset

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GuarddutyIpSetTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GuarddutyIpSetTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GuarddutyIpSetTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GuarddutyIpSetTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GuarddutyIpSetTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GuarddutyIpSetTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GuarddutyIpSetTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGuarddutyIpSetTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

