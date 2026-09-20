import java.util.*;

public class QueueUsingStack {
    Stack<Integer> s1 = new Stack<>();
    Stack<Integer> s2 = new Stack<>();

    void enqueue(int x){
        s1.add(x);
    }

    int dequeue(){
        if (s1.isEmpty() && s2.isEmpty()){
            System.out.println("Queue is empty");
            return -1;
        }

        if (s2.isEmpty()){
            while (!s1.isEmpty()) {
                s2.add(s1.pop());
            }
        }

        return s2.pop();
    }

    int peek(){
        if (s1.isEmpty() && s2.isEmpty()){
            System.out.println("Queue is empty");
            return -1;
        }

        if (s2.isEmpty()){
            while (!s1.isEmpty()) {
                s2.add(s1.pop());
            }
        }

        return s2.peek();
    }

    boolean isEmpty(){
        return s1.isEmpty() && s2.isEmpty();
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.println("Queue using Stack");
        QueueUsingStack queue = new QueueUsingStack();

        while (true) {
            System.out.println("\n1. Enqueue");
            System.out.println("2. Dequeue");
            System.out.println("3. Peek");
            System.out.println("4. Empty");
            System.out.println("5. Exit");
            System.out.print("Enter your choice: ");
            int choice = sc.nextInt();

            switch (choice) {
                case 1:
                    System.out.print("Enter element: ");
                    int x = sc.nextInt();
                    queue.enqueue(x);
                    System.out.print(x+" inserted");
                    break;
                
                case 2:
                    int removed = queue.dequeue();

                    if (removed != -1){
                        System.out.print(removed+"deleted");
                    }
                    break;
                case 3:
                    int front = queue.peek();

                    if (front != -1){
                        System.out.println("front "+front);
                    }
                    break;

                case 4:
                    if (queue.isEmpty()) {
                        System.out.println("Empty");
                    } else {
                        System.out.println("Not empty");
                    }
                    break;
                
                case 5:
                    System.out.println("Exiting...");
                    sc.close();
                    return;
            
                default:
                    System.out.println("Invalid");
                    break;
            }
        }
    }
}
