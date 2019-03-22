localhost:
	sh ./local/setup.sh mongodb 

gcp_k8s:
	sh ./k8s/setup.sh mongodb

stop:
	sh ./k8s/stop.sh
	sh ./local/stop.sh