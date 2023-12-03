import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class Day3_1 {
    public static void main(String[] args) throws FileNotFoundException {

        File file = new File(
                "/Users/oss5031/Documents/AdventoFCode/src/input_day3_1");

        char[][] schema = readSchema(file);

        int sum = processSchema(schema);

        System.out.println();

    }

    private static int processSchema(char[][] schema) {
        int sum = 0;
        int prevNumber = 0;

        for (int i = 0; i < schema.length; i++) {
            for (int j = 0; j < schema[i].length; j++) {
                char currChar = schema[i][j];
                if(isSymbol(currChar)){
                    sum += getNumbers(schema, i, j);
                }
            }
        }

        return sum;
    }

    private static int getNumbers(char[][] schema, int i, int j) {
        int sum = 0;

        //TODO - rip... muitos edge cases (no time available)

        return sum;
    }

    private static boolean isSymbol(char currChar) {

        boolean isSymbol = !Character.toString(currChar).matches("[0-9]|\\.");
        return isSymbol;
    }

    private static char[][] readSchema(File file) throws FileNotFoundException {

        Scanner sc = new Scanner(file);

        // Get bidimensional array dimensional
        String line = null;
        int lines = 0;
        while (sc.hasNextLine()){
            lines++;
            line = sc.nextLine();
        }

        int cols = line.toCharArray().length;

        sc = new Scanner(file);
        char[][] schema = new char[lines][cols];

        // Fill the schema
        int l  = 0, c = 0;
        while (sc.hasNextLine()){
            char[] currLine = sc.nextLine().toCharArray();
            for (char currChar : currLine) {
                schema[l][c] = currChar;
                c++;
            }
            l++;
            c = 0;
        }

        return schema;
    }
}
