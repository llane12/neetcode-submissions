class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        rows = []
        cols = []
        boxes = []
        for i in range(9):
            rows.append(set())
            cols.append(set())
            boxes.append(set())

        for r, row in enumerate(board):
            for c, cell in enumerate(row):
                if cell == ".":
                    continue
                b = ((r // 3) * 3) + (c // 3)
                if (cell in rows[r] or cell in cols[c] or cell in boxes[b]):
                    return False                
                rows[r].add(cell)
                cols[c].add(cell)
                boxes[b].add(cell)                
        return True