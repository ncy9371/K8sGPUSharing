/*
Copyright 2015 The Kubernetes Authors.

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

package v1

import (
	"bytes"
    "math/rand"
    "time"
    "unsafe"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
)

const (
    letterIdxBits = 5                    // 6 bits to represent a letter index
    letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
    letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
    letterBytes = "abcdefghijklmnopqrstuvwxyz"
    // letterIdxBits = 6                    // 6 bits to represent a letter index
    // letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
    // letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
    // letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MtgpuPod struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Status MtgpuPodStatus `json:"status,omitempty"`
	// +optional
	Spec corev1.PodSpec `json:"spec,omitempty"`
}

type MtgpuPodStatus struct {
	/*PodPhase          corev1.PodPhase
	ConfigFilePhase   ConfigFilePhase
	BoundDeviceID     string
	StartTime         *metav1.Time
	ContainerStatuses []corev1.ContainerStatus*/
	PodStatus     *corev1.PodStatus
	PodObjectMeta *metav1.ObjectMeta
	BoundDeviceID string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TestTypeList is a top-level list type. The client methods for lists are automatically created.
// You are not supposed to create a separated client for this one.
type MtgpuPodList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MtgpuPod `json:"items"`
}

func (this MtgpuPod) Print() {
	var buf bytes.Buffer
	buf.WriteString("\n================= MtgpuPod ==================")
	buf.WriteString("\nname: ")
	buf.WriteString(this.ObjectMeta.Namespace)
	buf.WriteString("/")
	buf.WriteString(this.ObjectMeta.Name)
	buf.WriteString("\nannotation:\n\tlsalab.nthu/gpu_request: ")
	buf.WriteString(this.ObjectMeta.Annotations["lsalab.nthu/gpu_request"])
	if this.Status.PodStatus != nil {
		buf.WriteString("\nstatus:\n\tPodStatus: ")
		buf.WriteString(string(this.Status.PodStatus.Phase))
	}
	buf.WriteString("\n\tGPUID: ")
	buf.WriteString(this.ObjectMeta.Annotations["lsalab.nthu/GPUID"])
	buf.WriteString("\n\tBoundDeviceID: ")
	buf.WriteString(this.Status.BoundDeviceID)
	buf.WriteString("\n=============================================")
	klog.Info(buf.String())
}

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go/31832326#31832326
var src = rand.NewSource(time.Now().UnixNano())
func NewGPUID(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
