import java.util.Scanner;

class Node {
    int data;
    Node next;

    Node(int data){
        this.data = data;
        this.next = null;
    }
}

public class addTwoNumbers {
    static Node insertAtEnd(Node head, int data){
        Node newNode = new Node(data);

        if (head == null) return newNode;

        Node temp = head;
        while (temp.next != null) {
            temp = temp.next;
        }
        temp.next = newNode;
        return head;
    }

    static void display(Node head){
        if (head == null){
            System.out.println("Empty");
            return;
        }

        Node temp = head;
        while (temp != null) {
            System.out.print(temp.data + " ");
            temp = temp.next;
        }
        System.out.println();
    }

    static Node adding(Node l1, Node l2){
        Node dummyNode = new Node(0);
        Node curr = dummyNode;
        int carry = 0;

        while (l1 != null || l2 != null || carry != 0)  {
            int val1 = (l1 != null) ? l1.data : 0;
            int val2 = (l2 != null) ? l2.data : 0;

            int sum = val1 + val2 + carry;
            carry = sum / 10;
            int digit = sum % 10;

            curr.next = new Node(digit);
            curr = curr.next;

            if(l1 != null) l1 = l1.next;
            if(l2 != null) l2 = l2.next;
        }

        return dummyNode.next;
    }
   
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        Node l1 = null;
        Node l2 = null;

        System.out.println("Enter first linked list elements (-1 to stop):");
        while(true)
        {
            int value = sc.nextInt();
            if(value == -1) break;
            l1 = insertAtEnd(l1, value);
        }

        System.out.println("Enter second linked list elements (-1 to stop):");
        while(true)
        {
            int value = sc.nextInt();
            if(value == -1) break;
            l2 = insertAtEnd(l2, value);
        }

        System.out.println("First Linked List:");
        display(l1);

        System.out.println("Second Linked List:");
        display(l2);

        Node result = adding(l1, l2);

        System.out.println("Result:");
        display(result);

    }
}
