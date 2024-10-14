import java.util.ArrayList;
import java.util.Collections;
import java.util.Scanner;

public class Visita {
//OVERVIEW: classe che modella la visita tra un Paziente e un Dottore
    private String nomeDottore;
    private String nomePaziente;

//CONSTRUCTOR
    public Visita(String d, String  p){
        this.nomeDottore = d;
        this.nomePaziente = p;
    }

//MAIN  
    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);
        
        ArrayList<Dottore> dottori = new ArrayList<>();
        ArrayList<Paziente> pazienti = new ArrayList<>();
        ArrayList<Visita> visite = new ArrayList<>();

        System.out.println("Inserisci medici nel formato `nome specializzazione parcella` (termina con CTRL+D)");
        while(s.hasNext()){
            String nome = s.next();
            String specializzazione = s.next();
            double parcella = s.nextDouble();
            dottori.add(new Dottore(nome, specializzazione, parcella));
        }
        //s.nextLine();

        System.out.println("Inserisci i pazienti nel formato `nome codice` (termina con CTRL+D)");
        while(s.hasNext()){
            String nome = s.next();
            String codice = s.next();
            pazienti.add(new Paziente(nome, codice));
        }

        System.out.println("Inserisci visite nel formato `nomeDottore codicePaziente` (termina con CTRL+D)");
        while(s.hasNext()){
            String nomeDottore = s.next();
            String nomePaziente = s.next();
            visite.add(new Visita(nomeDottore, nomePaziente));
        }

        System.out.println("I guadagni del mese sono:");
        for (Visita visita : visite) {
            int occurrences = Collections.frequency(visite, visita.nomeDottore);  
            double costoParcella = 0;
            for (Dottore d : dottori) {
                if(d.getNome().equals(visita.nomeDottore)){
                    costoParcella = d.getParcella();
                    break;
                }
            }  
            System.out.println(visita.nomeDottore + " " + costoParcella * occurrences);
        }

        
    }
    
}