import java.util.*;

public class evaluateRPN {
    static int RPN(String[] tokens){
        Stack<Integer> st = new Stack<>();
        for (String tocken : tokens) {
            if(tocken.equals("+")) st.push(st.pop() + st.pop());
            else if(tocken.equals("-")) st.push(st.pop() - st.pop());
            else if(tocken.equals("*")) st.push((int)st.pop() * (int)st.pop());
            else if(tocken.equals("/"))st.push((int) (int)st.pop() / (int)st.pop());
            else st.push(Integer.parseInt(tocken));
        }
        return st.pop();
    }

    public static void main(String[] args) {
        Scanner sc=new Scanner(System.in);
        System.out.print("Enter number of tokens ");
        int n = sc.nextInt();
        String tokens[]=new String[n];
        for(int i=0;i<n;i++){
            System.out.print("Enter token " + (i + 1)+": ");
            tokens[i]=sc.next();
        }
        System.out.println(RPN(tokens));
    }
}