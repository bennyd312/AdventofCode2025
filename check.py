def largest_digit(sub_bank: str) -> tuple[str]:
    """Find the largest digit in a subset of the bank, and return the batteries
    before it, the largest battery, and the batteries after it """
    b_max = "0"
    i_max = -1
    for i, battery in enumerate(sub_bank):
        if battery > b_max:
            b_max = battery
            i_max = i
        if battery == 9:
            break
    return (
        sub_bank[0:i_max],
        b_max,
        sub_bank[i_max+1:]
    )

def main(banks: list[str]):
    ans = 0
    for bank in banks:
        selected_batteries = []
        remaining_bank = bank
        while True:
            split_index = len(remaining_bank) - (12 - len(selected_batteries)) + 1
            possible_batteries = remaining_bank[0:split_index]
            unselectable_batteries = remaining_bank[split_index:]
            _, b_max, remainder = largest_digit(possible_batteries)
            selected_batteries.append(b_max)
            remaining_bank = remainder + unselectable_batteries
            if len(selected_batteries) == 12:
                break
        max_battery = "".join(selected_batteries)
        print(max_battery)
        ans += int(max_battery)
    print(f"Ans: {ans}")

data = []
with open("3a_input.txt",'r') as f:
    for line in f.readlines():
        data.append(line.rstrip())
main(data)