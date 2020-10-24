#!/bin/bash -e

export DICELOGGER_NS="dicelogger"
export LOKI_NS="loki"

minikube start

helm install dicelogger deploy/helm/dicelogger \
    -n "${DICELOGGER_NS}" \
    --create-namespace

helm install loki deploy/helm/loki-stack \
    -n "${LOKI_NS}" \
    --create-namespace \
    -f config/loki-overrides.yaml

echo -n "---> Admin password: "
kubectl get secret -n "${LOKI_NS}" loki-grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo

echo -n "---> Open grafana: kubectl port-forward --namespace ${LOKI_NS} service/loki-grafana 3000:80"