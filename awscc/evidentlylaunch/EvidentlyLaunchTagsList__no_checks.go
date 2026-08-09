// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package evidentlylaunch

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EvidentlyLaunchTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EvidentlyLaunchTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EvidentlyLaunchTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EvidentlyLaunchTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EvidentlyLaunchTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EvidentlyLaunchTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EvidentlyLaunchTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEvidentlyLaunchTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

