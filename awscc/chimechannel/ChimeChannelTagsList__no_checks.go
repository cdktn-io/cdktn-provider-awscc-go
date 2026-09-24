// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package chimechannel

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ChimeChannelTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ChimeChannelTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ChimeChannelTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ChimeChannelTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ChimeChannelTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ChimeChannelTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ChimeChannelTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewChimeChannelTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

