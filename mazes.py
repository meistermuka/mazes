import random


class Cell(object):

    def __init__(self, x_coord=None, y_coord=None, cell_data={}):
        self.x = x_coord
        self.y = y_coord
        self.data = cell_data
        self.north = None
        self.south = None
        self.east = None
        self.west = None

    def link(self, direction, linked_cell, bidi=True):
        if direction == "north" and bidi:
            self.north = linked_cell
            linked_cell.south = self

        if direction == "south" and bidi:
            self.south = linked_cell
            linked_cell.north = self

        if direction == "east" and bidi:
            self.east = linked_cell
            linked_cell.west = self

        if direction == "west" and bidi:
            self.west = linked_cell
            linked_cell.east = self


class Grid(object):

    def __init__(self, x_size, y_size):
        self.x_size = x_size
        self.y_size = y_size
        self.maze_type = ""
        self.cells = []
        self.reset_grid()

    def reset_grid(self):
        self.cells = []
        self.maze_type = ""

        # Initializing basic grid
        for row in range(self.y_size):
            row_list = []
            for col in range(self.x_size):
                row_list.append(Cell(x_coord=row, y_coord=col))

            self.cells.append(row_list)

    def print_grid(self):
        output = "+"
        body = "   "
        corner = "+"

        # Draw the top part
        for x in range(self.x_size):
            output += "---+"

        output += "\n"

        for row in range(self.y_size):
            top = "|"
            bottom = "+"
            for col in range(self.x_size):
                cell = self.cells[row][col]

                if cell.east is not None and col < self.x_size - 1:
                    top += body
                    top += " "
                else:
                    top += body
                    top += "|"

                if cell.south is not None and row < self.y_size - 1:
                    bottom += body
                    bottom += corner
                else:
                    bottom += "---"
                    bottom += corner

            output += top + "\n"
            output += bottom + "\n"

        print(self.maze_type)
        print(output)

    def sidewinder(self):
        for row in range(self.y_size):
            side_run = []
            for col in range(self.x_size):
                side_run.append(self.cells[row][col])

                at_east_bound = bool(col == self.x_size - 1)
                at_north_bound = bool(row == 0)
                should_closeout = bool(at_east_bound or (not at_north_bound and random.randrange(0, 2) == 0))

                index = random.randrange(0, len(side_run))
                sample = side_run[index]
                sample_x = sample.x
                sample_y = sample.y

                if should_closeout:
                    if sample_y > 0 and self.cells[row][col - 1]:
                        # Link north/south
                        self.cells[row][col].link("north", self.cells[row][col - 1])
                        side_run = []
                else:
                    # Link east/west
                    if sample_x < self.y_size - 1:
                        self.cells[row][col].link("east", self.cells[row +1][col])

        self.maze_type = "Sidewinder"

    def binary_tree(self):
        neighbours = []

        for row in range(self.y_size):
            for col in range(self.x_size):

                if row > 0:
                    neighbours.append(self.cells[row - 1][col])

                if col < self.x_size - 1:
                    neighbours.append(self.cells[row][col + 1])

                # Check if south is a valid cell

                index = random.randrange(0, len(neighbours))
                the_cell_x = neighbours[index].x
                the_cell_y = neighbours[index].y

                # If the cell to the west is valid
                if the_cell_y - 1 > -1:
                    self.cells[the_cell_x][the_cell_y].link("east", self.cells[the_cell_x][the_cell_y - 1])

                # IF the cell to the south is valid
                if the_cell_x + 1 < self.y_size:
                    self.cells[the_cell_x][the_cell_y].link("south", self.cells[the_cell_x + 1][the_cell_y])

        self.maze_type = "Binary Tree"


if __name__ == "__main__":

    maze = Grid(7, 5)
    maze.print_grid()
    maze.sidewinder()
    maze.print_grid()
    maze.reset_grid()
    maze.print_grid()
    maze.binary_tree()
    maze.print_grid()
