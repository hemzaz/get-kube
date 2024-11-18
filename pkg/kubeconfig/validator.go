package kubeconfig

import (
	"errors"

	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// ValidateConfig checks the overall validity of a kubeconfig.
func ValidateConfig(config *clientcmdapi.Config) error {
	if config == nil {
		return errors.New("kubeconfig is nil")
	}

	if len(config.Clusters) == 0 {
		return errors.New("no clusters defined in kubeconfig")
	}

	if len(config.AuthInfos) == 0 {
		return errors.New("no users defined in kubeconfig")
	}

	if len(config.Contexts) == 0 {
		return errors.New("no contexts defined in kubeconfig")
	}

	if config.CurrentContext == "" {
		return errors.New("current context is not set")
	}

	return nil
}

// ValidateCluster ensures a cluster entry is valid.
func ValidateCluster(cluster *clientcmdapi.Cluster) error {
	if cluster == nil {
		return errors.New("cluster is nil")
	}

	if cluster.Server == "" {
		return errors.New("cluster server URL is not set")
	}

	if len(cluster.CertificateAuthorityData) == 0 && cluster.CertificateAuthority == "" {
		return errors.New("cluster has no certificate authority data or reference")
	}

	return nil
}

// ValidateAuthInfo ensures a user entry is valid.
func ValidateAuthInfo(authInfo *clientcmdapi.AuthInfo) error {
	if authInfo == nil {
		return errors.New("auth info is nil")
	}

	if len(authInfo.ClientCertificateData) == 0 && authInfo.ClientCertificate == "" {
		return errors.New("user has no client certificate data or reference")
	}

	if len(authInfo.ClientKeyData) == 0 && authInfo.ClientKey == "" {
		return errors.New("user has no client key data or reference")
	}

	return nil
}

// ValidateContext ensures a context entry is valid.
func ValidateContext(context *clientcmdapi.Context) error {
	if context == nil {
		return errors.New("context is nil")
	}

	if context.Cluster == "" {
		return errors.New("context does not specify a cluster")
	}

	if context.AuthInfo == "" {
		return errors.New("context does not specify an auth info")
	}

	return nil
}
