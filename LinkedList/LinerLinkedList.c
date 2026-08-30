#include <stdio.h>
#include <stdlib.h>

struct Node {
    int coeff;
    int pow;
    struct Node* next;
} *head = NULL, *tail = NULL, *newNode = NULL, *temp = NULL;

struct Node* creatNode(int coeff, int pow){
    newNode = (struct Node*)malloc(sizeof(struct Node));
    newNode->coeff = coeff;
    newNode->pow = pow;
    newNode->next = NULL;
    return newNode;
}

void insertAtEnd(struct Node** head, struct Node** tail, int coeff, int pow){
    newNode = creatNode(coeff, pow);

    if (*head == NULL){
        *head = newNode;
        *tail = newNode;
        return;
    }

    (*tail)->next = newNode;
    *tail = newNode;
}

void display(struct Node** head){
    if (*head == NULL){
        printf("EMPTY\n");
        return;
    }

    struct Node* temp = *head;
    while(temp != NULL){
        printf("%dx^%d ", temp->coeff, temp->pow);
        if (temp->next != NULL){
            printf("+ ");
        }
        temp = temp->next;
    }
}

int main(){
    insertAtEnd(&head, &tail, 5, 2);
    insertAtEnd(&head, &tail, 3, 1);
    insertAtEnd(&head, &tail, 2, 0);

    display(&head);
    return 0;
}

