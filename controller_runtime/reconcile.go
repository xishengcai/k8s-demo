package controller_runtime

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)


type DemoReconcile struct {
	client.Client
}

func (r *DemoReconcile) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	panic("implement me")
}

// Reconcile 方法，内部实现具体的回调逻辑。
func (r *DemoReconcile)Reconclie(req reconcile.Request)(reconcile.Result, error){
	rs := appsv1.ReplicaSet{}
	if err := r.Get(context.TODO(),req.NamespacedName,rs);err != nil{
		return reconcile.Result{}, err
	}

	// TODO 逻辑代码
	return reconcile.Result{}, nil
}

func (r *DemoReconcile) InjectClient(c client.Client) error {
	r.Client = c
	return nil
}
