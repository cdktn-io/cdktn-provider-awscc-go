// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package directconnectlag

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DirectconnectLagTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DirectconnectLagTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DirectconnectLagTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DirectconnectLagTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DirectconnectLagTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DirectconnectLagTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DirectconnectLagTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDirectconnectLagTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

