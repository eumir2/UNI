import java.util.ArrayList;
import java.util.Scanner;

/*$ java Armageddon 0.02
67620000 74798
10140000 38014
50200000 103360
68300000 85857
26990000 34557
30310000 95752

Asteroide pericoloso a distanza 34557km del peso di 26990000kg
Lanciate Bruce Willis! */



public class Asteroide{
//OVERVIEW: classe che definisce l0oggetto asteroide con la sua massa e la sua distanza dalla terra

    private double massa;
    private double distanza;

    public static void main(String[] args) {
        double soglia = Double.parseDouble(args[0]);
        
        Scanner s = new Scanner(System.in);
        ArrayList<Asteroide> corpiCelesti = new ArrayList<Asteroide>();

        while(s.hasNext()){
            //if(s.nextLine().equals("=")) break;
            double tmpM = s.nextDouble();
            double tmpD = s.nextDouble();

            corpiCelesti.add(new Asteroide(tmpM, tmpD));

        }

        for (Asteroide a : corpiCelesti) {
            if(a.calcolaAttrazione() > soglia){
                System.out.println(a.toString());
            }
        }
        System.out.println("Lanciate Bruce Willis!");
    }




    //CONSTRUCTORS
    public Asteroide(){
    //MODIFIES: this
    //EFFECTS: istanzia un oggetto asteroide vuoto
        this.massa = 0.0;
        this.distanza = 0.0;
    }


    public Asteroide(double m, double d){
    //REQUIRES: massa e distanza devono essere due numeri positivi
    //MODIFIES: this
    //EFFECTS: istanzia un oggetto asteroide con massa e dimensione presi come parametro
        this.massa = m;
        this.distanza = d;
    }

    public double getMassa(){
        return massa;
    }

    public double getDistanza(){
        return distanza;
    }

    public void setMassa(double m){
        massa = m;
    }

    public void setDistanza(double d){
        distanza = d;
    }

    public String toString(){
        return "Asteroide pericoloso a distanza " + distanza + "km del peso di " + massa + "kg";
    }

    public double calcolaAttrazione(){
    //REQUIRES: l'oggetto asteroide deve essere istanziato correttamente e non vuoto
    //EFFECTS: calcola la forza gravitazionale esercitata tra l'asteroide e la terra secondo la formula matematica: forza = massa/distanza^2
        return massa / (distanza * distanza);
    }

}