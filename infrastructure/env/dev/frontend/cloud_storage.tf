resource "google_storage_bucket" "frontend" {
  name                        = "landin.tk"
  location                    = "ASIA"
  storage_class               = "COLDLINE"
  uniform_bucket_level_access = true
  force_destroy               = true
  website {
    main_page_suffix = "index.html"
    not_found_page   = "index.html"
  }
}

resource "google_storage_bucket_iam_binding" "frontend" {
  bucket = google_storage_bucket.frontend.name
  role   = "roles/storage.objectViewer"
  members = [
    "allUsers",
  ]
}
