/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type MemberObservation struct {

	// The ID of the GuardDuty member
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The status of the relationship between the member account and its primary account. More information can be found in Amazon GuardDuty API Reference.
	RelationshipStatus *string `json:"relationshipStatus,omitempty" tf:"relationship_status,omitempty"`
}

type MemberParameters struct {

	// AWS account ID for member account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/guardduty/v1beta1.Detector
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("account_id",true)
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Reference to a Detector in guardduty to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDRef *v1.Reference `json:"accountIdRef,omitempty" tf:"-"`

	// Selector for a Detector in guardduty to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDSelector *v1.Selector `json:"accountIdSelector,omitempty" tf:"-"`

	// The detector ID of the GuardDuty account where you want to create member accounts.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/guardduty/v1beta1.Detector
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DetectorID *string `json:"detectorId,omitempty" tf:"detector_id,omitempty"`

	// Reference to a Detector in guardduty to populate detectorId.
	// +kubebuilder:validation:Optional
	DetectorIDRef *v1.Reference `json:"detectorIdRef,omitempty" tf:"-"`

	// Selector for a Detector in guardduty to populate detectorId.
	// +kubebuilder:validation:Optional
	DetectorIDSelector *v1.Selector `json:"detectorIdSelector,omitempty" tf:"-"`

	// Boolean whether an email notification is sent to the accounts. Defaults to false.
	// +kubebuilder:validation:Optional
	DisableEmailNotification *bool `json:"disableEmailNotification,omitempty" tf:"disable_email_notification,omitempty"`

	// Email address for member account.
	// +kubebuilder:validation:Required
	Email *string `json:"email" tf:"email,omitempty"`

	// Message for invitation.
	// +kubebuilder:validation:Optional
	InvitationMessage *string `json:"invitationMessage,omitempty" tf:"invitation_message,omitempty"`

	// Boolean whether to invite the account to GuardDuty as a member. Defaults to false.
	// +kubebuilder:validation:Optional
	Invite *bool `json:"invite,omitempty" tf:"invite,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// MemberSpec defines the desired state of Member
type MemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MemberParameters `json:"forProvider"`
}

// MemberStatus defines the observed state of Member.
type MemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Member is the Schema for the Members API. Provides a resource to manage a GuardDuty member
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Member struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MemberSpec   `json:"spec"`
	Status            MemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MemberList contains a list of Members
type MemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Member `json:"items"`
}

// Repository type metadata.
var (
	Member_Kind             = "Member"
	Member_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Member_Kind}.String()
	Member_KindAPIVersion   = Member_Kind + "." + CRDGroupVersion.String()
	Member_GroupVersionKind = CRDGroupVersion.WithKind(Member_Kind)
)

func init() {
	SchemeBuilder.Register(&Member{}, &MemberList{})
}
