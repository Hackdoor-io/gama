defmodule Gama.Repo do
  use Ecto.Repo,
    otp_app: :gama,
    adapter: Ecto.Adapters.Postgres
end
