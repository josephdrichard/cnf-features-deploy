// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/openshift/ptp-operator/pkg/apis/ptp/v1.NodePtpDevice":           schema_pkg_apis_ptp_v1_NodePtpDevice(ref),
		"github.com/openshift/ptp-operator/pkg/apis/ptp/v1.NodePtpDeviceSpec":       schema_pkg_apis_ptp_v1_NodePtpDeviceSpec(ref),
		"github.com/openshift/ptp-operator/pkg/apis/ptp/v1.NodePtpDeviceStatus":     schema_pkg_apis_ptp_v1_NodePtpDeviceStatus(ref),
		"github.com/openshift/ptp-operator/pkg/apis/ptp/v1.PtpConfig":               schema_pkg_apis_ptp_v1_PtpConfig(ref),
		"github.com/openshift/ptp-operator/pkg/apis/ptp/v1.PtpConfigSpec":           schema_pkg_apis_ptp_v1_PtpConfigSpec(ref),
		"github.com/openshift/ptp-operator/pkg/apis/ptp/v1.PtpConfigStatus":         schema_pkg_apis_ptp_v1_PtpConfigStatus(ref),
		"github.com/openshift/ptp-operator/pkg/apis/ptp/v1.PtpOperatorConfig":       schema_pkg_apis_ptp_v1_PtpOperatorConfig(ref),
		"github.com/openshift/ptp-operator/pkg/apis/ptp/v1.PtpOperatorConfigSpec":   schema_pkg_apis_ptp_v1_PtpOperatorConfigSpec(ref),
		"github.com/openshift/ptp-operator/pkg/apis/ptp/v1.PtpOperatorConfigStatus": schema_pkg_apis_ptp_v1_PtpOperatorConfigStatus(ref),
	}
}

func schema_pkg_apis_ptp_v1_NodePtpDevice(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NodePtpDevice is the Schema for the nodeptpdevices API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/ptp-operator/pkg/apis/ptp/v1.NodePtpDeviceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/ptp-operator/pkg/apis/ptp/v1.NodePtpDeviceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/ptp-operator/pkg/apis/ptp/v1.NodePtpDeviceSpec", "github.com/openshift/ptp-operator/pkg/apis/ptp/v1.NodePtpDeviceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_ptp_v1_NodePtpDeviceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NodePtpDeviceSpec defines the desired state of NodePtpDevice",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_ptp_v1_NodePtpDeviceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NodePtpDeviceStatus defines the observed state of NodePtpDevice",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"devices": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/ptp-operator/pkg/apis/ptp/v1.PtpDevice"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/ptp-operator/pkg/apis/ptp/v1.PtpDevice"},
	}
}

func schema_pkg_apis_ptp_v1_PtpConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PtpConfig is the Schema for the ptpconfigs API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/ptp-operator/pkg/apis/ptp/v1.PtpConfigSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/ptp-operator/pkg/apis/ptp/v1.PtpConfigStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/ptp-operator/pkg/apis/ptp/v1.PtpConfigSpec", "github.com/openshift/ptp-operator/pkg/apis/ptp/v1.PtpConfigStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_ptp_v1_PtpConfigSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PtpConfigSpec defines the desired state of PtpConfig",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"profile": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/ptp-operator/pkg/apis/ptp/v1.PtpProfile"),
									},
								},
							},
						},
					},
					"recommend": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/ptp-operator/pkg/apis/ptp/v1.PtpRecommend"),
									},
								},
							},
						},
					},
				},
				Required: []string{"profile", "recommend"},
			},
		},
		Dependencies: []string{
			"github.com/openshift/ptp-operator/pkg/apis/ptp/v1.PtpProfile", "github.com/openshift/ptp-operator/pkg/apis/ptp/v1.PtpRecommend"},
	}
}

func schema_pkg_apis_ptp_v1_PtpConfigStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PtpConfigStatus defines the observed state of PtpConfig",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"matchList": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/ptp-operator/pkg/apis/ptp/v1.NodeMatchList"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/ptp-operator/pkg/apis/ptp/v1.NodeMatchList"},
	}
}

func schema_pkg_apis_ptp_v1_PtpOperatorConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PtpOperatorConfig is the Schema for the ptpoperatorconfigs API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/ptp-operator/pkg/apis/ptp/v1.PtpOperatorConfigSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/ptp-operator/pkg/apis/ptp/v1.PtpOperatorConfigStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/ptp-operator/pkg/apis/ptp/v1.PtpOperatorConfigSpec", "github.com/openshift/ptp-operator/pkg/apis/ptp/v1.PtpOperatorConfigStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_ptp_v1_PtpOperatorConfigSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PtpOperatorConfigSpec defines the desired state of PtpOperatorConfig",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"daemonNodeSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"daemonNodeSelector"},
			},
		},
	}
}

func schema_pkg_apis_ptp_v1_PtpOperatorConfigStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PtpOperatorConfigStatus defines the observed state of PtpOperatorConfig",
				Type:        []string{"object"},
			},
		},
	}
}