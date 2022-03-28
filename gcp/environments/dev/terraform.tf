terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.15.0"
    }
  }
  backend "gcs" {
    bucket = "landin-337813-tfstate"
    prefix = "env/dev"
  }
}
