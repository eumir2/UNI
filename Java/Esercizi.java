import java.util.Scanner;
import java.lang.Integer;

/*# Minimo, massimo e valor medio (Array)

Scrivere un programma che riceva in input da **riga di comando** un intero `n`.
Il programma deve leggere da **standard input** una sequenza di `n` valori interi e deve stampare a video il valore minimo, massimo e medio tra i valori letti.

Oltre al metodo `main`, devono essere definiti ed utilizzati almeno i seguenti metodi:
* `public static int minimo(int[] in)` che restituisce il minimo valore intero presente in `in`.
* `public static int massimo(int[] in)` che restituisce il massimo valore intero presente in `in`.
* `public static int massimo(int[] in)` che restituisce un valore reale pari alla media aritmetica dei valori interi presenti in `in`.
 */

public class Esercizi{
    public static void main(String[] args) {
        int n = Integer.parseInt(args[0]); 
        //System.out.println(n);

        Scanner s = new Scanner(System.in);
        int[] num = new int[n] ;
            for(int i = 0; i < n; i++){
                num[i] = s.nextInt();
            }
        //System.out.println(num);
        System.out.println("Minimo = " + minimo(num));
        System.out.println("Massimo = " + massimo(num));
        System.out.println("Media = " + medio(num));
    }

    public static int minimo(int[] in){
        int min = Integer.MAX_VALUE;

        for(int i = 0; i < in.length; i++){
            if (in[i] < min){
                min = in[i];
            }
        }
        return min;
    }


    public static int massimo(int[] in){
        int max = 0;
        for(int i = 0; i < in.length; i++){
            if (in[i] > max){
                max = in[i];
            }
        }
        return max;
    }

    public static float medio(int[] in){
        float somma = 0;
        for (int j : in){
            somma += j;
        }
        return somma/in.length;
    }
}
