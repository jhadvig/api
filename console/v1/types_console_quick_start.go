package v1

import (
	authorizationv1 "k8s.io/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConsoleQuickStart is an extension for guiding user thought various
// workflows in the OpenShift web console.
type ConsoleQuickStart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ConsoleQuickStartSpec `json:"spec"`
}

// ConsoleQuickStartSpec is the desired quick start configuration.
type ConsoleQuickStartSpec struct {
	// displayName is the display name of the Quick Start.
	DisplayName string `json:"displayName"`
	// icon is a base64 encoded image that are used as Quick Start's
	// +optional
	Icon string `json:"icon,omitempty"`
	// tags is an array of strings that describe the Quick Start.
	// +optional
	Tags []string `json:"tags,omitempty"`
	// durationMinutes describes approximately how many minutes it will take to complete the Quick Start.
	// +optional
	DurationMinutes int `json:"durationMinutes,omitempty"`
	// description is the description of the Quick Start. (includes markdown)
	// +optional
	Description string `json:"description,omitempty"`
	// prerequisites contains a list of Quick Start names that must be completed before this Quick Start.
	// +optional
	Prerequisites []string `json:"prerequisites,omitempty"`
	// introduction describes the purpose of the Quick Start. (includes markdown)
	Introduction string `json:"introduction"`
	// tasks contains list of tasks that user has to take in order to pass the Quick Start.
	Tasks []ConsoleQuickStartTask `json:"tasks"`
	// conclusion sums up the Quick Start and suggests the possible
	// next Quick Starts for user to take. (includes markdown)
	// +optional
	Conclusion string `json:"conclusion,omitempty"`
	// nextQuickStart is the name of the following Quick Start.
	// +optional
	NextQuickStart string `json:"nextQuickStart,omitempty"`
	// accessReviewResources containes arrays of resources that user's access
	// will be reviewed against in order for the user to complete the Quick Start.
	// The Quick Start will be hidden if any of the access reviews fail.
	// +optional
	AccessReviewResources []authorizationv1.ResourceAttributes `json:"accessReviewResources,omitempty"`
}

// ConsoleQuickStartTask is a single step in a Quick Start.
type ConsoleQuickStartTask struct {
	// title describes the task and will be rendered as a step header.
	Title string `json:"title"`
	// description describes the steps needed to complete the task. (includes markdown)
	Description string `json:"description"`
	// review contains instructions to validate the task is complete. The user will select 'Yes' or 'No'
	// using a radio button, which indicates whether the step was completed successfully.
	// +optional
	Review *ConsoleQuickStartTaskReview `json:"review,omitempty"`
	// recapitulation contains information about passed step.
	// +optional
	Recapitulation *ConsoleQuickStartTaskRecapitulation `json:"recapitulation,omitempty"`
}

// ConsoleQuickStartTaskReview contains instructions that validate a task was completed successfully.
type ConsoleQuickStartTaskReview struct {
	// instructions contains steps that user needs to take in order
	// to validate his work after going through a task. (includes markdown)
	Instructions string `json:"instructions"`
	// failedTaskHelp contains suggestions for a failed task review and is shown at the end of task.
	FailedTaskHelp string `json:"failedTaskHelp"`
}

// ConsoleQuickStartTaskRecapitulation contains information about a passed step.
type ConsoleQuickStartTaskRecapitulation struct {
	// success describes the succesfully passed task.
	Success string `json:"success"`
	// failed describes the unsuccessfully passed task.
	Failed string `json:"failed"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ConsoleQuickStartList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ConsoleQuickStart `json:"items"`
}
