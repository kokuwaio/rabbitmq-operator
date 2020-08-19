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

// RabbitmqShovelSpec defines the desired state of RabbitmqShovel
type RabbitmqShovelSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of RabbitmqShovel. Edit RabbitmqShovel_types.go to remove/update
	ClusterRef                    RabbitmqClusterRef `json:"clusterRef"`
	Vhost                         string             `json:"vhost"`
	Name                          string             `json:"name"`
	SourceURI                     string             `json:"sourceURI,omitempty"`
	SourceProtocol                string             `json:"sourceProtocol,omitempty"`
	SourceQueue                   string             `json:"sourceQueue,omitempty"`
	DestinationURI                string             `json:"destinationURI,omitempty"`
	DestinationProtocol           string             `json:"destinationProtocol,omitempty"`
	DestinationAddress            string             `json:"destinationAddress,omitempty"`
	DestinationAddForwardHeaders  bool               `json:"destinationAddForwardHeaders,omitempty"`
	AckMode                       string             `json:"ackMode,omitempty"`
	SrcDeleteAfter                string             `json:"srcDeleteAfter,omitempty"`
	AddForwardHeaders             bool               `json:"addForwardHeaders,omitempty"`
	DeleteAfter                   string             `json:"deleteAfter,omitempty"`
	DestinationAddTimestampHeader bool               `json:"destinationAddTimestampHeader,omitempty"`
	DestinationExchange           string             `json:"destinationExchange,omitempty"`
	DestinationExchangeKey        string             `json:"destinationExchangeKey,omitempty"`
	DestinationQueue              string             `json:"destinationQueue,omitempty"`
	SourceAddress                 string             `json:"sourceAddress,omitempty"`
	SourceExchange                string             `json:"sourceExchange,omitempty"`
	SourceExchangeKey             string             `json:"sourceExchangeKey,omitempty"`
	SourcePasswordSecretRef       PasswordSecretRef  `json:"sourceSecretRef,omitempty"`
	SourceUser                    string             `json:"sourceUser,omitempty"`
	DestinationPasswordSecretRef  PasswordSecretRef  `json:"destinationSecretRef,omitempty"`
	DestinationUser               string             `json:"destinationUser,omitempty"`
}

// RabbitmqShovelStatus defines the observed state of RabbitmqShovel
type RabbitmqShovelStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.status`
// +kubebuilder:printcolumn:name="Error",type=string,JSONPath=`.status.error`
// RabbitmqShovel is the Schema for the rabbitmqshovels API
type RabbitmqShovel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RabbitmqShovelSpec   `json:"spec,omitempty"`
	Status RabbitmqShovelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RabbitmqShovelList contains a list of RabbitmqShovel
type RabbitmqShovelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RabbitmqShovel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RabbitmqShovel{}, &RabbitmqShovelList{})
}
