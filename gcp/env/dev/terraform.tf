terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.15.0"
    }
  }
  backend "gcs" {
    bucket = "landin-dev-tfstate"
    prefix = "env/dev"
  }
}
