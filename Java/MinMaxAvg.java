import java.util.ArrayList;
import java.util.Scanner;

/*Scrivere un programma che riceva in input da standard input una sequenza di lunghezza arbitraria di valori interi e deve stampare a video il valore minimo, massimo e medio tra i valori letti.

Oltre al metodo main, devono essere definiti ed utilizzati almeno i seguenti metodi:

public static int minimo(ArrayList<Integer> in) che restituisce il minimo valore intero presente in in.
public static int massimo(ArrayList<Integer> in) che restituisce il massimo valore intero presente in in.
public static float media(ArrayList<Integer> in) che restituisce un valore reale pari alla media aritmetica dei valori interi presenti in in. */

public class MinMaxAvg{
    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);

        ArrayList<Integer> valori = new ArrayList<Integer>();

        while(s.hasNextInt()){
            valori.add(s.nextInt());
        }
        s.close();

        System.out.println("Minimo = " + minimo(valori));
        System.out.println("Massimo = " + massimo(valori));
        System.out.println("Media = " + medio(valori));
    }  

  public static int minimo(ArrayList<Integer> in){
        int min = Integer.MAX_VALUE;
        for (Integer val : in) {
            if(val < min)   min = val;
        }

        return min;
    }

    public static int massimo(ArrayList<Integer> in){
        int max = 0;
        for (Integer val : in) {
            if(val > max)   max = val;
        }

        return max;
    }

    public static float medio(ArrayList<Integer> in){
        float somma = 0;
        for (Integer val : in) somma += val;

        return somma/in.size();
    }
}