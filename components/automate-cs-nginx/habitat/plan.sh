#shellcheck disable=SC2034
#shellcheck disable=SC2154

pkg_name=automate-cs-nginx
pkg_origin=chef
vendor_origin="chef"
pkg_maintainer="Chef Software Inc. <support@chef.io>"
pkg_license=('Chef-MLSA')
# WARNING: Version managed by .expeditor/update_chef_server.sh
pkg_version="13.0.47"
pkg_deps=(
  chef/mlsa
  # TODO 2020-05-12: PIN PIN PIN
  #
  # All dependencies that are shared between this package and the
  # chef-server-* packages are pinned to the versions required by the
  # chef-server-* packages.
  #
  # Currently, the chef-server-* packages we are pinned to were
  # hand-built and any update to any of these three packages will
  # require them to be re-built.
  #
  # We cannot follow the latest chef-server-* packages until the
  # following issue is resolved:
  #
  # https://github.com/chef/automate/pull/2002#issuecomment-600343881
  #
  core/bundler/1.17.3/20200404130802
  core/curl/7.68.0/20200319191535
  core/ruby/2.5.7/20200404130135
  # WARNING: Version pin managed by .expeditor/update_chef_server.sh
  "${vendor_origin}/chef-server-nginx/13.0.47/20200421235836"
  "${vendor_origin}/chef-server-ctl/13.0.47/20200421235903"
)

pkg_bin_dirs=(bin)
pkg_exposes=(port)
pkg_exports=(
    [port]=service.port
)
pkg_binds=(
  [automate-cs-bookshelf]="http-port"
  [automate-cs-oc-erchef]="http-port"
  [automate-es-gateway]="http-port"
)
pkg_binds_optional=(
  [automate-gateway]="port"
)
pkg_description="NGINX configuration and content for Chef Server Automate Component"
pkg_upstream_url="https://www.chef.io/automate"

# TODO(ssd) 2018-07-27: This instance of nginx should not have to run
# as root. This is here so that we can reliably send access logs to
# stdout. Nginx does not natively have an option to send access logs
# to stdout. To get around this, we use /dev/stdout. Unfortunately,
# because of how we are started by hab, both /dev/stdout and
# /proc/self/fd/1 both eventually point to files that are owned by
# root. We could potentially get around this with some named pipes and
# `tail -f`, but since the nginx worker threads will drop priv's to
# the `hab` user I've opted to set our svc_user as root, in line with
# the other nginx instances in A2 at the moment:
pkg_svc_user="root"

pkg_scaffolding="${local_scaffolding_origin:-chef}/automate-scaffolding-go"
scaffolding_go_base_path=github.com/chef
scaffolding_go_repo_name=automate
scaffolding_go_import_path="${scaffolding_go_base_path}/${scaffolding_go_repo_name}/components/${pkg_name}"
scaffolding_go_binary_list=(
  "${scaffolding_go_import_path}/cmd/chef-server-ctl"
)

chef_automate_hab_binding_mode="relaxed"

do_prepare() {
  GO_LDFLAGS="-X main.BundlePath=$(pkg_path_for core/bundler)"
  GO_LDFLAGS="$GO_LDFLAGS -X main.RubyPath=$(pkg_path_for core/ruby)"
  GO_LDFLAGS="$GO_LDFLAGS -X main.ChefServerCtlPath=$(pkg_path_for chef/chef-server-ctl)"
  GO_LDFLAGS="$GO_LDFLAGS -X main.KnifePath=${pkg_prefix}/bin/knife"
  GO_LDFLAGS="$GO_LDFLAGS -X main.Version=${pkg_version}/${pkg_release}"
  export GO_LDFLAGS
}

do_install() {
  # Install chef-server-ctl shim
  do_default_install

  # Install knife shim
  wrapper_bin_path="${pkg_prefix}/bin"
  install "$PLAN_CONTEXT/bin/knife" "$wrapper_bin_path/knife"

  sed -i "s!__BUILDTIME_HAB_PKG_PATH_CHEF_SERVER_CTL__!$(pkg_path_for chef/chef-server-ctl)!g" "$wrapper_bin_path/knife"
  sed -i "s!__BUILDTIME_HAB_PKG_PATH_BUNDLER__!$(pkg_path_for core/bundler)!g" "$wrapper_bin_path/knife"
}
