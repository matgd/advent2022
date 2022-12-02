import sys
import pathlib
import string
import math
from dataclasses import dataclass

import networkx as nx

SCRIPT_DIR = pathlib.Path(__file__).parent.absolute()
sys.path.append((SCRIPT_DIR.parent / "utils").as_posix())

from utils import read_lines

@dataclass
class Node:
    x: int
    y: int
    height: int
    stepped: int = 0
    letter: str = ""

    def __hash__(self):
        return hash((self.x, self.y))

    def __repr__(self):
        return f"|y{self.y} x{self.x} h{self.height}|"
        

node_matrix: list[list[Node]] = []
lowest_elevation: list[Node] = []
start_node_coords = [0, 0]
end_node_coords = [0, 0]

def print_node_matrix():
    for row in node_matrix:
        for node in row:
            print(f"{node.height:2}", end=" ")
        print()
    
def print_node_matrix_steps(letter = False):
    for row in node_matrix:
        for node in row:
            print(f"{node.stepped:2}{node.letter if letter else ''}", end=" ")
        print()

def letter_to_height(letter: str):
    if letter == "S":
        letter = 'a'
    if letter == "E":
        letter = 'z'
    return string.ascii_lowercase.index(letter)

def day12(input_data: list[list[str]]):
    graph = nx.DiGraph()
    for y in range(len(input_data)):
        node_matrix.append([])
        for x in range(len(input_data[y])):
            height = letter_to_height(input_data[y][x])
            if input_data[y][x] == "S":
                start_node_coords = [x, y]
            if input_data[y][x] == "E":
                end_node_coords = [x, y]
            node = Node(x, y, height)
            node_matrix[y].append(node)
            node.letter = input_data[y][x]

            # Okay, this is not the best way to do this, but it works.
            # You can notice that every 'b' is x == 1 so we must start on 'x' == 0
            if node.letter == 'a' and x == 0:
                lowest_elevation.append(node)
            graph.add_node(node)


    for row in node_matrix:
        for node in row:
            if node.y - 1 < 0:
                up_node = None
            else:
                up_node = node_matrix[node.y - 1][node.x]

            if node.y > len(node_matrix) - 2:
                down_node = None
            else:
                down_node = node_matrix[node.y + 1][node.x]
            
            if node.x - 1 < 0:
                left_node = None
            else:
                left_node = node_matrix[node.y][node.x - 1]
            
            if node.x > len(node_matrix[0]) - 2:
                right_node = None
            else:
                right_node = node_matrix[node.y][node.x + 1]

            if left_node:
                weight = left_node.height - node.height
                if weight > 1:
                    weight = 10000000
                else:
                    weight = 1
                graph.add_edge(node, left_node, weight=weight)
            if right_node:
                weight = right_node.height - node.height
                if weight > 1:
                    weight = 10000000
                else:
                    weight = 1
                graph.add_edge(node, right_node, weight=weight)
            if up_node:
                weight = up_node.height - node.height
                if weight > 1:
                    weight = 10000000
                else:
                    weight = 1
                graph.add_edge(node, up_node, weight=weight)
            if down_node:
                weight = down_node.height - node.height
                if weight > 1:
                    weight = 10000000
                else:
                    weight = 1
                graph.add_edge(node, down_node, weight=weight)

    start_x = start_node_coords[0]
    start_y = start_node_coords[1]
    end_x = end_node_coords[0]
    end_y = end_node_coords[1]
    path = nx.shortest_path(graph, node_matrix[start_y][start_x], node_matrix[end_y][end_x], weight="weight")

    for i in range(len(path)):
        node = path[i]
        node.stepped = i+1

    print(len(path) - 1)
    paths_lenghts = []
    for i, starting_point in enumerate(lowest_elevation):
        path = nx.shortest_path(graph, starting_point, node_matrix[end_y][end_x], weight="weight")
        paths_lenghts.append(len(path) - 1)
        if i % 20 == 0:
            print(i, '/', len(lowest_elevation))
        
    print(min(paths_lenghts))
    
if __name__ == "__main__":
    day12(read_lines("input.txt"))
    