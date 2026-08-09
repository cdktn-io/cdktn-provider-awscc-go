// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package guarddutydetector

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GuarddutyDetectorTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GuarddutyDetectorTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GuarddutyDetectorTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GuarddutyDetectorTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GuarddutyDetectorTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GuarddutyDetectorTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GuarddutyDetectorTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGuarddutyDetectorTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

