defmodule Day2 do
  def star_2(file_path) do
    intervals =
      file_path
      |> File.read!()
      |> String.trim()
      |> String.split(",")

    isInvalid? = fn i ->
      s = Integer.to_string(i)
      len = String.length(s)

      max_n = div(len, 2)

      # handle len = 1 safely
      range = if max_n >= 1, do: 1..max_n, else: []

      Enum.filter(range, &(rem(len, &1) == 0))
      |> Enum.any?(fn n ->
        chunk = binary_part(s, 0, n)
        String.duplicate(chunk, div(len, n)) == s
      end)
    end

    total =
      intervals
      |> Enum.flat_map(fn interval ->
        [left, right] =
          interval
          |> String.split("-")
          |> Enum.map(&String.to_integer/1)

        left..right
      end)
      |> Enum.filter(isInvalid?)
      |> Enum.sum()

    IO.puts(total)
  end

  def star_1(file_path) do
    intervals =
      file_path
      |> File.read!()
      |> String.trim()
      |> String.split(",")

    mirrored_number? = fn i ->
      s = Integer.to_string(i)
      len = String.length(s)

      rem(len, 2) == 0 and
        String.slice(s, 0, div(len, 2)) == String.slice(s, div(len, 2), div(len, 2))
    end

    total =
      intervals
      |> Enum.flat_map(fn interval ->
        [left, right] =
          interval
          |> String.split("-")
          |> Enum.map(&String.to_integer/1)

        left..right
      end)
      |> Enum.filter(mirrored_number?)
      |> Enum.sum()

    IO.puts(total)
  end
end

Day2.star_1("input.txt")
Day2.star_2("input.txt")
