// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package healthagentdomain

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HealthagentDomainTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (h *jsiiProxy_HealthagentDomainTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (h *jsiiProxy_HealthagentDomainTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_HealthagentDomainTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_HealthagentDomainTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HealthagentDomainTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_HealthagentDomainTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewHealthagentDomainTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

