package NicePod

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
	"sigs.k8s.io/scheduler-plugins/pkg/apis/config"
)

type NicePod struct {
	handle     framework.FrameworkHandle
	whatisnice string
}

const Name = "NicePod"

func (n *NicePod) Name() string {
	return Name
}

func New(obj runtime.Object, h framework.FrameworkHandle) (framework.Plugin, error) {
	args, ok := obj.(*config.NicePodArgs)
	if !ok {
		return nil, fmt.Errorf("want args to be of type NicePodArgs, got %T", obj)
	}

	return &NicePod{
		handle:     h,
		whatisnice: &args.WhatIsNice
	}, nil
}

// PreFilter checks if the scheduled Pod has the nicepod-label value 'nice'. Because we only want 'nice' pods.
func (np *NicePod) PreFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod) (*framework.PreFilterResult, *framework.Status) {
	// If PreFilter fails, return framework.UnschedulableAndUnresolvable to avoid any preemption attempts.
	if (pod.labels[nicepod] != np.whatisnice) {
		return nil, framework.NewStatus(framework.UnschedulableAndUnresolvable, err.Error())
	}
	return nil, framework.NewStatus(framework.Success, "")
}

// PreFilterExtensions returns a PreFilterExtensions interface if the plugin implements one.
func (np *NicePod) PreFilterExtensions() framework.PreFilterExtensions {
	return nil
}