resource "google_cloud_run_service" "backend" {
  name                       = "backend"
  location                   = var.region
  autogenerate_revision_name = true
  template {
    spec {
      containers {
        # image = "${var.region}-docker.pkg.dev/${var.project_id}/${google_artifact_registry_repository.landin.repository_id}/backend:84df23a"
        image = "gcr.io/cloudrun/hello"
        # env {
        #   name = "MYSQL_USER"
        #   value_from {
        #     secret_key_ref {
        #       name = data.google_secret_manager_secret.landin["sql_user_name"].secret_id
        #       key  = "latest"
        #     }
        #   }
        # }
        # env {
        #   name = "MYSQL_PASSWORD"
        #   value_from {
        #     secret_key_ref {
        #       name = data.google_secret_manager_secret.landin["sql_user_password"].secret_id
        #       key  = "latest"
        #     }
        #   }
        # }
        # env {
        #   name  = "MYSQL_HOST"
        #   value = data.google_sql_database_instance.landin.private_ip_address
        # }
        # env {
        #   name  = "MYSQL_PORT"
        #   value = "3306"
        # }
        # env {
        #   name  = "MYSQL_DATABASE"
        #   value = var.db.google_sql_database.landin.name
        # }
      }
      service_account_name = data.google_service_account.landin["backend"].email
    }
    metadata {
      annotations = {
        "run.googleapis.com/vpc-access-connector" = google_vpc_access_connector.backend.name
        "run.googleapis.com/vpc-access-egress"    = "all"
      }
    }
  }
  # lifecycle {
  #   ignore_changes = [
  #     template[0].spec[0].containers[0].image
  #   ]
  # }
}
