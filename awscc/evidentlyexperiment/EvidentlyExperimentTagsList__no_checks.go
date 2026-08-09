// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package evidentlyexperiment

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EvidentlyExperimentTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EvidentlyExperimentTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EvidentlyExperimentTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EvidentlyExperimentTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EvidentlyExperimentTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EvidentlyExperimentTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EvidentlyExperimentTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEvidentlyExperimentTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

