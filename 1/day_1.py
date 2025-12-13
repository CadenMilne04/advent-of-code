def star_1():
    start_num = 50
    max_nums = 100
    times = 0

    with open("input.txt") as f:
        for line in f:
            line = line.strip()
            match line[0]:
                case "L":
                    start_num = (start_num - int(line[1:])) % max_nums
                case "R":
                    start_num = (start_num + int(line[1:])) % max_nums

            # print(line)
            # print(start_num)
            if start_num == 0:
                times += 1

        print(times)

def star_2():
    pos = 50
    max_nums = 100
    total_hits = 0

    with open("input.txt") as f:
        for line in f:
            line = line.strip()
            if not line:
                continue

            direction = line[0]
            amount = int(line[1:])

            for _ in range(amount):
                if direction == "L":
                    pos = (pos - 1) % max_nums
                else:  # R
                    pos = (pos + 1) % max_nums

                if pos == 0:
                    total_hits += 1

    print(total_hits)

def main():
    star_1()
    star_2()


if __name__ == "__main__":
    main()
