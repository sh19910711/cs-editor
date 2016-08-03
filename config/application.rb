require_relative 'boot'

require 'rails/all'

# Require the gems listed in Gemfile, including any gems
# you've limited to :test, :development, or :production.
Bundler.require(*Rails.groups)

module Editor
  class Application < Rails::Application
    # Settings in config/environments/* take precedence over those specified here.
    # Application configuration should go into files in config/initializers
    # -- all .rb files in that directory are automatically loaded.

    # create workspace if not exists
    config.workspace = "./workspace"
    config.before_initialize do
      unless Dir.exists?(config.workspace)
        FileUtils.mkdir_p config.workspace
      end
    end

    config.web_console.whitelisted_ips = '0.0.0.0/0' if Rails.env.development?
  end
end
