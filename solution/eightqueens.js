const check = (index, solution)=>{
    for (let i = 0; i< index; i++){
        if (solution[i] == solution[index] || solution[i] + (index - i) == solution[index] || solution[i] - (index - i) == solution[index] ){
            return false
        }
    }
    return true
}

const eightqueens = (index, solution)=>{
    if (index == 8){
        console.log(solution.join(''))
    } else {
        for (let i = 1; i < 9; i++){
            solution[index] = i
            if (check(index,solution)){
                eightqueens(index+1, solution)
            }
        }
        eightqueens[index] = 0
    }
}

let solution = new Array(8).fill(0)
let index = 0
eightqueens(index, solution)