import java.util.Date;
import java.util.Scanner;
import java.util.concurrent.TimeUnit;

public final class Evento {
//OVERVIEW: modella un evento di un calendario, definito da una data specificata e dal suo nome
//          Modelliamo il dato come immutabile
    private final Date data;
    private final String nome;

public static void main(String[] args) {
    Scanner s = new Scanner(System.in);
    
    System.out.println("Inserisci data del primo evento:");
    String[] tmp1 = s.next().split("/");
    int g1 = Integer.parseInt(tmp1[0]);
    int m1 = Integer.parseInt(tmp1[1]) - 1;
    int a1 = Integer.parseInt(tmp1[2]) - 1900;

    System.out.println("Inserisci il nome del primo evento:");
    String n1 = s.next();

    System.out.println("Inserisci data del secondo evento:");
    String[] tmp2 = s.next().split("/");
    int g2 = Integer.parseInt(tmp2[0]);
    int m2 = Integer.parseInt(tmp2[1]) - 1;
    int a2 = Integer.parseInt(tmp2[2]) - 1900;

    System.out.println("Inserisci il nome del secondo evento:");
    String n2 = s.next();

    Evento e1 = new Evento(new Date(a1,m1,g1), n1);
    Evento e2 = new Evento(new Date(a2,m2,g2), n2);

    if(e1.equals(e2)) System.out.println("Sono uguali");
}


//CONSTRUCTORS
    public Evento(Date data, String nome) throws IllegalArgumentException{
    //MODIFIES: this
    //EFFECTS: inizializza l'oggetto Evento, se la data e/o la stringa sono nulli lancia IllegalArgumentException
        if  (nome == null) throw new IllegalArgumentException("Nome nullo"); 
        if  (data == null) throw new IllegalArgumentException("Data Nulla");
        if  (nome == "") throw new IllegalArgumentException("Nome vuoto");

        this.data = (Date)(data.clone());
        this.nome = nome;

        assert repOk();
    }


//METHODS
    public String getNome(){
    //EFFECTS: restituisce il nome dell evento
        return new String(nome);
    }

    public Date getData(){
    //EFFECTS: restituisce la data dell'evento
        return (Date)(this.data.clone());
    }

    public Evento copiaEvento(int n){
    //EFFECTS: metodo che creae restituisce un nuovo evento con lo stesso nome ma a `n` giorni di distanza,
        Date newDate = new Date(this.data.getTime() + TimeUnit.DAYS.toMillis(n));
        return new Evento(newDate, this.nome);
    } 

    public String toString(){
        return "Evento:(nome:  " + this.nome + ", data: " + this.data + ")\n";
    }

    private boolean repOk(){
        if(data == null) return false;
        if(nome == null) return false;
        if(nome.equals("")) return false;
        
        return true;
    }

    @Override
    public boolean equals(Object obj) {
        if(obj == null) return false;
        if(! (obj instanceof Evento)) return false;

        Evento tmp = (Evento)obj;

        if(tmp.data.equals(this.data) && (tmp.nome.equals(this.nome))) return true;

        return false;
    }


}
