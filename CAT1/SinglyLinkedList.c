#include <stdio.h>
#include <stdlib.h>

struct Node {
    int data;
    struct Node* next;
};


void insertBegin(struct Node** head, int data){
    struct Node* newNode = (struct Node*)malloc(sizeof(struct Node));
    newNode->data = data;
    newNode->next = *head;
    *head = newNode;
}

void insertEnd(struct Node** head, int data){
    struct Node* newNode = (struct Node*)malloc(sizeof(struct Node));
    newNode->data = data;
    newNode->next = NULL;

    if (*head == NULL){
        *head = newNode;
        return;
    }

    struct Node* temp = *head;
    while(temp->next != NULL){
        temp = temp->next;
    }
    temp->next = newNode;

}

void deleteBegin(struct Node** head){
    if (*head == NULL){
        printf("EMPTY\n");
        return;
    }

    struct Node* temp = *head;
    *head = (*head)->next;
    free(temp);
}

void DeleteEnd(struct Node** head){
    if (*head == NULL){
        printf("EMPTY\n");
        return;
    }

    // one node
    if ((*head)->next == NULL){
        free(*head);
        *head = NULL;
        return;
    }

    struct Node* temp = *head;
    while(temp->next->next != NULL){
        temp = temp->next;
    }
    free(temp->next);
    temp->next = NULL;


}

void deleteNode(struct Node** head, int value){
    if (*head == NULL){
        printf("EMPTY\n");
        return;
    }

    if ((*head)->data == value){
        struct Node* temp = *head;
        *head = (*head)->next;
        free(temp);
        return;
    }

    struct Node* temp = *head;
    while(temp->next !=NULL && temp->next->data != value){
        temp = temp->next;
    }

    if (temp->next == NULL){
        printf("Not found\n");
        return;
    }

     struct Node* delNode = temp->next;
     temp->next = temp->next->next,
     free(delNode);
     return;
}

void display(struct Node* head){
    struct Node* temp = head;

    while (temp != NULL) {
        printf("%d -> ", temp->data);
        temp = temp->next;
    }
    printf("NULL\n");
}

void reverse(struct Node** head){
    struct Node* prev = NULL;
    struct Node* curr = *head;
    struct Node* next = NULL;

    while(curr != NULL){
        next = curr->next;
        curr->next = prev;
        prev = curr;
        curr = next;
    }
    *head = prev;
}

int main(){
    struct Node* head = NULL;

    insertBegin(&head, 30);
    insertBegin(&head, 20);
    insertBegin(&head, 10);

    printf("After inserting at front:\n");
    display(head);

    insertEnd(&head, 100);
    insertEnd(&head, 200);
    insertEnd(&head, 300);
    printf("After inserting at end:\n");
    display(head);

    // deleteBegin(&head);
    // printf("After deleting at front:\n");
    // display(head);

    // DeleteEnd(&head);
    // printf("After deleting at end:\n");
    // display(head);

    // deleteNode(&head, 200);
    //  printf("After deleting node 40:\n");
    // display(head);

    reverse(&head);
    display(head);
    
}