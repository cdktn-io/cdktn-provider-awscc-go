// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ivschannel

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IvsChannelTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IvsChannelTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IvsChannelTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IvsChannelTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IvsChannelTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IvsChannelTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IvsChannelTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIvsChannelTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

