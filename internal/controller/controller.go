package controller

import (
	"errors"
	"log"
	"os"

	v1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	opv1 "carroll.codes/portfolio-operator/api/v1"
	"carroll.codes/portfolio-operator/internal/config"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
	port     = opv1.Portfolio{}
)

func init() {
	utilruntime.Must(opv1.AddToScheme(scheme))
	utilruntime.Must(v1.AddToScheme(scheme))
}

func Main() {
	log.Println("starting controller...")
	var (
		restConfig *rest.Config
		err        error
	)

	if _, err := os.Stat(config.Instance.KUBECONFIG); errors.Is(err, os.ErrNotExist) { // if kube config doesn't exist, try incluster config
		restConfig, err = rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
	} else {
		restConfig, err = clientcmd.BuildConfigFromFlags("", config.Instance.KUBECONFIG)
		if err != nil {
			panic(err.Error())
		}
	}

	ctrl.SetLogger(zap.New())

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		panic(err.Error())
	}
	log.Println("kubeconfig established")

	mgr, err := ctrl.NewManager(restConfig, ctrl.Options{
		Scheme: scheme,
		Metrics: metricsserver.Options{
			BindAddress: ":8443",
		},
	})
	if err != nil {
		log.Println("unable to start manager")
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	err = ctrl.NewControllerManagedBy(mgr).
		For(&v1.Ingress{}).
		Complete(&ingressReconciler{
			Client:     mgr.GetClient(),
			scheme:     mgr.GetScheme(),
			kubeClient: clientset,
		})
	if err != nil {
		setupLog.Error(err, "unable to create controller")
		os.Exit(1)
	}

	setupLog.Info("starting manager")

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "error running manager")
		os.Exit(1)
	}
}
