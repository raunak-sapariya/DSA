#include <stdio.h>
#include <stdlib.h>

struct Node {
    char data;
    struct Node* next;
} *head = NULL, *tail = NULL, *newNode = NULL, *temp = NULL;

void insertAtEnd(char data) {
    newNode = (struct Node*)malloc(sizeof(struct Node));
    newNode->data = data;
    newNode->next = NULL;

    if (head == NULL) {
        head = newNode;
        tail = newNode;
        newNode->next = head;
    } else {
        tail->next = newNode;
        tail = newNode;
        tail->next = head;
    }
}

void deleteAtEnd() {
    if (head == NULL) {
        printf("The circular linked list is empty. Cannot delete.\n");
        return;
    }

    if (head == tail) {
        free(head);
        head = NULL;
        tail = NULL;
    } else {
        temp = head;
        while (temp->next != tail) {
            temp = temp->next;
        }
        free(tail);
        tail = temp;
        tail->next = head;
    }
}

void display() {
    if (head == NULL) {
        printf("The circular linked list is empty.\n");
        return;
    }

    temp = head;
    printf("\nCircular Linked List: ");
    while (1) {
        printf("%c ", temp->data);
        temp = temp->next;
        if (temp == head) {
            break;
        }
    }
}

int isPalindrome() {
    char str[100];
    int i = 0;

    if (head == NULL) {
        return 1;
    }

    temp = head;
    while (1) {
        str[i++] = temp->data;
        temp = temp->next;
        if (temp == head) {
            break;
        }
    }
    

    int j = 0;
    int k = i - 1;
    while (j < k) {
        if (str[j] != str[k]) {
            return 0;
        }
        j++;
        k--;
    }

    return 1;
}

int main() {
    insertAtEnd('a');
    insertAtEnd('b');
    insertAtEnd('c');
    insertAtEnd('b');
    insertAtEnd('a');
    display();

    deleteAtEnd();
    display();

    if (isPalindrome()) {
        printf("\nThe circular linked list is a palindrome.\n");
    } else {
        printf("\nThe circular linked list is not a palindrome.\n");
    }

    return 0;
}