# This file is responsible for configuring your application
# and its dependencies with the aid of the Mix.Config module.
#
# This configuration file is loaded before any dependency and
# is restricted to this project.

# General application configuration
use Mix.Config

config :gama,
  ecto_repos: [Gama.Repo]

# Configures the endpoint
config :gama, GamaWeb.Endpoint,
  url: [host: "localhost"],
  secret_key_base: "SS9sxBg3uJ/Gl+X+pIPr15VSEqCjHniIWI1N0ASgBzE79wcnbAbVhmq+9349CRJV",
  render_errors: [view: GamaWeb.ErrorView, accepts: ~w(json)],
  pubsub: [name: Gama.PubSub, adapter: Phoenix.PubSub.PG2],
  live_view: [signing_salt: "sAeQIiUO"]

# Configures Elixir's Logger
config :logger, :console,
  format: "$time $metadata[$level] $message\n",
  metadata: [:request_id]

# Use Jason for JSON parsing in Phoenix
config :phoenix, :json_library, Jason

# Import environment specific config. This must remain at the bottom
# of this file so it overrides the configuration defined above.
import_config "#{Mix.env()}.exs"
