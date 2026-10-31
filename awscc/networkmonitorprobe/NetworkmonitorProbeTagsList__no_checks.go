// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package networkmonitorprobe

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkmonitorProbeTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetworkmonitorProbeTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkmonitorProbeTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkmonitorProbeTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkmonitorProbeTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkmonitorProbeTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkmonitorProbeTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkmonitorProbeTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

