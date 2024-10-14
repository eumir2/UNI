import java.util.ArrayList;

/*Scrivere un programma che legga da riga di comando un numero intero n e stampi a video la corrispondente tavola pitagorica n x n.

Oltre al metodo main(), devono essere definiti ed utilizzati almeno i seguenti metodi:

public static ArrayList<ArrayList<Integer> > creaTavolaPitagorica(int n) che restituisce una ArrayList di ArrayList in cui sono memorizzati i valori di una tavola pitagorica n x n;
public static void stampaTavolaPitagorica(ArrayList<ArrayList<Integer> > s) che stampa la tavola pitagorica corrispondete ai valori memorizzati s. */


public class TavolaPitagoricaArrayList {
    public static void main(String[] args) {
        int val = Integer.parseInt(args[0]);

        stampaTavolaPitagorica(creaTavolaPitagorica(val));
    }


    public static ArrayList<ArrayList<Integer>> creaTavolaPitagorica(int n){

        ArrayList<ArrayList<Integer>> tabella = new ArrayList<ArrayList<Integer>>();

        for(int i = 1; i <= n; i++){
            ArrayList<Integer> riga = new ArrayList<Integer>();
            for(int j = 1 ; j <= n; j++){
                riga.add(j*i);
                //System.out.println(j*i); 
            }

            tabella.add(riga);
            //System.out.println(colonne);
            //riga.clear();
            //System.out.println(tabella);
        }
        return tabella;
    }

    public static void stampaTavolaPitagorica(ArrayList<ArrayList<Integer>> s){
        //System.out.println(s);
        for (ArrayList<Integer> riga : s) {
            for(int val : riga){
                System.out.print(val + "\t");
            }
            System.out.println();
            System.out.println();
    
        }
    }
}
