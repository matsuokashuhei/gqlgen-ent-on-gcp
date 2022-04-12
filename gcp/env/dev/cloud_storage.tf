# resource "google_storage_bucket" "frontend" {
#   name          = join("-", [var.project.id, "frontend"])
#   force_destroy = true
#   location      = "ASIA"
#   storage_class = "MULTI_REGIONAL"
# }

# data "google_iam_policy" "frontend" {
#   binding {
#     role = "roles/storage.objectViewer"
#     members = [
#       "allUsers",
#     ]
#   }
# }

# resource "google_storage_bucket_iam_policy" "frontend" {
#   bucket      = google_storage_bucket.frontend.name
#   policy_data = data.google_iam_policy.frontend.policy_data
# }


resource "google_storage_bucket" "slack_notifier" {
  name          = join("-", [var.project.id, "slack-notifier"])
  location      = "ASIA"
  force_destroy = true
}

resource "google_storage_bucket_object" "slack_notifier" {
  for_each = toset(["manifest.yaml"])

  name   = each.key
  source = "slack_notifier/${each.key}"
  bucket = google_storage_bucket.slack_notifier.name
}

# resource "google_storage_bucket" "frontend" {
#   name          = join("-", [var.project.id, "frontend", "again"])
#   force_destroy = true
#   location      = "ASIA"
#   storage_class = "MULTI_REGIONAL"
# }

# resource "google_storage_bucket_iam_binding" "frontend" {
#   bucket = google_storage_bucket.frontend.name
#   role   = "roles/storage.objectViewer"
#   members = [
#     "allUsers",
#   ]
# }
