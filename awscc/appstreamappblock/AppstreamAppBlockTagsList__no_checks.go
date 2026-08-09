// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appstreamappblock

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppstreamAppBlockTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppstreamAppBlockTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppstreamAppBlockTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppstreamAppBlockTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppstreamAppBlockTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppstreamAppBlockTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppstreamAppBlockTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppstreamAppBlockTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

