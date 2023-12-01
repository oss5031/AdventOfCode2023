import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class Day1_2 {

    static final String[] numberString = { "one", "two", "three", "four", "five", "six", "seven", "eight", "nine" };
    static final char[] numberChar = { '1', '2', '3', '4', '5', '6', '7', '8', '9' };

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

        boolean useSpellingLetter = true;

        // Deal with spelling letters
        int digitPosition = Integer.MAX_VALUE;
        char firstDigitLetter = ' ';

        for (int i = 0; i < numberString.length; i++) {
            int index = line.indexOf(numberString[i]);
            if(index != -1 && index < digitPosition){
                digitPosition = index;
                firstDigitLetter = numberChar[i];
            }
        }

        //Deal with normal digits
        char[] chars = line.toCharArray();
        for (int i = 0; i < chars.length && i < digitPosition; i++) {
            if (isDigit(chars[i])){
                digitPosition = i;
                useSpellingLetter = false;
                break;
            }
        }

        return useSpellingLetter
                ? firstDigitLetter
                : chars[digitPosition];
    }

    private static char getLastDigit(String line) {

        boolean useSpellingLetter = true;

        // Deal with spelling letters
        int digitPosition = Integer.MIN_VALUE;
        char firstDigitLetter = ' ';

        for (int i = 0; i < numberString.length; i++) {
            int index = line.lastIndexOf(numberString[i]);
            if(index != -1 && index > digitPosition){
                digitPosition = index;
                firstDigitLetter = numberChar[i];
            }
        }

        //Deal with normal digits
        char[] chars = line.toCharArray();
        for(int i = chars.length-1; i >= 0 && i > digitPosition; i--) {
            if (isDigit(chars[i])){
                digitPosition = i;
                useSpellingLetter = false;
                break;
            }
        }

        return useSpellingLetter
                ? firstDigitLetter
                : chars[digitPosition];
    }

    private static boolean isDigit(char c) {
        return Character.isDigit(c);
    }
}
