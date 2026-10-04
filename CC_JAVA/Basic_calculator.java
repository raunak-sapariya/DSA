import java.util.*;
public class Basic_calculator {
    static int BasicCalculator(String s)
    {
        int result=0;
        int sign=1;
        int number=0;
        Stack<Integer> st=new Stack<>();
        for(int i=0;i<s.length();i++)
        {
            char ch = s.charAt(i);
            if(Character.isDigit(ch))
            {
                number = number *10 + (ch-'0');
            }
            else if(ch=='+')
            {
                result+=sign*number;
                sign=1;
                number=0;
            }
            else if(ch=='-')
            {
                result+=sign*number;
                sign=-1;
                number=0;
            }
            else if(ch=='(')
            {
                st.push(result);
                st.push(sign);
                result=0;
                sign=1;
            }
            else if(ch==')')
            {
                result+=sign*number;
                result *=st.pop();
                result +=st.pop();
                number=0;
            }

        }
        result+=sign*number;
        return result;
    }
    public static void main(String[] args) {
        Scanner sc=new Scanner(System.in);
        String s=sc.next();
        System.out.println(BasicCalculator(s));
    }
}