import java.util.*;

class StackUsingQueue {

    Queue<Integer> queue = new LinkedList<>();

    public void push(int x){
        queue.add(x);

        int size = queue.size();
        // move all element
        for (int i = 0; i < size-1; i++) {
            queue.add(queue.remove());
        }
    }

    public int pop(){
        if (queue.isEmpty()) {
            System.out.println("Stack is empty");
            return -1;
        }
        return queue.remove();
    }

    public int top(){
        if (queue.isEmpty()) {
            System.out.println("Stack is empty");
            return -1;
        }
        return queue.peek();
    }

    public boolean empty(){
        return queue.isEmpty();
    }    
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        StackUsingQueue stack = new StackUsingQueue();

        while (true) {
            System.out.println("1. Push");
            System.out.println("2. Pop");
            System.out.println("3. Top");
            System.out.println("4. Empty");
            System.out.println("5. Exit");
            System.out.print("Enter your choice: ");
            int choice = sc.nextInt();

            switch (choice) {
                case 1:
                    System.out.print("Enter element to push: ");
                    int element = sc.nextInt();
                    stack.push(element);
                    break;
                case 2:
                    int poppedElement = stack.pop();
                    if (poppedElement != -1) {
                        System.out.println("Popped element: " + poppedElement);
                    }
                    break;
                case 3:
                    int topElement = stack.top();
                    if (topElement != -1) {
                        System.out.println("Top element: " + topElement);
                    }
                    break;
                case 4:
                    boolean isEmpty = stack.empty();
                    System.out.println("Is stack empty? " + isEmpty);
                    break;
                case 5:
                    System.out.println("Exiting...");
                    sc.close();
                    return;
                default:
                    System.out.println("Invalid choice! Please try again.");
            }
        }
    }
}