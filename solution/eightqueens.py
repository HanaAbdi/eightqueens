def check(index, solution):
    for i in range(0,index):
        if solution[i] == solution[index] or solution[i] + (index - i) == solution[index] or solution[i] - (index - i) == solution[index]:
            return False
    return True

def eightqueens(index, solution):
    if index == 8:
        for n in solution:
            print(n,end = '')
        print()
    else:
        for i in range(1,9):
            solution[index] = i
            if check(index, solution):
                eightqueens(index+1,solution)
        solution[index]= 0

if __name__ == '__main__':
    solution = [0]*8
    index = 0
    eightqueens(index, solution)

