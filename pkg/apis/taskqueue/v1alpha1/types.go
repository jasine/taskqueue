package v1alpha1

import (
	"k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TaskQueue is a specification for a TaskQueue resource
type TaskQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TaskQueueSpec   `json:"spec"`
	Status TaskQueueStatus `json:"status"`
}

// TaskQueueSpec is the spec for a TaskQueue resource
type TaskQueueSpec struct {
	DeploymentName string   `json:"deploymentName"`
	Replicas       *int32   `json:"replicas"`
	Jobs           []v1.Job `json:"jobs"`
}

// TaskQueueStatus is the status for a TaskQueue resource
type TaskQueueStatus struct {
	AvailableReplicas int32        `json:"availableReplicas"`
	//runnig job name in this queue
	CurrentJob        string       `json:"currentJob" protobuf:"bytes,6,opt,name=currentJob"`
	Message           string       `json:"message,omitempty" protobuf:"bytes,6,opt,name=message"`
	StartTime         *metav1.Time `json:"startTime,omitempty" protobuf:"bytes,2,opt,name=startTime"`
	CompletionTime    *metav1.Time `json:"completionTime,omitempty" protobuf:"bytes,3,opt,name=completionTime"`
	Active            int32        `json:"active,omitempty" protobuf:"varint,4,opt,name=active"`
	Succeeded         int32        `json:"succeeded,omitempty" protobuf:"varint,5,opt,name=succeeded"`
	Failed            int32        `json:"failed,omitempty" protobuf:"varint,6,opt,name=failed"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TaskQueueList is a list of TaskQueue resources
type TaskQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []TaskQueue `json:"items"`
}
