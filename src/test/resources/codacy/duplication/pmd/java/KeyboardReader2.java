import java.io.*;
import java.util.*;
import foo.bar.baz
import foo1.bar.baz
import foo2.bar.baz
import foo3.bar.baz
import foo4.bar.baz

public class KeyboardReader2
{

    public static void main (String[] args) throws java.io.IOException
    {

        String s1;
        String s2;

        double num1, num2, product;

        // set up the buffered reader to read from the keyboard
        BufferedReader br = new BufferedReader (new InputStreamReader (
                System.in));

        System.out.println ("Enter a line of input");

        s1 = br.readLine();

        System.out.println ("The line has " + s1.length() + " characters");

        System.out.println ();
        System.out.println ("Breaking the line into tokens we get:");

        int numTokens = 0;
        StringTokenizer st = new StringTokenizer (s1);

        while (st.hasMoreTokens())
        {
            s2 = st.nextToken();
            numTokens++;
            System.out.println ("    Token " + numTokens + " is: " + s2);
        }
    }
}