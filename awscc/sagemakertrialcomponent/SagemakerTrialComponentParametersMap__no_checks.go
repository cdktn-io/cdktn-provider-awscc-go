// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sagemakertrialcomponent

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SagemakerTrialComponentParametersMap) validateGetParameters(key *string) error {
	return nil
}

func (s *jsiiProxy_SagemakerTrialComponentParametersMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (s *jsiiProxy_SagemakerTrialComponentParametersMap) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SagemakerTrialComponentParametersMap) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SagemakerTrialComponentParametersMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SagemakerTrialComponentParametersMap) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func validateNewSagemakerTrialComponentParametersMapParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

