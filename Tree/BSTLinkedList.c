#include <stdio.h>
#include <stdlib.h>

struct BinaryTree {
    int data;
    struct BinaryTree* left;
    struct BinaryTree* right;
}*root=NULL, *tail=NULL, *newNode=NULL, *temp=NULL;

struct BinaryTree* insert(struct BinaryTree* root, int value){
    newNode = (struct BinaryTree*)malloc(sizeof(struct BinaryTree));
    newNode->data = value;
    newNode->left = NULL;
    newNode->right = NULL;

    if (root == NULL){
        root = newNode;
        return root;
    }

    if (value < root -> data){
        root->left = insert(root->left, value);
    } else {
        root->right = insert(root->right, value);
    }
    return root;
}

void display(struct BinaryTree* root){
    if (root == NULL){
        return;
    }
    display(root->left);
    printf("%d ", root->data);
    display(root->right);
}

void displayInOrder(struct BinaryTree* root){
    if (root == NULL){
        return;
    }
    displayInOrder(root->left);
    printf("%d ", root->data);
    displayInOrder(root->right);
}

void displayPreOrder(struct BinaryTree* root){
    if (root == NULL){
        return;
    }
    printf("%d ", root->data);
    displayPreOrder(root->left);
    displayPreOrder(root->right);
}

void displayPostOrder(struct BinaryTree* root){
    if (root == NULL){
        return;
    }
    displayPostOrder(root->left);
    displayPostOrder(root->right);
    printf("%d ", root->data);
}

void displayLevelOrder(struct BinaryTree* root){
    if (root == NULL){
        return;
    }
    struct BinaryTree* queue[100];
    int front = 0, rear = 0;

    queue[rear++] = root;

    while (front < rear){
        struct BinaryTree* current = queue[front++];
        printf("%d ", current->data);

        if (current->left != NULL){
            queue[rear++] = current->left;
        }
        if (current->right != NULL){
            queue[rear++] = current->right;
        }
    }
}

int main() {
    int n, value;

    printf("Enter the number of nodes: ");
    scanf("%d", &n);

    for (int i = 0; i < n; i++) {
        printf("Enter value for node %d: ", i + 1);
        scanf("%d", &value);
        root = insert(root, value);
    }

    printf("In-order traversal of the BST: ");
    displayInOrder(root);
    printf("\n");

    printf("Pre-order traversal of the BST: ");
    displayPreOrder(root);
    printf("\n");

    printf("Post-order traversal of the BST: ");
    displayPostOrder(root);
    printf("\n");
    
    // printf("Tree structure:\n");
    // displayLikeTree(root, 0);

    return 0;
}