import sys
import pathlib
import json

SCRIPT_DIR = pathlib.Path(__file__).parent.absolute()
sys.path.append((SCRIPT_DIR.parent / "utils").as_posix())

from utils import read_lines

def in_order(packet_a: list[any], packet_b: list[any]) -> bool:
    for i in range(max(len(packet_a), len(packet_b))):
        try:
            el_a = packet_a[i]
        except IndexError:
            return True

        try:
            el_b = packet_b[i]
        except IndexError:
            return False


        if type(el_a) == list and type(el_b) == int:
            res = in_order(el_a, [el_b])
            if res is not None:
                return res
        
        if type(el_a) == int and type(el_b) == list:
            res = in_order([el_a], el_b)
            if res is not None:
                return res

        if type(el_a) == int and type(el_b) == int:
            if el_a > el_b:
                return False
            elif el_a < el_b:
                return True
            else:
                continue
        
        if type(el_a) == list and type(el_b) == list:
            res = in_order(el_a, el_b)
            if res is not None:
                return res
            
    return None


def day13(input_data: list[list[str]]):
    index = 0
    in_right_order = []
    for i in range(0, len(input_data), 3):
        index += 1
        packet_A = json.loads(input_data[i])
        packet_B = json.loads(input_data[i+1])
        if in_order(packet_A, packet_B):
            in_right_order.append(index)

    print("Day 13, pt. 1:", sum(in_right_order))

def day13p2(input_data: list[list[str]]):
    packets = []
    for line in input_data:
        if line:
            packets.append(json.loads(line))
    packets.append([[2]])
    packets.append([[6]])
    
    # Bubble sort
    for i in range(len(packets)):
        for j in range(len(packets) - i - 1):
            if not in_order(packets[j], packets[j+1]):
                packets[j], packets[j+1] = packets[j+1], packets[j]

    print("Day 13, pt. 2:", (packets.index([[2]])+1) * (packets.index([[6]])+1))



if __name__ == "__main__":
    day13(read_lines(SCRIPT_DIR / "input.txt"))
    day13p2(read_lines(SCRIPT_DIR / "input.txt"))
    