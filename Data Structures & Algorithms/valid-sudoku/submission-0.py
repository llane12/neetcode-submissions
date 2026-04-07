class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        rows = []
        cols = []
        for i in range(9):
            rows.append(set())
            cols.append(set())
        
        boxes = []
        for i in range(3):
            boxes.append([])
            for j in range(3):
                boxes[i].append(set())

        for rowIdx, row in enumerate(board):
            for colIdx, cell in enumerate(row):
                if cell == ".":
                    continue

                if cell in rows[rowIdx]:
                    return False
                rows[rowIdx].add(cell)

                if cell in cols[colIdx]:
                    return False
                cols[colIdx].add(cell)

                if cell in boxes[rowIdx // 3][colIdx // 3]:
                    return False
                boxes[rowIdx // 3][colIdx // 3].add(cell)
                
        return True