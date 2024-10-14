/*Scrivere un programma che legga da riga di comando un numero intero n e stampi a video la corrispondente tavola pitagorica n x n.

Oltre al metodo main(), devono essere definiti ed utilizzati almeno i seguenti metodi:

public static int[][] creaTavolaPitagorica(int n) che restituisce una matrice in cui sono memorizzati i valori di una tavola pitagorica n x n;
public static void stampaTavolaPitagorica(int[][] s) che stampa la tavola pitagorica corrispondete ai valori memorizzati s. */

public class TavolaPitagorica{
    public static void main(String[] args) {
        
        int n = Integer.parseInt(args[0]);

        stampaTavolaPitagorica(creaTavolaPitagorica(n));     

    }

    public static int[][] creaTavolaPitagorica(int n){
        int[][] matrice = new int[n][n];

        for(int i = 1; i <= n; i++){
            for(int j = 1 ; j <= n; j++){
                matrice[i-1][j-1] = j*i; 
            }
        }
        return matrice;
    }

    public static void stampaTavolaPitagorica(int[][] s){
        for(int i = 0; i < s.length; i++){
            for(int j = 0 ; j < s.length; j++){
                System.out.print(s[i][j]+ "\t");
            }
            System.out.println();
            System.out.println();
        }
    }
}