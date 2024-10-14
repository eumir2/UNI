import java.util.ArrayList;
import java.util.Date;
import java.util.Scanner;


public class Calendario {
//OVERVIEW: classe che gestisce un calendario di oggetti di tipo Evento
    private ArrayList<Evento> eventi;

    public static void main(String[] args) {
        Calendario c = new Calendario();
        Scanner s = new Scanner(System.in);
        System.out.println("Inserisci +/-/* NomeEvento gg/mm/aaaa [offset] (termina con CTRL+D)");
       
        while(s.hasNext()){
            String[] parametri = s.nextLine().split(" ");
            //System.out.println(parametri[0] + parametri[1] + parametri[2] + "°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°"); 
            String[] data = parametri[2].split("/");
                    int g = Integer.parseInt(data[0]);
                    int m = Integer.parseInt(data[1]); 
                    int a = Integer.parseInt(data[2]);
            Date d = new Date(a,m,g);
            Evento tmp = new Evento(d,parametri[1]);
            switch(parametri[0]){
                case "+" :
                    c.addEvento(tmp); 
                    break;
                case "-" :
                    c.removeEvento(d,tmp.getNome());
                }
            if(parametri.length == 4)
                    c.cloneEvento(new Evento(d, parametri[1]), Integer.parseInt(parametri[3]));

           parametri = null;
        }
        s.close();
        
        System.out.println(c.eventi.toString());
    }


//CONSTRUCTORS
    public Calendario(){
    //MODIFIES: this
    //EFFECTS: inizializza la classe calendario a un caledario vuoto
        this.eventi = new ArrayList<Evento>();
    }

    public Calendario(ArrayList<Evento> eventi){
    //MODIFIES: this
    //EFFECTS: inizializza la classe calendario contenente gli eventi passati come parametro
         this.eventi = new ArrayList<Evento>();
        for (Evento e : eventi) {
            this.eventi.add(e);
        }
    }

//METHODS
    public void addEvento(Evento e){
    //MODIFIES: this
    //EFFECTS: metodo che aggiunge un evento al calendario, se gia presente stampa messaggio di errore.        boolean flag = true;
        boolean flag = true;
        for (Evento evento : this.eventi) {
            if(evento.equals(e)){
                System.out.println("Evento già presente");
                flag = false;
                break;
            }
        }
        if(flag){
            this.eventi.add(e);
            //System.out.println("aggiungendo elemento");
        }
            
    }

    public void removeEvento(Date data, String nome){
    //MODIFIES: this
    //EFFECTS: rimuove un evento dal calendario dati il suo nome e la sua data 
        boolean flag = true;
        int index = 0;
        Evento tmp = new Evento(data,nome);
        for (Evento evento : this.eventi) {
            if(evento.equals(tmp)){
                index = this.eventi.indexOf(evento); 
                flag = false;
                break;
            }
        }
        if(flag)
            System.out.println("Evento non presente");
        else
            this.eventi.remove(index);
    }

    public ArrayList<Evento> getEventi(Date data){
    //EFFECTS: data una data, restituisce un elenco degli eventi in quella data
       ArrayList<Evento> tmp = new ArrayList<Evento>();   
        for (Evento evento : this.eventi) {
            if(evento.getData().equals(data)){
                tmp.add(evento);
            }
        }
        return tmp;
    }

    public void cloneEvento(Evento e, int offset){
    //MODIFIES: this
    //EFFECTS: metodo per copiare un `Evento` esistente a `n` giorni di distanza.
        int flag = this.eventi.indexOf(e);

        if(flag != -1){
            Evento clone = this.eventi.get(flag).copiaEvento(offset);
            this.addEvento(clone);
            return;
        }
        System.out.println("Evento non presente");
    }

    public String toString(){
        String tmp = "Calendario : (\n";
        for (Evento evento : this.eventi) {
            tmp += evento.toString() + "\n";
        }
        return tmp += "\n )";
    }



}
