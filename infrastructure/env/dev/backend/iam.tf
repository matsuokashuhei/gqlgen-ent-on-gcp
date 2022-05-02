data "google_service_account" "landin" {
  for_each   = var.project.google_service_account
  account_id = each.value.account_id
}
# resource "google_service_account" "backend" {
#   account_id   = "backend"
#   display_name = "backend"
# }

# resource "google_project_iam_member" "backend" {
#   for_each = toset([
#     "roles/secretmanager.secretAccessor",
#   ])
#   project = data.google_project.landin.project_id
#   role    = each.key
#   member  = "serviceAccount:${google_service_account.backend.email}"
# }

data "google_iam_policy" "backend" {
  # for_each = {
  #   "roles/secretmanager.secretAccessor" = {
  #     members = ["serviceAccount:${google_service_account.backend.email}", ]
  #   }
  #   "roles/run.invoker" = {
  #     members = ["allUsers"]
  #   }
  # }
  # binding {
  #   role    = each.key
  #   members = each.value.members
  # }
  binding {
    role    = "roles/run.invoker"
    members = ["allUsers"]
  }
}

resource "google_cloud_run_service_iam_policy" "backend" {
  location    = google_cloud_run_service.backend.location
  project     = google_cloud_run_service.backend.project
  service     = google_cloud_run_service.backend.name
  policy_data = data.google_iam_policy.backend.policy_data
}
