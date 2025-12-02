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
    start_num = 50
    max_nums = 100
    times = 0

    with open("input.txt") as f:
        for line in f:
            line = line.strip()
            match line[0]:
                case "L":
                    times += (int(line[1:]) + (max_nums - start_num - 1)) // max_nums
                    start_num = (start_num - int(line[1:])) % max_nums
                case "R":
                    times += (int(line[1:]) + start_num) // max_nums
                    start_num = (start_num + int(line[1:])) % max_nums

        print(times)

def main():
    star_1()
    star_2()


if __name__ == "__main__":
    main()
