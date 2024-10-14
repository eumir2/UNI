/*Scrivere un programma che legga da **riga di comando** un numero decimale oppure un numero romano. Se il numero letto è decimale convertirlo in un numero romano e viceversa.

Progettare, specificare e implementare delle procedure appropriate a tale scopo.

Di quali procedure dovrei avere bisogno? Di quali vincoli dovrò tenere conto nella specifica di queste procedure? Posso rendere le procedure totali risolvendo questi vincoli? Come?
 */
package NumeriRomani;

import java.util.HashMap;
import java.util.Map;

import javax.print.attribute.HashDocAttributeSet;

public class NumRom{

    public static void main(String[] args) {
        int num = romanToDec(args[0]);
       // System.out.println(num);
        if(num != 0){
            System.out.println(num);
        }else{
            System.out.println(decToRoman(Integer.parseInt(args[0])));
        }
    }

    public static int romanToDec(String roman){
    //EFFECTS:  return an int whitch is the roman string converted to decimal
    //          if not valid roman, returns 0
    Map<String, Integer> simboliValori= new HashMap<String, Integer>();
    simboliValori.put("I", 1);
    simboliValori.put("IV", 4);
    simboliValori.put("V", 5);
    simboliValori.put("IX", 9);
    simboliValori.put("X", 10);
    simboliValori.put("XL", 40);
    simboliValori.put("L", 50);
    simboliValori.put("XC", 90);
    simboliValori.put("C", 100);
    simboliValori.put("CD", 400);
    simboliValori.put("D", 500);
    simboliValori.put("CM", 900);
    simboliValori.put("M", 1000);


    int result = 0;
    int i = 0;
    while(i < roman.length()){
        if(i < roman.length()-1 && simboliValori.get(roman.substring(i, i + 2)) != null){
            result += simboliValori.get(roman.substring(i, i + 2));
            i += 2;
        }
        else if(simboliValori.get(roman.substring(i, i + 1)) != null){
            result += simboliValori.get(roman.substring(i, i + 1));
            i += 1;
        }else{
            return 0;
        }
    }
    return result;
    } 
    
    public static String decToRoman(int dec){
    //REQUIRES: dec has to be a valid decimal number (above 0 and lower than 3999)
    //EFFECTS:  return the roman version of the decimal    
        //I = 1  V = 5  x = 10 L = 50  C = 100  D = 500 M = 1000

    int[] valori = {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1};
    String[] simboli = {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"};
    String result = ""; 
    int current = dec;

    for (int i = 0; i < simboli.length; i++) {
        while(current-valori[i] >= 0){
            current -= valori[i];
            result += simboli[i];
        }
    }
    return result;
    }

}

