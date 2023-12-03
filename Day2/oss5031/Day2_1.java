import java.io.File;
import java.io.FileNotFoundException;
import java.util.Dictionary;
import java.util.Hashtable;
import java.util.Scanner;

public class Day2_1 {

    static Dictionary<String, Integer> dict= new Hashtable<>();

    public static void main(String[] args) throws FileNotFoundException {
        //Bag: 12 red cubes, 13 green cubes, and 14 blue cubes
        dict.put("red", 12);
        dict.put("green", 13);
        dict.put("blue", 14);

        File file = new File(
                "/Users/oss5031/Documents/AdventoFCode/src/input_day2_1");
        Scanner sc = new Scanner(file);

        int sumOfGameIds = 0;
        while (sc.hasNextLine()){
            //System.out.println(sc.nextLine());
            String line = sc.nextLine();

            String[] splittedLine = line.split(":");

            String gameId = splittedLine[0].replaceAll("Game ","");
            String[] gameSets = splittedLine[1].split(";");

            if(isGameValid(gameSets)){
                sumOfGameIds += Integer.parseInt(gameId);
            }

        }

        System.out.println(sumOfGameIds);
    }

    private static boolean isGameValid(String[] gameSets) {

        for (String set: gameSets) {
            String[] colors = set.split(",");
            for (String color: colors) {
                int colorNumber = Integer.parseInt(color.replaceAll("[^0-9]", ""));
                String colorName = color.split(" ")[2];

                int bagValue = dict.get(colorName);
                if(bagValue < colorNumber){
                    return false;
                }
            }
        }

        return true;
    }


}
