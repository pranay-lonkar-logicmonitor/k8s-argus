package permission

import (
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var (
	enable  = true
	disable = false

	client *kubernetes.Clientset

	deploymentPermissionFlag *bool
)

// Init is a function than init the permission context
func Init(k8sClient *kubernetes.Clientset) {
	client = k8sClient
}

// HasDeploymentPermissions is a function that check if the deployment resource has permissions
func HasDeploymentPermissions() bool {
	if deploymentPermissionFlag != nil {
		return *deploymentPermissionFlag
	}
	_, err := client.AppsV1beta2().Deployments(v1.NamespaceAll).List(metav1.ListOptions{})
	if err != nil {
		deploymentPermissionFlag = &disable
		log.Errorf("Failed to list deployments: %+v", err)
	} else {
		deploymentPermissionFlag = &enable
	}
	return *deploymentPermissionFlag
}
