import java.util.*;

class Node {
    int data;
    Node next;

    Node(int data){
        this.data = data;
        this.next = null;
    }
}

public class partitionlist {
    static Node head = null;

    static void insertAtEnd(int data){
        Node newNode = new Node(data);

        if (head == null){
            head = newNode;
            return;
        }

        Node temp = head;
        while (temp.next != null) {
            temp = temp.next;
        }

        temp.next = newNode;
    }

    static void display(){
        if (head == null){
            System.out.println("Empty");
            return;
        }

        Node temp = head;
        while (temp != null) {
            System.out.println(temp.data);
            temp = temp.next;
        }
        System.out.println();
    }

    static Node PartitionList(int x){
        Node less = new Node(0);
        Node lessTail = less;

        Node greater  = new Node(0);
        Node greaterTail = greater;

        Node curr = head;
        while (curr != null) {
            if(curr.data <x){
                lessTail.next = curr;
                lessTail = lessTail.next;
            } else{
                greaterTail.next = curr;
                greaterTail = greaterTail.next; 
            }

            curr = curr.next;
        }

        lessTail.next = greater.next;
        greaterTail.next = null;
        return less.next;
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        System.out.println("Enter element (-1 to stop)");
        while (true) {
            int val = sc.nextInt();
            if (val == -1){
                break;
            }

            insertAtEnd(val);
        }

        System.out.println("Original");
        display();

        System.out.println("Enter x");
        int x = sc.nextInt();
        head = PartitionList(x);
        System.out.println("After Partitioning the Linked List");
        display();
    }
}
