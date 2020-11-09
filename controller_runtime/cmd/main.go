package main

import (
	_ "context"
	"k8s-demo/controller_runtime"
	"os"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	_ "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	_ "sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func main() {
	// ①
	mgr, err := manager.New(config.GetConfigOrDie(), manager.Options{})
	if err != nil {
		os.Exit(1)
	}

	// ②
	err = builder.
		ControllerManagedBy(mgr).
		For(&appsv1.ReplicaSet{}).
		Owns(&corev1.Pod{}).
		Complete(&controller_runtime.DemoReconcile{})
	if err != nil {
		os.Exit(1)
	}

	// ③
	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		os.Exit(1)
	}
}
