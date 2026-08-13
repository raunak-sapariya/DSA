#include <stdio.h>
#include <stdlib.h>

struct Node {
    int data;
    struct Node* next;
    struct Node* prev;
};

void insertAtEnd(struct Node** head, struct Node** tail, int data) {
    struct Node* newNode = (struct Node*)malloc(sizeof(struct Node));
    newNode->data = data;
    newNode->next = NULL;
    newNode->prev = *tail;

    if (*head == NULL) {
        *head = newNode;
        *tail = newNode;
    } else {
        (*tail)->next = newNode;
        newNode->prev = *tail;
        *tail = newNode;
    }

}

void insertBegin(struct Node** head, struct Node** tail, int data){
    struct Node* newNode = (struct Node*)malloc(sizeof(struct Node));
    newNode->data = data;
    newNode->next = NULL;
    newNode->prev = *tail;

    if (*head == NULL){
        *head = newNode;
        *tail = newNode;
    } else {
        (*head)->prev = newNode;
        newNode->next = *head;
        *head = newNode;        
    }

}

void Display(struct Node* head) {
    struct Node* current = head;
    while (current != NULL) {
        printf("%d ", current->data);
        current = current->next;
    }
    printf("\n");
    
} 

int main() {
    struct Node* head = NULL;
    struct Node* tail = NULL;

    insertAtEnd(&head, &tail, 10);
    insertAtEnd(&head, &tail, 20);
    insertAtEnd(&head, &tail, 30);

    printf("Doubly Linked List after inserting at end: ");
    Display(head);

    insertBegin(&head, &tail, 5);
    printf("Doubly Linked List after inserting at beginning: ");
    Display(head);

    return 0;
}








