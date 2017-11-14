import random

DIMENSION = 10

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

            if cell.east != None:
                top += body
                top += " "
            else:
                top += body
                top += "|"

            if cell.south != None:
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

    binary_tree(cells)
    printMap(cells)
