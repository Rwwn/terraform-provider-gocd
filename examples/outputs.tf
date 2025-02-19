output "encrypted_value" {
  value     = gocd_encrypt_value.new_value.encrypted_value
  sensitive = true
}

output "sample_config_repo" {
  value = data.gocd_config_repository.sample_config_repo.material
}

output "password_file_config" {
  value = data.gocd_auth_config.password_file_config.properties
}

output "ec2_cluster_profile" {
  value = data.gocd_cluster_profile.ec2_cluster_profile.properties
}

output "sample_environment" {
  value = data.gocd_environment.sample_environment.pipelines
}

output "sample_ec2" {
  value = data.gocd_elastic_agent_profile.sample_ec2.properties
}

output "yaml_plugin_settings" {
  value = data.gocd_plugin_setting.yaml_plugin_settings.configuration
}

output "sample_kube_secret_config" {
  value = data.gocd_secret_config.sample_kube_secret_config
}