// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2listener


type Elasticloadbalancingv2ListenerMutualAuthentication struct {
	// Indicates whether trust store CA certificate names are advertised.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticloadbalancingv2_listener#advertise_trust_store_ca_names Elasticloadbalancingv2Listener#advertise_trust_store_ca_names}
	AdvertiseTrustStoreCaNames *string `field:"optional" json:"advertiseTrustStoreCaNames" yaml:"advertiseTrustStoreCaNames"`
	// Indicates whether expired client certificates are ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticloadbalancingv2_listener#ignore_client_certificate_expiry Elasticloadbalancingv2Listener#ignore_client_certificate_expiry}
	IgnoreClientCertificateExpiry interface{} `field:"optional" json:"ignoreClientCertificateExpiry" yaml:"ignoreClientCertificateExpiry"`
	// The client certificate handling method.
	//
	// Options are ``off``, ``passthrough`` or ``verify``. The default value on initial resource creation is ``off``. After mutual authentication is turned on, you must explicitly set the ``Mode`` to ``off`` to turn it off; removing the property from your template will not turn it off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticloadbalancingv2_listener#mode Elasticloadbalancingv2Listener#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The Amazon Resource Name (ARN) of the trust store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticloadbalancingv2_listener#trust_store_arn Elasticloadbalancingv2Listener#trust_store_arn}
	TrustStoreArn *string `field:"optional" json:"trustStoreArn" yaml:"trustStoreArn"`
}

