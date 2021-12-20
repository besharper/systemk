module github.com/virtual-kubelet/systemk

go 1.16

require (
	github.com/ForAllSecure/rootfs_builder v0.0.0-20200521031332-adbe259b64bf
	github.com/andreyvit/diff v0.0.0-20170406064948-c7f18ee00883
	github.com/coreos/go-systemd/v22 v22.3.2
	github.com/ghodss/yaml v1.0.0
	github.com/gorilla/mux v1.8.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/virtual-kubelet/virtual-kubelet v1.5.1-0.20210601190559-68347d4ed102
	k8s.io/api v0.23.0
	k8s.io/apimachinery v0.23.0
	k8s.io/client-go v0.23.0
	k8s.io/klog/v2 v2.30.0
	k8s.io/kubectl v0.21.1
	k8s.io/kubelet v0.23.0
)
