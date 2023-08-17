/*
  NOTE: Creation of cluster level SSH Key type Repository Credential
*/

terraform {
  required_providers {
    tanzu-mission-control = {
      source = "vmware/tanzu-mission-control"
    }
  }
}

# Create cluster group
resource "tanzu-mission-control_cluster_group" "create_cluster_group" {
  name = "demo-cluster-group"
}

# Attach a Tanzu Mission Control cluster with k8s cluster kubeconfig provided
# The provider would create the cluster entry and apply the deployment link manifests on to the k8s kubeconfig provided.
resource "tanzu-mission-control_cluster" "attach_cluster_with_kubeconfig" {
  management_cluster_name = "attached"     // Default: attached
  provisioner_name        = "attached"     // Default: attached
  name                    = "demo-cluster" // Required

  attach_k8s_cluster {
    kubeconfig_file = "<kube-config-path>" // Required
    description     = "optional description about the kube-config provided"
  }

  meta {
    description = "description of the cluster"
    labels      = { "key" : "value" }
  }

  spec {
    cluster_group = tanzu-mission-control_cluster_group.create_cluster_group.name // Default: default
  }

  ready_wait_timeout = "15m" # Default: waits until 3 min for the cluster to become ready
  // The deployment link and the command needed to be run to attach this cluster would be provided in the output.status.execution_cmd
}

# Create cluster level Repository Credential
resource "tanzu-mission-control_repository_credential" "create_cluster_source_secret_ssh" {
  name = "tf-secret" # Required

  scope {
    cluster {
      cluster_name            = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.name                    # Required
      provisioner_name        = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.provisioner_name        # Default: attached
      management_cluster_name = tanzu-mission-control_cluster.attach_cluster_with_kubeconfig.management_cluster_name # Default: attached
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    ssh_key {
      identity    = "testidentity"    # Required
      known_hosts = "testknown_hosts" # Required
    }
  }
}