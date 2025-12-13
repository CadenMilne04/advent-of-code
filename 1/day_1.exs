# star1.exs
defmodule Dial do
  @max 100

  def star_1(file_path) do
    {_final_pos, zero_count} =
      File.stream!(file_path)
      |> Stream.map(&String.trim/1)
      |> Stream.reject(&(&1 == ""))
      |> Enum.reduce({50, 0}, fn line, {pos, count} ->
        direction = String.first(line)
        amount = String.slice(line, 1..-1//1) |> String.to_integer()
        new_pos = rotate(pos, direction, amount)
        count = if new_pos == 0, do: count + 1, else: count
        {new_pos, count}
      end)

    IO.puts("Star 1: zero count = #{zero_count}")
  end

  def star_2(file_path) do
    {_final_pos, zero_count} =
      File.stream!(file_path)
      |> Stream.map(&String.trim/1)
      |> Stream.reject(&(&1 == ""))
      |> Enum.reduce({50, 0}, fn line, {pos, count} ->
        direction = String.first(line)
        amount = String.slice(line, 1..-1//1) |> String.to_integer()
        {new_pos, hits} = rotate_step_by_step(pos, direction, amount)
        {new_pos, count + hits}
      end)

    IO.puts("Star 2: zero count = #{zero_count}")
  end

  # Rotate the dial one step at a time and count every time 0 is passed
  defp rotate_step_by_step(pos, direction, amount) do
    Enum.reduce(1..amount, {pos, 0}, fn _, {current_pos, hits} ->
      new_pos =
        case direction do
          "R" -> rem(current_pos + 1, @max)
          "L" -> rem(current_pos - 1 + @max, @max)
        end

      new_hits = if new_pos == 0, do: hits + 1, else: hits
      {new_pos, new_hits}
    end)
  end

  defp rotate(pos, "L", amount), do: rem(pos - amount + @max, @max)
  defp rotate(pos, "R", amount), do: rem(pos + amount, @max)
end

# Run it
Dial.star_1("input.txt")
Dial.star_2("input.txt")
