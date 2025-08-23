#!/usr/bin/env bash

echo "Installing Helm chart for myapp..."

cd ..
cd 6.2-helm/myapp

helm install myapp . 

echo "Installing Helm chart for myapp./.\./.\./.\......."

sleep 3

echo "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"

sleep 3

oc status 

curl $(oc get route myapp -o jsonpath='{.spec.host}') \n

sleep 3 

echo "Deleting the app"

sleep 90

helm uninstall myapp