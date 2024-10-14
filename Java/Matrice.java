import java.util.Iterator;
import java.util.Scanner;

public class Matrice implements Iterable<Iterator<Integer>>{

    //OVERVIEW: classe che modella una matrice di interi di dimensione `n * m`
    private int[][] matrice;

    

//CONSTRUCTORS
    public Matrice(int righe, int colonne) {
    //MODIFIES: this
    //EFFECTS: istanzia una matrice con righe e colonne
       this.matrice =new int[righe][colonne];
    }        
    
//METHODS
    public int getRighe() {
        return this.matrice.length;
    }

    public int getColonne() {
        return this.matrice[0].length;
    }
    

    public Iterator<Iterator<Integer>> iterator(){
    //MODIFIES: this.matrice
    //EFFECTS: metodo che ritorna un iteratore sull'intera matrice
         return new Iterator<Iterator<Integer>>(){
            private int righe = 0;
            @Override
            public boolean hasNext() {
                return this.righe < Matrice.this.getRighe();
            }

            @Override
            public Iterator<Integer> next() {

                Iterator<Integer> iteratorRighe= new  Iterator<Integer>(){
                    private int colonne = 0;
                    @Override
                    public boolean hasNext() {
                        return this.colonne < Matrice.this.getColonne();
                    }

                    @Override
                    public Integer next() {
                        int tmp = Matrice.this.matrice[righe-1][colonne];
                        colonne += 1;
                        return tmp;
                    }                    
                };
                this.righe += 1;
                return iteratorRighe;
            }  
         };
    }


//MAIN
    public static void main(String[] args) {
        int righe = Integer.parseInt(args[0]);
        int colonne = Integer.parseInt(args[1]);

        Matrice m = new Matrice(righe, colonne);

        Scanner s = new Scanner(System.in);

        for(int i = 0; i < righe; i++){
            System.out.println("Inserisci i " + colonne + " numeri (separati da spazio) della riga " + (i+1));
            String[] valori = s.nextLine().split(" ");

            for(int j = 0; j < valori.length; j++){
                m.matrice[i][j] = Integer.parseInt(valori[j]);
            }
        }

        System.out.println("----------------------------");

        Iterator<Iterator<Integer>> righeIterator = m.iterator();
        Iterator<Integer> iteratorColonne = null;
        //Iterator<Integer> colonneIterator = righeIterator.next();

        while(righeIterator.hasNext()){
            iteratorColonne = righeIterator.next();
            while(iteratorColonne.hasNext()){
                System.out.print(iteratorColonne.next() + " ");
            }
            System.out.println();
        }
    }
}

