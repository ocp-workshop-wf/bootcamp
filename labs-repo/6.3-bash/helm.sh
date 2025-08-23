#!/usr/bin/env bash
export APP_NAME=myapp
echo "Installing Helm chart for myapp..."

cd ..
cd 6.2-helm/myapp

helm install $APP_NAME . 

echo "Installing Helm chart for $APP_NAME............"

sleep 3

echo "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"

sleep 3

oc status 

curl $(oc get route myapp -o jsonpath='{.spec.host}')

sleep 3 


echo "Deleting $APP_NAME"

sleep 90

helm uninstall $APP_NAME