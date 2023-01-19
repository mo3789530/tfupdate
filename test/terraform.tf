terraform {
  required_version = "~> 1.1.0"
  required_providers {
    docker = {
        host = "unix:///var/run/docker.sock"
    }
  }
}