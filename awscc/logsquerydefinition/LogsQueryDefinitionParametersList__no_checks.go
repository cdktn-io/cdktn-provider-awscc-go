// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package logsquerydefinition

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogsQueryDefinitionParametersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LogsQueryDefinitionParametersList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LogsQueryDefinitionParametersList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LogsQueryDefinitionParametersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LogsQueryDefinitionParametersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LogsQueryDefinitionParametersList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LogsQueryDefinitionParametersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLogsQueryDefinitionParametersListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

