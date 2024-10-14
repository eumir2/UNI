import java.util.ArrayList;
import java.util.Scanner;

public class Percorso {
//OVERVIEW: classe che modella un percorso costituito da oggetti di tipo Punto
    public  ArrayList<Punto> percorso;

//CONSTRUCTORS
    public Percorso(ArrayList<Punto> p){
    //MODIFIES: this
    //EFFECTS: inizializza un percorso di punti vuoto
        this.percorso = p;
    }

//METHODS
    public void addPunto(Punto p){
    //MODIFIES: this
    //EFFECTS: aggiunge un punto al percorso in coda alla sequenza
        //System.out.println(p.toString());
        this.percorso.add(p);
    }

    public void removePunto(){
    //MODIFIES: this
    //EFFECTS: rimuove l'ultimo punto del percorso
        this.percorso.remove(percorso.size());
    }

    public double getDistanza(){
    //EFFECTS: calcola la lunghezza totale del percorso.
        double dTmp = 0.0;
        for (int i = 0; i < this.percorso.size() -1; i++) {
            dTmp += percorso.get(i).distanza(percorso.get(i+1));
        }

        return dTmp;
    }

//MAIN
    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);
        
        Punto tmp = new Punto();

        System.out.println("Inserisci le coordinate di un punto per riga nel formato <x y> (termina con CTRL+D)");
        ArrayList<Punto> punti = new ArrayList<Punto>();
        while(s.hasNext()){
            double x = s.nextDouble();
            double y = Double.parseDouble(s.next());

            tmp.setXY(x, y);
            //System.out.println(tmp.toString());

            punti.add(tmp);
        }
        Percorso p = new Percorso(punti);
        s.close();

        for (int i = 0; i < p.percorso.size()-1; i++) {
            System.out.println("Tratto " + (i+1) + ": distanza " + p.percorso.get(i).distanza(p.percorso.get(i+1)));
        }


        System.out.println("Totale: " + p.getDistanza());
    }

}
