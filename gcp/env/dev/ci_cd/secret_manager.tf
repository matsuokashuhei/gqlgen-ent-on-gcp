data "google_secret_manager_secret" "landin" {
  for_each  = toset(["slack_webhook_url", "sql_user_name", "sql_user_password"])
  secret_id = each.key
}

data "google_secret_manager_secret_version" "landin" {
  for_each = toset(["slack_webhook_url", "sql_user_name", "sql_user_password"])
  secret   = data.google_secret_manager_secret.landin[each.key].id
}
