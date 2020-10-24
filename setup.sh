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
    --set fluent-bit.enabled=true \
    --set fluent-bit.config.batchWait="1s" \
    --set fluent-bit.config.batchSize="1048576" \
    --set promtail.enabled=false \
    --set grafana.enabled=true \
    --set loki.serviceName=loki.loki.svc.cluster.local \

echo -n "---> Admin password:"
kubectl get secret -n "${LOKI_NS}" loki-grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo

echo -n "---> Open grafana"
kubectl port-forward --namespace loki service/loki-grafana 3000:80