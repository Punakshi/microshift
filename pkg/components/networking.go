package components

import (
	"path/filepath"

	"github.com/openshift/microshift/pkg/assets"
	"github.com/openshift/microshift/pkg/config"
	"k8s.io/klog/v2"
)

func startPatu(cfg *config.MicroshiftConfig, kubeconfigPath string) error {
	var (
		sa = []string{
			"assets/components/patu/serviceaccount.yaml",
		}
		cr = []string{
			"assets/components/patu/clusterrole.yaml",
		}
		crb = []string{
			"assets/components/patu/clusterrolebinding.yaml",
		}
		cm = []string{
			"assets/components/patu/configmap.yaml",
		}
		apps = []string{
			"assets/components/patu/daemonset.yaml",
		}
	)

	if err := assets.ApplyServiceAccounts(sa, kubeconfigPath); err != nil {
		klog.Warningf("Failed to apply serviceAccount %v %v", sa, err)
		return err
	}
	if err := assets.ApplyClusterRoles(cr, kubeconfigPath); err != nil {
		klog.Warningf("Failed to apply clusterRole %v %v", cr, err)
		return err
	}
	if err := assets.ApplyClusterRoleBindings(crb, kubeconfigPath); err != nil {
		klog.Warningf("Failed to apply clusterRoleBinding %v %v", crb, err)
		return err
	}
	params := assets.RenderParams{
		"ClusterCIDR":    cfg.Cluster.ClusterCIDR,
		"ServiceCIDR":    cfg.Cluster.ServiceCIDR,
		"KubeconfigPath": kubeconfigPath,
		"KubeconfigDir":  filepath.Join(cfg.DataDir, "/resources/kubeadmin"),
	}
	if err := assets.ApplyConfigMaps(cm, kubeconfigPath); err != nil {
		klog.Warningf("Failed to apply configMap %v %v", cm, err)
		return err
	}
	if err := assets.ApplyDaemonSets(apps, renderOVNKManifests, params, kubeconfigPath); err != nil {
		klog.Warningf("Failed to apply apps %v %v", apps, err)
		return err
	}
	return nil
}
