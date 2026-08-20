// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package inspectorv2connector

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_Inspectorv2ConnectorTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_Inspectorv2ConnectorTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_Inspectorv2ConnectorTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Inspectorv2ConnectorTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Inspectorv2ConnectorTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Inspectorv2ConnectorTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Inspectorv2ConnectorTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewInspectorv2ConnectorTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

