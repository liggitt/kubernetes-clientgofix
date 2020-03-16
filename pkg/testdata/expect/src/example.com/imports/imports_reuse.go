package misc

import (
	mycontext "context"

	corev1 "k8s.io/api/core/v1"
	mymetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

func Imports_Reuse_Existing_Aliases(c kubernetes.Interface, ctx mycontext.Context) {
	pod := &corev1.Pod{}
	c.CoreV1().Pods("").Create(mycontext.TODO(), pod, mymetav1.CreateOptions{})
	c.CoreV1().Pods("").Delete(mycontext.TODO(), "", mymetav1.DeleteOptions{})
}
