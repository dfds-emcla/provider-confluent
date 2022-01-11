package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// ACLBlock
// type ACLBlock struct {
// 	Action         string   `json:"action"`
// 	ClusterScope   string   `json:"clusterScope"`
// 	ConsumerGroup  string   `json:"consumerGroup"`
// 	Operations     []string `json:"operations"`
// 	Prefix         string   `json:"prefix"`
// 	ServiceAccount string   `json:"serviceAccount"`
// 	Topic          string   `json:"topic"`
// 	Environment    string   `json:"environment"`
// 	Cluster        string   `json:"cluster"`
// }

// ACLBlock
type ACLBlock struct {
	Operation    string `json:"operation"`
	PatternType  string `json:"pattern_type"`
	Permission   string `json:"permission"`
	Principal    string `json:"principal"`
	ResourceName string `json:"resource_name"`
	ResourceType string `json:"resource_type"`
	Environment  string `json:"environment"`
	Cluster      string `json:"cluster"`
}

// ACLParameters are the configurable fields of a ACL.
type ACLParameters struct {
	ACLBlockList []ACLBlock `json:"aclBlockList"`
}

// ACLObservation are the observable fields of a ACL.
type ACLObservation struct {
	ACLBlockObservationList []ACLBlock `json:"aclBlockObservationList"`
}

// ACL Spec defines the desired state of a ACL.
type ACLSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ACLParameters `json:"forProvider"`
}

// ACL Status represents the observed state of a ACL.
type ACLStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ACLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Schema is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,confluent}
type ACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ACLSpec   `json:"spec"`
	Status            ACLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ACL List contains a list of ACL
type ACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ACL `json:"items"`
}

// ACL type metadata.
var (
	ACLKind             = reflect.TypeOf(ACL{}).Name()
	ACLGroupKind        = schema.GroupKind{Group: Group, Kind: ACLKind}.String()
	ACLKindAPIVersion   = ACLKind + "." + SchemeGroupVersion.String()
	ACLGroupVersionKind = SchemeGroupVersion.WithKind(ACLKind)
)

func init() {
	SchemeBuilder.Register(&ACL{}, &ACLList{})
}
