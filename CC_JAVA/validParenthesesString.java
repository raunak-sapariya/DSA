import java.util.Scanner;
import java.util.Stack;

public class validParenthesesString {
    static boolean isValid(String s){
        Stack<Integer> open = new Stack<>();
        Stack<Integer> star = new Stack<>();

        for (int i = 0; i < s.length(); i++) {
            Character ch = s.charAt(i);
            if(ch=='(') open.push(i);
            else if(ch=='*')star.push(i);
            else{
                if (!open.isEmpty()) open.pop();
                else if (!star.isEmpty()) star.pop();
                else return false;
            }
        }

        while (!open.isEmpty() && !star.isEmpty()) {
            if (open.peek() < star.peek()){
                open.pop();
                star.pop();
            } else return false;
        }
    return open.isEmpty();
    }
    public static void main(String[] args) {
        Scanner sc=new Scanner(System.in);
        String s = sc.next();
        System.out.println(isValid(s));
    }
}
