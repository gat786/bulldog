# Bulldog

A simple CLI Golang application which you can use to scrape your kubernetes
cluster for resources you are looking for.

This tool can list and give you entire manifests that you have deployed on
your cluster. It can help in investigating objects that are residing on your
cluster as well as you can use this tool for monitoring purposes for example
this can run as a job everyday at some specified time and store the objects on your
cluster and send you a list of files with each manifests that you can then use to
do some data analysis on those or to find naughty incorrectly defined manifests
without manually having to go through each namespace and doing a kubectl get.

