/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	backblaze "github.com/ihyoudou/go-backblaze"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type BucketLifecycle struct {
	DaysFromUploadingToHiding int    `json:"daysfromuploadingtohiding,omitempty"`
	DaysFromHidingToDeleting  int    `json:"daysfromhidingtodelete,omitempty"`
	FileNamePrefix            string `json:"filenameprefix,omitempty"`
}

type BucketWriteConnectionSecretToRef struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Acl defines if content of this bucket should be public or private
	Acl                              string                           `json:"acl,omitempty"`
	BucketLifecycle                  []backblaze.LifecycleRule        `json:"bucketlifecycle,omitempty"`
	BucketWriteConnectionSecretToRef BucketWriteConnectionSecretToRef `json:"writeconnectionsecrettoref,omitempty"`
}

// BucketStatus defines the observed state of Bucket
type BucketStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	AtProvider BucketSpec `json:"atprovider,omitempty"`
	Reconciled bool       `json:"reconciled,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Bucket is the Schema for the buckets API
// +kubebuilder:subresource:status
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BucketSpec   `json:"spec,omitempty"`
	Status BucketStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BucketList contains a list of Bucket
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}
