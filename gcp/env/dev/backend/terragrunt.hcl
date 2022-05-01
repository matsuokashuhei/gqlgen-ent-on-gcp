include {
  path = find_in_parent_folders()
}

dependency "db" {
  config_path = "../db"
}

dependency "dns" {
  config_path = "../dns"
}

dependency "network" {
  config_path = "../network"
}

dependency "project" {
  config_path = "../project"
}

inputs = {
  db      = dependency.db.outputs.db
  dns     = dependency.dns.outputs.dns
  network = dependency.network.outputs.network
  project = dependency.project.outputs.project
}
