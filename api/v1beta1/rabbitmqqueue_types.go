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

type RabbitmqClusterRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type RabbitmqQueueSetting struct {
	Type string `json:"type"`
	// kubebuilder:default=false
	Durable    bool              `json:"durable"`
	AutoDelete bool              `json:"auto_delete,omitempty"`
	Arguments  map[string]string `json:"arguments,omitempty"`
}

// RabbitmqQueueSpec defines the desired state of RabbitmqQueue
type RabbitmqQueueSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of RabbitmqQueue. Edit RabbitmqQueue_types.go to remove/update
	// kubebuilder:default=/
	Vhost      string               `json:"vhost,omitempty"`
	Name       string               `json:"name,omitempty"`
	ClusterRef RabbitmqClusterRef   `json:"clusterRef"`
	Settings   RabbitmqQueueSetting `json:"settings"`
}

// RabbitmqQueueStatus defines the observed state of RabbitmqQueue
type RabbitmqQueueStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.status`
// +kubebuilder:printcolumn:name="Error",type=string,JSONPath=`.status.error`
// RabbitmqQueue is the Schema for the rabbitmqqueues API
type RabbitmqQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RabbitmqQueueSpec   `json:"spec,omitempty"`
	Status RabbitmqQueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RabbitmqQueueList contains a list of RabbitmqQueue
type RabbitmqQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RabbitmqQueue `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RabbitmqQueue{}, &RabbitmqQueueList{})
}
