resource "google_cloudbuild_trigger" "landin" {
  for_each = toset(["frontend", "backend"])
  name     = join("-", ["github", each.key, "trigger"])
  tags     = [each.key]
  github {
    owner = data.github_user.landin.username
    name  = data.github_repository.landin.name
    push {
      branch       = "setup-gcp"
      invert_regex = false
    }
  }
  filename = "${each.key}/cloudbuild.yaml"
  included_files = [
    "${each.key}/**/*"
  ]
}
