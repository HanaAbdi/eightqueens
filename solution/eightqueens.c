#include <stdio.h>

int check(int index, int solution[8]){
    for (int i = 0; i < index; i++){
        if (solution[i] == solution[index] || solution[i] + (index - i) == solution[index] || solution[i] - (index - i) == solution[index] )
            return(0);
    }
    return(1);
}

void eightqueens(int index,int solution[8]){
    if (index == 8){
        for (int i = 0; i<8; i++)
            printf("%d",solution[i]);
        printf("\n");
    }
    else{
        for (int i = 1; i <= 8; i++){
            solution[index] = i;
            if (check(index, solution))
                eightqueens(index+1,solution);
        }
    }
}

int main(void){
    int solution[8] = {0};
    int index = 0;

    eightqueens(index, solution);
    return(0);
}