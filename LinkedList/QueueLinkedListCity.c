#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Node{
    char city[100];
    struct Node* next;
} *head = NULL, *tail = NULL, *newNode = NULL, *temp = NULL;

void enqueue(char city[]){
    newNode = (struct Node*)malloc(sizeof(struct Node));
    strcpy(newNode->city, city);
    newNode ->next = NULL;

    if (head == NULL){
        head = newNode;
        tail = newNode;
    }
    else{
        tail->next = newNode;
        tail = newNode;
    }
}

char *dequeue(){
    if (head == NULL)
    {
        printf("EMPTY");
        return NULL;
    }

    temp = head;
    head = head->next;

    if (head == NULL) {
        tail = NULL;
    }

    char *city = malloc(100);
    strcpy(city, temp->city);
    free(temp);
    return city;
}

int isValid(char city[]){
    if (city[0] != 'A' &&
        city[0] != 'E' &&
        city[0] != 'I' &&
        city[0] != 'O' &&
        city[0] != 'U') {

        return 0;
    }

    int length = strlen(city);
    if (length < 6) return 0;

    return 1;
}

void display(){
    if (head == NULL) {
        printf("The Queue linked list is empty.\n");
        return;
    }

    temp = head;
    while (temp != NULL) {
        printf("%s -> ", temp->city);
        temp = temp->next;
    }
    printf("NULL\n");
}

int main(){
    enqueue("Agra");
    enqueue("Mumbai");
    enqueue("Indore");
    enqueue("Oslo");
    enqueue("Ernakulam");

    printf("Original Queue:\n");
    display();

    int count = 0;
    temp = head;
    while (temp != NULL) {
       count ++;
       temp = temp->next;
    }

    printf("\nNumber of cities in the Queue: %d\n", count);
    for (int i = 0 ; i < count; i++) {
        char *city = dequeue();
        if (isValid(city)) {
            enqueue(city);
        }
        free(city);
    }

    printf("\nResultant Queue:\n");
    display();

}