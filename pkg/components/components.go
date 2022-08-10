package components

import (
	"path/filepath"

	"github.com/openshift/microshift/pkg/config"
	"k8s.io/klog/v2"
)

func StartComponents(cfg *config.MicroshiftConfig) error {
	if err := startServiceCAController(cfg, filepath.Join(cfg.DataDir, "/resources/kubeadmin/kubeconfig")); err != nil {
		klog.Warningf("Failed to start service-ca controller: %v", err)
		return err
	}

	if err := startCSIPLugin(cfg.DataDir + "/resources/kubeadmin/kubeconfig"); err != nil {
		klog.Warningf("Failed to start csi plugin: %v", err)
		return err
	}

	if err := startIngressController(cfg, filepath.Join(cfg.DataDir, "/resources/kubeadmin/kubeconfig")); err != nil {
		klog.Warningf("Failed to start ingress router controller: %v", err)
		return err
	}
	if err := startDNSController(cfg, filepath.Join(cfg.DataDir, "/resources/kubeadmin/kubeconfig")); err != nil {
		klog.Warningf("Failed to start DNS controller: %v", err)
		return err
	}

	if err := startPatu(cfg, cfg.DataDir+"/resources/kubeadmin/kubeconfig"); err != nil {
		klog.Warningf("Failed to start PATU: %v", err)
		return err
	}
	return nil
}
