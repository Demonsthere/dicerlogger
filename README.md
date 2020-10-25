# dicerlogger


## Setup
- run `setup.sh`, copy GRAFANA_PASSWORD
- wait for all pods in `loki` namespace to start
- forward grafana port as printed by the setup script
- open [grafana](localhost:80)
- login as `admin:$GRAFANA_PASSWORD`
- import [dashboard](config/grafana-dashboard.json)
- explore query `{custom_error_level=~"critical|error|fatal"}`
