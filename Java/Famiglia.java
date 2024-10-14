import java.util.ArrayList;
import java.util.InputMismatchException;
import java.util.Scanner;

public class Famiglia {
//OVERVIEW: classe che modella una famiglia con un certo numero di persone e un reddito complessivo
    private double reddito;
    private int dimensione;
    public static void main(String[] args) {
        double costoMedioVitto = Double.parseDouble(args[0]);
        int costoMedioAlloggio = Integer.parseInt(args[1]);

        ArrayList<Famiglia> famiglie = new ArrayList<Famiglia>();

        Scanner s = new Scanner(System.in);

        while(true){
            
            System.out.println("Inserisci il reddito e la dimensione di una famiglia (Ctrl+D per terminare la lettura)");
               //System.out.println(r);
                if(!s.hasNext()) break;
                //String r = "dsfdf";
                String r = s.next();
                int d = s.nextInt();
                double rr = Double.parseDouble(r);
                Famiglia f = null;

                try {
                    f = new Famiglia(rr, d);
                } catch (InputMismatchException e) {
                    System.out.println(e.getMessage());
                }

                try {
                    if(f.sottoSogliaPoverta(costoMedioVitto, costoMedioAlloggio)){
                        famiglie.add(f);
                    }
                }catch (InputMismatchException e) {
                    System.out.println(e.getMessage());
                }
                break;
            }    
                

        int i = 1;
        System.out.println(famiglie.size() + " famiglie sotto la soglia di povertà:");
        for (Famiglia f : famiglie) {
            System.out.print("Famiglia " + i + ": "+ f.toString());
            i++;
        }
    }

    //CONSTRUCTORS
    public Famiglia(double reddito, int dimensione) throws InputMismatchException{
    //EFFECTS: inizializza un nuovo oggetto di tipo `Famiglia` e lancia una `InputMismatchException` se `dimensione` è non-positiva
        if(dimensione <= 0){
            throw new InputMismatchException("Valore famiglia non valido");  
        }  
        else{
            this.reddito = reddito;
            this.dimensione = dimensione;
        }
    }

    public boolean sottoSogliaPoverta(double costoCasa, double costoCibo) throws InputMismatchException{
    /*EFFECTS: restituisce `true` se la somma tra `costoCasa` e `costoCibo` (moltiplicato per la `dimensione` della famiglia) è maggiore della metà del `reddito` della famiglia
                lancia una `InputMismatchException` se costoCasa o costoCibo non sono positivi;*/
        if(costoCasa <= 0 || costoCibo <= 0){
            throw new InputMismatchException("Valori del costo non compatibili");
        }else{
            double costoCiboTmp = costoCibo * this.dimensione;
            double tmp = this.reddito / 2;
            if((costoCasa + costoCiboTmp) > tmp) return true;
            else return false; 
        }    
    }

    public String toString(){
        return dimensione + "persone con reddito complessivo di" + reddito;
    }

}
