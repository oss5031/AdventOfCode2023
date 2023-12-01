import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class Day1_1 {

    public static void main(String[] args) throws FileNotFoundException {

        // pass the path to the file as a parameter
        File file = new File(
                "/Users/oss5031/Documents/AdventoFCode/src/input_day1_1");
        Scanner sc = new Scanner(file);

        int sum = 0;
        while (sc.hasNextLine()){
            //System.out.println(sc.nextLine());
            String line = sc.nextLine();
            char firstDigit = getFirstDigit(line);
            char lastDigit = getLastDigit(line);
            int digitsConcat = Integer.parseInt(firstDigit + "" + lastDigit);

            sum += digitsConcat;
        }

        System.out.println(sum);
    }
    
    private static char getFirstDigit(String line) {
        for (char c : line.toCharArray()) {
            if (isDigit(c)){
                return c;
            }
        }
        return ' ';
    }

    private static char getLastDigit(String line) {
        char[] chars = line.toCharArray();
        for(int i = chars.length-1; i>=0; i--)
            if (isDigit(chars[i])){
                return chars[i];
            }

        return ' ';
    }

    private static boolean isDigit(char c) {
        return Character.isDigit(c);
    }
}
