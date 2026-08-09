// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package evidentlyfeature

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EvidentlyFeatureTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EvidentlyFeatureTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EvidentlyFeatureTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EvidentlyFeatureTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EvidentlyFeatureTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EvidentlyFeatureTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EvidentlyFeatureTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEvidentlyFeatureTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

