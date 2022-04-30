data "google_service_account" "landin" {
  for_each   = var.project.google_service_account
  account_id = each.value.account_id
}
