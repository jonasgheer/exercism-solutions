defmodule Lasagna do
  def expected_minutes_in_oven(), do: 40

  def remaining_minutes_in_oven(minutes_in_oven), do: expected_minutes_in_oven - minutes_in_oven

  def preparation_time_in_minutes(num_layers), do: num_layers * 2

  def total_time_in_minutes(num_layers, minutes_in_oven), do: minutes_in_oven + preparation_time_in_minutes(num_layers)

  def alarm(), do: "Ding!"
end
