module github.com/openshift-kni/cnf-features-deploy

go 1.13

require (
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/ignition v0.35.0
	github.com/emicklei/go-restful v2.12.0+incompatible // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/ishidawataru/sctp v0.0.0-20191218070446-00ab2ac2db07
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.1
	github.com/openshift-kni/performance-addon-operators v0.0.0-20201021185419-3a3fec891a60
	github.com/openshift/api v3.9.1-0.20191213091414-3fbf6bcf78e8+incompatible
	github.com/openshift/client-go v0.0.0-20200326155132-2a6cd50aedd0
	github.com/openshift/cluster-node-tuning-operator v0.0.0-20200914165052-a39511828cf0
	github.com/openshift/machine-config-operator v4.2.0-alpha.0.0.20190917115525-033375cbe820+incompatible
	github.com/openshift/ptp-operator v0.0.0-20201020160817-8db64deb14c7
	github.com/openshift/sriov-network-operator v0.0.0-20201021182021-9de9f4ce8413
	github.com/spf13/cobra v1.0.0
	go4.org v0.0.0-20200411211856-f5505b9728dd // indirect
	k8s.io/api v0.18.9
	k8s.io/apiextensions-apiserver v0.18.9
	k8s.io/apimachinery v0.18.9
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/kubelet v0.18.3
	k8s.io/kubernetes v1.18.3
	k8s.io/utils v0.0.0-20200603063816-c1c6865ac451
	kubevirt.io/qe-tools v0.1.6
	sigs.k8s.io/controller-runtime v0.6.2
)

// Pinned to kubernetes-1.18.0
replace (
	k8s.io/api => k8s.io/api v0.18.0
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.0
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.0
	k8s.io/apiserver => k8s.io/apiserver v0.18.0
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.18.0
	k8s.io/client-go => k8s.io/client-go v0.18.0
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.18.0
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.18.0
	k8s.io/code-generator => k8s.io/code-generator v0.18.0
	k8s.io/component-base => k8s.io/component-base v0.18.0
	k8s.io/cri-api => k8s.io/cri-api v0.18.0
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.18.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.18.0
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.18.0
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.18.0
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.18.0
	k8s.io/kubectl => k8s.io/kubectl v0.18.0
	k8s.io/kubelet => k8s.io/kubelet v0.18.0
	k8s.io/kubernetes => k8s.io/kubernetes v1.18.0
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.18.0
	k8s.io/metrics => k8s.io/metrics v0.18.0
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.18.0
)

// Other pinned deps
replace (
	github.com/cri-o/cri-o => github.com/cri-o/cri-o v1.18.1
	//github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309 // Required by Helm
	github.com/go-log/log => github.com/go-log/log v0.1.0
	github.com/mtrmac/gpgme => github.com/mtrmac/gpgme v0.1.1
	github.com/openshift/api => github.com/openshift/api v0.0.0-20200526144822-34f54f12813a // release-4.5
	github.com/openshift/client-go => github.com/openshift/client-go v0.0.0-20200521150516-05eb9880269c // release-4.5
	github.com/openshift/cluster-node-tuning-operator => github.com/openshift/cluster-node-tuning-operator v0.0.0-20200408190329-b227599f61b0 // release-4.5
	github.com/openshift/library-go => github.com/openshift/library-go v0.0.0-20200421122923-c1de486c7d47 // fix bitbucket dependency https://github.com/openshift/library-go/pull/776
	github.com/openshift/machine-config-operator => github.com/openshift/machine-config-operator v0.0.1-0.20200618004043-7b1eb84e0083 // release-4.5
	golang.org/x/tools => golang.org/x/tools v0.0.0-20191206213732-070c9d21b343
)

// Test deps
replace (
	github.com/openshift-kni/performance-addon-operators => github.com/openshift-kni/performance-addon-operators v0.0.0-20201026132112-25109e719e3d // release-4.7
	github.com/openshift/ptp-operator => github.com/openshift/ptp-operator v0.0.0-20201020160817-8db64deb14c7 // release-4.7
	github.com/openshift/sriov-network-operator => github.com/openshift/sriov-network-operator v0.0.0-20201021182021-9de9f4ce8413 // release-4.7
)
