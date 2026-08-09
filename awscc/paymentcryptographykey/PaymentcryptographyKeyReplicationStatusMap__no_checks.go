// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package paymentcryptographykey

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PaymentcryptographyKeyReplicationStatusMap) validateGetParameters(key *string) error {
	return nil
}

func (p *jsiiProxy_PaymentcryptographyKeyReplicationStatusMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (p *jsiiProxy_PaymentcryptographyKeyReplicationStatusMap) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PaymentcryptographyKeyReplicationStatusMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PaymentcryptographyKeyReplicationStatusMap) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func validateNewPaymentcryptographyKeyReplicationStatusMapParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

