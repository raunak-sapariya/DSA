
import java.util.*;

// public class minStack {
//     class pair {
//         int first;
//         int second;

//         pair(int first, int second) {
//             this.first = first;
//             this.second = second;
//         }
//     }
//     Stack<pair> stack = new Stack<>();

//     void push(int x) {
//         if (stack.isEmpty()) {
//             stack.push(new pair(x, x));
//         } else {
//             stack.push(new pair(x, Math.min(x, stack.peek().second)));
//         }
//     }

//     void pop() {
//         if (!stack.isEmpty()) {
//             stack.pop();
//         }
//     }

//     int top() {
//         if (!stack.isEmpty()) {
//             return stack.peek().first;
//         }
//         return -1;
//     }

//     int getMin() {
//         if (!stack.isEmpty()) {
//             return stack.peek().second;
//         }
//         return -1; 
//     }

//     public static void main(String[] args) {
//         minStack minStack = new minStack();
//         minStack.push(-2);
//         minStack.push(0);
//         minStack.push(-3);
//         System.out.println(minStack.getMin());
//         minStack.pop();
//         System.out.println(minStack.top());
//         System.out.println(minStack.getMin());
//     }
// }


class minStack {
    // if it modify min 2 * value - prevMin & prevMin = 2 * value - newvalue
    int min = Integer.MAX_VALUE;
    Stack<Integer> stack = new Stack<>();

    void push(int x) {
        if (stack.isEmpty()) {
            stack.push(x);
            min = x;
        } else {
            if (x < min) {
                stack.push(2 * x - min);
                min = x;
            } else {
                stack.push(x);
            }
        }
    }

    int pop() {
        if (stack.isEmpty()) {
            System.out.println("Stack is empty");
            return -1;
        }

        int top = stack.pop();
        if (top < min) {
            int prevMin = min;
            min = 2 * min - top;
            return prevMin;
        } else {
            return top;
        }
    }

    int top() {
        if (stack.isEmpty()) {
            System.out.println("Stack is empty");
            return -1;
        }

        int top = stack.peek();
        if (top < min) {
            return min;
        } else {
            return top;
        }
    }

    int getMin() {
        if (stack.isEmpty()) {
            System.out.println("Stack is empty");
            return -1;
        }
        return min;
    }

    public static void main(String[] args) {
        minStack minStack = new minStack();
        System.out.println("Min Stack!");
        Scanner sc = new Scanner(System.in);
        while (true) {
            System.out.println("1. Push");
            System.out.println("2. Pop");
            System.out.println("3. Top");
            System.out.println("4. Get Min");
            System.out.println("5. Exit");
            System.out.print("Enter your choice: ");
            int choice = sc.nextInt();

            switch (choice) {
                case 1:
                    System.out.print("Enter element to push: ");
                    int element = sc.nextInt();
                    minStack.push(element);
                    break;
                case 2:
                    int poppedElement = minStack.pop();
                    if (poppedElement != -1) {
                        System.out.println("Popped element: " + poppedElement);
                    }
                    break;
                case 3:
                    int topElement = minStack.top();
                    if (topElement != -1) {
                        System.out.println("Top element: " + topElement);
                    }
                    break;
                case 4:
                    int minElement = minStack.getMin();
                    if (minElement != -1) {
                        System.out.println("Minimum element: " + minElement);
                    }
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

