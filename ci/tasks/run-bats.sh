#!/usr/bin/env bash

set -e

cp /usr/local/bin/bosh /usr/local/bin/bosh2
gem install bundle

echo "Setting up artifacts..."
cp ./stemcell/*.tgz stemcell.tgz

# path to the stemcell you want to use for testing
export BAT_STEMCELL="${PWD}/stemcell.tgz"

# path to the bat yaml file which is used to generate the deployment manifest
export BAT_DEPLOYMENT_SPEC="${PWD}/oci-config/bat.yml"

# BOSH CLI executable path
export BAT_BOSH_CLI=`which bosh2`

export BAT_DNS_HOST=8.8.8.8

# the name of infrastructure that is used by bosh deployment. Examples: oci, aws, vsphere, openstack, warden.
export BAT_INFRASTRUCTURE=oci

# the type of networking being used: dynamic or manual.
export BAT_NETWORKING="manual"

# Run tests with --fail-fast and skip cleanup in case of failure (optional)
# export BAT_DEBUG_MODE=

env_vars="${PWD}/deployment/director-env-vars.yml"
creds_yml="${PWD}/deployment/creds.yml"

export BOSH_ENVIRONMENT="$(bosh2 int ${env_vars} --path /external_ip)"
export BOSH_CLIENT=admin
export BOSH_CLIENT_SECRET="$(bosh2 int ${creds_yml} --path /admin_password)"
export BOSH_CA_CERT="$(bosh2 int ${creds_yml} --path /director_ssl/ca)"

# the path to ssh key, used by OS specs to ssh into BOSH VMs
export BAT_PRIVATE_KEY="$(bosh2 int ${creds_yml} --path /cpikeys/private_key)"

# BOSH_GW_PRIVATE_KEY is used  by `bosh ssh` to proxy through a gateway
# Needed when BATs VMs are only accessible
# from the director(gateway)
ssh_key_path=/tmp/bat_private_key
echo "${BAT_PRIVATE_KEY}" > ${ssh_key_path}
chmod 600 ${ssh_key_path}
export BOSH_GW_PRIVATE_KEY=${ssh_key_path}


echo "Using BOSH CLI version..."
bosh2 -v

echo "Targeting BOSH director..."
bosh2 login

pushd bats
  echo "Installing gems..."
  bundle install

  echo "Running BOSH Acceptance Tests..."
  bundle exec rspec spec --tag ~vip_networking --tag ~dynamic_networking --tag ~root_partition --tag ~raw_ephemeral_storage --tag ~changing_static_ip --tag ~network_reconfiguration --tag ~dns --tag ~multiple_manual_networks --tag ~persistent_disk
popd
