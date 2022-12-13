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

    print(sum(in_right_order))

if __name__ == "__main__":
    day13(read_lines(SCRIPT_DIR / "input_example.txt"))
    day13(read_lines(SCRIPT_DIR / "input.txt"))
    