resource "google_secret_manager_secret" "landin" {
  for_each  = var.secret_manager
  secret_id = each.key
  labels = {
    label = each.value.label
  }
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "landin" {
  for_each    = var.secret_manager
  secret      = google_secret_manager_secret.landin[each.key].id
  secret_data = each.value.value
}
