oc delete all -lapp=gitea
oc delete cm gitea-config
oc delete pvc gitea-repositories
oc delete all -lapp=gitea-postgresql
oc delete pvc gitea-postgres-data
oc delete secret postgresql
