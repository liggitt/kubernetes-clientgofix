package testpackages

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

func TestHandleGet(t *testing.T) {
	var c dynamic.Interface
	c.Resource(schema.GroupVersionResource{}).Get("", metav1.GetOptions{})
}

func TestHandleList(t *testing.T) {
	var c dynamic.Interface
	listOpts := metav1.ListOptions{
		LabelSelector: action.GetListRestrictions().Labels.String(),
		FieldSelector: action.GetListRestrictions().Fields.String(),
	}
	c.Resource(schema.GroupVersionResource{}).List(listOpts)
}
