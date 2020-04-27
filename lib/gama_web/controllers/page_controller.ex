defmodule GamaWeb.PageController do
  use GamaWeb, :controller

  @base_url "https://www.hackdoor.io"
  @utm_source "?utm_source=gama"
  @hashid_length 8
  @valid_targets ["tags", "articles", "authors", "patrons"]

  def index(conn, _params) do
    conn
    |> put_status(:moved_permanently)
    |> redirect(external: to_hackdoor)
  end

  def toExternalResource(conn, %{"target" => target, "id" => id, "path" => path}) do
    case hash_target(target) do
      {:ok, h} -> 
        hashid = Hashids.encode(h, get_id_as_int(id))
        hck_path = Enum.join([rename_target(target), hashid, path], "/")
        
        conn
        |> put_status(:moved_permanently)
        |> redirect(external: to_hackdoor(hck_path ))
      true ->
        conn
        |> put_status(:moved_permanently)
        |> redirect(external: to_hackdoor)
    end
  end

  def catchAllPath(conn, %{"path" => path}) do
    conn
    |> put_status(:moved_permanently)
    |> redirect(external: to_hackdoor(Enum.join(path, "/")))
  end

  defp to_hackdoor(),     do: @base_url <> @utm_source
  defp to_hackdoor(path), do: @base_url <> "/" <> path <> @utm_source

  defp hash_target(target) do
    cond do
      Enum.member?(@valid_targets, target) -> 
        {:ok, Hashids.new([ salt: get_hashids_salt(rename_target(target)), min_len: @hashid_length ])}
      true ->
        {:ko}
    end
  end

  defp get_id_as_int(id) do
    { int, _ } = Integer.parse(id)
    int
  end

  defp get_hashids_salt(target) do
    target
    |> String.capitalize
    |> (&(&1 <> "_SALT")).()
    |> System.get_env
  end

  defp rename_target(target) do
    case target do
      "tags" -> "topics"
      _   -> target
    end
  end

end