package cmd

import (
	"github.com/cybozu/neco-containers/admission/hooks"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)

	// +kubebuilder:scaffold:scheme
}

func run(addr string, port int) error {
	ctrl.SetLogger(zap.Logger(config.development))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: config.metricsAddr,
		LeaderElection:     false,
		Host:               addr,
		Port:               port,
		CertDir:            config.certDir,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		return err
	}

	// register webhook handlers
	// admission.NewDecoder never returns non-nil error
	dec, _ := admission.NewDecoder(scheme)
	wh := mgr.GetWebhookServer()
	wh.Register("/validate-projectcalico-org-networkpolicy", hooks.NewCalicoNetworkPolicyValidator(mgr.GetClient(), dec, 1000))
	wh.Register("/mutate-projectcontour-io-httpproxy", hooks.NewContourHTTPProxyMutator(mgr.GetClient(), dec))

	// +kubebuilder:scaffold:builder

	// pre-cache objects
	if _, err := mgr.GetCache().GetInformer(&corev1.Namespace{}); err != nil {
		setupLog.Error(err, "unable to setup informer for namespaces")
		return err
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		return err
	}
	return nil
}
