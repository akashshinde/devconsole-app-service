package watcher

import (
	"fmt"
	"github.com/redhat-developer/app-service/kubeclient"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type Watcher struct {
	Client *kubeclient.KubeClient
	ResultStream chan watch.Event
	Namespace string
}

func NewWatcher(namespace string) *Watcher {
	w := new(Watcher)
	w.Client = kubeclient.NewKubeClient()
	w.ResultStream = make(chan watch.Event)
	w.Namespace = namespace
	return w
}

func (w Watcher) StartWatcher() {
	dWatcher, _ := w.Client.CoreClient.AppsV1().Deployments(w.Namespace).Watch(v1.ListOptions{})
	podWatcher, _ := w.Client.CoreClient.CoreV1().Pods(w.Namespace).Watch(v1.ListOptions{})
	dcWatcher, _ := w.Client.OcClient.DeploymentConfigs(w.Namespace).Watch(v1.ListOptions{})
	go SendToChannel(dWatcher, w.ResultStream)
	go SendToChannel(podWatcher, w.ResultStream)
	go SendToChannel(dcWatcher, w.ResultStream)
}

func (w Watcher) ListenWatcher() {
	for {
		obj := <- w.ResultStream
		fmt.Printf("\nObj is %+v", obj)
	}
}

func SendToChannel(w watch.Interface, ch chan watch.Event)  {
	for {
		v := <- w.ResultChan()
		ch <- v
	}
}
