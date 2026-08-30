#include <stdio.h>
#include <stdlib.h>

struct Node {
    int matrix[3][3];
    struct Node* next;
    struct Node* prev;
}*head=NULL, *tail=NULL, *newNode=NULL, *temp=NULL;

void insertAtEnd(int matrix[3][3]){
    newNode = (struct Node*)malloc(sizeof(struct Node));
    for (int i=0;i <3;i++){
        for (int j=0;j <3;j++){
            newNode->matrix[i][j] = matrix[i][j];
        }
    }
    if (head == NULL){
        head = newNode;
        tail = newNode;
        newNode->prev = NULL;
    } else {
        temp = tail;
        temp->next = newNode;
        tail = newNode;
        tail->prev = temp;
    }
}

void display(){
    if (head == NULL){
        printf("empty");
    }

    temp = head;
    while (temp != NULL){
        for(int i = 0; i < 3; i++){
            for(int j =0; j < 3; j++){
                printf("%d ", temp->matrix[i][j]);
            }
        printf("\n");
        }
        temp = temp->next;
        printf("\n");
    }
}

void search(int data){
    if (head == NULL){
        printf("empty");
    }

    temp = head;
    while (temp != NULL){
        for(int i = 0; i < 3; i++){
            for(int j =0; j < 3; j++){
                if (data == temp->matrix[i][j]){
                    printf("Found at [%d][%d]",i,j);
                }
            }
        }
        temp = temp->next;
    }

}

int main(){
    int arr1[3][3] = {{1,2,3},{4,5,6},{7,8,9}};
    int arr2[3][3] = {{10,20,30},{40,50,60},{70,80,90}};
    insertAtEnd(arr1);
    insertAtEnd(arr2);
    display();
    int target = 3;
    search(target);
    printf("END");
}
