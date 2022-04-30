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
  service_account = data.google_service_account.landin["cloudbuild"].id
  filename        = "${each.key}/cloudbuild.yaml"
  included_files = [
    "${each.key}/**/*"
  ]
  substitutions = {
    _NODE_VERSION      = "16"
    _BUCKET_NAME       = data.google_storage_bucket.frontend.name
    _REPOSITORY_ID     = var.backend.google_artifact_registry_repository.landin.repository_id
    _REGION            = var.region
    _CLOUD_SQL_URL     = "mysql://${data.google_secret_manager_secret_version.landin["sql_user_name"].secret_data}:${data.google_secret_manager_secret_version.landin["sql_user_password"].secret_data}@tcp(${data.google_sql_database_instance.landin.private_ip_address}:3306)/${var.db.google_sql_database.landin.name}"
    _POOL_NAME         = google_cloudbuild_worker_pool.landin.id
    _MYSQL_USER_ID     = data.google_secret_manager_secret_version.landin["sql_user_name"].name
    _MYSQL_PASSWORD_ID = data.google_secret_manager_secret_version.landin["sql_user_password"].name
    _MYSQL_HOST        = data.google_sql_database_instance.landin.private_ip_address
    _MYSQL_DATABASE    = var.db.google_sql_database.landin.name
    _VPC_CONNECTOR     = var.backend.google_vpc_access_connector.backend.name
    _LOG_BUCKET        = google_storage_bucket.cloudbuild.name
  }
}

# resource "google_service_account" "cloudbuild" {
#   account_id   = "cloudbuild"
#   display_name = "cloudbuild"
# }

# resource "google_project_iam_member" "cloudbuild" {
#   for_each = toset([
#     "roles/artifactregistry.writer",
#     "roles/cloudbuild.builds.editor",
#     "roles/cloudsql.client",
#     "roles/compute.networkAdmin",
#     "roles/iam.serviceAccountUser",
#     "roles/logging.logWriter",
#     "roles/run.admin",
#     "roles/storage.admin",
#   ])
#   project = data.google_project.landin.project_id
#   role    = each.key
#   member  = "serviceAccount:${google_service_account.cloudbuild.email}"
# }

resource "google_cloudbuild_worker_pool" "landin" {
  name     = "landin"
  location = var.region
  network_config {
    peered_network = data.google_compute_network.landin.id
  }
}

# data "google_iam_policy" "cloudbuild" {
#   binding {k
#     role = "roles/secretmanager.secretAccessor"
#     members = [
#       "serviceAccount:${google_service_account.cloudbuild.email}",
#       "serviceAccount:backend@studio-landin-dev.iam.gserviceaccount.com",
#     ]
#   }
# }

# resource "google_secret_manager_secret_iam_policy" "cloudbuild" {
#   for_each    = toset(["sql_user_name", "sql_user_password"])
#   project     = data.google_secret_manager_secret.landin[each.key].project
#   secret_id   = data.google_secret_manager_secret.landin[each.key].secret_id
#   policy_data = data.google_iam_policy.cloudbuild.policy_data
# }

# resource "google_secret_manager_secret_iam_binding" "cloudbuild" {
#   for_each  = toset(["sql_user_name", "sql_user_password"])
#   project   = data.google_secret_manager_secret.landin[each.key].project
#   secret_id = data.google_secret_manager_secret.landin[each.key].secret_id
#   role      = "roles/secretmanager.secretAccessor"
#   members = [
#     "serviceAccount:${google_service_account.cloudbuild.email}",
#   ]
# }
