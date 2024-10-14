import java.util.HashMap;
import java.util.Scanner;

public class Stanza {
    //OVERVIEW: classe che implementa una stanza con la propria capictà massima

    private int capienza;
    private int persone;


    public static void main(String[] args) {

        int personePresenti = 0;
        
        Stanza[] stanze = new Stanza[args.length];

        for (int i = 0; i < args.length; i++) {
            stanze[i] = new Stanza(Integer.parseInt(args[i]));    
        }

        System.out.println("Create " + stanze.length + " stanze");

        Scanner s = new Scanner(System.in);
        while(s.hasNext()){
            int nStanza = s.nextInt();

            if(nStanza > stanze.length){
                System.out.println("Stanza "+ nStanza + " non esiste");
                s.next();
            }
            else{
                String metodo = s.next();
                //System.out.println(metodo + "@@@@@@@@@@@@@@@@@@@");
                try{
                    if(metodo.equals("+"))
                        stanze[nStanza - 1].add();
                    else if(metodo.equals("-"))
                    stanze[nStanza - 1].remove();
                }catch(IllegalStateException e){
                    System.out.println(e.getMessage() + nStanza);
                }
            }
        }

        for (int i = 0; i < stanze.length; i++) {
            personePresenti += stanze[i].persone;
        }

        System.out.println("Numero totale di persone presenti: " + personePresenti);
    }


    //CONSTRUCTORS
    public Stanza(int capienza)throws IllegalArgumentException {
    //MODIFIES: this
    //EFFECTS: istanzia una nuovo oggetto Stanza con la sua capienza

        if(capienza <= 0)
            throw new IllegalArgumentException("La capienza non è valida");

        this.capienza = capienza;
        this.persone = 0;
    }

    //METHODS
    public void add()throws IllegalStateException{
    //MODIFIES: this.persone
    //EFFECTS: aggiunge una persona alla stanza, se la stanza è piena lancia IllegalStateException
        //System.out.println("metodo add richiamato");
        persone++;

        if(this.persone > this.capienza){
            persone = capienza;
            throw new IllegalStateException("Capienza massima raggiunta nella stanza ");
        }        
    }

    public void remove()throws IllegalStateException{
    //MODIFIES: this.persone
    //EFFECTS: toglie una persona alla stanza, se la stanza è vuota lancia IllegalStateException
   // System.out.println("metodo remove richiamato");
        persone--;

        if(this.persone < 0){
            persone = 0;
            throw new IllegalStateException("Nessuno nella stanza ");
        }
    }


    
}
