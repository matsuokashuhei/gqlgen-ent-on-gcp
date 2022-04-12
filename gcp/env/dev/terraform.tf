terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.15.0"
    }
    google-beta = {
      source  = "hashicorp/google-beta"
      version = "4.15.0"
    }
    github = {
      source  = "integrations/github"
      version = "4.23.0"
    }
  }
  backend "gcs" {
    bucket = "landin-dev-346809-tfstate"
    prefix = "env/dev"
  }
}
