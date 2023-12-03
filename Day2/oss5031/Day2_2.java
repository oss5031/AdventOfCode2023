import java.io.File;
import java.io.FileNotFoundException;
import java.util.*;

public class Day2_2 {

    static Dictionary<String, Integer> dict= new Hashtable<>();

    public static void main(String[] args) throws FileNotFoundException {
        initDict();

        File file = new File(
                "/Users/oss5031/Documents/AdventoFCode/src/input_day2_1");
        Scanner sc = new Scanner(file);

        int power = 0;
        while (sc.hasNextLine()){
            //System.out.println(sc.nextLine());
            String line = sc.nextLine();

            String[] splittedLine = line.split(":");

            String gameId = splittedLine[0].replaceAll("Game ","");
            String[] gameSets = splittedLine[1].split(";");

            power += getSetPower(gameSets);
            initDict();
        }

        System.out.println(power);
    }

    private static void initDict() {
        dict.put("red", 0);
        dict.put("green", 0);
        dict.put("blue", 0);
    }

    private static int getSetPower(String[] gameSets) {

        for (String set: gameSets) {
            String[] colors = set.split(",");
            for (String color: colors) {
                int colorNumber = Integer.parseInt(color.replaceAll("[^0-9]", ""));
                String colorName = color.split(" ")[2];

                int bagValue = dict.get(colorName);
                if(bagValue < colorNumber){
                    dict.put(colorName, colorNumber);
                }
            }
        }

        return calculatePower();
    }

    private static int calculatePower() {

        int fullPower = 1;

        for (int power : Collections.list(dict.elements())) {
            fullPower *= power;
        }

        return fullPower;
    }
}
