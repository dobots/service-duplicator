package service_duplicator

import (
	"context"
//        "fmt"
        "log"
//        "strings"
	"net/http"
//	"net/http/httputil"
//	"net/url"
//	"regexp"
//	"time"

//	appsv1 "k8s.io/api/apps/v1"
//	core_v1 "k8s.io/api/core/v1"

//	autoscalingv1 "k8s.io/api/autoscaling/v1"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
//	"k8s.io/client-go/tools/cache"
)

// Config the plugin configuration.
type Config struct {
        Header             string      `json:"header,omitempty" yaml:"Header" mapstructure:"Header" default:"X-Forwarded-User"`
        Namespace             string      `json:"namespace,omitempty" yaml:"Namespace" mapstructure:"Namespace" default:"default"`
}

type KubernetesProvider struct {
	Client kubernetes.Interface
}

func newKubernetesProvider() (*KubernetesProvider, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &KubernetesProvider{
		Client: client,
	}, nil
}


// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
                Header:  "X-Forwarded-User",
                Namespace: "Asimovo-userspace",
	}
}

type ServiceDup struct {
	config *Config
        provider *KubernetesProvider
	next   http.Handler
	name   string
}

// New created a new ServiceDup plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
        provider, err := newKubernetesProvider()
	if err != nil {
		return nil, err
	}
	return &ServiceDup{
		config: config,
                provider: provider,
		next:   next,
		name:   name,
	}, nil
}

func (a *ServiceDup) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
     //setup login to Kubernetes API
     //Determine current hostname
     //Get full config of my associated Service (should be the generic service)
     //Get all yamls that are associated with my service
     //Search and replace all occurances of the hostname -> servicehost to my current hostname
	//Modify the Middleware list
	//(optionally)Modify the amount of instances?
     //Create all new elements
     services := a.provider.Client.CoreV1().Services(a.config.Namespace)
     log.Printf("services: %s", services)
     a.next.ServeHTTP(rw, req)
}
