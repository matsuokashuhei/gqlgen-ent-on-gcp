data "google_secret_manager_secret" "landin" {
  for_each  = toset(["sql_user_name", "sql_user_password"])
  secret_id = each.key
}

data "google_secret_manager_secret_version" "landin" {
  for_each = toset(["sql_user_name", "sql_user_password"])
  secret   = data.google_secret_manager_secret.landin[each.key].id
}

# data "google_iam_policy" "backend_secret_accessor" {
#   binding {
#     role = "roles/secretmanager.secretAccessor"
#     members = [
#       "serviceAccount:${google_service_account.backend.email}",
#     ]
#   }
# }

# resource "google_secret_manager_secret_iam_policy" "backend" {
#   for_each    = toset(["sql_user_name", "sql_user_password"])
#   project     = data.google_secret_manager_secret.landin[each.key].project
#   secret_id   = data.google_secret_manager_secret.landin[each.key].secret_id
#   policy_data = data.google_iam_policy.backend["roles/secretmanager.secretAccessor"].policy_data
# }

# resource "google_secret_manager_secret_iam_binding" "cloudbuild" {
#   for_each  = toset(["sql_user_name", "sql_user_password"])
#   project   = data.google_secret_manager_secret.landin[each.key].project
#   secret_id = data.google_secret_manager_secret.landin[each.key].secret_id
#   role      = "roles/secretmanager.secretAccessor"
#   members = [
#     "serviceAccount:${google_service_account.backend.email}",
#   ]
# }
