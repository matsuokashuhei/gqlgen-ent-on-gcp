include {
  path = find_in_parent_folders()
}

dependency "dns" {
  config_path = "../dns"
}

inputs = {
  dns = dependency.dns.outputs.dns
}
