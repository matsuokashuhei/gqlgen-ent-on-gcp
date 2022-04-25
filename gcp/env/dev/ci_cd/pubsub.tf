resource "google_pubsub_topic" "cloud_builds" {
  name = "cloud-builds"
}

resource "google_pubsub_subscription" "cloud_builds" {
  name  = "cloud-builds"
  topic = google_pubsub_topic.cloud_builds.name
  push_config {
    push_endpoint = google_cloud_run_service.slack_notifier.status[0].url
    oidc_token {
      service_account_email = google_service_account.slack_notifier.email
    }
  }
}
