// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package chimeappinstance

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ChimeAppInstanceTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ChimeAppInstanceTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ChimeAppInstanceTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ChimeAppInstanceTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ChimeAppInstanceTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ChimeAppInstanceTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ChimeAppInstanceTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewChimeAppInstanceTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

