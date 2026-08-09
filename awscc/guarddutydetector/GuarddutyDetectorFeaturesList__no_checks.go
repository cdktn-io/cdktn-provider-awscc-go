// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package guarddutydetector

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GuarddutyDetectorFeaturesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GuarddutyDetectorFeaturesList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GuarddutyDetectorFeaturesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GuarddutyDetectorFeaturesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GuarddutyDetectorFeaturesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GuarddutyDetectorFeaturesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GuarddutyDetectorFeaturesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGuarddutyDetectorFeaturesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

