defmodule Day3 do
  def star2(file_name) do
    file_name
    |> File.read!()
    |> String.trim()
    |> String.split("\n")
    |> Enum.map(&max_12_digit_joltage/1)
    |> Enum.sum()
    |> IO.puts()
  end

  # Compute the largest possible 12-digit number
  # by removing digits while preserving order
  defp max_12_digit_joltage(str) do
    digits =
      str
      |> String.graphemes()
      |> Enum.map(&String.to_integer/1)

    k = 12
    drop = length(digits) - k

    {stack, _} =
      Enum.reduce(digits, {[], drop}, fn d, {stack, drop_left} ->
        {stack, drop_left} = pop_smaller(stack, d, drop_left)
        {[d | stack], drop_left}
      end)

    stack
    |> Enum.reverse()
    |> Enum.take(k)
    |> Enum.join()
    |> String.to_integer()
  end

  # Pop smaller digits if we can still drop digits
  defp pop_smaller([h | t], d, drop_left) when drop_left > 0 and h < d do
    pop_smaller(t, d, drop_left - 1)
  end

  defp pop_smaller(stack, _d, drop_left), do: {stack, drop_left}

  def star1(file_name) do
    cells =
      file_name
      |> File.read!()
      |> String.trim()
      |> String.split("\n")

    cell_max = fn str ->
      digits = String.graphemes(str) |> Enum.map(&String.to_integer/1)

      {max_num, _} =
        Enum.reduce(Enum.reverse(digits), {0, -1}, fn digit, {current_max, max_right} ->
          if max_right == -1 do
            {current_max, digit}
          else
            new_max = max(current_max, digit * 10 + max_right)
            {new_max, max(digit, max_right)}
          end
        end)

      [max_num]
    end

    battery_total =
      cells
      |> Enum.flat_map(fn cell ->
        cell_max.(cell)
      end)
      |> Enum.sum()

    IO.puts(battery_total)
  end
end

Day3.star1("input.txt")
Day3.star2("input.txt")
