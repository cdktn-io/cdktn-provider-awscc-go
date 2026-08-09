// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package eksaccessentry

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EksAccessEntryAccessPoliciesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EksAccessEntryAccessPoliciesList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EksAccessEntryAccessPoliciesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EksAccessEntryAccessPoliciesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EksAccessEntryAccessPoliciesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksAccessEntryAccessPoliciesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EksAccessEntryAccessPoliciesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEksAccessEntryAccessPoliciesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

