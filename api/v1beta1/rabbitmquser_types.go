/*


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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RabbitmqUserSpec defines the desired state of RabbitmqUser
type RabbitmqUserSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of RabbitmqUser. Edit RabbitmqUser_types.go to remove/update
	Name       string              `json:"name,omitempty"`
	ClusterRef RabbitmqClusterRef  `json:"clusterRef"`
	Settings   RabbitmqUserSetting `json:"settings"`
}

type RabbitmqUserSetting struct {
	Name string `json:"name"`
	/*Tags:             "",
	Password:         "",
	PasswordHash:     "",
	HashingAlgorithm: "",*/
}

// RabbitmqUserStatus defines the observed state of RabbitmqUser
type RabbitmqUserStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

// +kubebuilder:object:root=true

// RabbitmqUser is the Schema for the rabbitmqusers API
type RabbitmqUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RabbitmqUserSpec   `json:"spec,omitempty"`
	Status RabbitmqUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RabbitmqUserList contains a list of RabbitmqUser
type RabbitmqUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RabbitmqUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RabbitmqUser{}, &RabbitmqUserList{})
}
