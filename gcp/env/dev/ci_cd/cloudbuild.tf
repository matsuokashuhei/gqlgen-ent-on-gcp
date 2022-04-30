resource "google_cloudbuild_trigger" "landin" {
  for_each = toset(["frontend", "backend"])
  name     = join("-", ["github", each.key, "trigger"])
  tags     = [each.key]
  github {
    owner = var.github["username"]
    name  = var.github["repository"]
    push {
      branch = var.github["branch"]
    }
  }
  service_account = google_service_account.cloudbuild.id
  filename        = "${each.key}/cloudbuild.yaml"
  included_files = [
    "${each.key}/**/*"
  ]
  substitutions = {
    _NODE_VERSION  = "16"
    _BUCKET_NAME   = data.google_storage_bucket.frontend.name
    _REPOSITORY_ID = var.backend.google_artifact_registry_repository.landin.repository_id
    _REGION        = var.project.region
    _CLOUD_SQL_URL = "mysql://${data.google_secret_manager_secret_version.landin["sql_user_name"].secret_data}:${data.google_secret_manager_secret_version.landin["sql_user_password"].secret_data}@tcp(${data.google_sql_database_instance.landin.private_ip_address}:3306)/${var.db.google_sql_database.landin.name}"
    _POOL_NAME     = google_cloudbuild_worker_pool.landin.id
  }
}

resource "google_service_account" "cloudbuild" {
  account_id   = "cloudbuild"
  display_name = "cloudbuild"
}

resource "google_project_iam_member" "cloudbuild" {
  for_each = toset([
    "roles/artifactregistry.writer",
    "roles/cloudbuild.builds.editor",
    "roles/cloudsql.client",
    "roles/compute.networkAdmin",
    "roles/iam.serviceAccountUser",
    "roles/logging.logWriter",
    "roles/run.admin",
    "roles/storage.admin",
  ])
  project = data.google_project.landin.project_id
  role    = each.key
  member  = "serviceAccount:${google_service_account.cloudbuild.email}"
}

resource "google_cloudbuild_worker_pool" "landin" {
  name     = "landin"
  location = var.project.region
  network_config {
    peered_network = data.google_compute_network.landin.id
  }
}
