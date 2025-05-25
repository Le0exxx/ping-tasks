export TAG ?= latest
export DOCKER_REGISTRY ?= le0exxx

build_all: \
	docker_build \
	docker_tag \
	docker_push

docker_build: 
	docker build -t ping-tasks-webservice .

docker_tag:
	docker tag ping-tasks-webservice ${DOCKER_REGISTRY}/ping-tasks-webservice:${TAG}

docker_push:
	docker push ${DOCKER_REGISTRY}/ping-tasks-webservice:${TAG}

helm_all: \
	helm_ingress \
	helm_app

helm_ingress:
	helm upgrade --install ingress-nginx ingress-nginx \
	  --repo https://kubernetes.github.io/ingress-nginx \
	  --namespace ingress-nginx --create-namespace \
	  --version 4.10.0
	echo "Sleep 10 seconds to wait ingress ready"
	sleep 10

helm_app:
	helm install ping-tasks-webservice ./helm \
	--namespace ping-tasks-webservice --create-namespace

helm_uninstall:
	helm uninstall ping-tasks-webservice -n ping-tasks-webservice
	helm uninstall ingress-nginx  -n ingress-nginx
	