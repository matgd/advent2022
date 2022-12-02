def read_lines(path):
    with open(path, 'r') as f:
        lines = f.readlines()
    return [line.strip() for line in lines]
