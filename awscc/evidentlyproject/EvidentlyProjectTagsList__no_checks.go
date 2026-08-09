// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package evidentlyproject

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EvidentlyProjectTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EvidentlyProjectTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EvidentlyProjectTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EvidentlyProjectTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EvidentlyProjectTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EvidentlyProjectTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EvidentlyProjectTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEvidentlyProjectTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

