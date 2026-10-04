import java.util.*;
public class validateStackSequences {
    static boolean isValidate(int[] pushed,int[] popped){
        Stack<Integer> st = new Stack<>();
        int j = 0;
        for (int x : pushed) {
            st.add(x);
            while (!st.isEmpty() && j < popped.length && st.peek()==popped[j]) {
                st.pop();
                j++;
            }
        }
        return st.isEmpty();
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.print("Enter number of element: ");
        int n = sc.nextInt();

        int pushed[]=new int[n];
        System.out.println("\nEnter pushed element");
        for(int i=0;i<n;i++){
            System.out.print("Enter pushed element " + (i + 1)+": ");
            pushed[i]=sc.nextInt();
        }

        int popped[]=new int[n];
        System.out.println("\nEnter poped element");
        for(int i=0;i<n;i++){
            System.out.print("Enter poped element " + (i + 1)+": ");
            popped[i]=sc.nextInt();
        }

        System.out.println(isValidate(pushed, popped));

    }

}
