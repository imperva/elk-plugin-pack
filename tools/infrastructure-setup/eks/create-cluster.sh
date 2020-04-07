#!/usr/bin/env bash

eksctl create cluster \
        --name lsar-eks-cluster \
        --version 1.14 \
        --region us-east-2 \
        --nodegroup-name lsar-cluster-workers \
        --node-type m5ad.2xlarge \
        --nodes     4 \
        --nodes-min 3 \
        --nodes-max 8 \
        --alb-ingress-access \
        --node-ami auto
