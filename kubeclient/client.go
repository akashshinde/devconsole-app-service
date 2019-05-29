package kubeclient

import (
	"flag"
	ocappsclient "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

type KubeClient struct {
   CoreClient	kubernetes.Interface
   OcClient 	ocappsclient.AppsV1Interface
}

func NewKubeClient() *KubeClient  {
	var err error
	kc := new(KubeClient)
	config := getKubeConfig()
	kc.CoreClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	kc.OcClient, err = ocappsclient.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return kc;
}

func getKubeConfig() *rest.Config {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	return config
}