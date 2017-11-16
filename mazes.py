import random

DIMENSION = 4

class Cell(object):

    def __init__(self, x_coord=None, y_coord=None, cell_data={}):
        self.x = x_coord
        self.y = y_coord
        self.data = cell_data
        self.north = None
        self.south = None
        self.east = None
        self.west = None


def printMap(cells):
    output = "+"
    body = "   "
    corner = "+"

    # Draw the top part
    for x in range(DIMENSION):
        output += "---+"

    output += "\n"

    for row in range(DIMENSION):
        top = "|"
        bottom = "+"
        for col in range(DIMENSION):
            cell = cells[row][col]

            if cell.east != None and col < DIMENSION-1:
                top += body
                top += " "
            else:
                top += body
                top += "|"

            if cell.south != None and row < DIMENSION-1:
                bottom += body
                bottom += corner
            else:
                bottom += "---"
                bottom += corner

        output += top + "\n"
        output += bottom + "\n"

        top = ""
        bottom = ""

    print(output)


def sidewinder(cells):

    for row in range(DIMENSION):
        side_run = []
        for col in range(DIMENSION):
            side_run.append(cells[row][col])

            at_east_bound = bool(col == DIMENSION-1)
            at_north_bound = bool(row == 0)
            should_closeout = bool(at_east_bound or (not at_north_bound and random.randrange(0, 2) == 0))

            index = random.randrange(0, len(side_run))
            sample = side_run[index]
            sample_x = sample.x
            sample_y = sample.y

            if should_closeout:
                if sample_y > 0 and cells[row][col-1]:
                    # Link north/south
                    cells[row][col].north = cells[row][col-1]
                    cells[row][col-1].south = cells[row][col]
                    side_run = []
            else:
                # Link east/west
                if sample_x < DIMENSION-1:
                    cells[row][col].east = cells[row+1][col]
                    cells[row+1][col].west = cells[row][col]

def binary_tree(cells):

    neighbours = []

    for row in range(DIMENSION):
        for col in range(DIMENSION):

            if row > 0:
                neighbours.append(cells[row-1][col])

            if col < DIMENSION-1:
                neighbours.append(cells[row][col+1])

            # Check if south is a valid cell

            index = random.randrange(0, len(neighbours))
            the_cell_x = neighbours[index].x
            the_cell_y = neighbours[index].y
            # If the cell to the west is valid
            if the_cell_y-1 > -1:
                cells[the_cell_x][the_cell_y].east = cells[the_cell_x][the_cell_y-1]
                cells[the_cell_x][the_cell_y-1].west = cells[the_cell_x][the_cell_y]

            # IF the cell to the south is valid
            if the_cell_x+1 < DIMENSION:
                cells[the_cell_x][the_cell_y].south = cells[the_cell_x+1][the_cell_y]
                cells[the_cell_x+1][the_cell_y].north = cells[the_cell_x][the_cell_y]


if __name__ == "__main__":

    cells = []
    # Initializing basic grid
    for row in range(DIMENSION):
        row_list = []
        for col in range(DIMENSION):
            row_list.append(Cell(x_coord=row, y_coord=col))

        cells.append(row_list)

    #binary_tree(cells)
    sidewinder(cells)
    printMap(cells)
