VERSION?=v1.19.4
VERSION_CCM?=v1.19.2
REGISTRY?=k8s.gcr.io
IMAGE?=keppel.eu-de-1.cloud.sap/ccloud
ARCH=amd64
CNI_PLUGINS_RELEASE=v0.8.5

all: build push

build:
	docker build -t ${IMAGE}/kube-apiserver:${VERSION} --build-arg IMAGE=${REGISTRY}/kube-apiserver:${VERSION} -f Dockerfile.apiserver .
	docker build -t ${IMAGE}/kube-scheduler:${VERSION} --build-arg IMAGE=${REGISTRY}/kube-scheduler:${VERSION} .
	docker build -t ${IMAGE}/kube-controller-manager:${VERSION} --build-arg IMAGE=${REGISTRY}/kube-controller-manager:${VERSION} .
	docker build -t ${IMAGE}/kube-proxy:${VERSION} --build-arg IMAGE=${REGISTRY}/kube-proxy:${VERSION} .
	docker build -t ${IMAGE}/kubelet:${VERSION} --build-arg KUBERNETES_VERSION=${VERSION} --build-arg ARCH=${ARCH} --build-arg CNI_PLUGINS_RELEASE=${CNI_PLUGINS_RELEASE} - < Dockerfile.kubelet
	docker build -t ${IMAGE}/openstack-cloud-controller-manager:${VERSION_CCM} --build-arg IMAGE=k8scloudprovider/openstack-cloud-controller-manager:${VERSION_CCM} .

push:
	docker push ${IMAGE}/kube-apiserver:${VERSION}
	docker push ${IMAGE}/kube-scheduler:${VERSION}
	docker push ${IMAGE}/kube-controller-manager:${VERSION}
	docker push ${IMAGE}/kube-proxy:${VERSION}
	docker push ${IMAGE}/kubelet:${VERSION}
	docker push ${IMAGE}/openstack-cloud-controller-manager:${VERSION_CCM}

.PHONY: all build push
